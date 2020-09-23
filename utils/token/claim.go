package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gisvr/defi-common/utils/ueth"
	"github.com/satori/go.uuid"
)

const (
	DefaultTokenTTL = time.Hour * 72
	IssueBefore     = time.Second * 5
)

type claimSet struct {
	StdClaims *jwt.StandardClaims `json:"StdClaims"`
}

func NewClaims(options ...claimOption) *claimSet {
	claims := &claimSet{
		StdClaims: newStandardClaims(time.Now()),
	}
	for _, opt := range options {
		opt(claims)
	}
	return claims
}

func newStandardClaims(now time.Time) *jwt.StandardClaims {
	id := uuid.NewV4().String()
	return &jwt.StandardClaims{
		Id:        id,
		Subject:   "address",
		Audience:  "eth/eos",
		Issuer:    "dex.top",
		NotBefore: 12345, // userId
		// token的发行时间提前5秒，避免： Token used before issued
		IssuedAt:  now.Add(-IssueBefore).Unix(),
		ExpiresAt: now.Add(DefaultTokenTTL).Unix(),
	}
}

func (claims *claimSet) Valid() error {
	var addr = claims.StdClaims.Subject
	if ueth.IsHexAddr(addr) && ueth.IsNormalizedHexAddr(addr) {
		return claims.StdClaims.Valid()
	}
	return claims.StdClaims.Valid()
}

func (claims *claimSet) Id() string {
	return claims.StdClaims.Id
}

/**
 * 获取Subject存放的当前用户ETH地址，或EOS帐号。
 */
func (claims *claimSet) Trader() string {
	return claims.StdClaims.Subject
}

/**
 * 获取存放在Audience中的chain信息。
 */
func (claims *claimSet) Chain() string {
	return claims.StdClaims.Audience
}

/**
 * 获取存放在NotBefore中的用户id。
 */
func (claims *claimSet) UserID() int64 {
	return claims.StdClaims.NotBefore
}

type claimOption func(*claimSet)

func ID(id string) claimOption {
	return func(claims *claimSet) {
		claims.StdClaims.Id = id
	}
}

/**
 * 将地址/帐号信息放到Subject字段中。
 */
func Trader(addr string) claimOption {
	return func(claims *claimSet) {
		claims.StdClaims.Subject = addr
	}
}

/**
 * 将chain数据放入到Audience字段中。
 */
func ChainData(chain string) claimOption {
	return func(claims *claimSet) {
		claims.StdClaims.Audience = chain
	}
}

/**
 * 将用户ID放入到NotBefore字段中。
 */
func UserID(userId int64) claimOption {
	return func(claims *claimSet) {
		claims.StdClaims.NotBefore = userId
	}
}

func IssuedAtTimeSec(timeSec int64) claimOption {
	return func(claims *claimSet) {
		claims.StdClaims.IssuedAt = timeSec
	}
}

// Expire at the give timestamp (in second).
func ExpiresAtTimeSec(timeSec int64) claimOption {
	return func(claims *claimSet) {
		claims.StdClaims.ExpiresAt = timeSec
	}
}
