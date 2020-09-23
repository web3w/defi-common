package token

import (
	"testing"
	"time"
)

func BenchmarkGenerateToken(b *testing.B) {
	g := NewGenerator()
	for n := 0; n < b.N; n++ {
		_, _ = g.GenerateToken(NewClaims(UserID(1),
			ExpiresAtTimeSec(time.Now().Add(time.Hour).Unix()),
			IssuedAtTimeSec(time.Now().Unix())))
	}
}

func BenchmarkVerifyToken(b *testing.B) {
	g := NewGenerator()
	token, _ := g.GenerateToken(NewClaims(UserID(2)))
	marshaledKey := g.MarshaledVerifyKey()
	verifier := NewVerifierFromMarshaledKey(marshaledKey)
	for n := 0; n < b.N; n++ {
		_, _ = verifier.VerifyToken(token)
	}
}
