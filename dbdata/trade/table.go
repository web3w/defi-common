package dbtrade

//-------------------------- 订单等信息相关表 -----------------------------------

/**
 * 各链订单号记录： order_ids_sequence
 */
type OrderIdsSequence struct {
	ChainToken string
	LastId     int64
}

/**
 * 用户操作记录： xxx_operate_data
 */
type OperateData struct {
	Id             int64
	Type           int8
	Trader         string
	OperateOrderId int64
	Data           string
	Status         int8
	CreateTime     int64
	UpdateTime     int64
}

/**
 * 用于各链订单记录： xxx_orders
 */
type Orders struct {
	OrderId           int64
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
	AmountLeftE8      uint64
	AccumulatedFundE8 uint64
	OrderSign         string
	CreateTime        int64
	UpdateTime        int64
	ExpireTime        int64
	Status            int8
}

/**
 * 币种对撮合序列： general_match_sequence
 */
type GeneralMatchSequence struct {
	Id               int64
	CashTokenSymbol  string
	StockTokenSymbol string
	ChainToken       string
	Sequence         int64
}

/**
 * 用于各链撮合记录： xxx_order_match
 */
type OrderMatch struct {
	Id               int64
	OrderId          int64
	OrderTrader      string
	PriceE8          uint64
	AmountE8         uint64
	PayAmountE8      uint64
	FeeRateE4        int32
	Action           int8
	Type             int8
	MatchOrderId     int64
	MatchOrderTrader string
	MatchPriceE8     uint64
	MatchAmountE8    uint64
	MatchPayAmountE8 uint64
	MatchFeeRateE4   int32
	CashTokenSymbol  string
	StockTokenSymbol string
	Status           int8
	Sequence         int64
	CreateTime       int64
}

/**
 * 用于各链成交明细： xxx_order_pending
 */
type OrderPending struct {
	Id               int64
	Trader           string
	OrderId          int64
	MatchOrderId     int64
	Direction        int8
	Side             int8
	Type             int8
	StockTokenSymbol string
	CashTokenSymbol  string
	PriceE8          uint64
	AmountE8         uint64
	TotalE8          uint64
	FeeRateE4        uint64
	FeeTokenSymbol   string
	FeeE8            uint64
	CreateTime       int64
}

//-------------------------- 系统配置相关表 -----------------------------------
/**
 * 交易区配置表： zone_config
 */
type ZoneConfig struct {
	Id              int64
	CashTokenId     int64
	CashTokenSymbol string
	ChainToken      string
	MarketAddr      string
	MakerFeeRateE4  uint64
	TakerFeeRateE4  uint64
	WithdrawFeeE4   uint64
	UpdateTime      uint64
}

/**
 * 币种信息表： coin_type
 */
type CoinType struct {
	Id                  int64
	TokenSymbol         string
	TokenName           string
	ChainToken          string
	ChainDecimals       int32
	Website             string
	TokenAddr           string
	TokenUrl            string
	TokenLogo           string
	InitPrice           string
	TotalAmount         string
	Circulation         string
	DescEn              string
	DescZh              string
	DepositMinAmountE8  int64
	WithdrawMinAmountE8 int64
	Weight              int32
	TradingStartTimeSec int64
	DelistTimeSec       int64
	EnableTimeSec       int64
	CreateTime          int64
	UpdateTime          int64
}

/**
 * 币种对数据： coin_pair表
 */
type CoinPair struct {
	Id               int64
	CashTokenId      int64
	CashTokenSymbol  string
	StockTokenId     int64
	StockTokenSymbol string
	ChainToken       string
	ValidDecimals    int32
	OrderMinAmountE8 int64
	OrderMinValueE8  int64
	Status           int32
	CreateTime       int64
	UpdateTime       int64
}

//-------------------------- 充值提现相关表 -----------------------------------

/**
 * 充值记录： xxx_deposits表。
 */
type Deposit struct {
	Id           int64
	DepositIndex int64
	BlockNumber  int64
	TxHash       string
	Trader       string
	TokenName    string
	AmountE8     uint64
	BalanceE8    uint64
	Status       int8
	CreateTime   int64
}

/**
 * 提现记录： xxx_withdraws表。
 */
type Withdraw struct {
	Id            int64
	WithdrawIndex int64
	BlockNumber   int64
	TxHash        string
	Trader        string
	TokenName     string
	AmountE8      uint64
	FeeE8         uint64
	BalanceE8     uint64
	Status        int8
	CreateTime    int64
	UpdateTime    int64
}
