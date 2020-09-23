package dbuser

/**
 * 用户资产表： user_asset
 */
type UserAsset struct {
	Id          int64
	Trader      string
	TokenId     int64
	TokenSymbol string
	ChainToken  string
	TotalE8     uint64
	FreezeE8    uint64
}
