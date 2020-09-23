package dex2

import (
	"encoding/hex"
	"github.com/gisvr/defi-common/utils/ueth"
	"github.com/gisvr/defi-common/utils/utest"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignatureTestSuite struct {
	utest.RequireSuite
}

// The hook of `go test`
func TestSignatureTestSuite(t *testing.T) {
	utest.Run(t, new(SignatureTestSuite))
}

func (t *SignatureTestSuite) SetupTest() {}

func (t *SignatureTestSuite) TestGoldenData1() {
	marketAddr := common.HexToAddress("0x221fEe57689Dec269481Df9BA38F5E15F322Dd10")
	nonce := uint64(1520217059531)
	order := &Order{
		PairId:        100,
		Action:        0,
		Ioc:           0,
		PriceE8:       0.5e8,
		AmountE8:      1e8,
		ExpireTimeSec: 1520303459,
	}
	privKey, err := ueth.HexToECDSAPrivKey(
		"24fce097d74931f6cddc1dbf0ef28dea1828d6332bed5cb45fae4046a7f8542a")
	t.NoError(err)
	traderAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	bytesToSign := GetOrderBytesToSign(marketAddr, nonce, order)
	hash := crypto.Keccak256(bytesToSign)
	sig, err := crypto.Sign(hash, privKey)
	t.NoError(err)

	t.True(VerifyOrderSig(marketAddr, nonce, order, traderAddr, sig))
	t.Equal("0x9081c1a1e309ab07e6086d6ec06303dc4a1c3276de9276b749b327f1d2815b0774894d4c92e03"+
		"93ec8acc840c4eecd7677ac2278e09af7a47139df467033608600", "0x"+hex.EncodeToString(sig))

	// fmt.Println("marketAddr:", marketAddr.Hex())
	// fmt.Println("nonce:", nonce)
	// fmt.Println("order:", *order)
	// fmt.Println("traderAddr:", traderAddr.Hex())
	// fmt.Println("bytesToSign:", "0x"+hex.EncodeToString(bytesToSign))
	// fmt.Println("sig:", "0x"+hex.EncodeToString(sig))
	// fmt.Println("----------------------------------------")
}

func (t *SignatureTestSuite) TestGoldenData2() {
	marketAddr := common.HexToAddress("0x771a4c54EaA7618582a11A2E8a025CA3e0A0B530")
	nonce := uint64(1520304377652)
	order := &Order{
		PairId:        101,
		Action:        0,
		Ioc:           0,
		PriceE8:       1e8,
		AmountE8:      0.1e8,
		ExpireTimeSec: 1520390777,
	}
	privKey, err := ueth.HexToECDSAPrivKey(
		"0x24fce097d74931f6cddc1dbf0ef28dea1828d6332bed5cb45fae4046a7f8542a")
	t.NoError(err)
	traderAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	bytesToSign := GetOrderBytesToSign(marketAddr, nonce, order)
	hash := crypto.Keccak256(bytesToSign)
	sig, err := crypto.Sign(hash, privKey)
	t.NoError(err)

	t.True(VerifyOrderSig(marketAddr, nonce, order, traderAddr, sig))
	t.Equal("0xdc44ce0f2bb962b90d4c990c07c574349bc46f326c15da9488fb74eb0c7c102b0df945f8cd5a5"+
		"8f0ce8422ba365d9e4345139eb9da28b27b2217ef9c6a14fd3700", "0x"+hex.EncodeToString(sig))
}

func (t *SignatureTestSuite) TestGoldenData3_Typed() {
	marketAddr := common.HexToAddress("0xd7e0DE6BAAD35e8812713dA4a9DC0b95eA6590b4")
	nonce := uint64(1522837528308)
	order := &Order{
		PairId:        100,
		Action:        0,
		Ioc:           0,
		PriceE8:       5000000000012345678,
		AmountE8:      10000000000123456789,
		ExpireTimeSec: 1522923928,
	}
	privKey, err := ueth.HexToECDSAPrivKey(
		"0x8DB6CC914CE26C37850130A0CCD61B3F103F43600E48FED229C4B3DB50FAE263")
	t.NoError(err)
	traderAddr := crypto.PubkeyToAddress(privKey.PublicKey)

	bytesToSignWithTypes := GetOrderBytesToSignWithTypes(marketAddr, nonce, order)
	hashValues := crypto.Keccak256(bytesToSignWithTypes)
	hash := crypto.Keccak256(OrderHashTypes, hashValues)

	sig, err := crypto.Sign(hash, privKey)
	t.NoError(err)

	t.True(VerifyOrderSig(marketAddr, nonce, order, traderAddr, sig))
	t.Equal(
		"0x7724c9427bdb43c230de790cdd4ee9237564f50f99770b7082f44a4b7686e495786dabdee79"+
			"6a6ca308b29939bf284a267bb3292f2c5a0a8d0bf3af724dc2d2300",
		"0x"+hex.EncodeToString(sig))
}

func (t *SignatureTestSuite) TestClosedLoop() {
	for i := 0; i < 100; i++ {
		marketAddr := t.getRandAddress()
		nonce := rand.Uint64()
		order := t.getRandOrder()
		privKey, err := crypto.GenerateKey()
		t.NoError(err)
		traderAddr := crypto.PubkeyToAddress(privKey.PublicKey)

		bytesToSign := GetOrderBytesToSign(marketAddr, nonce, order)
		hash := crypto.Keccak256(bytesToSign)
		sig, err := crypto.Sign(hash, privKey)
		t.NoError(err)

		t.True(VerifyOrderSig(marketAddr, nonce, order, traderAddr, sig))

		bytesToSignWithTypes := GetOrderBytesToSignWithTypes(marketAddr, nonce, order)
		hashValues := crypto.Keccak256(bytesToSignWithTypes)
		hash = crypto.Keccak256(OrderHashTypes, hashValues)
		sig, err = crypto.Sign(hash, privKey)
		t.NoError(err)

		t.True(VerifyOrderSig(marketAddr, nonce, order, traderAddr, sig))
	}
}

func (t *SignatureTestSuite) getRandAddress() common.Address {
	buf := make([]byte, common.AddressLength)
	rand.Read(buf)
	return common.BytesToAddress(buf)
}

func (t *SignatureTestSuite) getRandOrder() *Order {
	order := new(Order)
	order.PairId = rand.Uint32()
	order.Action = uint8(rand.Uint32() & 1)
	order.Ioc = uint8(rand.Uint32() & 1)
	order.PriceE8 = rand.Uint64()
	order.AmountE8 = rand.Uint64()
	order.ExpireTimeSec = rand.Uint64()
	return order
}
