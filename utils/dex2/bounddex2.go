package dex2

import (
	"context"
	"fmt"
	"github.com/gisvr/defi-common/utils/ueth"
	"github.com/gisvr/defi-common/utils/ulog"
	"github.com/gisvr/defi-common/utils/utime"
	"math/big"
	"time"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// This supplements the auto-generated Dex2 bindings like Dex2Caller, Dex2Transactor.
type BoundDex2 struct {
	backend      bind.ContractBackend
	dex2Addr     common.Address
	rpcTimeoutMs int64
}

func NewBoundDex2(
	backend bind.ContractBackend, dex2Addr common.Address, rpcTimeout time.Duration) *BoundDex2 {

	return &BoundDex2{
		backend:      backend,
		dex2Addr:     dex2Addr,
		rpcTimeoutMs: int64(rpcTimeout / time.Millisecond),
	}
}

func (bd2 *BoundDex2) ExeStatusAtBlock(blockNumber int64) (ExeStatus, error) {
	result := ExeStatus{}

	opts := &bind.CallOpts{}
	input, err := Dex2Abi.Pack("exeStatus")
	if err != nil {
		return result, err
	}

	msg := geth.CallMsg{From: opts.From, To: &bd2.dex2Addr, Data: input}

	height := big.NewInt(blockNumber)
	output, err := bd2.backend.CallContract(bd2.newRpcCtx(), msg, height)
	if err != nil {
		return result, err
	}
	if len(output) == 0 {
		// Make sure we have a contract to operate on, and bail out otherwise.
		code, err := bd2.backend.CodeAt(bd2.newRpcCtx(), bd2.dex2Addr, height)
		if err != nil {
			return result, err
		}
		if len(code) == 0 {
			return result, bind.ErrNoCode
		}
	}

	err = Dex2Abi.Unpack(&result, "exeStatus", output)
	return result, err
}

type TokenInfoFull struct {
	TokenCode   uint16
	Symbol      string
	TokenAddr   common.Address
	ScaleFactor uint64
	MinDeposit  *big.Int
}

// Get info of all tokens that have been set in the market, including the lastest minDeposit.
func (bd2 *BoundDex2) GetTokenInfoMaps() (
	byCodes map[uint16]TokenInfoFull, bySymbols map[string]TokenInfoFull, retErr error) {

	query := geth.FilterQuery{
		Addresses: []common.Address{bd2.dex2Addr},
		Topics:    [][]common.Hash{{Dex2Abi.Events["SetTokenInfoEvent"].ID()}},
	}
	logs, err := ueth.QueryLogs(bd2.backend, bd2.newRpcCtx(), query)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to query SetTokenInfoEvent: %v", err)
	}

	byCodes = make(map[uint16]TokenInfoFull)
	bySymbols = make(map[string]TokenInfoFull)
	for _, log := range logs {
		x, err := DecodeEvent(&log)
		if err != nil {
			return nil, nil, err
		}
		event, ok := x.(SetTokenInfoEvent)
		if !ok {
			ulog.Panicf("Expecting %T, but got %T", event, x)
		}

		byCodes[event.TokenCode] = TokenInfoFull(event)
		bySymbols[event.Symbol] = TokenInfoFull(event)
	}
	return byCodes, bySymbols, nil
}

// Returns trader addresses that have been deposited to. This is a comprehensive list of trader
// addresses that possibly have non-zero balances of tokens.
func (bd2 *BoundDex2) GetDepositedTraderAddrs() ([]common.Address, error) {
	query := geth.FilterQuery{
		Addresses: []common.Address{bd2.dex2Addr},
		Topics:    [][]common.Hash{{Dex2Abi.Events["DepositEvent"].ID()}},
	}
	logs, err := ueth.QueryLogs(bd2.backend, bd2.newRpcCtx(), query)
	if err != nil {
		return nil, fmt.Errorf("Failed to query DepositEvent: %v", err)
	}

	var traderAddrs []common.Address
	has := make(map[common.Address]bool)
	for _, log := range logs {
		x, err := DecodeEvent(&log)
		if err != nil {
			return nil, err
		}
		event, ok := x.(DepositEvent)
		if !ok {
			ulog.Panicf("Expecting %T, but got %T", event, x)
		}

		if !has[event.Trader] {
			has[event.Trader] = true
			traderAddrs = append(traderAddrs, event.Trader)
		}
	}
	return traderAddrs, nil
}

func (bd2 *BoundDex2) newRpcCtx() context.Context {
	return utime.CtxWithTimeoutMs(bd2.rpcTimeoutMs)
}
