package pbu

import (
	"fmt"

	pb "github.com/gisvr/defi-common/proto/dex2"
)

func NewTradeEvent(
	eventIndex, commitIndex int64, logicTimeMs int64, body interface{}) *pb.TradeEvent {

	event := &pb.TradeEvent{
		EventIndex:  eventIndex,
		CommitIndex: commitIndex,
		LogicTimeMs: logicTimeMs,
	}

	switch x := body.(type) {
	case *pb.SetFeeRatesEvent:
		event.Body = &pb.TradeEvent_SetFeeRates{SetFeeRates: x}
	case *pb.AddDepositEvent:
		event.Body = &pb.TradeEvent_AddDeposit{AddDeposit: x}
	case *pb.InitiateWithdrawEvent:
		event.Body = &pb.TradeEvent_InitiateWithdraw{InitiateWithdraw: x}
	case *pb.PlaceOrderEvent:
		event.Body = &pb.TradeEvent_PlaceOrder{PlaceOrder: x}
	case *pb.CancelOrderEvent:
		event.Body = &pb.TradeEvent_CancelOrder{CancelOrder: x}
	case *pb.MatchOrdersEvent:
		event.Body = &pb.TradeEvent_MatchOrders{MatchOrders: x}
	case *pb.ExpiredOrdersEvent:
		event.Body = &pb.TradeEvent_ExpiredOrders{ExpiredOrders: x}
	default:
		panic(fmt.Errorf("unknown event body type %T", x))
	}

	return event
}
