package orderbook

import (
	"github.com/gisvr/deif-common/dbdata/trade"
	pb "github.com/gisvr/deif-common/proto/dex2"
	"github.com/gisvr/deif-common/utils/ulog"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

// An implementation of `OrderBook` using black-red-tree based map.
type RbtOrderBook struct {
	tradePairId string

	logicTimeSec int64

	asks        *rbt.Tree
	bids        *rbt.Tree
	pricesByIds map[int64]uint64 // used to get the order key from the order id
	expiryQue   *ExpiryQueue
}

func NewRbtOrderBook(tradePair string) *RbtOrderBook {
	return &RbtOrderBook{
		tradePairId: tradePair,

		logicTimeSec: 0,

		asks:        rbt.NewWith(AskComparator),
		bids:        rbt.NewWith(BidComparator),
		pricesByIds: make(map[int64]uint64),
		expiryQue:   NewExpiryQueue(),
	}
}

func (book *RbtOrderBook) PushLogicTimeSec(logicTimeSec int64) []*dbtrade.Orders {
	if logicTimeSec < book.logicTimeSec {
		ulog.Warningf("Ignored backward time pushing from %v to %v.", book.logicTimeSec, logicTimeSec)
		return nil
	}

	book.logicTimeSec = logicTimeSec

	var expiredOrders []*dbtrade.Orders
	for book.expiryQue.Len() != 0 {
		item := book.expiryQue.Top()
		if item.expireTimeSec > logicTimeSec {
			break
		}

		order := book.GetOrderById(item.orderId)
		if order == nil {
			ulog.Panicf("unexpected - expired order not found. %+v", item)
		}
		book.Delete(item.orderId)
		expiredOrders = append(expiredOrders, order)
	}
	return expiredOrders
}

func (book *RbtOrderBook) Add(order *dbtrade.Orders) {
	if book.tradePairId != dbtrade.GetTradePairOfOrder(order) {
		ulog.Panicf("OrderBook %v adding order %+v", book.tradePairId, order)
	}
	if !book.isOpenOrder(order) {
		ulog.Panic("Adding a non-open order: ", order)
	}

	savedSize := book.Size()
	switch order.Action {
	case dbtrade.ORDERS_ACTION_BUY: // bid
		book.bids.Put(GetKey(order), order)
	case dbtrade.ORDERS_ACTION_SELL: // ask
		book.asks.Put(GetKey(order), order)
	default:
		ulog.Panic("invalid order action")
	}

	book.pricesByIds[order.OrderId] = order.PriceE8
	book.expiryQue.Add(order.OrderId, order.ExpireTime)
	book.assertSize(savedSize + 1)
}

func (book *RbtOrderBook) Delete(orderId int64) bool {
	orderKey, ok := book.getOrderKeyFromId(orderId)
	if !ok {
		return false
	}

	savedSize := book.Size()
	book.asks.Remove(orderKey)
	if book.Size()+1 != savedSize {
		book.bids.Remove(orderKey)
		if book.Size()+1 != savedSize {
			ulog.Panic("Size sanity check failed when deleting an order")
		}
	}

	delete(book.pricesByIds, orderId)
	book.expiryQue.Remove(orderId)
	book.assertSize(savedSize - 1)
	return true
}

func (book *RbtOrderBook) Update(order *dbtrade.Orders) {
	if book.tradePairId != dbtrade.GetTradePairOfOrder(order) {
		ulog.Panicf("OrderBook %v updating order %+v", book.tradePairId, order)
	}
	if !book.isOpenOrder(order) {
		ulog.Panic("Updating a non-open order: ", order)
	}

	savedSize := book.Size()
	switch order.Action {
	case dbtrade.ORDERS_ACTION_BUY: // bid
		book.bids.Put(GetKey(order), order)
	case dbtrade.ORDERS_ACTION_SELL: // ask
		book.asks.Put(GetKey(order), order)
	default:
		ulog.Panic("invalid order action")
	}
	if book.Size() != savedSize {
		ulog.Panic("size sanity check failed when updating an order")
	}
}

//------------------------------ Read-only Section: ------------------------------------

func (book *RbtOrderBook) Size() int {
	return book.asks.Size() + book.bids.Size()
}

func (book *RbtOrderBook) GetTopBuyOrder() (*dbtrade.Orders, bool) {
	if book.bids.Empty() {
		return nil, false
	}
	return book.bids.Left().Value.(*dbtrade.Orders), true
}

func (book *RbtOrderBook) GetTopSellOrder() (*dbtrade.Orders, bool) {
	if book.asks.Empty() {
		return nil, false
	}
	return book.asks.Left().Value.(*dbtrade.Orders), true
}

func (book *RbtOrderBook) isOpenOrder(order *dbtrade.Orders) bool {
	return order.AmountLeftE8 > 0 && order.ExpireTime > book.logicTimeSec &&
		(order.Status == dbtrade.ORDERS_STATUS_UNFILLED || order.Status == dbtrade.ORDERS_STATUS_PARTIALLY_FILLED)
}

func (book *RbtOrderBook) assertSize(expected int) {
	if book.Size() != expected || len(book.pricesByIds) != expected ||
		book.expiryQue.Len() != expected {
		ulog.Panicf("Size sanity check failed. %v vs. %v, %v, %v: ",
			expected, book.Size(), len(book.pricesByIds), book.expiryQue.Len())
	}
}

func (book *RbtOrderBook) getOrderKeyFromId(orderId int64) (orderKey OrderKey, ok bool) {
	priceE8, ok := book.pricesByIds[orderId]
	if !ok {
		return OrderKey{}, false
	}
	return OrderKey{PriceE8: priceE8, OrderId: orderId}, true
}

func (book *RbtOrderBook) GetOrdersOfTrader(traderAddr string) []*dbtrade.Orders {
	result := book.getOrdersOfTrader(book.asks, traderAddr)
	result = append(result, book.getOrdersOfTrader(book.bids, traderAddr)...)
	return result
}

func (book *RbtOrderBook) getOrdersOfTrader(orders *rbt.Tree, traderAddr string) []*dbtrade.Orders {
	var result []*dbtrade.Orders
	for itr := orders.Iterator(); itr.Next(); {
		order := itr.Value().(*dbtrade.Orders)
		if order.Trader == traderAddr {
			result = append(result, order)
		}
	}
	return result
}

func (book *RbtOrderBook) GetOrderById(orderId int64) (result *dbtrade.Orders) {
	orderKey, ok := book.getOrderKeyFromId(orderId)
	if !ok {
		return nil
	}

	if order, ok := book.asks.Get(orderKey); ok {
		result = order.(*dbtrade.Orders)
	} else if order, ok := book.bids.Get(orderKey); ok {
		result = order.(*dbtrade.Orders)
	}
	return result
}

func (book *RbtOrderBook) ExportAllOrdersTo(slice []*dbtrade.Orders) []*dbtrade.Orders {
	slice = book.exportAllOrdersToSlice(book.bids, slice)
	slice = book.exportAllOrdersToSlice(book.asks, slice)
	return slice
}

func (book *RbtOrderBook) exportAllOrdersToSlice(orders *rbt.Tree, slice []*dbtrade.Orders) []*dbtrade.Orders {
	for itr := orders.Iterator(); itr.Next(); {
		order := itr.Value().(*dbtrade.Orders)
		slice = append(slice, order)
	}
	return slice
}

func (book *RbtOrderBook) GetDepth(size int, mergingDecimals int) (
	askEntries []*pb.DepthEntry, bidEntries []*pb.DepthEntry) {

	askEntries = book.getDepthEntries(book.asks, size, mergingDecimals, ASK)
	bidEntries = book.getDepthEntries(book.bids, size, mergingDecimals, BID)
	return
}

func (book *RbtOrderBook) getDepthEntries(
	orders *rbt.Tree, size int, mergingDecimals int, depthType string) []*pb.DepthEntry {

	if size == 0 {
		return nil
	}

	var mergingUint uint64
	switch mergingDecimals {
	case 2:
		mergingUint = 1e6
	case 4:
		mergingUint = 1e4
	case 5:
		mergingUint = 1e3
	case 6:
		mergingUint = 1e2
	case 7:
		mergingUint = 10
	default:
		mergingUint = 1
	}

	var entries []*pb.DepthEntry
	var curEntry *pb.DepthEntry
	for itr := orders.Iterator(); itr.Next(); {
		order := itr.Value().(*dbtrade.Orders)

		if curEntry != nil {
			var isCurEntry bool
			if depthType == ASK {
				isCurEntry = order.PriceE8 <= curEntry.PriceE8
			} else { // BID
				isCurEntry = order.PriceE8 >= curEntry.PriceE8
			}

			if isCurEntry {
				curEntry.AmountE8 += order.AmountLeftE8
				if curEntry.AmountE8 < order.AmountLeftE8 || curEntry.AmountE8 > MAX_DEPTH_ENTRY_AMOUNT_E8 {
					// Overflowed or too large. Reset it to MAX_DEPTH_ENTRY_AMOUNT
					curEntry.AmountE8 = MAX_DEPTH_ENTRY_AMOUNT_E8
				}
				curEntry.TotalValue += float64(order.AmountLeftE8) / 1e8 * float64(order.PriceE8) / 1e8
				continue
			}
		}

		if curEntry != nil {
			entries = append(entries, curEntry)
			curEntry = nil
			if len(entries) >= size {
				break
			}
		}

		var curEntryPriceE8 uint64
		if depthType == ASK {
			curEntryPriceE8 = ((order.PriceE8 + mergingUint - 1) / mergingUint) * mergingUint
		} else { // BID
			curEntryPriceE8 = (order.PriceE8 / mergingUint) * mergingUint
		}
		curEntry = &pb.DepthEntry{
			PriceE8:    curEntryPriceE8,
			AmountE8:   order.AmountLeftE8,
			TotalValue: float64(order.AmountLeftE8) / 1e8 * float64(order.PriceE8) / 1e8,
		}
	}

	if curEntry != nil {
		entries = append(entries, curEntry)
	}
	return entries
}
