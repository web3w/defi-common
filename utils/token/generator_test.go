package token

import (
	"fmt"
	"testing"
)

func TestGenerator_GenerateToken(t *testing.T) {
	claimSet := NewClaims(ID("11"),
		Trader("0x36B1a29E0bbD47dFe9dcf7380F276e86da90c4c2"),
		ChainData("ETH"),
		UserID(12121),
		ExpiresAtTimeSec(1553418929))

	g := NewGenerator()
	jwtToken, _ := g.GenerateToken(claimSet)
	fmt.Println(jwtToken)

}
