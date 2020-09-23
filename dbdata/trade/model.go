package dbtrade

/**
 * OperateData中用户下单数据。
 */
type Order2Place struct {
	Trader            string
	UserFrom          string
	Action            int8
	Type              int8
	StockTokenId      int64
	StockTokenSymbol  string
	CashTokenId       int64
	CashTokenSymbol   string
	PriceE8           uint64
	AmountTotalE8     uint64
	AccumulatedFundE8 uint64
	OrderSign         string
	ExpireTime        int64
}

/**
 * OperateData中用户撤单数据。
 */
type Order2Cancel struct {
	Trader           string
	UserFrom         string
	StockTokenSymbol string
	CashTokenSymbol  string
	OrderId          int64
	OrderSign        string
}

/**
 * 币种简要信息。
 */
type CoinTypeSimple struct {
	Id          int64
	TokenSymbol string
	TokenName   string
	ChainToken  string
	TokenAddr   string
	ScaleFactor uint64
}

/**
 * 撮合数据传输用。
 */
type MatchDataTransfer struct {
	CashTokenSymbol  string
	StockTokenSymbol string

	// maker支付币种
	MakerPayTokenId int64
	MakerPayToken   string
	// maker获取币种
	MakerGetTokenId int64
	MakerGetToken   string
	// taker支付币种
	TakerPayTokenId int64
	TakerPayToken   string
	// taker获取币种
	TakerGetTokenId int64
	TakerGetToken   string

	// 撮合价格，成交数量，成交总金额。
	MatchPriceE8            uint64
	MatchStockTokenAmountE8 uint64
	MatchCashTokenAmount    uint64

	MakerFeeRateE4 uint64
	TakerFeeRateE4 uint64

	// maker支付数量
	MakerPayAmountE8 uint64
	// maker解冻数量
	MakerUnlockAmountE8 uint64
	// maker手续费数量
	MakerFeeE8 uint64
	// maker获取数量
	MakerGetAmount uint64

	// taker支付数量
	TakerPayAmountE8 uint64
	// taker解冻数量
	TakerUnlockAmountE8 uint64
	// taker手续费数量
	TakerFeeE8 uint64
	// taker获取数量
	TakerGetAmountE8 uint64

	// maker订单方向
	MakerDirection int8
	// taker订单方向
	TakerDirection int8

	SeqDbId       uint64
	MatchOrderSeq int64
}
