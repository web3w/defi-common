package dex2

import (
	"encoding/binary"
	"fmt"
	"github.com/gisvr/defi-common/utils/ubig"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// u = (u << 8) | v
func PushUint8(u *big.Int, v uint8) {
	u.Or(u.Lsh(u, 8), big.NewInt(int64(v)))
}

// u = (u << 16) | v
func PushUint16(u *big.Int, v uint16) {
	u.Or(u.Lsh(u, 16), big.NewInt(int64(v)))
}

// u = (u << 32) | v
func PushUint32(u *big.Int, v uint32) {
	u.Or(u.Lsh(u, 32), big.NewInt(int64(v)))
}

// u = (u << 64) | v
func PushUint64(u *big.Int, v uint64) {
	u.Or(u.Lsh(u, 64), new(big.Int).SetUint64(v))
}

// u = (u << 160) | v
func PushAddress(u *big.Int, a common.Address) {
	u.Or(u.Lsh(u, 160), a.Hash().Big())
}

func Uint64ToBigEndianBytes(value uint64) []byte {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, value)
	return bs
}

func Uint32ToBigEndianBytes(value uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, value)
	return bs
}

func GetTokenAccountKey(traderAddr common.Address, tokenCode uint16) *big.Int {
	key := new(big.Int).Lsh(big.NewInt(int64(tokenCode)), 160)
	key.Or(key, traderAddr.Hash().Big())
	return key
}

func GetOrderKey(traderAddr common.Address, nonce uint64) *big.Int {
	key := new(big.Int).Lsh(new(big.Int).SetUint64(nonce), 160)
	key.Or(key, traderAddr.Hash().Big())
	return key
}

// returns `u & 0xFFFFFFFFFFFFFFFF` and `u = u >> 64`
func PopUint64(u *big.Int) uint64 {
	result := ubig.And(u, ubig.U64(0xFFFFFFFFFFFFFFFF)).Uint64()
	u.Rsh(u, 64)
	return result
}

// returns `u & 0xFFFFFFFF` and `u = u >> 32`
func PopUint32(u *big.Int) uint32 {
	result := ubig.And(u, ubig.U64(0xFFFFFFFF)).Uint64()
	u.Rsh(u, 32)
	return uint32(result)
}

// returns `u & 0xFFFF` and `u = u >> 16`
func PopUint16(u *big.Int) uint16 {
	result := ubig.And(u, ubig.U64(0xFFFF)).Uint64()
	u.Rsh(u, 16)
	return uint16(result)
}

// returns `u & 0xFF` and `u = u >> 8`
func PopUint8(u *big.Int) uint8 {
	result := ubig.And(u, ubig.U64(0xFF)).Uint64()
	u.Rsh(u, 8)
	return uint8(result)
}

func PopUint160(u *big.Int) *big.Int {
	mask := ubig.Sub(ubig.Exp(ubig.U64(2), ubig.U64(160)), ubig.U64(1))
	result := ubig.And(u, mask)
	u.Rsh(u, 160)
	return result
}

// Returns an error if the scale factor is invalid to Dex2 SC; otherwise returns the corresponding
// decimals of the token and nil error.
func CheckTokenScaleFactor(scaleFactor uint64) ( /*decimals*/ int, error) {
	str := fmt.Sprintf("%d", scaleFactor)
	if str != "1"+strings.Repeat("0", len(str)-1) {
		return 0, fmt.Errorf("scale factor %v is not a power of 10.", scaleFactor)
	}
	decimals := len(str) - 1
	if decimals > 18 {
		return 0, fmt.Errorf("scale factor %v is too large", scaleFactor)
	}
	return decimals, nil
}
