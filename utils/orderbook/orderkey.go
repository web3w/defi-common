package orderbook

import (
	"github.com/gisvr/defi-common/dbdata/trade"
)

// The key of an order is <price, orderId>.
//
// This is a value type. Usually do not use pointer of it.
type OrderKey struct {
	PriceE8 uint64
	OrderId int64
}

func GetKey(order *dbtrade.Orders) OrderKey {
	return OrderKey{
		PriceE8: order.PriceE8,
		OrderId: order.OrderId,
	}
}

// For the ask prices, the lower the better.
func AskComparator(av, bv interface{}) int {
	a := av.(OrderKey)
	b := bv.(OrderKey)

	// Note that the prices are uint64. Thus cannot use subtraction.
	switch {
	case a.PriceE8 < b.PriceE8:
		return -1
	case a.PriceE8 > b.PriceE8:
		return 1
	default:
		switch {
		case a.OrderId < b.OrderId:
			return -1
		case a.OrderId > b.OrderId:
			return 1
		default:
			return 0
		}
	}
}

// For the bid prices, the lower the better.
func BidComparator(av, bv interface{}) int {
	a := av.(OrderKey)
	b := bv.(OrderKey)

	// Note that the prices are uint64. Thus cannot use subtraction.
	switch {
	case a.PriceE8 > b.PriceE8:
		return -1
	case a.PriceE8 < b.PriceE8:
		return 1
	default:
		switch {
		case a.OrderId < b.OrderId:
			return -1
		case a.OrderId > b.OrderId:
			return 1
		default:
			return 0
		}
	}
}
