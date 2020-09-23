package pbu

import (
	pb "github.com/gisvr/defi-common/proto/dex2"
)

// Check whether an order is partially filled BASED ON AMOUNT LEFT in it.
func IsPartiallyFilled(order *pb.Order) bool {
	return order.AmountLeftE8 < order.AmountTotalE8
}

// Check whether an order is unfilled BASED ON AMOUNT LEFT in it.
func IsUnfilled(order *pb.Order) bool {
	return order.AmountLeftE8 == order.AmountTotalE8
}
