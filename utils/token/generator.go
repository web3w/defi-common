package token

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"crypto/x509"
	"encoding/base64"

	"github.com/dgrijalva/jwt-go"
	"github.com/gisvr/defi-common/utils/ulog"
)

type generator struct {
	privateKey       *ecdsa.PrivateKey
	marshaledPubkey  string
	marshaledPrivkey string
}

func NewGenerator() Generator {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return &generator{
		privateKey:       privKey,
		marshaledPubkey:  marshalPubkey(privKey.Public().(*ecdsa.PublicKey)),
		marshaledPrivkey: marshalPrivKey(privKey),
	}
}

func NewGeneratorFromMarshaledKey(marshaledKey string) Generator {
	privKey := unmarshalPrivKey(marshaledKey)
	return &generator{
		privateKey:       privKey,
		marshaledPubkey:  marshalPubkey(privKey.Public().(*ecdsa.PublicKey)),
		marshaledPrivkey: marshalPrivKey(privKey),
	}
}

func marshalPubkey(pubK *ecdsa.PublicKey) string {
	bs, err := x509.MarshalPKIXPublicKey(pubK)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bs)
}

func marshalPrivKey(privK *ecdsa.PrivateKey) string {
	bs, err := x509.MarshalECPrivateKey(privK)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bs)
}

func unmarshalPrivKey(key string) *ecdsa.PrivateKey {
	bs, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	privKey, err := x509.ParseECPrivateKey(bs)
	if err != nil {
		panic(err)
	}
	return privKey
}

func (tg *generator) MarshaledVerifyKey() string {
	return tg.marshaledPubkey
}

func (tg *generator) MarshaledSignKey() string {
	return tg.marshaledPrivkey
}

func (tg *generator) GenerateToken(claims Claims) (string, error) {
	if err := claims.Valid(); err != nil {
		ulog.Error(err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signed, err := token.SignedString(tg.privateKey)
	if err != nil {
		ulog.Error(err)
		return "", err
	}
	return signed, nil
}
