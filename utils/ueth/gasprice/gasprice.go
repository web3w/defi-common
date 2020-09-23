package gasprice

import (
	"github.com/gisvr/deif-common/utils/ubig"
	"github.com/gisvr/deif-common/utils/ulog"
	"github.com/gisvr/deif-common/utils/utime"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	lastGasStationReadingValue     atomic.Value
	lastEthNodeSuggestedPriceValue atomic.Value

	//defaultStdGasPrice  = big.NewInt(5e9)
	//defaultFastGasPrice = big.NewInt(10e9)
	//
	//maxGasPrice = big.NewInt(200e9)

	// 将默认的各级别手续费调低。
	defaultStdGasPrice  = big.NewInt(3e9)
	defaultFastGasPrice = big.NewInt(6e9)

	maxGasPrice = big.NewInt(50e9)
)

var (
	enableInternetSourceOnce sync.Once
	enableEthNodeSourceOnce  sync.Once
	waitForInternetSource    = sync.NewCond(&sync.Mutex{})
	internetSourceReady      = false
)

// Mainnet only
func EnableInternetSource() {
	enableInternetSourceOnce.Do(func() {
		ulog.Info("Internet source of gas price enabled")
		go watchInternetSource()
	})
}

// Mainnet only
func WaitForInternetSource() {
	waitForInternetSource.L.Lock()
	for !internetSourceReady {
		waitForInternetSource.Wait()
	}
	waitForInternetSource.L.Unlock()
}

func EnableEthNodeSource(ethClient *ethclient.Client) {
	enableEthNodeSourceOnce.Do(func() {
		ulog.Info("EthNode source of gas price enabled")
		go watchEthNodeSource(ethClient)
	})
}

func GetAverageGasPrice() (result *big.Int) {
	if x := lastGasStationReadingValue.Load(); x != nil {
		result = x.(gasStationReading).Average
	} else if x := lastEthNodeSuggestedPriceValue.Load(); x != nil {
		result = x.(*big.Int)
	} else {
		result = defaultStdGasPrice
	}

	// result *= 1.001 to make the gas price more competitive than the exact average.
	result = ubig.Add(result, ubig.Quo(result, ubig.U64(1000)))

	if result.Cmp(maxGasPrice) > 0 {
		result = maxGasPrice
	}
	return ubig.Clone(result) // clone it just in case it is modified outside
}

func GetFastGasPrice() (result *big.Int) {
	if x := lastGasStationReadingValue.Load(); x != nil {
		result = x.(gasStationReading).Fast
	} else if x := lastEthNodeSuggestedPriceValue.Load(); x != nil {
		suggested := x.(*big.Int)
		result = ubig.Add(suggested, big.NewInt(1e9))
	} else {
		result = defaultFastGasPrice
	}

	// result *= 1.001 to make the gas price more competitive than the exact average.
	result = ubig.Add(result, ubig.Quo(result, ubig.U64(1000)))

	if result.Cmp(maxGasPrice) > 0 {
		result = maxGasPrice
	}
	return ubig.Clone(result) // clone it just in case it is modified outside
}

func watchInternetSource() {
	for {
		reading, err := FetchGasStationReading()
		if err != nil {
			ulog.Errorf("Failed to get price from GasStation: %v", err)
			time.Sleep(15 * time.Second)
			continue
		}

		lastGasStationReadingValue.Store(reading)
		if !internetSourceReady {
			waitForInternetSource.L.Lock()
			internetSourceReady = true
			waitForInternetSource.L.Unlock()
			waitForInternetSource.Signal()
		}
		ulog.Infof("Last GasStation price reading: %+v", reading)
		time.Sleep(120 * time.Second)
	}
}

func watchEthNodeSource(ethClient *ethclient.Client) {
	for {
		suggested, err := ethClient.SuggestGasPrice(utime.CtxWithTimeoutMs(8e3))
		if err != nil {
			ulog.Errorf("Failed to ethClient.SuggestGasPrice: %v", err)
			time.Sleep(15 * time.Second)
			continue
		}

		lastEthNodeSuggestedPriceValue.Store(suggested)
		ulog.Info("Last suggested gas price: ", suggested)
		time.Sleep(120 * time.Second)
	}
}
