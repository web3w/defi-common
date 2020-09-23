package main

import (
	"fmt"
	"time"

	"github.com/gisvr/defi-common/utils/token"
)

func main() {
	g := token.NewGenerator()
	// marshal generator sign key
	signKey := g.MarshaledSignKey()
	// unmarshal key & construct a generator
	g = token.NewGeneratorFromMarshaledKey(signKey)
	claims := token.NewClaims(token.UserID(314159), token.IssuedAtTimeSec(time.Now().Unix()))

	jwtToken, err := g.GenerateToken(claims)

	if err != nil {
		panic(err)
	}

	marshaledKey := g.MarshaledVerifyKey()
	v := token.NewVerifierFromMarshaledKey(marshaledKey)
	claimSet, err := v.VerifyToken(jwtToken)
	if err != nil {
		panic(err)
	}
	id := claimSet.UserID()
	fmt.Println(id)
}
