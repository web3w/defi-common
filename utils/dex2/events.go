package dex2

import (
	"github.com/gisvr/defi-common/utils/contract"
	"github.com/gisvr/defi-common/utils/ueth"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	dex2LogDecoder *ueth.EvmLogDecoder

	DeployMarketEventTopic     common.Hash
	InitiateWithdrawEventTopic common.Hash
)

func init() {
	var err error
	dex2LogDecoder, err = ueth.NewEvmLogDecoder(contract.Dex2ABI)
	if err != nil {
		panic(err)
	}

	DeployMarketEventTopic = Dex2Abi.Events["DeployMarketEvent"].ID()
	InitiateWithdrawEventTopic = Dex2Abi.Events["InitiateWithdrawEvent"].ID()
}

// Decode an EVM log to a strongly typed DEx2 event. Returns an error if the log cannot be decoded
// as a DEx2 event (e.g. a log emitted from another smart contract).
func DecodeEvent(log *types.Log) (interface{}, error) {
	eventName, args, err := dex2LogDecoder.Decode(log)
	if err != nil {
		return nil, err
	}

	switch eventName {
	case "DeployMarketEvent":
		return DeployMarketEvent{}, nil
	case "ChangeMarketStatusEvent":
		return ChangeMarketStatusEvent{
			Status: args["status"].(uint8),
		}, nil
	case "SetTokenInfoEvent":
		minDeposit := args["minDeposit"].(*big.Int)
		if minDeposit.Cmp(big.NewInt(0)) == 0 {
			// Avoid returning a big 0 that internally has a non-nil empty slice, making it not
			// DeepEqual to big.NewInt(0), which internally has a nil slice.
			minDeposit = big.NewInt(0)
		}
		return SetTokenInfoEvent{
			TokenCode:   args["tokenCode"].(uint16),
			Symbol:      args["symbol"].(string),
			TokenAddr:   args["tokenAddr"].(common.Address),
			ScaleFactor: args["scaleFactor"].(uint64),
			MinDeposit:  minDeposit,
		}, nil
	case "SetWithdrawAddrEvent":
		return SetWithdrawAddrEvent{
			Trader:       args["trader"].(common.Address),
			WithdrawAddr: args["withdrawAddr"].(common.Address),
		}, nil
	case "DepositEvent":
		return DepositEvent{
			Trader:       args["trader"].(common.Address),
			TokenCode:    args["tokenCode"].(uint16),
			Symbol:       args["symbol"].(string),
			AmountE8:     args["amountE8"].(uint64),
			DepositIndex: args["depositIndex"].(uint64),
		}, nil
	case "WithdrawEvent":
		return WithdrawEvent{
			Trader:      args["trader"].(common.Address),
			TokenCode:   args["tokenCode"].(uint16),
			Symbol:      args["symbol"].(string),
			AmountE8:    args["amountE8"].(uint64),
			LastOpIndex: args["lastOpIndex"].(uint64),
		}, nil
	case "TransferFeeEvent":
		return TransferFeeEvent{
			TokenCode: args["tokenCode"].(uint16),
			AmountE8:  args["amountE8"].(uint64),
			ToAddr:    args["toAddr"].(common.Address),
		}, nil
	case "ConfirmDepositEvent":
		return ConfirmDepositEvent{
			Trader:    args["trader"].(common.Address),
			TokenCode: args["tokenCode"].(uint16),
			BalanceE8: args["balanceE8"].(uint64),
		}, nil
	case "InitiateWithdrawEvent":
		return InitiateWithdrawEvent{
			Trader:            args["trader"].(common.Address),
			TokenCode:         args["tokenCode"].(uint16),
			AmountE8:          args["amountE8"].(uint64),
			PendingWithdrawE8: args["pendingWithdrawE8"].(uint64),
		}, nil
	case "MatchOrdersEvent":
		return MatchOrdersEvent{
			Trader1: args["trader1"].(common.Address),
			Nonce1:  args["nonce1"].(uint64),
			Trader2: args["trader2"].(common.Address),
			Nonce2:  args["nonce2"].(uint64),
		}, nil
	case "HardCancelOrderEvent":
		return HardCancelOrderEvent{
			Trader: args["trader"].(common.Address),
			Nonce:  args["nonce"].(uint64),
		}, nil
	case "SetFeeRatesEvent":
		return SetFeeRatesEvent{
			MakerFeeRateE4:    args["makerFeeRateE4"].(uint16),
			TakerFeeRateE4:    args["takerFeeRateE4"].(uint16),
			WithdrawFeeRateE4: args["withdrawFeeRateE4"].(uint16),
		}, nil
	case "SetFeeRebatePercentEvent":
		return SetFeeRebatePercentEvent{
			Trader:           args["trader"].(common.Address),
			FeeRebatePercent: args["feeRebatePercent"].(uint8),
		}, nil
	default:
		logger.Panic("unexpected DEx2 event name: ", eventName)
	}
	panic("Unreachable")
}
