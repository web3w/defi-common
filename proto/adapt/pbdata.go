package adapt

import (
	"fmt"
	"github.com/gisvr/defi-common/dbdata/trade"
	be "github.com/gisvr/defi-common/proto/dex2"
	pb "github.com/gisvr/defi-common/proto/dex2fe"
	"github.com/gisvr/defi-common/utils/numeric"
	"strconv"
)

func FromBeTokenBalance(beBalance *be.TokenBalance, withdrawing float64) *pb.TokenBalance {
	return &pb.TokenBalance{
		TokenId:     beBalance.TokenId,
		Total:       numeric.Uint64E8ToString(beBalance.TotalE8),
		Active:      numeric.Uint64E8ToString(beBalance.TotalE8 - beBalance.LockedE8),
		Locked:      numeric.Uint64E8ToString(beBalance.LockedE8),
		Withdrawing: strconv.FormatFloat(withdrawing, 'f', 8, 64),
	}
}

func FromBeOrder(beOrder *be.Order) *pb.Order {
	var filledAveragePrice float64
	amountFilledE8 := beOrder.AmountTotalE8 - beOrder.AmountLeftE8
	if amountFilledE8 > 0 {
		filledAveragePrice = float64(beOrder.AccumulatedFundE8) / float64(amountFilledE8)
	}

	return &pb.Order{
		OrderId:            beOrder.OrderId,
		BookId:             beOrder.BookId,
		PairId:             dbtrade.GetTradePairSymbol(beOrder.CashTokenId, beOrder.StockTokenId),
		Action:             beOrder.Action,
		Price:              numeric.Uint64E8ToString(beOrder.PriceE8),
		AmountTotal:        numeric.Uint64E8ToString(beOrder.AmountTotalE8),
		AmountFilled:       numeric.Uint64E8ToString(amountFilledE8),
		FilledAveragePrice: fmt.Sprintf("%.8f", filledAveragePrice),
		Status:             beOrder.Status,
		CreateTimeMs:       beOrder.CreateTimeMs,
		UpdateTimeMs:       beOrder.UpdateTimeMs,
		ExpireTimeSec:      beOrder.ExpireTimeSec,
		Nonce:              beOrder.Nonce,
	}
}

func FromBeDepth(beDepth *be.Depth) *pb.Depth {
	depth := new(pb.Depth)
	depth.PairId = dbtrade.GetTradePairSymbol(beDepth.CashTokenId, beDepth.StockTokenId)
	depth.TimeMs = beDepth.TimeMs

	var askSumValue float64
	for _, depthEntry := range beDepth.Asks {
		askSumValue += depthEntry.TotalValue
		depth.Asks = append(depth.Asks, FromBeDepthEntry(depthEntry, askSumValue))
	}

	var bidSumValue float64
	for _, depthEntry := range beDepth.Bids {
		bidSumValue += depthEntry.TotalValue
		depth.Bids = append(depth.Bids, FromBeDepthEntry(depthEntry, bidSumValue))
	}
	return depth
}

func FromBeDepthEntry(beDepthEntry *be.DepthEntry, sumValue float64) *pb.DepthEntry {
	return &pb.DepthEntry{
		Price:      numeric.Uint64E8ToString(beDepthEntry.PriceE8),
		Amount:     numeric.Uint64E8ToString(beDepthEntry.AmountE8),
		TotalValue: fmt.Sprintf("%.8f", beDepthEntry.TotalValue),
		SumValue:   fmt.Sprintf("%.8f", sumValue),
	}
}
