// Code generated by protoc-gen-go. DO NOT EDIT.
// source: general.proto

package dex2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 行为： 买入、卖出、取消。
type GeneralOrder_Action int32

const (
	GeneralOrder_Buy    GeneralOrder_Action = 0
	GeneralOrder_Sell   GeneralOrder_Action = 1
	GeneralOrder_Cancel GeneralOrder_Action = 2
)

var GeneralOrder_Action_name = map[int32]string{
	0: "Buy",
	1: "Sell",
	2: "Cancel",
}

var GeneralOrder_Action_value = map[string]int32{
	"Buy":    0,
	"Sell":   1,
	"Cancel": 2,
}

func (x GeneralOrder_Action) String() string {
	return proto.EnumName(GeneralOrder_Action_name, int32(x))
}

func (GeneralOrder_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{2, 0}
}

// 订单类型： 限价单、市价单。
type GeneralOrder_OrderType int32

const (
	GeneralOrder_Fixed  GeneralOrder_OrderType = 0
	GeneralOrder_Market GeneralOrder_OrderType = 1
)

var GeneralOrder_OrderType_name = map[int32]string{
	0: "Fixed",
	1: "Market",
}

var GeneralOrder_OrderType_value = map[string]int32{
	"Fixed":  0,
	"Market": 1,
}

func (x GeneralOrder_OrderType) String() string {
	return proto.EnumName(GeneralOrder_OrderType_name, int32(x))
}

func (GeneralOrder_OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{2, 1}
}

// Transition:
//           Pending  =>        Unfilled | Rejected
//          Unfilled  => PartiallyFilled | Cancelled | Expired
//   PartiallyFilled  =>          Filled | PartiallyCancelled | PartiallyExpired
type GeneralOrder_Status int32

const (
	GeneralOrder_NewOrder        GeneralOrder_Status = 0
	GeneralOrder_Unfilled        GeneralOrder_Status = 2
	GeneralOrder_PartiallyFilled GeneralOrder_Status = 3
	//--- inactive status: ---
	GeneralOrder_Rejected           GeneralOrder_Status = 4
	GeneralOrder_Filled             GeneralOrder_Status = 5
	GeneralOrder_Cancelled          GeneralOrder_Status = 6
	GeneralOrder_PartiallyCancelled GeneralOrder_Status = 7
	GeneralOrder_Expired            GeneralOrder_Status = 8
	GeneralOrder_PartiallyExpired   GeneralOrder_Status = 9
)

var GeneralOrder_Status_name = map[int32]string{
	0: "NewOrder",
	2: "Unfilled",
	3: "PartiallyFilled",
	4: "Rejected",
	5: "Filled",
	6: "Cancelled",
	7: "PartiallyCancelled",
	8: "Expired",
	9: "PartiallyExpired",
}

var GeneralOrder_Status_value = map[string]int32{
	"NewOrder":           0,
	"Unfilled":           2,
	"PartiallyFilled":    3,
	"Rejected":           4,
	"Filled":             5,
	"Cancelled":          6,
	"PartiallyCancelled": 7,
	"Expired":            8,
	"PartiallyExpired":   9,
}

func (x GeneralOrder_Status) String() string {
	return proto.EnumName(GeneralOrder_Status_name, int32(x))
}

func (GeneralOrder_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{2, 2}
}

// 通用市场配置。
type GeneralMarketConfig struct {
	MakerFeeRateE4 uint64 `protobuf:"varint,1,opt,name=maker_fee_rate_e4,json=makerFeeRateE4,proto3" json:"maker_fee_rate_e4,omitempty"`
	TakerFeeRateE4 uint64 `protobuf:"varint,2,opt,name=taker_fee_rate_e4,json=takerFeeRateE4,proto3" json:"taker_fee_rate_e4,omitempty"`
	WithdrawFeeE4  uint64 `protobuf:"varint,3,opt,name=withdraw_fee_e4,json=withdrawFeeE4,proto3" json:"withdraw_fee_e4,omitempty"`
	// The tokens that are can be used as cash in this market.
	CashTokens *CoinTypeSimple `protobuf:"bytes,5,opt,name=cash_tokens,json=cashTokens,proto3" json:"cash_tokens,omitempty"`
	// The tokens that can be traded as stock in this market.
	StockTokens []*CoinTypeSimple `protobuf:"bytes,6,rep,name=stock_tokens,json=stockTokens,proto3" json:"stock_tokens,omitempty"`
	// The tokens that can neither be used as cash, nor be used as stock.
	DisabledTokens       []*CoinTypeSimple `protobuf:"bytes,7,rep,name=disabled_tokens,json=disabledTokens,proto3" json:"disabled_tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GeneralMarketConfig) Reset()         { *m = GeneralMarketConfig{} }
func (m *GeneralMarketConfig) String() string { return proto.CompactTextString(m) }
func (*GeneralMarketConfig) ProtoMessage()    {}
func (*GeneralMarketConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{0}
}

func (m *GeneralMarketConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralMarketConfig.Unmarshal(m, b)
}
func (m *GeneralMarketConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralMarketConfig.Marshal(b, m, deterministic)
}
func (m *GeneralMarketConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralMarketConfig.Merge(m, src)
}
func (m *GeneralMarketConfig) XXX_Size() int {
	return xxx_messageInfo_GeneralMarketConfig.Size(m)
}
func (m *GeneralMarketConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralMarketConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralMarketConfig proto.InternalMessageInfo

func (m *GeneralMarketConfig) GetMakerFeeRateE4() uint64 {
	if m != nil {
		return m.MakerFeeRateE4
	}
	return 0
}

func (m *GeneralMarketConfig) GetTakerFeeRateE4() uint64 {
	if m != nil {
		return m.TakerFeeRateE4
	}
	return 0
}

func (m *GeneralMarketConfig) GetWithdrawFeeE4() uint64 {
	if m != nil {
		return m.WithdrawFeeE4
	}
	return 0
}

func (m *GeneralMarketConfig) GetCashTokens() *CoinTypeSimple {
	if m != nil {
		return m.CashTokens
	}
	return nil
}

func (m *GeneralMarketConfig) GetStockTokens() []*CoinTypeSimple {
	if m != nil {
		return m.StockTokens
	}
	return nil
}

func (m *GeneralMarketConfig) GetDisabledTokens() []*CoinTypeSimple {
	if m != nil {
		return m.DisabledTokens
	}
	return nil
}

// 币种简要信息
type CoinTypeSimple struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TokenSymbol          string   `protobuf:"bytes,2,opt,name=token_symbol,json=tokenSymbol,proto3" json:"token_symbol,omitempty"`
	TokenName            string   `protobuf:"bytes,3,opt,name=token_name,json=tokenName,proto3" json:"token_name,omitempty"`
	ChainToken           string   `protobuf:"bytes,4,opt,name=chain_token,json=chainToken,proto3" json:"chain_token,omitempty"`
	TokenAddr            string   `protobuf:"bytes,5,opt,name=token_addr,json=tokenAddr,proto3" json:"token_addr,omitempty"`
	ScaleFactor          uint64   `protobuf:"varint,6,opt,name=scale_factor,json=scaleFactor,proto3" json:"scale_factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoinTypeSimple) Reset()         { *m = CoinTypeSimple{} }
func (m *CoinTypeSimple) String() string { return proto.CompactTextString(m) }
func (*CoinTypeSimple) ProtoMessage()    {}
func (*CoinTypeSimple) Descriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{1}
}

func (m *CoinTypeSimple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoinTypeSimple.Unmarshal(m, b)
}
func (m *CoinTypeSimple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoinTypeSimple.Marshal(b, m, deterministic)
}
func (m *CoinTypeSimple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinTypeSimple.Merge(m, src)
}
func (m *CoinTypeSimple) XXX_Size() int {
	return xxx_messageInfo_CoinTypeSimple.Size(m)
}
func (m *CoinTypeSimple) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinTypeSimple.DiscardUnknown(m)
}

var xxx_messageInfo_CoinTypeSimple proto.InternalMessageInfo

func (m *CoinTypeSimple) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CoinTypeSimple) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *CoinTypeSimple) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *CoinTypeSimple) GetChainToken() string {
	if m != nil {
		return m.ChainToken
	}
	return ""
}

func (m *CoinTypeSimple) GetTokenAddr() string {
	if m != nil {
		return m.TokenAddr
	}
	return ""
}

func (m *CoinTypeSimple) GetScaleFactor() uint64 {
	if m != nil {
		return m.ScaleFactor
	}
	return 0
}

// 通用订单。
type GeneralOrder struct {
	// The id of the order, which is a positive value. An earlier order has a lower id.
	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// 下单人
	Trader string                 `protobuf:"bytes,2,opt,name=trader,proto3" json:"trader,omitempty"`
	BookId int64                  `protobuf:"varint,3,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Action GeneralOrder_Action    `protobuf:"varint,4,opt,name=action,proto3,enum=dex.GeneralOrder_Action" json:"action,omitempty"`
	Type   GeneralOrder_OrderType `protobuf:"varint,5,opt,name=type,proto3,enum=dex.GeneralOrder_OrderType" json:"type,omitempty"`
	// 交易币id和symbol。
	StockTokenId     uint64 `protobuf:"varint,6,opt,name=stock_token_id,json=stockTokenId,proto3" json:"stock_token_id,omitempty"`
	StockTokenSymbol string `protobuf:"bytes,7,opt,name=stock_token_symbol,json=stockTokenSymbol,proto3" json:"stock_token_symbol,omitempty"`
	// 计价币id和symbol。
	CashTokenId     uint64 `protobuf:"varint,8,opt,name=cash_token_id,json=cashTokenId,proto3" json:"cash_token_id,omitempty"`
	CashTokenSymbol string `protobuf:"bytes,9,opt,name=cash_token_symbol,json=cashTokenSymbol,proto3" json:"cash_token_symbol,omitempty"`
	ChainToken      string `protobuf:"bytes,10,opt,name=chain_token,json=chainToken,proto3" json:"chain_token,omitempty"`
	// 挂单价格。
	PriceE8 uint64 `protobuf:"varint,11,opt,name=price_e8,json=priceE8,proto3" json:"price_e8,omitempty"`
	// 挂单总数量。
	AmountTotalE8 uint64 `protobuf:"varint,12,opt,name=amount_total_e8,json=amountTotalE8,proto3" json:"amount_total_e8,omitempty"`
	// 挂单总资产。
	AccumulatedFundE8 uint64 `protobuf:"varint,13,opt,name=accumulated_fund_e8,json=accumulatedFundE8,proto3" json:"accumulated_fund_e8,omitempty"`
	// 订单签名。
	OrderSign string `protobuf:"bytes,14,opt,name=order_sign,json=orderSign,proto3" json:"order_sign,omitempty"`
	// Do NOT use create time to check which order is accepted earlier. Use order id instead.
	CreateTime int64               `protobuf:"varint,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ExpireTime int64               `protobuf:"varint,17,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	UpdateTime int64               `protobuf:"varint,16,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Status     GeneralOrder_Status `protobuf:"varint,18,opt,name=status,proto3,enum=dex.GeneralOrder_Status" json:"status,omitempty"`
	// 剩余未成交数量。
	AmountLeftE8 uint64 `protobuf:"varint,19,opt,name=amount_left_e8,json=amountLeftE8,proto3" json:"amount_left_e8,omitempty"`
	// 订单来源
	UserFrom             string   `protobuf:"bytes,20,opt,name=user_from,json=userFrom,proto3" json:"user_from,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralOrder) Reset()         { *m = GeneralOrder{} }
func (m *GeneralOrder) String() string { return proto.CompactTextString(m) }
func (*GeneralOrder) ProtoMessage()    {}
func (*GeneralOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{2}
}

func (m *GeneralOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralOrder.Unmarshal(m, b)
}
func (m *GeneralOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralOrder.Marshal(b, m, deterministic)
}
func (m *GeneralOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralOrder.Merge(m, src)
}
func (m *GeneralOrder) XXX_Size() int {
	return xxx_messageInfo_GeneralOrder.Size(m)
}
func (m *GeneralOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralOrder.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralOrder proto.InternalMessageInfo

func (m *GeneralOrder) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *GeneralOrder) GetTrader() string {
	if m != nil {
		return m.Trader
	}
	return ""
}

func (m *GeneralOrder) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *GeneralOrder) GetAction() GeneralOrder_Action {
	if m != nil {
		return m.Action
	}
	return GeneralOrder_Buy
}

func (m *GeneralOrder) GetType() GeneralOrder_OrderType {
	if m != nil {
		return m.Type
	}
	return GeneralOrder_Fixed
}

func (m *GeneralOrder) GetStockTokenId() uint64 {
	if m != nil {
		return m.StockTokenId
	}
	return 0
}

func (m *GeneralOrder) GetStockTokenSymbol() string {
	if m != nil {
		return m.StockTokenSymbol
	}
	return ""
}

func (m *GeneralOrder) GetCashTokenId() uint64 {
	if m != nil {
		return m.CashTokenId
	}
	return 0
}

func (m *GeneralOrder) GetCashTokenSymbol() string {
	if m != nil {
		return m.CashTokenSymbol
	}
	return ""
}

func (m *GeneralOrder) GetChainToken() string {
	if m != nil {
		return m.ChainToken
	}
	return ""
}

func (m *GeneralOrder) GetPriceE8() uint64 {
	if m != nil {
		return m.PriceE8
	}
	return 0
}

func (m *GeneralOrder) GetAmountTotalE8() uint64 {
	if m != nil {
		return m.AmountTotalE8
	}
	return 0
}

func (m *GeneralOrder) GetAccumulatedFundE8() uint64 {
	if m != nil {
		return m.AccumulatedFundE8
	}
	return 0
}

func (m *GeneralOrder) GetOrderSign() string {
	if m != nil {
		return m.OrderSign
	}
	return ""
}

func (m *GeneralOrder) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *GeneralOrder) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *GeneralOrder) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *GeneralOrder) GetStatus() GeneralOrder_Status {
	if m != nil {
		return m.Status
	}
	return GeneralOrder_NewOrder
}

func (m *GeneralOrder) GetAmountLeftE8() uint64 {
	if m != nil {
		return m.AmountLeftE8
	}
	return 0
}

func (m *GeneralOrder) GetUserFrom() string {
	if m != nil {
		return m.UserFrom
	}
	return ""
}

func init() {
	proto.RegisterEnum("dex.GeneralOrder_Action", GeneralOrder_Action_name, GeneralOrder_Action_value)
	proto.RegisterEnum("dex.GeneralOrder_OrderType", GeneralOrder_OrderType_name, GeneralOrder_OrderType_value)
	proto.RegisterEnum("dex.GeneralOrder_Status", GeneralOrder_Status_name, GeneralOrder_Status_value)
	proto.RegisterType((*GeneralMarketConfig)(nil), "dex.GeneralMarketConfig")
	proto.RegisterType((*CoinTypeSimple)(nil), "dex.CoinTypeSimple")
	proto.RegisterType((*GeneralOrder)(nil), "dex.GeneralOrder")
}

func init() { proto.RegisterFile("general.proto", fileDescriptor_a913b1a5d8940539) }

var fileDescriptor_a913b1a5d8940539 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0xc7, 0xad, 0x83, 0x29, 0x71, 0x74, 0xa2, 0x57, 0x41, 0x3e, 0x06, 0xc1, 0x87, 0xba, 0x42,
	0xd1, 0xa6, 0x41, 0x20, 0x07, 0xae, 0x51, 0xe8, 0xa2, 0x37, 0x89, 0x21, 0x15, 0x06, 0xda, 0xb4,
	0xa0, 0xdc, 0x9b, 0xde, 0x10, 0x2b, 0xee, 0x48, 0xde, 0x8a, 0xe4, 0x0a, 0xcb, 0x25, 0x2c, 0x3d,
	0x4c, 0x9f, 0xa5, 0x57, 0x7d, 0x88, 0xbe, 0x4d, 0xb1, 0xb3, 0xd4, 0x21, 0x35, 0x7c, 0x23, 0x70,
	0xff, 0xf3, 0x9b, 0xe1, 0x9c, 0x96, 0x82, 0xde, 0x0a, 0x73, 0xd4, 0x3c, 0x1d, 0x6f, 0xb4, 0x32,
	0x8a, 0x35, 0x04, 0x6e, 0x47, 0x7f, 0xd5, 0x61, 0xf8, 0xa3, 0x93, 0x7f, 0xe6, 0x7a, 0x8d, 0xe6,
	0x56, 0xe5, 0x4b, 0xb9, 0x62, 0xdf, 0xc2, 0x45, 0xc6, 0xd7, 0xa8, 0xe3, 0x25, 0x62, 0xac, 0xb9,
	0xc1, 0x18, 0x6f, 0xc2, 0xda, 0x65, 0xed, 0x4d, 0x33, 0xea, 0x93, 0x61, 0x86, 0x18, 0x71, 0x83,
	0xd3, 0x1b, 0x8b, 0x9a, 0x27, 0x68, 0xdd, 0xa1, 0xe6, 0x73, 0xf4, 0x6b, 0x18, 0x3c, 0x4a, 0xf3,
	0x20, 0x34, 0x7f, 0x24, 0x1a, 0x6f, 0xc2, 0x06, 0x81, 0xbd, 0xbd, 0x3c, 0x43, 0xcb, 0xdd, 0x40,
	0x27, 0xe1, 0xc5, 0x43, 0x6c, 0xd4, 0x1a, 0xf3, 0x22, 0x3c, 0xbf, 0xac, 0xbd, 0xe9, 0x5c, 0x0f,
	0xc7, 0x02, 0xb7, 0xe3, 0x5b, 0x25, 0xf3, 0xfb, 0xdd, 0x06, 0xe7, 0x32, 0xdb, 0xa4, 0x18, 0x81,
	0xe5, 0xee, 0x09, 0x63, 0xdf, 0x43, 0xb7, 0x30, 0x2a, 0x59, 0xef, 0xdd, 0xbc, 0xcb, 0xc6, 0x73,
	0x6e, 0x1d, 0x02, 0x2b, 0xbf, 0x1f, 0x60, 0x20, 0x64, 0xc1, 0x17, 0x29, 0x8a, 0xbd, 0x6b, 0xeb,
	0x79, 0xd7, 0xfe, 0x9e, 0x75, 0xde, 0xa3, 0xbf, 0x6b, 0xd0, 0xff, 0x1c, 0x61, 0x7d, 0xa8, 0x4b,
	0x51, 0x75, 0xab, 0x2e, 0x05, 0xfb, 0x12, 0xba, 0x14, 0x37, 0x2e, 0x76, 0xd9, 0x42, 0xa5, 0xd4,
	0x1c, 0x3f, 0xea, 0x90, 0x36, 0x27, 0x89, 0xfd, 0x1f, 0xc0, 0x21, 0x39, 0xcf, 0x90, 0x9a, 0xe2,
	0x47, 0x3e, 0x29, 0x9f, 0x78, 0x86, 0xec, 0x0b, 0xe8, 0x24, 0x0f, 0x5c, 0xe6, 0x2e, 0xbf, 0xb0,
	0x49, 0x76, 0x20, 0x89, 0xd2, 0x38, 0xfa, 0x73, 0x21, 0x34, 0x35, 0x6c, 0xef, 0xff, 0x41, 0x08,
	0x6d, 0x33, 0x28, 0x12, 0x9e, 0x62, 0xbc, 0xe4, 0x89, 0x51, 0x3a, 0xf4, 0x28, 0xb7, 0x0e, 0x69,
	0x33, 0x92, 0x46, 0xff, 0xb4, 0xa0, 0x5b, 0x6d, 0xc2, 0x2f, 0x5a, 0xa0, 0x66, 0xaf, 0xa0, 0xad,
	0xec, 0x43, 0x5c, 0xd5, 0xd2, 0x88, 0x5a, 0x74, 0xbe, 0x13, 0xec, 0x25, 0x78, 0x46, 0x73, 0x81,
	0xba, 0x2a, 0xa5, 0x3a, 0xb1, 0xff, 0x41, 0x6b, 0xa1, 0xd4, 0xda, 0x7a, 0x34, 0xc8, 0xc3, 0xb3,
	0xc7, 0x3b, 0xc1, 0xde, 0x83, 0xc7, 0x13, 0x23, 0x95, 0x4b, 0xbd, 0x7f, 0x1d, 0x52, 0x67, 0x4f,
	0x5f, 0x37, 0xfe, 0x40, 0xf6, 0xa8, 0xe2, 0xd8, 0x15, 0x34, 0xcd, 0x6e, 0x83, 0x54, 0x4a, 0xff,
	0xfa, 0xf5, 0x53, 0x9e, 0x7e, 0x6d, 0xd3, 0x23, 0x02, 0xd9, 0x57, 0xd0, 0x3f, 0x99, 0xbe, 0x4d,
	0xc1, 0x15, 0xd9, 0x3d, 0x8e, 0xfa, 0x4e, 0xb0, 0x77, 0xc0, 0x4e, 0xa9, 0x6a, 0x20, 0x2d, 0xaa,
	0x22, 0x38, 0x92, 0xd5, 0x54, 0x46, 0xd0, 0x3b, 0xee, 0xa1, 0x0d, 0xd9, 0x76, 0x7d, 0x3b, 0x2c,
	0xdd, 0x9d, 0x60, 0x6f, 0xe1, 0xe2, 0x84, 0xa9, 0x02, 0xfa, 0x14, 0x70, 0x70, 0xe0, 0xaa, 0x78,
	0xff, 0x19, 0x23, 0x3c, 0x19, 0xe3, 0x2b, 0x68, 0x6f, 0xb4, 0x4c, 0x30, 0xc6, 0x49, 0xd8, 0xa1,
	0x77, 0xb5, 0xe8, 0x3c, 0x9d, 0xd8, 0xbb, 0xc3, 0x33, 0x55, 0xe6, 0x26, 0x36, 0xca, 0xf0, 0xd4,
	0x12, 0x5d, 0x77, 0x77, 0x9c, 0x7c, 0x6f, 0xd5, 0xe9, 0x84, 0x8d, 0x61, 0xc8, 0x93, 0xa4, 0xcc,
	0xca, 0x94, 0x1b, 0x14, 0xf1, 0xb2, 0xcc, 0x85, 0x65, 0x7b, 0xc4, 0x5e, 0x9c, 0x98, 0x66, 0x65,
	0x2e, 0xa6, 0x13, 0xbb, 0x39, 0x6e, 0xcc, 0x85, 0x5c, 0xe5, 0x61, 0xdf, 0x6d, 0x0e, 0x29, 0x73,
	0xb9, 0xca, 0x29, 0x65, 0x8d, 0xf6, 0x56, 0x1b, 0x99, 0x61, 0x38, 0xa0, 0xb1, 0x82, 0x93, 0xee,
	0xa5, 0x5b, 0x4d, 0xdc, 0x6e, 0xa4, 0xae, 0x80, 0x0b, 0x07, 0x38, 0x69, 0x0f, 0x94, 0x1b, 0x71,
	0x88, 0x10, 0x38, 0xc0, 0x49, 0x04, 0xbc, 0x07, 0xaf, 0x30, 0xdc, 0x94, 0x45, 0xc8, 0x9e, 0x5b,
	0x8e, 0x39, 0xd9, 0xa3, 0x8a, 0xb3, 0xb3, 0xae, 0x7a, 0x91, 0xe2, 0xd2, 0xd8, 0xf2, 0x86, 0x6e,
	0xd6, 0x4e, 0xfd, 0x09, 0x97, 0x66, 0x3a, 0x61, 0xaf, 0xc1, 0x2f, 0x0b, 0xfb, 0x5d, 0xd2, 0x2a,
	0x0b, 0x5f, 0x50, 0x61, 0x6d, 0x2b, 0xcc, 0xb4, 0xca, 0x46, 0xdf, 0x80, 0xe7, 0x36, 0x8e, 0xb5,
	0xa0, 0xf1, 0xb1, 0xdc, 0x05, 0x67, 0xac, 0x0d, 0xcd, 0x39, 0xa6, 0x69, 0x50, 0x63, 0x00, 0xde,
	0x2d, 0xcf, 0x13, 0x4c, 0x83, 0xfa, 0x68, 0x04, 0xfe, 0x61, 0xd5, 0x98, 0x0f, 0xe7, 0x33, 0xb9,
	0x45, 0x11, 0x9c, 0x59, 0xc6, 0x7d, 0x31, 0x83, 0xda, 0xe8, 0xcf, 0x1a, 0x78, 0x2e, 0x45, 0xd6,
	0x85, 0xf6, 0x27, 0x7c, 0x24, 0x8f, 0xe0, 0xcc, 0x9e, 0x7e, 0xcb, 0x97, 0x32, 0x4d, 0x51, 0x04,
	0x75, 0x36, 0x84, 0xc1, 0xaf, 0x5c, 0x1b, 0xc9, 0xd3, 0x74, 0x37, 0x73, 0x62, 0xc3, 0x22, 0x11,
	0xfe, 0x81, 0x89, 0x41, 0x11, 0x34, 0x6d, 0xd4, 0xca, 0x72, 0xce, 0x7a, 0xe0, 0xbb, 0x2c, 0xec,
	0xd1, 0x63, 0x2f, 0x81, 0x1d, 0xbc, 0x8f, 0x7a, 0x8b, 0x75, 0xa0, 0x35, 0xa5, 0x6e, 0x8b, 0xa0,
	0xcd, 0x5e, 0x40, 0x70, 0x80, 0xf6, 0xaa, 0xff, 0xf1, 0xdd, 0xef, 0x6f, 0x57, 0xd2, 0x8c, 0x17,
	0x72, 0xa1, 0xb6, 0xe3, 0x44, 0x65, 0x57, 0x02, 0xb7, 0x46, 0x6d, 0xae, 0x12, 0x95, 0x65, 0x2a,
	0x1f, 0xaf, 0xa4, 0xb9, 0xa2, 0xff, 0x04, 0xab, 0x5f, 0x2f, 0x3c, 0x7a, 0xfe, 0xee, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa4, 0xf9, 0x25, 0xfc, 0x30, 0x06, 0x00, 0x00,
}