// Parse operation sequence.

package dex2

import (
	"bytes"
	"fmt"
	"github.com/gisvr/defi-common/utils/ubig"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Must with 0x prefix.
func ParseOpsFromTxInputHex(txInputHex string) (string, error) {
	data, err := hexutil.Decode(txInputHex)
	if err != nil {
		return "", err
	}
	if len(data) < 4 {
		return "", fmt.Errorf("Invalid data length")
	}
	method, err := Dex2Abi.MethodById(data[0:4])
	if err != nil {
		return "", err
	}
	if method.Name != "exeSequence" {
		return "", fmt.Errorf("Unexpected method name: %v", method.Name)
	}

	var input ExeSequenceInput
	err = method.Inputs.Unpack(&input, data[4:])
	if err != nil {
		return "", err
	}
	if len(input.Body) != len(data)/32-1-2 /*num-u256s - header - body-array-metadata*/ {
		return "", fmt.Errorf("Expecting len(Body) %v, but got %v", len(data)/32-1-2, len(input.Body))
	}
	return ParseOpsFromU256(input.Header, input.Body)
}

func ParseHeader(header *big.Int) (uint64, uint64) {
	header = new(big.Int).Set(header) // make a copy to avoid modifying the argument
	beginIndex := PopUint64(header)
	newLogicTimeSec := PopUint64(header)
	return beginIndex, newLogicTimeSec
}

// `header` can be empty string, in which case it skips parsing header.
// Returns partial result when there is an error.
//
// It does NOT require the length of the hex chars to be exactly 64.
func ParseOpsFromHex(header string, body []string) (string, error) {
	var headerBig *big.Int
	if len(header) != 0 {
		headerBig = ubig.MustHex(header)
	}
	bodyBig := make([]*big.Int, len(body))
	for i := range body {
		bodyBig[i] = ubig.MustHex(body[i])
	}
	return ParseOpsFromU256(headerBig, bodyBig)
}

// `header` can be nil, in which case it skips parsing header.
// Returns partial result when there is an error. The argument will NOT be modified.
func ParseOpsFromU256(header *big.Int, originalBody []*big.Int) (string, error) {
	// make a copy to avoid modifying the argument
	header = new(big.Int).Set(header)
	u256s := make([]*big.Int, len(originalBody))
	for i, u := range originalBody {
		u256s[i] = new(big.Int).Set(u)
	}

	out := new(bytes.Buffer)
	if len(u256s) == 0 {
		return "", fmt.Errorf("empty body")
	}

	if header != nil {
		// <newLogicTimeSec>(64) <beginIndex>(64)
		beginIndex := PopUint64(header)
		newLogicTimeSec := PopUint64(header)
		if header.BitLen() != 0 {
			return out.String(), fmt.Errorf("invalid header")
		}
		fmt.Fprintln(out, "newLogicTimeSec:", newLogicTimeSec)
		fmt.Fprintln(out, "beginIndex:", beginIndex)
	}
	fmt.Fprintln(out, "len(body):", len(u256s))

	for i := 0; i < len(u256s); {
		bits := u256s[i]
		opcode := PopUint16(bits)
		if (opcode >> 8) != 0xDE {
			return out.String(), fmt.Errorf("wrong magic number")
		}

		consumed := 1
		var err error
		switch opcode {
		case 0xDE01:
			err = parseConfirmDepositOp(out, bits)
		case 0xDE02:
			err = parseInitiateWithdrawOp(out, bits)
		case 0xDE03:
			consumed, err = parseMatchOrdersOp(out, bits, u256s[i+1:])
		case 0xDE04:
			err = parseHardCancelOrderOp(out, bits)
		case 0xDE05:
			err = parseSetFeeRatesOp(out, bits)
		case 0xDE06:
			err = parseSetFeeRebatePercentOp(out, bits)
		default:
			return out.String(), fmt.Errorf("invalid opcode %#x", opcode)
		}

		if err != nil {
			return out.String(), err
		}
		i += consumed
		if i < len(u256s) {
			fmt.Fprintln(out, "")
		}
	}
	return out.String(), nil
}

func parseConfirmDepositOp(out *bytes.Buffer, bits *big.Int) error {
	fmt.Fprintln(out, "operation ConfirmDeposit:")
	fmt.Fprintln(out, "  depositIndex: ", bits)
	return nil
}

func parseInitiateWithdrawOp(out *bytes.Buffer, bits *big.Int) error {
	fmt.Fprintln(out, "operation InitiateWithdraw:")
	// <amountE8>(64) <tokenCode>(16) <traderAddr>(160) <opcode>(16)
	fmt.Fprintf(out, "  traderAddr: %#x\n", PopUint160(bits))
	fmt.Fprintln(out, "  tokenCode:", PopUint16(bits))
	fmt.Fprintln(out, "  amountE8:", PopUint64(bits))
	return nil
}

func parseMatchOrdersOp(out *bytes.Buffer, bits *big.Int, rest []*big.Int) (consumed int, err error) {
	fmt.Fprintln(out, "operation MatchOrder:")
	v1 := PopUint8(bits)
	if v1 == 0 {
		fmt.Fprintln(out, "  makerOrder (existing):")
		if len(rest) < 1 {
			return 0, fmt.Errorf("not enough inputs for matching order")
		}
		consumed += 1
		if err := parseOrderOpOperand(out, bits, nil); err != nil {
			return consumed, err
		}
	} else {
		fmt.Fprintf(out, "  makerOrder (new, v1=%v):\n", v1)
		if v1 != 27 && v1 != 28 {
			return 0, fmt.Errorf("invalid v1: %v", v1)
		}
		if len(rest) < 4 {
			return 0, fmt.Errorf("not enough inputs for matching order")
		}
		consumed += 4
		if err := parseOrderOpOperand(out, bits, rest); err != nil {
			return consumed, err
		}
		rest = rest[3:]
	}

	bits, rest = rest[0], rest[1:]
	v2 := PopUint8(bits)
	if v2 == 0 {
		fmt.Fprintln(out, "  takerOrder (existing):")
		consumed += 1
		if err := parseOrderOpOperand(out, bits, nil); err != nil {
			return consumed, err
		}
	} else {
		fmt.Fprintf(out, "  takerOrder (new, v2=%v):\n", v2)
		if v2 != 27 && v2 != 28 {
			return 0, fmt.Errorf("invalid v2: %v", v2)
		}
		if len(rest) < 3 {
			return 0, fmt.Errorf("not enough inputs for matching order")
		}
		consumed += 4
		if err := parseOrderOpOperand(out, bits, rest); err != nil {
			return consumed, err
		}
	}

	return consumed, nil
}

// Precondition: `rest` is nil or has at least 3 elements.
func parseOrderOpOperand(out *bytes.Buffer, bits *big.Int, rest []*big.Int) error {
	fmt.Fprintf(out, "    trader: %#x\n", PopUint160(bits))
	fmt.Fprintln(out, "    nonce:", PopUint64(bits))
	if bits.BitLen() != 0 {
		return fmt.Errorf("extra bits in orderKey: %s", bits)
	}
	if rest == nil { // existing order
		return nil
	}

	// <expireTimeSec>(64) <amountE8>(64) <priceE8>(64) <ioc>(8) <action>(8) <pairId>(32)
	bits = rest[0]
	fmt.Fprintln(out, "    pairId  :", PopUint32(bits))
	fmt.Fprintln(out, "    action  :", PopUint8(bits))
	fmt.Fprintln(out, "    ioc     :", PopUint8(bits))
	fmt.Fprintln(out, "    priceE8 :", PopUint64(bits))
	fmt.Fprintln(out, "    amountE8:", PopUint64(bits))
	fmt.Fprintln(out, "    expire  :", PopUint64(bits))
	if bits.BitLen() != 0 {
		return fmt.Errorf("extra data in order bits: %#x", bits)
	}

	if rest[1].BitLen() == 0 {
		return fmt.Errorf("signature uint256 r is zero")
	}
	fmt.Fprintf(out, "    s       : %#x\n", rest[1])
	if rest[2].BitLen() == 0 {
		return fmt.Errorf("signature uint256 s is zero")
	}
	fmt.Fprintf(out, "    t       : %#x\n", rest[2])
	return nil
}

func parseHardCancelOrderOp(out *bytes.Buffer, bits *big.Int) error {
	fmt.Fprintln(out, "operation HardCancel:")
	fmt.Fprintln(out, "  <to be implemented>")
	return nil
}

func parseSetFeeRatesOp(out *bytes.Buffer, bits *big.Int) error {
	fmt.Fprintln(out, "operation SetFeeRates:")
	fmt.Fprintln(out, "  <to be implemented>")
	return nil
}

func parseSetFeeRebatePercentOp(out *bytes.Buffer, bits *big.Int) error {
	fmt.Fprintln(out, "operation SetFeeRebatePercent:")
	fmt.Fprintln(out, "  <to be implemented>")
	return nil
}
