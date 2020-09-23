package dbtrade

import (
	"github.com/gisvr/deif-common/utils/ulog"
	"strings"
)

func GetTradePairSymbol(cashTokenSymbol, stockTokenSymbol string) string {
	return cashTokenSymbol + "_" + stockTokenSymbol
}

func GetTradePairCode(cashTokenCode, stockTokenCode uint16) uint32 {
	return uint32(cashTokenCode)<<16 | uint32(stockTokenCode)
}

func PairCodeToCashStockCode(pairCode uint32) (cashTokenCode, stockTokenCode uint16) {
	return uint16(pairCode >> 16), uint16(pairCode)
}

func GetTradePairOfOrder(order *Orders) string {
	return GetTradePairSymbol(order.CashTokenSymbol, order.StockTokenSymbol)
}

// Check whether an order is partially filled BASED ON AMOUNT LEFT in it.
func IsPartiallyFilled(order *Orders) bool {
	return order.AmountLeftE8 < order.AmountTotalE8
}

// Check whether an order is unfilled BASED ON AMOUNT LEFT in it.
func IsUnfilled(order *Orders) bool {
	return order.AmountLeftE8 == order.AmountTotalE8
}

/**
 * 从币种对中拆解出币种。
 */
func PairIdToTokenIds(tokenPair string) (cashTokenId, stockTokenId string) {
	ids := strings.Split(tokenPair, "_")
	if len(ids) != 2 {
		ulog.Errorln("pair id format error")
		return
	}
	cashTokenId = ids[0]
	stockTokenId = ids[1]
	return
}
