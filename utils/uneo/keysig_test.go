package uneo

import "testing"

func TestVerifyMsg(t *testing.T) {
	r := VerifyMsg("038e52dcf7da9ee0c5dcded5636356f014cc02bf850f01c0663c0d57788abcc5af",
		"61e43ed84c790987a18a01dde0338d96eb7c5897ca9db53d5c53e14cf183d2018622226ead4973bcdf04e7bed3295966de8ada369fe81f6abb5177d8cde9ee43",
		"123456789", "")

	t.Log(r)
}

func TestVerifyLoginMsg(t *testing.T) {
	r := VerifyLoginMsg("038e52dcf7da9ee0c5dcded5636356f014cc02bf850f01c0663c0d57788abcc5af",
		"e180bd21cb5cd5849aa00a00376dd7183cb7eda13514ac929279132a94204f92fe1ef200afce1960b6ac91054a94cf44cdd48613f3d56076a34436266f8c52fc",
		"AMiHZGwY1tBbAbGZD1JvsEDbcPheV9qKpF", "1568618465", "86400", "413fe4e3462f7918773bdb0fe95e7b05")

	t.Log(r)
}

func TestVerifyMakeOrderMsg(t *testing.T) {
	r := VerifyMakeOrderMsg("038e52dcf7da9ee0c5dcded5636356f014cc02bf850f01c0663c0d57788abcc5af",
		"64929a803b8465dd9b32e9d6c264dbd6a2697471178a19e0b67d74d39bbd4d6436e91d2dbfa3be6b267f6eaaf6c01c0520460ca59a2e2d3e4e946da4a239fc54",
		"SEA", "NEO", "0", "1500000000", "25000000", "0", "1564580627", "0", "")

	t.Log(r)
}

func TestVerifyWithdrawMsg(t *testing.T) {
	r := VerifyWithdrawMsg("038e52dcf7da9ee0c5dcded5636356f014cc02bf850f01c0663c0d57788abcc5af",
		"02574bda7aaff085b2210a537d15b475305ebd0711aff4cc548fcd7835e4a4c4e17f0ab7be9d3ec4b3135921e341fbb4723c2782f0fc8bbf9d8539b06b12de69",
		"AMiHZGwY1tBbAbGZD1JvsEDbcPheV9qKpF", "NEO", "1500000000", "1564580627", "")

	t.Log(r)
}
