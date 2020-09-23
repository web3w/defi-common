package token

import (
	"strings"
	"testing"

	"time"

	"math/rand"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
	g Generator
	v Verifier
}

func (t *TokenTestSuite) SetupTest() {
	t.g = NewGenerator()
	t.v = NewVerifierFromMarshaledKey(t.g.MarshaledVerifyKey())
}

func TestRun_TokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}

func (t *TokenTestSuite) TestTokenGeneration() {
	// Generate Token for user
	userClaimSet := NewClaims(UserID(3), IssuedAtTimeSec(time.Now().Unix()))
	token, err := t.g.GenerateToken(userClaimSet)
	if err != nil {
		panic(err)
	}

	cs, err := t.v.VerifyToken(token)
	t.NoError(err)

	t.Equal(userClaimSet, cs)

	// Generate Token for trader addr
	addrClaimSet := NewClaims(Trader("0xFc571a5fA85fd82393FbD3Ff9d74583a000C174c"))
	addrToken, err := t.g.GenerateToken(addrClaimSet)
	if err != nil {
		panic(err)
	}

	acs, err := t.v.VerifyToken(addrToken)
	t.NoError(err)

	t.Equal(addrClaimSet, acs)
}

func (t *TokenTestSuite) TestBadSignature() {
	claimSet := NewClaims(UserID(314159), ID(uuid.NewV4().String()))
	token, err := t.g.GenerateToken(claimSet)
	t.NoError(err)

	_, e := t.v.VerifyToken(token + strconv.Itoa(rand.Int()))
	t.Error(e)
}

func quickClock(sec int) func() time.Time {
	return func() time.Time {
		now := time.Now()
		then := now.Add(time.Second * time.Duration(sec))
		return then
	}
}

func (t *TokenTestSuite) TestExpiration() {
	claimSet := NewClaims(UserID(271818), ExpiresAtTimeSec(time.Now().Add(time.Second*3).Unix()))
	token, e := t.g.GenerateToken(claimSet)
	t.NoError(e)
	cs, e := t.v.VerifyToken(token)
	t.NoError(e)
	t.Equal(claimSet, cs)

	// to make token expire without waiting
	jwt.TimeFunc = quickClock(10)
	_, err := t.v.VerifyToken(token)
	t.Error(err)
}

func (t *TokenTestSuite) TestTokenSubject() {
	claimSet := NewClaims(UserID(0))
	_, e := t.g.GenerateToken(claimSet)
	t.Error(e)
}

func modifyTokenExpiration(g Generator, t string) string {
	v := NewVerifierFromMarshaledKey(g.MarshaledVerifyKey())
	parts := strings.Split(t, ".")
	claims, err := v.VerifyToken(t)
	if err != nil {
		panic(err)
	}
	sub := claims.UserID()

	//id, err := strconv.ParseUint(sub, 10, 64)
	//if err != nil {
	//	ulog.Panic(err)
	//}
	cs := NewClaims(UserID(sub), ExpiresAtTimeSec(time.Now().Add(time.Hour).Unix()))
	newToken, _ := g.GenerateToken(cs)
	newClaimStr := strings.Split(newToken, ".")[1]
	parts[1] = newClaimStr
	return strings.Join(parts, ".")
}

func (t *TokenTestSuite) TestHackedTokenVerification() {
	claimSet := NewClaims(UserID(314159))
	token, e := t.g.GenerateToken(claimSet)
	t.NoError(e)
	newToken := modifyTokenExpiration(t.g, token)
	_, err := t.v.VerifyToken(newToken)
	t.Error(err)
}

func (t *TokenTestSuite) TestValidationWithDifferentKey() {
	claimSet := NewClaims(UserID(271818))
	newGen := NewGenerator()
	token, e := newGen.GenerateToken(claimSet)
	t.NoError(e)
	_, err := t.v.VerifyToken(token)
	t.Error(err)
}
