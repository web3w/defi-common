package dexinit

import (
	"github.com/gisvr/deif-common/utils/ueth"
	"github.com/gisvr/deif-common/utils/ulog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Env string

const (
	Mainnet Env = "mainnet"
	Testnet Env = "testnet"
	Devnet  Env = "devnet"
	Local   Env = "local"
	Testing Env = "testing"
)

// Returns `ok = false` if dexcfg.env config is missing or is not a canonical env.
func GetDexcfgEnv() (env Env, ok bool) {
	env = Env(viper.GetString("dexcfg.env"))
	switch env {
	case Mainnet:
		return Mainnet, true
	case Testnet:
		return Testnet, true
	case Devnet:
		return Devnet, true
	case Local:
		return Local, true
	case Testing:
		return Testing, true
	default:
		return env, false
	}
}

// It checks normalization in case it is got somewhere directly via viper and used as an id/key.
func MustGetMarketAddrFromDexcfg() common.Address {
	marketAddr := viper.GetString("dexcfg.dex2Contract.marketAddr")
	if marketAddr == "" {
		ulog.Panic("dexcfg.dex2Contract.marketAddr is missing.")
	}
	if !ueth.IsNormalizedHexAddr(marketAddr) {
		ulog.Panic("dexcfg.dex2Contract.marketAddr is not normalized: ", marketAddr)
	}
	return common.HexToAddress(marketAddr)
}

func MustGetChainIdFromDexcfg() int64 {
	chainId := viper.GetInt64("dexcfg.dex2Contract.chainId")
	if chainId <= 0 {
		ulog.Panic("Invalid dexcfg.dex2Contract.chainId: ", chainId)
	}
	return chainId
}

func DialGrpcUsingDexcfgCert(grpcServiceEndpoint string) (*grpc.ClientConn, error) {
	certFile := viper.GetString("dexcfg.grpc.tls.cert")
	if certFile == "" {
		return grpc.Dial(grpcServiceEndpoint, grpc.WithInsecure())
	}

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, err
	}
	return grpc.Dial(grpcServiceEndpoint, grpc.WithTransportCredentials(creds))
}

func MustMatchDexEnvAndChainId(chainId *big.Int) {
	MustMatchEnvAndChainId(viper.GetString("dexcfg.env"), chainId)
}

func MustMatchEnvAndChainId(env string, chainId *big.Int) {
	if !chainId.IsInt64() {
		ulog.Fatal("Invalid chainId: ", chainId)
	}
	chainName := ueth.ChainNameById(chainId.Int64())

	if (env == "mainnet") != (chainName == "Mainnet") {
		ulog.Fatalf("Env and ChainId mismatch. %v, %v", env, chainId)
	}
	if (env == "devnet" || env == "testnet") && chainName == "Unknown" {
		ulog.Fatalf("Env and ChainId mismatch. %v, %v", env, chainId)
	}
	if env == "local" && chainName != "Unknown" {
		ulog.Fatalf("Env and ChainId mismatch. %v, %v", env, chainId)
	}
}
