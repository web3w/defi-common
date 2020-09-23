package calc

import (
	"errors"
	"github.com/gisvr/deif-common/dbdata/trade"
	"math"
	"math/big"

	"github.com/gisvr/deif-common/utils/ubig"
	"github.com/gisvr/deif-common/utils/ulog"
)

// Get the amount (multiplied by 1e8) of cash required to buy the given amount of stock at the given
// price. In short, it calculates `stockAmountE8 * priceE8 / 1e8` in overflow-safe way.
func GetPaymentE8(stockAmountE8, priceE8 uint64) (uint64, error) {
	stockAmountE8Big := new(big.Int).SetUint64(stockAmountE8)
	priceE8Big := new(big.Int).SetUint64(priceE8)

	productE16 := new(big.Int).Mul(stockAmountE8Big, priceE8Big)
	paymentE8 := new(big.Int).Quo(productE16, big.NewInt(1e8))

	if !paymentE8.IsUint64() {
		return math.MaxUint64, errors.New("amount overflow")
	}
	return paymentE8.Uint64(), nil
}

func MustGetPaymentE8(stockAmountE8, priceE8 uint64, symbol string) uint64 {
	result, err := GetPaymentE8(stockAmountE8, priceE8)
	if err != nil {
		ulog.Panic("error in getPaymentE8: ", err)
	}
	return result
}

// Calculates the fee (multiplied by 1e8).
//
// - `amountE8` is the original amount (multiplied by 1e8)
// - `feeRateE4` is is fee rate (multiplied by 1e4)
func CalcFeeE8(amountE8 uint64, feeRateE4 uint64) uint64 {
	if feeRateE4 > 1e4 {
		ulog.Panic("feeRateE4 too large: ", feeRateE4)
	}
	// 先不做最小手续费限制
	return ubig.Quo(ubig.MulU64(amountE8, feeRateE4), big.NewInt(1e4)).Uint64()
}

// `order` is the order state before executing the deal.
func CalcUnlockCashAmountForDealE8(order *dbtrade.Orders, dealExeAmountE8 uint64, symbol string) uint64 {
	if order.AmountLeftE8 < dealExeAmountE8 {
		ulog.Panicf("Order amount left %v is less than the deal amount %v.",
			order.AmountLeftE8, dealExeAmountE8)
	}

	curLockedE8 := MustGetPaymentE8(order.AmountLeftE8, order.PriceE8, symbol)
	newLockedE8 := MustGetPaymentE8(order.AmountLeftE8-dealExeAmountE8, order.PriceE8, symbol)
	if curLockedE8 < newLockedE8 {
		ulog.Panicf("Locked amount must be smaller after a deal. %v, %v", curLockedE8, newLockedE8)
	}
	return curLockedE8 - newLockedE8
}

/**
 * 根据精度位数返回相应的值。
 */
func Decimal2E(decimals int32) uint64 {
	return uint64(math.Pow10(int(decimals)))
}

/**
 * 除以1E8后转为float类型。
 */
func E8ToFloat(num int64) float64 {
	return float64(num) / 1e8
}
