package dbtrade

// 主链符号。
const (
	CHAIN_TOKEN_ETH  = "ETH"
	CHAIN_TOKEN_EOS  = "EOS"
	CHAIN_TOKEN_NEO  = "NEO"
	CHAIN_TOKEN_IOST = "IOST"
	CHAIN_TOKEN_NULS = "NULS"
)

const (
	PairInfoKeyPrefix          = "eventdaemon:pair-info:"     // e.g. "eventdaemon:pair-info:ETH_ADX"
	TradeRecordKeyPrefix       = "eventdaemon:trade-records:" // e.g. "eventdaemon:trade-records:ETH_ADX"
	TradeRecordsReservationNum = 50
	DepthKeyPrefix             = "eventdaemon:depth:" // e.g. "eventdaemon:depth:ETH_ADX"

	/**
	 * 剩余要购买的交易币数量低于该值，订单直接置为完成。
	 */
	ORDER_FILL_AMOUNT = 1000000
)

const (
	/**
	 * xxx_operate_data表记录类型： 下新订单。
	 */
	OPERATE_TYPE_PLACE = 0
	/**
	 * xxx_operate_data表记录类型： 取消订单。
	 */
	OPERATE_TYPE_CANCEL = 1
	/**
	 * xxx_operate_data表记录类型： 其他操作。
	 */
	OPERATE_TYPE_OTHER = 2
)

const (
	/**
	 * xxx_operate_data表状态： 新记录。
	 */
	OPERATE_STATUS_NEW = 0
	/**
	 * xxx_operate_data表状态： 已处理。
	 */
	OPERATE_STATUS_HANDLED = 1
	/**
	 * xxx_operate_data表状态： 资产不足失败。
	 */
	OPERATE_STATUS_ASSET = 2
	/**
	 * xxx_operate_data表状态： 处理失败。
	 */
	OPERATE_STATUS_FAILURE = -1
)

// --------------------------------------------------------------

const (
	/**
	 * xxx_orders表订单行为： 买单。
	 */
	ORDERS_ACTION_BUY = 0
	/**
	 * xxx_orders表订单行为： 卖单。
	 */
	ORDERS_ACTION_SELL = 1
)

const (
	/**
	 * xxx_orders表订单类型： 限价单。
	 */
	ORDERS_TYPE_FIXED = 0
	/**
	 * xxx_orders表订单类型： 市价单。
	 */
	ORDERS_TYPE_MARKET = 0
)

const (
	/**
	 * xxx_orders表订单状态： 新订单
	 */
	ORDERS_STATUS_NEW = 0
	/**
	 * xxx_orders表订单状态： 未成交
	 */
	ORDERS_STATUS_UNFILLED = 2
	/**
	 * xxx_orders表订单状态： 部分成交
	 */
	ORDERS_STATUS_PARTIALLY_FILLED = 3
	/**
	 * xxx_orders表订单状态： 已拒绝
	 */
	ORDERS_STATUS_REJECTED = 4
	/**
	 * xxx_orders表订单状态： 已成交
	 */
	ORDERS_STATUS_FILLED = 5
	/**
	 * xxx_orders表订单状态： 已取消
	 */
	ORDERS_STATUS_CANCELLED = 6
	/**
	 * xxx_orders表订单状态： 部分取消
	 */
	ORDERS_STATUS_PARTIALLY_CANCELLED = 7
	/**
	 * xxx_orders表订单状态： 已过期
	 */
	ORDERS_STATUS_EXPIRED = 8
	/**
	 * xxx_orders表订单状态： 部分过期
	 */
	ORDERS_STATUS_PARTIALLY_EXPIRED = 9
)

// --------------------------------------------------------------

const (
	/**
	 * xxx_order_match表记录类型： 撮合
	 */
	MATCH_TYPE_MATCH = 0
	/**
	 * xxx_order_match表记录类型： 撤单
	 */
	MATCH_TYPE_CANCEL = 1
)

const (
	/**
	 * xxx_order_match表记录状态： 未处理
	 */
	MATCH_STATUS_NEW = 0
	/**
	 * xxx_order_match表记录状态： 已上链
	 */
	MATCH_STATUS_HANDLED = 1
	/**
	 * xxx_order_match表记录状态： 已确认
	 */
	MATCH_STATUS_CONFIRM = 2
)

// --------------------------------------------------------------

const (
	/**
	 * xxx_order_pending表方向： 买
	 */
	PENDING_DIRECTION_BUY = 0
	/**
	 * xxx_order_pending表方向： 卖
	 */
	PENDING_DIRECTION_SELL = 1
)

const (
	/**
	 * xxx_order_pending表角色： maker
	 */
	PENDING_SIDE_MAKER = 0
	/**
	 * xxx_order_pending表角色： taker
	 */
	PENDING_SIDE_TAKER = 1
)
const (
	/**
	 * xxx_order_pending表类型： 限价单。
	 */
	PENDING_TYPE_FIXED = 0
	/**
	 * xxx_order_pending表类型： 市价单。
	 */
	PENDING_TYPE_MARKET = 0
)

// --------------------------------------------------------------

const (
	/**
	 * 转账类型： 充值。
	 */
	TRANSFER_DEPOSIT = 0
	/**
	 * 转账类型： 提现。
	 */
	TRANSFER_WITHDRAW = 1

	/**
	 * 转账状态字符串： 确认中。
	 */
	TRANSFER_STATUS_CONFIRMING = "Confirming"
	/**
	 * 转账状态字符串： 待处理。
	 */
	TRANSFER_STATUS_WAITING = "Waiting"
	/**
	 * 转账状态字符串： 已广播。
	 */
	TRANSFER_STATUS_SEND = "Send"
	/**
	 * 转账状态字符串： 已完成。
	 */
	TRANSFER_STATUS_DONE = "Done"
)

const (
	/**
	 * 充值类型： 未确认。
	 */
	DEPOSIT_STATUS_CONFIRMING = 0
	/**
	 * 充值类型： 已确认。
	 */
	DEPOSIT_STATUS_CONFIRM = 1
	/**
	 * 充值类型： 已到账。
	 */
	DEPOSIT_STATUS_DONE = 2
)

const (
	/**
	 * 提现类型： 待处理。
	 */
	WITHDRAW_STATUS_WAIT = 0
	/**
	 * 提现类型： 已广播。
	 */
	WITHDRAW_STATUS_SEND = 1
	/**
	 * 提现类型： 已确认。
	 */
	WITHDRAW_STATUS_DONE = 2
)

// --------------------------------------------------------------
