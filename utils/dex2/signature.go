package dex2

import (
	"crypto/ecdsa"
	"github.com/gisvr/defi-common/utils/ueth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// The returned bytes are to be hashed (using keccak256) for signing. The content are the
// concatenation of the following (uints are in big-endian byte order):
//
// 1. String "DEx2 Order: "(96)  (Note the trailing whitespace)
// 2. <market address>(160)      (For preventing cross-market replay attack)
// 3. <nonce>(64) <expireTimeSec>(64) <amountE8>(64) <priceE8>(64) <ioc>(8) <action>(8) <pairId>(32)
func GetOrderBytesToSign(marketAddr common.Address, nonce uint64, order *Order) []byte {
	bs := make([]byte, 0, 98)
	bs = append(bs, []byte("\x19Ethereum Signed Message:\n70")...)
	bs = append(bs, []byte("DEx2 Order: ")...)
	bs = append(bs, marketAddr.Bytes()...)
	bs = append(bs, Uint64ToBigEndianBytes(nonce)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.ExpireTimeSec)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.AmountE8)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.PriceE8)...)
	bs = append(bs, order.Ioc)
	bs = append(bs, order.Action)
	bs = append(bs, Uint32ToBigEndianBytes(order.PairId)...)

	if len(bs) != 98 { // 784 bits
		logger.Panic("The byte length of signing an order must be 98, but got ", len(bs))
	}
	return bs
}

func SignOrder(marketAddr common.Address, nonce uint64, order *Order, privKey *ecdsa.PrivateKey) (
	[]byte, error) {

	bytesToSign := GetOrderBytesToSign(marketAddr, nonce, order)
	hash := crypto.Keccak256(bytesToSign)
	return crypto.Sign(hash, privKey)
}

// The returned bytes are to be hashed (using keccak256) for signing. The content are the
// concatenation of the following (uints are in big-endian byte order):
//
// 1. String "DEx2 Order"(80)
// 2. <market address>(160)      (For preventing cross-market replay attack)
// 3. <nonce>(64) <expireTimeSec>(64) <amountE8>(64) <priceE8>(64) <ioc>(8) <action>(8) <pairId>(32)
func GetOrderBytesToSignWithTypes(marketAddr common.Address, nonce uint64, order *Order) []byte {
	bs := make([]byte, 0, 68)
	bs = append(bs, []byte("DEx2 Order")...)
	bs = append(bs, marketAddr.Bytes()...)
	bs = append(bs, Uint64ToBigEndianBytes(nonce)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.ExpireTimeSec)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.AmountE8)...)
	bs = append(bs, Uint64ToBigEndianBytes(order.PriceE8)...)
	bs = append(bs, order.Ioc)
	bs = append(bs, order.Action)
	bs = append(bs, Uint32ToBigEndianBytes(order.PairId)...)

	if len(bs) != 68 { // 544 bits
		logger.Panic("The byte length of signing an order must be 68, but got ", len(bs))
	}
	return bs
}

// `sig` is in the [R || S || V] format where V is 0 or 1.
func VerifyOrderSig(marketAddr common.Address, nonce uint64, order *Order,
	traderAddr common.Address, sig []byte) bool {

	bytesToSign := GetOrderBytesToSign(marketAddr, nonce, order)
	if ueth.VerifySig(bytesToSign, sig, traderAddr) {
		return true
	}

	bytesToSignWithTypes := GetOrderBytesToSignWithTypes(marketAddr, nonce, order)
	hashValues := crypto.Keccak256(bytesToSignWithTypes)

	hash := crypto.Keccak256(OrderHashTypes, hashValues)
	return ueth.VerifySigByHash(hash, sig, traderAddr)
}
