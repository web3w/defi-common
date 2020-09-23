package orderbook_test

import (
	"github.com/gisvr/deif-common/dbdata/trade"
	"testing"
	"time"

	pb "github.com/gisvr/deif-common/proto/dex2"
	"github.com/gisvr/deif-common/utils/numeric"
	"github.com/gisvr/deif-common/utils/orderbook"
	"github.com/gisvr/deif-common/utils/utest"
	"github.com/gisvr/deif-common/utils/utime"
)

type OrderBookTest struct {
	utest.RequireSuite
	nextOrderId int64
}

// The hook of `go test`
func TestRun_OrderBookTest(t *testing.T) {
	utest.Run(t, &OrderBookTest{nextOrderId: 1000000})
}

func (t *OrderBookTest) TestExpireOrders() {
	ob := orderbook.NewRbtOrderBook("ETH_ADX")

	baseTimeSec := time.Now().Unix()
	t.ElementsMatch(nil, ob.PushLogicTimeSec(baseTimeSec))

	order1 := t.makeOrder(dbtrade.ORDERS_ACTION_BUY, 100, 1000, baseTimeSec+40)
	order2 := t.makeOrder(dbtrade.ORDERS_ACTION_BUY, 200, 1000, baseTimeSec+30)
	order3 := t.makeOrder(dbtrade.ORDERS_ACTION_BUY, 300, 2000, baseTimeSec+20)
	order4 := t.makeOrder(dbtrade.ORDERS_ACTION_BUY, 400, 2500, baseTimeSec+10)
	order5 := t.makeOrder(dbtrade.ORDERS_ACTION_BUY, 150, 1000, baseTimeSec+15)

	order6 := t.makeOrder(dbtrade.ORDERS_ACTION_SELL, 500, 1000, baseTimeSec+20)
	order7 := t.makeOrder(dbtrade.ORDERS_ACTION_SELL, 600, 1000, baseTimeSec+40)
	order8 := t.makeOrder(dbtrade.ORDERS_ACTION_SELL, 700, 2000, baseTimeSec+20)
	order9 := t.makeOrder(dbtrade.ORDERS_ACTION_SELL, 800, 2500, baseTimeSec+10)
	order10 := t.makeOrder(dbtrade.ORDERS_ACTION_SELL, 750, 1000, baseTimeSec+15)

	ob.Add(order3)
	ob.Add(order4)
	ob.Add(order5)
	ob.Add(order9)
	ob.Add(order10)
	ob.Add(order1)
	ob.Add(order2)
	ob.Add(order6)
	ob.Add(order8)
	ob.Add(order7)

	// Adding an already expired order will panic.
	t.Panics(func() {
		ob.Add(t.makeOrder(dbtrade.ORDERS_ACTION_BUY, 999, 1000, baseTimeSec-1))
	})

	t.ElementsMatch(
		[]*dbtrade.Orders{order1, order2, order3, order4, order5, order6, order7, order8, order9, order10},
		ob.ExportAllOrdersTo(nil))

	buy, ok := ob.GetTopBuyOrder()
	t.True(ok)
	t.Equal(order4, buy)

	sell, ok := ob.GetTopSellOrder()
	t.True(ok)
	t.Equal(order6, sell)

	t.True(ob.Delete(order5.OrderId))

	//----------------------------------------------------------------------------
	expired := ob.PushLogicTimeSec(baseTimeSec + 25)
	t.ElementsMatch([]*dbtrade.Orders{order3, order4, order6, order8, order9, order10}, expired)

	buy, ok = ob.GetTopBuyOrder()
	t.True(ok)
	t.Equal(order2, buy)

	sell, ok = ob.GetTopSellOrder()
	t.True(ok)
	t.Equal(order7, sell)

	t.ElementsMatch([]*dbtrade.Orders{order1, order2, order7}, ob.ExportAllOrdersTo(nil))

	//----------------------------------------------------------------------------
	expired = ob.PushLogicTimeSec(baseTimeSec + 28)
	t.ElementsMatch(nil, expired)
	expired = ob.PushLogicTimeSec(baseTimeSec + 30)
	t.ElementsMatch([]*dbtrade.Orders{order2}, expired)
	expired = ob.PushLogicTimeSec(baseTimeSec + 40)
	t.ElementsMatch([]*dbtrade.Orders{order1, order7}, expired)

	t.ElementsMatch([]*dbtrade.Orders{}, ob.ExportAllOrdersTo(nil))
	_, ok = ob.GetTopBuyOrder()
	t.False(ok)
	_, ok = ob.GetTopSellOrder()
	t.False(ok)
}

func (t *OrderBookTest) TestMergingOrders() {
	ob := orderbook.NewRbtOrderBook("ETH_ADX")

	baseTimeSec := time.Now().Unix()
	t.ElementsMatch(nil, ob.PushLogicTimeSec(baseTimeSec))

	order1 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_BUY, "0.10879001", "500", baseTimeSec+5)
	order2 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_BUY, "0.10865500", "300", baseTimeSec+5)
	order3 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_BUY, "0.10861189", "200", baseTimeSec+5)
	order4 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_BUY, "0.10861108", "200", baseTimeSec+5)
	order5 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_BUY, "0.10859001", "200", baseTimeSec+5)
	order6 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_BUY, "0.08854400", "100", baseTimeSec+5)

	order7 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "0.10889021", "300", baseTimeSec+5)
	order8 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "0.10889021", "100", baseTimeSec+5)
	order9 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "0.10889501", "2000", baseTimeSec+5)
	order10 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "0.10889531", "2000", baseTimeSec+5)
	order11 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "0.10899001", "300", baseTimeSec+5)
	order12 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "1.09999999", "150", baseTimeSec+5)
	order13 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "1.10000000", "30", baseTimeSec+5)
	order14 := t.makerOrderWithStringValue(dbtrade.ORDERS_ACTION_SELL, "1.10000001", "90", baseTimeSec+5)

	ob.Add(order3)
	ob.Add(order4)
	ob.Add(order5)
	ob.Add(order9)
	ob.Add(order10)
	ob.Add(order1)
	ob.Add(order2)
	ob.Add(order6)
	ob.Add(order8)
	ob.Add(order7)
	ob.Add(order13)
	ob.Add(order11)
	ob.Add(order14)
	ob.Add(order12)

	askEntryValue1 := t.calcDepthEntryValue("0.10889021", "400")
	askEntryValue2 := t.calcDepthEntryValue("0.10889501", "2000")
	askEntryValue3 := t.calcDepthEntryValue("0.10889531", "2000")
	askEntryValue4 := t.calcDepthEntryValue("0.10899001", "300")
	askEntryValue5 := t.calcDepthEntryValue("1.09999999", "150")
	askEntryValue6 := t.calcDepthEntryValue("1.10000000", "30")
	askEntryValue7 := t.calcDepthEntryValue("1.10000001", "90")

	bidEntryValue1 := t.calcDepthEntryValue("0.10879001", "500")
	bidEntryValue2 := t.calcDepthEntryValue("0.10865500", "300")
	bidEntryValue3 := t.calcDepthEntryValue("0.10861189", "200")
	bidEntryValue4 := t.calcDepthEntryValue("0.10861108", "200")
	bidEntryValue5 := t.calcDepthEntryValue("0.10859001", "200")
	bidEntryValue6 := t.calcDepthEntryValue("0.08854400", "100")

	askEntries, bidEntries := ob.GetDepth(7, 8) // no merging
	t.checkDepthEntries(askEntries,
		"0.10889021", "400", askEntryValue1,
		"0.10889501", "2000", askEntryValue2,
		"0.10889531", "2000", askEntryValue3,
		"0.10899001", "300", askEntryValue4,
		"1.09999999", "150", askEntryValue5,
		"1.10000000", "30", askEntryValue6,
		"1.10000001", "90", askEntryValue7)
	t.checkDepthEntries(bidEntries,
		"0.10879001", "500", bidEntryValue1,
		"0.10865500", "300", bidEntryValue2,
		"0.10861189", "200", bidEntryValue3,
		"0.10861108", "200", bidEntryValue4,
		"0.10859001", "200", bidEntryValue5,
		"0.08854400", "100", bidEntryValue6)

	askEntries, bidEntries = ob.GetDepth(7, 7) // merging with 7 decimals
	t.checkDepthEntries(askEntries,
		"0.10889030", "400", askEntryValue1,
		"0.10889510", "2000", askEntryValue2,
		"0.10889540", "2000", askEntryValue3,
		"0.10899010", "300", askEntryValue4,
		"1.10000000", "180", askEntryValue5+askEntryValue6,
		"1.10000010", "90", askEntryValue7)
	t.checkDepthEntries(bidEntries,
		"0.10879000", "500", bidEntryValue1,
		"0.10865500", "300", bidEntryValue2,
		"0.10861180", "200", bidEntryValue3,
		"0.10861100", "200", bidEntryValue4,
		"0.10859000", "200", bidEntryValue5,
		"0.08854400", "100", bidEntryValue6)

	askEntries, bidEntries = ob.GetDepth(5, 6) // merging with 6 decimals
	t.checkDepthEntries(askEntries,
		"0.10889100", "400", askEntryValue1,
		"0.10889600", "4000", askEntryValue2+askEntryValue3,
		"0.10899100", "300", askEntryValue4,
		"1.10000000", "180", askEntryValue5+askEntryValue6,
		"1.10000100", "90", askEntryValue7)
	t.checkDepthEntries(bidEntries,
		"0.10879000", "500", bidEntryValue1,
		"0.10865500", "300", bidEntryValue2,
		"0.10861100", "400", bidEntryValue3+bidEntryValue4,
		"0.10859000", "200", bidEntryValue5,
		"0.08854400", "100", bidEntryValue6)

	askEntries, bidEntries = ob.GetDepth(5, 5) // merging with 5 decimals
	t.checkDepthEntries(askEntries,
		"0.10890000", "4400", askEntryValue1+askEntryValue2+askEntryValue3,
		"0.10900000", "300", askEntryValue4,
		"1.10000000", "180", askEntryValue5+askEntryValue6,
		"1.10001000", "90", askEntryValue7)
	t.checkDepthEntries(bidEntries,
		"0.10879000", "500", bidEntryValue1,
		"0.10865000", "300", bidEntryValue2,
		"0.10861000", "400", bidEntryValue3+bidEntryValue4,
		"0.10859000", "200", bidEntryValue5,
		"0.08854000", "100", bidEntryValue6)

	askEntries, bidEntries = ob.GetDepth(5, 4) // merging with 4 decimals
	t.checkDepthEntries(askEntries,
		"0.10890000", "4400", askEntryValue1+askEntryValue2+askEntryValue3,
		"0.10900000", "300", askEntryValue4,
		"1.10000000", "180", askEntryValue5+askEntryValue6,
		"1.10010000", "90", askEntryValue7)
	t.checkDepthEntries(bidEntries,
		"0.10870000", "500", bidEntryValue1,
		"0.10860000", "700", bidEntryValue2+bidEntryValue3+bidEntryValue4,
		"0.10850000", "200", bidEntryValue5,
		"0.08850000", "100", bidEntryValue6)

	askEntries, bidEntries = ob.GetDepth(5, 2) // merging with 2 decimals
	t.checkDepthEntries(askEntries,
		"0.11000000", "4700", askEntryValue1+askEntryValue2+askEntryValue3+askEntryValue4,
		"1.10000000", "180", askEntryValue5+askEntryValue6,
		"1.11000000", "90", askEntryValue7)
	t.checkDepthEntries(bidEntries,
		"0.10000000", "1400", bidEntryValue1+bidEntryValue2+bidEntryValue3+bidEntryValue4+bidEntryValue5,
		"0.08000000", "100", bidEntryValue6)
}

// Check depth entries.

// `args`: depthEntryPriceA, depthEntryAmountA, depthEntryTotalValueA,
//         depthEntryPriceB, depthEntryAmountB, depthEntryTotalValueB, ...
func (t *OrderBookTest) checkDepthEntries(actualDepthEntries []*pb.DepthEntry, args ...interface{}) {
	var expectedDepthEntries []*pb.DepthEntry
	for i := 0; i < len(args); i += 3 {
		expectedDepthEntries = append(expectedDepthEntries, &pb.DepthEntry{
			PriceE8:    t.mustStringToUint64E8(args[i].(string)),
			AmountE8:   t.mustStringToUint64E8(args[i+1].(string)),
			TotalValue: args[i+2].(float64),
		})
	}

	t.Equal(len(expectedDepthEntries), len(actualDepthEntries))
	t.Subset(actualDepthEntries, expectedDepthEntries)
	t.Subset(expectedDepthEntries, actualDepthEntries)
}

func (t *OrderBookTest) makerOrderWithStringValue(
	action int8, price, amount string, expireTimeSec int64) *dbtrade.Orders {

	priceE8 := t.mustStringToUint64E8(price)
	amountE8 := t.mustStringToUint64E8(amount)
	return t.makeOrder(action, priceE8, amountE8, expireTimeSec)
}

func (t *OrderBookTest) mustStringToUint64E8(str string) uint64 {
	vUint64, err := numeric.StringToUint64E8(str)
	t.NoError(err)
	return vUint64
}

func (t *OrderBookTest) calcDepthEntryValue(entryAmount, entryPrice string) float64 {
	entryAmountE8 := t.mustStringToUint64E8(entryAmount)
	entryPriceE8 := t.mustStringToUint64E8(entryPrice)

	return float64(entryPriceE8) / 1e8 * float64(entryAmountE8) / 1e8
}

func (t *OrderBookTest) makeOrder(action int8, priceE8, amountE8 uint64,
	expireTimeSec int64) *dbtrade.Orders {
	t.nextOrderId++

	return &dbtrade.Orders{
		OrderId:           t.nextOrderId - 1,
		Trader:            "0xFAC12345B496bB74e6D0cBc2D1a7061B80881f7c",
		Action:            action,
		CashTokenSymbol:   "ETH",
		StockTokenSymbol:  "ADX",
		PriceE8:           priceE8,
		AmountTotalE8:     amountE8,
		CreateTime:        utime.NowMs(),
		ExpireTime:        expireTimeSec,
		OrderSign:         "",
		AmountLeftE8:      amountE8,
		Status:            dbtrade.ORDERS_STATUS_UNFILLED,
		UpdateTime:        utime.NowMs(),
		AccumulatedFundE8: 0,
	}
}
