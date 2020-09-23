package uneo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gisvr/deif-common/utils/ulog"
	"github.com/CityOfZion/neo-go/pkg/crypto"
	"math/big"
)

func VerifyMsg(pubStr, sig, msg, salt string) bool {
	pubKey, err := crypto.NewPublicKeyFromString(pubStr)

	if err != nil {
		ulog.Error("pubKey parse error ", err)
		return false
	}

	eccPubKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X: pubKey.X,
		Y: pubKey.Y,
	}

	h := sha256.New()
	h.Write([]byte(msg))

	sigHex, _ := hex.DecodeString(sig)
	r := new(big.Int).SetBytes(sigHex[:32])
	s := new(big.Int).SetBytes(sigHex[32:])

	rs := ecdsa.Verify(eccPubKey, h.Sum(nil), r, s)

	if !rs {
		//NEOLine 验签调整 https://github.com/NeoNextClub/neoline/commit/300254f4400d64ed29667b31e04fecf130ca967d
		msgByte := []byte(salt + msg)
		msgByte = append(new(big.Int).SetInt64(int64(len(msgByte))).Bytes(), msgByte...)
		msgByte = append([]byte{0x01,0x00,0x01,0xf0}, msgByte...)
		msgByte = append(msgByte, []byte{0x00, 0x00}...)
		h := sha256.New()
		h.Write(msgByte)
		rs = ecdsa.Verify(eccPubKey, h.Sum(nil), r, s)
	}

	return rs
}

func VerifyLoginMsg(pubStr, sig, account, time, duration, salt string) bool {
	msg := fmt.Sprintf("DEx2: Authenticate Trader <%v> at time <%v> for <%v> Seconds",
		account, time, duration)
	return VerifyMsg(pubStr, sig, msg, salt)
}

func VerifyMakeOrderMsg(pubStr, sig, stock, cash, action, amount, price, tp, timestamp, expire, salt string) bool {
	// DEx2 Order: <stock> <cash> <action> <amount> <price> <type> <timestamp> <expiretime>
	msg := fmt.Sprintf("DEx2 Order: <%v> <%v> <%v> <%v> <%v> <%v> <%v> <%v>",
		stock, cash, action, amount, price, tp, timestamp, expire)
	return VerifyMsg(pubStr, sig, msg, salt)
}

func VerifyWithdrawMsg(pubStr, sig, account, asset, amount, time, salt string) bool {
	msg := fmt.Sprintf("DEx2: Trader <%v> Withdraws <%v> <%v> at Timestamp %v",
		account, asset, amount, time)
	return VerifyMsg(pubStr, sig, msg, salt)
}