package main

import (
	"fmt"
	"math"
)

func main() {

	//creds, err := credentials.NewClientTLSFromFile(certFile, "")
	//if err != nil {
	//	ulog.Panic(err)
	//}
	//fmt.Println(creds)

	fmt.Println(math.Pow10(3))

	//privateStr := "581e0e12af8222581ec80354161013eb4a142a5a038a643509894338eb071a6b7745238316d02c101d8006d475bd1f696fb21d3ac4860d2cb07f0052b1a22189abfb0f8a38eb3f794062a3fed7fb05fc9371c93ff93ebaf86f3bda39d9b80e3b33bcaac9a6d6a7e928323c781d12eaebbbea0022b7221c3ba73b260724d70543b269166df2b0c55220b671e6ce00af6bd0787dd07c9e2b5edf33762421101a4cc65151e835c6ff1c83f30936a3a7085d53f8e15bd02f8a23ff81eed77e7c2b1ad36aafde09ff8a6a5b6615b8d15ebd8cb8653f137e991d8fef7898f1bedbcc3375df6dced3c75bf06c107122aac934e889561f56483cef810bf8dc90b6841bc70b3c6b7c240bbf87a1ceca825814100b16e8a94d2ed823c81f931a0e75b6537b93aec470e1ee20f8dd1083c77a056f74dee94e4f2e477da43b96290999cdf6d23ae2f4474fa294385807eaef368f5e63feff5c4c7942f868a336e0"
	//str, _ := cipher.Decrypt(privateStr, "aoeuaoeu")
	//fmt.Println("private decoded: ", str)
	//
	//privateKey, _ :=ueth.HexToECDSAPrivKey("0xc1f728fb4b5c1c1693592ff78c0ec0ffd5c3e4f498710703557141d9375514dd")
	//address := ueth.AddrOfPrivKey(privateKey)
	//fmt.Println("address: ", address.String())

	//fmt.Println(fmt.Sprintf("%v_%v", 12345, time.Now().Unix()))

	//marketAddr := common.HexToAddress("0xf8ea41950583020c590A0383Eeb184801b051fD6")
	//traderAddr := common.HexToAddress("0x853dd39d83F45Cf68CFDba40db10AF9B427C2a0C")
	//sigBytes := common.FromHex("0xa526cc626fbd2382df5b967f692795937e454f1bd75415018262dcd1ee381ecd013e8c3f571c1131674f571e3e6b4b74c283b1b3a1068c003a6e55e992aa6f9c01")
	//order := &dex2.Order{
	//	PairId:       (uint32(0) << 16) | uint32(1),
	//	Action:        0,
	//	Ioc:           0, // TODO: support this
	//	PriceE8:       10000000,
	//	AmountE8:      800000000,
	//	ExpireTimeSec: uint64(1553839809),
	//}
	//
	//flag := dex2.VerifyOrderSig(marketAddr, uint64(1553753409370), order, traderAddr, sigBytes)
	//fmt.Println(flag)
	//
	//str := "0xb3104b4b9da82025e8b9f8fb28b3553ce2f67069"
	//fmt.Println(ueth.NormalizeHexAddress(str))
	//
	//// DEx2: Authenticate Trader <0x853dd39d83f45cf68cfdba40db10af9b427c2a0c> at time <1553662750> for <86400> Seconds
	//
	//flag := VerifyTradeAddrAuthSig("0x853dd39d83f45cf68cfdba40db10af9b427c2a0c", 1553662291, 86400,
	//	common.FromHex("0x3a3a2cddbb7dbe04e3fef803b14c22bf11468bc2e5afa0c1639d698186f6e3b74d344059ca8576afc390a31f9b955c5568adfcbf74ff7d2d3a5c75ee3ff3cce301"))
	//fmt.Println(flag)
}

//func VerifyTradeAddrAuthSig(traderAddr string, timestampSec int64, durationSec int64, sig []byte) bool {
//	// The letter case of traderAddr is arbitrary, as long as the sig is valid.
//	msg := fmt.Sprintf(`DEx2: Authenticate Trader <%v> at time <%v> for <%v> Seconds`, traderAddr, timestampSec, durationSec)
//	signedStr := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(msg), msg)
//	return ueth.VerifySig([]byte(signedStr), sig, common.HexToAddress(traderAddr))
//}
