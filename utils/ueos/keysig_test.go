package ueos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyMsgSig(t *testing.T) {
	endPoint := "https://eos.greymass.com"
	account := "mushanzilike"
	msg := "1552911136" + "mushanzilike" + "web-zPbdG8RJ" + "TokenPocket" //time + account + uuid + ref
	sig := "SIG_K1_KXkkC5T8NP4FmqWD57Dab8NN3CnaG6SkJwFeZZxFv4xLEs17nhfNV7gier5oHXvr6dGuSLkNPfgb1N6HSeJv6V1ZVo3L2z"

	rs := VerifyMsgSig(endPoint, account, msg, sig)

	assert.True(t, rs)
}

func TestVerifyDextopMsgSig(t *testing.T) {
	endPoint := "https://api-kylin.eoslaomao.com"
	account := "dextopalpha1"
	sig := "SIG_K1_KeTLxtbfSWY45DXsmHex51UhjLkUYD7qw1r8Cd9ZWnWoeW36km9ddwfrYryoiLi8SdSBHf3Ugw8JcjLksbxBeZtM9eFS4C"

	rs := VerifyDextopMsgSig(endPoint, account, "1553495946", "86400", sig)

	assert.True(t, rs)
}

func TestVerifyWithdrawMsg(t *testing.T) {
	endPoint := "https://api-kylin.eoslaomao.com"
	account := "dextopalpha1"
	sig := "SIG_K1_K7DQr4B9Zr2HiCVKMXSVi5Kb9fv37qQ3aZ1md5LAGWegjKxo21aATWmyz9dztvMSdQUPbs3KVrMTxgqMwTU6VLeVH3ThV3"

	rs := VerifyWithdrawMsg(endPoint, account, "1.0000 EOS", "1553658236", sig)

	assert.True(t, rs)
}

func TestVerifyMyKeySig(t *testing.T) {
	endPoint := "https://eos.greymass.com"
	account := "mykeydex1234"
	sig := "SIG_K1_K3MPiRhaW8pNPdCKPXfpSW95TqHkhcne7ei8Ywn9FhyHHifktfgixDsPoRpfhyZ9Ep1q6i5ndtX5ir8c39spqTVYdS4yuo"

	rs := VerifyDextopMsgSig(endPoint, account, "1567665744", "86400", sig)

	assert.True(t, rs)
}