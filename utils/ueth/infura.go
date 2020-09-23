package ueth

import (
	"context"
	"github.com/gisvr/deif-common/utils/ulog"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Noop for unsupported chain id.
func BroadcastTxViaInfura(ctx context.Context, chainId *big.Int, tx *types.Transaction) {
	if !chainId.IsInt64() {
		return
	}

	// 使用自行申请的key替换掉原有的"ZOfDzLOSnNxb7l9ayaHt"。
	var url string
	switch chainId.Int64() {
	case 1:
		url = "https://mainnet.infura.io/v3/257218f7cb5f4419891312becd996ff1"
	case 3:
		url = "https://ropsten.infura.io/v3/257218f7cb5f4419891312becd996ff1"
	case 4:
		url = "https://rinkeby.infura.io/v3/257218f7cb5f4419891312becd996ff1"
	case 42:
		url = "https://kovan.infura.io/v3/257218f7cb5f4419891312becd996ff1"
	default:
		return
	}

	ethClient, err := ethclient.DialContext(ctx, url)
	if err != nil {
		ulog.Errorf("Failed to dial Infura chainId %v, url %v. Err: %v", chainId, url, err)
		return
	}

	err = ethClient.SendTransaction(ctx, tx)
	if err != nil {
		ulog.Errorf("Failed to SendTransaction via Infura for chainId %v, tx %v. Err: %v",
			chainId, tx.Hash().Hex(), err)
		return
	}

	ulog.Info("Tx broadcast via Infura: ", tx.Hash().Hex())
	return
}
