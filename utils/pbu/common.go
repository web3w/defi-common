package pbu

import (
	"github.com/gisvr/defi-common/utils/ulog"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

const (
	// The maximum nonce that a user request is allowed to use.
	MAX_NONCE = 5e18
)

func GetTradePairCode(cashTokenCode, stockTokenCode uint16) uint32 {
	return uint32(cashTokenCode)<<16 | uint32(stockTokenCode)
}

func PairCodeToCashStockCode(pairCode uint32) (cashTokenCode, stockTokenCode uint16) {
	return uint16(pairCode >> 16), uint16(pairCode)
}

func MustGetBeautifiedJsonPb(msg proto.Message) string {
	marshaler := jsonpb.Marshaler{EmitDefaults: true, Indent: "  "}
	str, err := marshaler.MarshalToString(msg)
	if err != nil {
		ulog.Panic("Error in MustGetBeautifiedJsonPb: ", err)
	}
	return str
}
