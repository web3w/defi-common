package dex2

import (
	"github.com/gisvr/defi-common/utils/contract"
	"github.com/gisvr/defi-common/utils/ulog"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	logger = ulog.NewLogger()

	Dex2Abi abi.ABI

	OrderHashTypes []byte
)

func init() {
	var err error
	Dex2Abi, err = abi.JSON(strings.NewReader(contract.Dex2ABI))
	if err != nil {
		panic(err)
	}

	OrderHashTypes = crypto.Keccak256(
		[]byte("string title"), []byte("address market_address"), []byte("uint64 nonce"),
		[]byte("uint64 expire_time_sec"), []byte("uint64 amount_e8"), []byte("uint64 price_e8"),
		[]byte("uint8 immediate_or_cancel"), []byte("uint8 action"),
		[]byte("uint16 cash_token_code"), []byte("uint16 stock_token_code"))
}
