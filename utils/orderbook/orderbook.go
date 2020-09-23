package orderbook

import (
	"git.bibox.com/dextop/common.git/dbdata/trade"
	pb "git.bibox.com/dextop/common.git/proto/dex2"
)

const MAX_DEPTH_ENTRY_AMOUNT_E8 = 1e9*1e8 - 1 // original amount 999,999,999.99999999
const ASK = "Ask"
const BID = "Bid"

// An OrderBook stores all the open orders of a particular trading pair.
//
// `Orderbook` is NOT responsible for updating the fieds (AmountLeft, Status, etc.) of the orders.
// An order should be completely updated before placing back to the orderbook via `Update` method.
//
// The `Add`, `Update` method can be called only with "Open" order, i.e. amount left > 0 and has not
// been cancelled or expired. In terms of status, the order must be Unfilled or PartiallyFilled.
type OrderBook interface {
	// Push the logic timestamp of the orderbook to `timeSec`. Orders expire on or before `timeSec`
	// are deleted and returned.
	//
	// Note: the returned orders are NOT updated (field Status, UpdateTime, etc.).
	PushLogicTimeSec(timeSec int64) []*dbtrade.Orders

	// Adds a new order.
	//
	// Precondition:
	// - The order must be open (see above for definition of "open").
	Add(order *dbtrade.Orders)

	// Deletes an order. Returns false if the order is not found.
	Delete(orderId int64) bool

	// Updates an order.
	//
	// Precondition:
	// - The order must exist and `order` must be open (see above for definition of "open").
	Update(order *dbtrade.Orders)

	//------------------------------ Read-only Section: ------------------------------------

	// How many orders in the orderbook.
	Size() int

	// Gets the buy order to be matched first. Returns (nil, false) if there is no buy order.
	GetTopBuyOrder() (buy *dbtrade.Orders, ok bool)

	// Gets the sell order to be matched first. Returns (nil, false) if there is no sell order.
	GetTopSellOrder() (sell *dbtrade.Orders, ok bool)

	// Gets all orders of a trader.
	GetOrdersOfTrader(traderAddr string) []*dbtrade.Orders

	// Gets order with the given id. Returns nil if not found.
	GetOrderById(orderId int64) *dbtrade.Orders

	// Gets depth. `size` is the number of levels.
	GetDepth(size int, mergingDecimals int) (askEntries []*pb.DepthEntry, bidEntries []*pb.DepthEntry)

	// Exports (by appending) all orders in the orderbook to `slice` (in unspecified order). Returns
	// the result (appended) slice.
	ExportAllOrdersTo(slice []*dbtrade.Orders) []*dbtrade.Orders
}
