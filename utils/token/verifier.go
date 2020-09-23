package token

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type verifier struct {
	publicKey crypto.PublicKey
}

func unmarshalPubkey(pk string) crypto.PublicKey {
	bs, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		panic(err)
	}
	pubkey, err := x509.ParsePKIXPublicKey(bs)
	if err != nil {
		panic(err)
	}
	return pubkey.(crypto.PublicKey)
}

func NewVerifierFromMarshaledKey(marshaledKey string) Verifier {
	return &verifier{publicKey: unmarshalPubkey(marshaledKey)}
}

func (ver *verifier) parse(jwtToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(jwtToken, &claimSet{}, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !(ok && method.CurveBits == 256) {
			return nil, &jwt.ValidationError{
				Inner:  fmt.Errorf("unexpected signing method: %v", token.Header["alg"]),
				Errors: jwt.ValidationErrorSignatureInvalid}
		}
		return ver.publicKey, nil
	})
}

func (tv *verifier) VerifyToken(jwtStr string) (Claims, error) {
	token, err := tv.parse(jwtStr)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*claimSet)

	if !ok {
		return nil, &jwt.ValidationError{
			Inner:  errors.New("unexpected claim type"),
			Errors: jwt.ValidationErrorUnverifiable}
	}
	return claims, nil
}
