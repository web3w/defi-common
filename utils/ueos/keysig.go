package ueos

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gisvr/deif-common/utils/ulog"
	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

/**
 * 对直接对信息签名的方式验签。
 */
func VerifyMsgSig(endPoint, account, msg string, sigStr string) bool {
	resp, err := eos.New(endPoint).GetAccount(eos.AN(account))
	if err != nil {
		ulog.Info(err)
		return false
	}

	if len(resp.Permissions) == 0 {
		return false
	}

	sig, err := ecc.NewSignature(sigStr)
	if err != nil {
		return false
	}

	h := sha256.New()
	h.Write([]byte(msg))

	if len(resp.Permissions[0].RequiredAuth.Keys) == 0 {
		//MyKey
		if len(resp.Permissions[0].RequiredAuth.Accounts) == 0 ||
			resp.Permissions[0].RequiredAuth.Accounts[0].Permission.Actor != "mykeymanager" {
			return false
		}
		resp2, err := eos.New(endPoint).GetTableRows(eos.GetTableRowsRequest{
			Code: "mykeymanager",
			Table: "keydata",
			Scope: account,
			Limit: 1,
			LowerBound: "3",
			JSON: true,
		})
		if err != nil {
			ulog.Info(err)
			return false
		}
		var rows []interface{}
		_ = json.Unmarshal(resp2.Rows, &rows)
		pubStr := rows[0].(map[string]interface{})["key"].(map[string]interface{})["pubkey"].(string)
		fmt.Println(pubStr)
		pub, _ := ecc.NewPublicKey(pubStr)
		return sig.Verify(h.Sum(nil), pub)
	}

	for _, perm := range resp.Permissions {
		for _, key := range perm.RequiredAuth.Keys {

			pub := key.PublicKey

			if sig.Verify(h.Sum(nil), pub) {
				return true
			}
		}
	}
	return false
}

func VerifyDextopMsgSig(endPoint, account, time, duration, sigStr string) bool {
	msgPre := fmt.Sprintf("DEx2: Authenticate Trader <%s> at time <%s> for <%s> Seconds",
		account, time, duration)

	msgSign := ""
	for i := 0; i < len(msgPre); i += 12 {
		to := i + 12
		if to > len(msgPre) {
			to = len(msgPre)
		}
		msgSign += msgPre[i:to]
		msgSign += " "
	}
	msgSign = msgSign[:len(msgSign)-1]
	return VerifyMsgSig(endPoint, account, msgSign, sigStr)
}

func VerifyWithdrawMsg(endPoint, account, asset, time, sigStr string) bool {
	msgPre := fmt.Sprintf("DEx2: Trader <%v> Withdraws <%v> at Timestamp %v",
		account, asset, time)

	msgSign := ""
	for i := 0; i < len(msgPre); i += 12 {
		to := i + 12
		if to > len(msgPre) {
			to = len(msgPre)
		}
		msgSign += msgPre[i:to]
		msgSign += " "
	}
	msgSign = msgSign[:len(msgSign)-1]
	return VerifyMsgSig(endPoint, account, msgSign, sigStr)
}
