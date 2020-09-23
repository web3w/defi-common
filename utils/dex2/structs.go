package dex2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//----------------------- Structs in Dex2 Smart Contract -----------------------

type TokenInfo struct {
	Symbol      string
	TokenAddr   common.Address
	ScaleFactor uint64
	MinDeposit  *big.Int
}

type TraderInfo struct {
	WithdrawAddr     common.Address
	FeeRebatePercent uint8
}

type TokenAccount struct {
	BalanceE8         uint64
	PendingWithdrawE8 uint64
}

type Order struct {
	PairId        uint32
	Action        uint8
	Ioc           uint8
	PriceE8       uint64
	AmountE8      uint64
	ExpireTimeSec uint64
}

type Deposit struct {
	TraderAddr      common.Address
	TokenCode       uint16
	PendingAmountE8 uint64
}

type ExeStatus struct {
	LogicTimeSec       uint64
	LastOperationIndex uint64
}

//----------------------- Composit Map Key Types -------------------------------

type TokenAccountKey struct {
	TraderAddr common.Address
	TokenCode  uint16
}

func (key *TokenAccountKey) Big() *big.Int {
	return GetTokenAccountKey(key.TraderAddr, key.TokenCode)
}

//---------------------- External Method Inputs --------------------------------

type DepositEthInput struct {
	TraderAddr common.Address
}

type DepositTokenInput struct {
	TraderAddr     common.Address
	TokenCode      uint16
	OriginalAmount *big.Int
}

type WithdrawEthInput struct {
	TraderAddr common.Address
}

type WithdrawTokenInput struct {
	TraderAddr common.Address
	TokenCode  uint16
}

type ExeSequenceInput struct {
	Header *big.Int
	Body   []*big.Int
}

//----------------------------- Events -----------------------------------------

type DeployMarketEvent struct {
}

type ChangeMarketStatusEvent struct {
	Status uint8
}

type SetTokenInfoEvent struct {
	TokenCode   uint16
	Symbol      string
	TokenAddr   common.Address
	ScaleFactor uint64
	MinDeposit  *big.Int
}

type SetWithdrawAddrEvent struct {
	Trader       common.Address
	WithdrawAddr common.Address
}

type DepositEvent struct {
	Trader       common.Address
	TokenCode    uint16
	Symbol       string
	AmountE8     uint64
	DepositIndex uint64
}

type WithdrawEvent struct {
	Trader      common.Address
	TokenCode   uint16
	Symbol      string
	AmountE8    uint64
	LastOpIndex uint64
}

type TransferFeeEvent struct {
	TokenCode uint16
	AmountE8  uint64
	ToAddr    common.Address
}

type ConfirmDepositEvent struct {
	Trader    common.Address
	TokenCode uint16
	BalanceE8 uint64
}

type InitiateWithdrawEvent struct {
	Trader            common.Address
	TokenCode         uint16
	AmountE8          uint64
	PendingWithdrawE8 uint64
}

type MatchOrdersEvent struct {
	Trader1 common.Address
	Nonce1  uint64
	Trader2 common.Address
	Nonce2  uint64
}

type HardCancelOrderEvent struct {
	Trader common.Address
	Nonce  uint64
}

type SetFeeRatesEvent struct {
	MakerFeeRateE4    uint16
	TakerFeeRateE4    uint16
	WithdrawFeeRateE4 uint16
}

type SetFeeRebatePercentEvent struct {
	Trader           common.Address
	FeeRebatePercent uint8
}
