package dex2

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/gisvr/defi-common/utils/utime"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// TODO:
//   OpSequence's APIs are too heavy. Consider use it only for testing only. We need a lighter
// operation sequence builder with lower level API for cross-repo code reusing.
type OpSequence struct {
	marketAddr      common.Address
	newLogicTimeSec uint64
	beginIndex      uint64
	numOps          int

	body []*big.Int
}

func NewOpSequence(
	marketAddr common.Address, newLogicTimeSec uint64, beginIndex uint64) *OpSequence {

	utime.AssertRealTimeSec(int64(newLogicTimeSec))
	return &OpSequence{
		marketAddr:      marketAddr,
		newLogicTimeSec: newLogicTimeSec,
		beginIndex:      beginIndex,
		numOps:          0,
		body:            nil,
	}
}

func (seq *OpSequence) Header() *big.Int {
	return NewOpSeqHeader(seq.newLogicTimeSec, seq.beginIndex)
}

func (seq *OpSequence) Body() []*big.Int {
	return seq.body
}

func (seq *OpSequence) NumOps() int {
	return seq.numOps
}

// <depositIndex>(64) <opcode>(16)
func (seq *OpSequence) AddConfirmDeposit(depositIndex uint64) {
	seq.numOps++
	u := new(big.Int)
	PushUint64(u, depositIndex)
	PushUint16(u, 0xDE01)
	seq.body = append(seq.body, u)
}

// <amountE8>(64) <tokenCode>(16) <traderAddr>(160) <opcode>(16)
func (seq *OpSequence) AddInitiateWithdraw(
	traderAddr common.Address, tokenCode uint16, amountE8 uint64) {

	seq.numOps++
	u := new(big.Int)
	PushUint64(u, amountE8)
	PushUint16(u, tokenCode)
	PushAddress(u, traderAddr)
	PushUint16(u, 0xDE02)
	seq.body = append(seq.body, u)
}

// `makerOrder` is nil means that the maker order is already in storage, in which case
// `makerPrivKey` is unused (thus can be nil). Ditto for `takerOrder`.
func (seq *OpSequence) AddMatchOrders(
	maker *bind.TransactOpts, makerPrivKey *ecdsa.PrivateKey, makerNonce uint64, makerOrder *Order,
	taker *bind.TransactOpts, takerPrivKey *ecdsa.PrivateKey, takerNonce uint64, takerOrder *Order) {

	seq.numOps++
	makerUints := encodeTraderOrder(maker, makerPrivKey, seq.marketAddr, makerNonce, makerOrder)
	takerUints := encodeTraderOrder(taker, takerPrivKey, seq.marketAddr, takerNonce, takerOrder)
	PushUint16(makerUints[0], 0xDE03)
	seq.body = append(seq.body, makerUints...)
	seq.body = append(seq.body, takerUints...)
}

// `makerOrder` is nil means that the maker order is already in storage, in which case
// `makerPrivKey` is unused (thus can be nil). Ditto for `takerOrder`.
func (seq *OpSequence) AddTypesSignedMatchOrders(
	maker *bind.TransactOpts, makerPrivKey *ecdsa.PrivateKey, makerNonce uint64, makerOrder *Order,
	taker *bind.TransactOpts, takerPrivKey *ecdsa.PrivateKey, takerNonce uint64, takerOrder *Order) {

	seq.numOps++
	makerUints := encodeTypesSignedTraderOrder(maker, makerPrivKey, seq.marketAddr, makerNonce, makerOrder)
	takerUints := encodeTypesSignedTraderOrder(taker, takerPrivKey, seq.marketAddr, takerNonce, takerOrder)
	PushUint16(makerUints[0], 0xDE03)
	seq.body = append(seq.body, makerUints...)
	seq.body = append(seq.body, takerUints...)
}

// <nonce>(64) <traderAddr>(160) <opcode>(16)
func (seq *OpSequence) AddHardCancelOrder(traderAddr common.Address, nonce uint64) {
	seq.numOps++

	u := new(big.Int)
	PushUint64(u, nonce)
	PushAddress(u, traderAddr)
	PushUint16(u, 0xDE04)
	seq.body = append(seq.body, u)
}

// <withdrawFeeRateE4>(16) <takerFeeRateE4>(16) <makerFeeRateE4>(16) <opcode>(16)
func (seq *OpSequence) AddSetFeeRates(makerFeeRateE4, takerFeeRateE4, withdrawFeeRateE4 uint16) {
	seq.numOps++
	u := new(big.Int)
	PushUint16(u, withdrawFeeRateE4)
	PushUint16(u, takerFeeRateE4)
	PushUint16(u, makerFeeRateE4)
	PushUint16(u, 0xDE05)
	seq.body = append(seq.body, u)
}

// <feeRebatePercent>(8) <traderAddr>(160) <opcode>(16)
func (seq *OpSequence) AddSetFeeRebatePercent(traderAddr common.Address, feeRebatePercent uint8) {
	seq.numOps++
	u := new(big.Int)
	PushUint8(u, feeRebatePercent)
	PushAddress(u, traderAddr)
	PushUint16(u, 0xDE06)
	seq.body = append(seq.body, u)
}

func NewOpSeqHeader(newLogicTimeSec uint64, beginIndex uint64) *big.Int {
	header := new(big.Int).SetUint64(newLogicTimeSec)
	PushUint64(header, beginIndex)
	return header
}

// `order` is nil means that this order is already in storage, in which case `sig` ignored.
// The `sig` has format [R || S || V] where V is 0 or 1.
func EncodeOrderForMatching(
	traderAddr common.Address, nonce uint64, order *Order, sig []byte) ([]*big.Int, error) {

	var r, s *big.Int
	var v uint8 = 0

	if order != nil {
		if len(sig) != 65 || sig[64] > 1 {
			return nil, fmt.Errorf("invalid sig format %v", sig)
		}
		r = new(big.Int).SetBytes(sig[:32])
		s = new(big.Int).SetBytes(sig[32:64])
		// Refer: https://bitcoin.stackexchange.com/questions/38351/ecdsa-v-r-s-what-is-v
		v = sig[64] + 27
	}

	u1 := new(big.Int)
	PushUint64(u1, nonce)
	PushAddress(u1, traderAddr)
	PushUint8(u1, v)

	result := []*big.Int{u1}
	if order != nil {
		u2 := new(big.Int)
		PushUint64(u2, order.ExpireTimeSec)
		PushUint64(u2, order.AmountE8)
		PushUint64(u2, order.PriceE8)
		PushUint8(u2, order.Ioc)
		PushUint8(u2, order.Action)
		PushUint32(u2, order.PairId)

		result = append(result, u2)
		result = append(result, r)
		result = append(result, s)
	}
	return result, nil
}

// `order` is nil means that this order is already in storage.
func encodeTraderOrder(
	trader *bind.TransactOpts, traderPrivKey *ecdsa.PrivateKey, marketAddr common.Address,
	nonce uint64, order *Order) []*big.Int {

	var sig []byte
	if order != nil {
		bytesToSign := GetOrderBytesToSign(marketAddr, nonce, order)
		hash := crypto.Keccak256(bytesToSign)
		var err error
		sig, err = crypto.Sign(hash, traderPrivKey)
		if err != nil {
			panic(err)
		}
	}

	result, err := EncodeOrderForMatching(trader.From, nonce, order, sig)
	if err != nil {
		panic(err)
	}
	return result
}

// `order` is nil means that this order is already in storage.
func encodeTypesSignedTraderOrder(
	trader *bind.TransactOpts, traderPrivKey *ecdsa.PrivateKey, marketAddr common.Address,
	nonce uint64, order *Order) []*big.Int {

	var sig []byte
	if order != nil {
		bytesToSign := GetOrderBytesToSignWithTypes(marketAddr, nonce, order)
		hashValues := crypto.Keccak256(bytesToSign)

		hash := crypto.Keccak256(OrderHashTypes, hashValues)

		var err error
		sig, err = crypto.Sign(hash, traderPrivKey)
		if err != nil {
			panic(err)
		}
	}

	result, err := EncodeOrderForMatching(trader.From, nonce, order, sig)
	if err != nil {
		panic(err)
	}
	return result
}
