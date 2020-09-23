package unuls

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gisvr/defi-common/utils/ulog"
	"github.com/gisvr/defi-common/utils/unuls/btcec"
)

func VerifyRpcSingedMsg(msg string, signed string, address string) bool {
	signedBytes, _ := hex.DecodeString(signed)
	//解析txHexStr，获得txHex，pubKeyStr，signedStr
	decoder := NewDecoder(signedBytes)

	decoder.ParseUint8()
	pubKeyHexLength, _ := decoder.ParseUint8()
	pubKeyHexBytes, _ := decoder.ParseBytesWithLength(int(pubKeyHexLength))
	//fmt.Println("pubKeyHexBytes:", pubKeyHexBytes)

	signedDataLength, _ := decoder.ParseUint8()
	signedData , _ := decoder.ParseBytesWithLength(int(signedDataLength))
	//fmt.Println("signedData:", signedData)

	pubKey, err := btcec.ParsePubKey(pubKeyHexBytes, btcec.S256())
	if err != nil {
		return false
	}
	//创建signature
	signature, err := btcec.ParseSignature(signedData, btcec.S256())
	if err != nil {
		return false
	}
	msgBytes, _ := hex.DecodeString(msg)
	sum := sha256.Sum256(msgBytes)
	sum = sha256.Sum256(sum[:])
	if !signature.Verify(sum[:], pubKey) {
		return false
	}

	pubKeyHash160 := Ripemmd160OfSha256(pubKey)

	addressFromPubKey := NewAddress(TESTNET_CHAIN_ID, DEFAULT_ADDRESS_TYPE, pubKeyHash160)

	if address != addressFromPubKey.string(){
		 return false
	}
	return true
}


func VerifyTx(signedTx string) bool {

	var txBytesLength uint64

	signedBytes, _ := hex.DecodeString(signedTx)
	//解析txHexStr，获得txHex，pubKeyStr，signedStr
	decoder := NewDecoder(signedBytes)

	txType, err := decoder.ParseBytesWithLength(2)
	txBytesLength += 2
	ulog.Debugln("type:", hex.EncodeToString(txType))

	time, err := decoder.ParseBytesWithLength(4)
	txBytesLength += 4
	ulog.Debugln("time:", hex.EncodeToString(time))

	remark, length, lengthBytesLen, _ := decoder.ParseByteByLength()
	txBytesLength += length
	txBytesLength += uint64(lengthBytesLen)
	ulog.Debugln("remark:", hex.EncodeToString(remark))

	txData, length, lengthBytesLen, _ :=decoder.ParseByteByLength()
	txBytesLength += length
	txBytesLength += uint64(lengthBytesLen)
	ulog.Debugln("txData:", hex.EncodeToString(txData))

	coinData, length, lengthBytesLen, _ :=decoder.ParseByteByLength()
	txBytesLength += length
	txBytesLength += uint64(lengthBytesLen)
	ulog.Debugln("coinData:", hex.EncodeToString(coinData))

	decoder.ParseUint8()
	pubKeyHexLength, _ := decoder.ParseUint8()
	pubKeyHexBytes, _ := decoder.ParseBytesWithLength(int(pubKeyHexLength))
	ulog.Debugln("pubKeyHexBytes:", pubKeyHexBytes)

	signedDataLength, _ := decoder.ParseUint8()
	signedData , _ := decoder.ParseBytesWithLength(int(signedDataLength))

	ulog.Debugln("signedData:", signedData)
	//通过pubKey得到PublicKey
	//通过signedStr得到sig.R sig.S

	pubKey, err := btcec.ParsePubKey(pubKeyHexBytes, btcec.S256())
	if err != nil {
		return false
	}
	//创建signature
	signature, err := btcec.ParseSignature(signedData, btcec.S256())
	if err != nil {
		return false
	}
	//txHex的hash
	msg := signedBytes[0:txBytesLength]
	sum := sha256.Sum256(msg)
	sum = sha256.Sum256(sum[:])

	return signature.Verify(sum[:], pubKey)
}