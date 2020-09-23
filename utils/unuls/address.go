package unuls

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"git.bibox.com/dextop/common.git/utils/unuls/btcec"
	"golang.org/x/crypto/ripemd160"
)

const (
	ADDRESS_LENGTH = 23

	MAINNET_CHAIN_ID = 1;
	TESTNET_CHAIN_ID = 2;
	MAINNET_DEFAULT_ADDRESS_PREFIX = "NULS";
	TESTNET_DEFAULT_ADDRESS_PREFIX = "tNULS";

	DEFAULT_ADDRESS_TYPE = 1;

)

type Address struct {
	prefix string
	chainId int
	addressStr string
	addressType byte
	hash160 []byte
	addressBytes []byte
}

func GetAddressStringFrom(priv *btcec.PrivateKey) (string) {
	pubKeyHash160 := Ripemmd160OfSha256(priv.PubKey())

	address := NewAddress(TESTNET_CHAIN_ID, DEFAULT_ADDRESS_TYPE, pubKeyHash160)

	return address.string()
}

func  Ripemmd160OfSha256(pub *btcec.PublicKey) ([]byte) {
	pubkeyBytes := pub.SerializeCompressed()

	sha256Bytes := sha256.Sum256(pubkeyBytes)
	tempBytes := make([]byte, 0)
	tempBytes = append(tempBytes, sha256Bytes[:]...)

	hasher := ripemd160.New()
	hasher.Write(tempBytes)
	hashBytes := hasher.Sum(nil)

	return hashBytes
}

func NewAddress(chainId int, addressType byte, hash160 []byte) (*Address) {
	addr := new(Address)
	addr.chainId = chainId
	addr.addressType = addressType
	addr.hash160 = hash160
	addr.prefix = addr.getPrefix()
	addr.addressBytes = addr.calcAddressbytes()

	return addr
}

func (addr *Address) calcAddressbytes()  ([]byte) {
	body := make([]byte, 0)

	chainIdTemp := int16(addr.chainId)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, &chainIdTemp)

	body = append(body, bytesBuffer.Bytes()...)

	body = append(body, addr.addressType)

	body = append(body, addr.hash160...)

	return body
}

func  (addr *Address) getPrefix() (string) {
	if addr.chainId == MAINNET_CHAIN_ID {
		return MAINNET_DEFAULT_ADDRESS_PREFIX
	} else if addr.chainId == TESTNET_CHAIN_ID {
		return  TESTNET_DEFAULT_ADDRESS_PREFIX
	}

	return ""
}

func (addr *Address)getStringAddressByBytes() (addrStr string){
	if (len(addr.addressBytes) == 0) {
		return "";
	}
	if (len(addr.addressBytes) != ADDRESS_LENGTH) {
		return "";
	}

	resultBytes := make([]byte, 0)

	//resultBytes = append(resultBytes, []byte(addr.prefix)...)
	resultBytes = append(resultBytes, addr.addressBytes...)

	xor := getXOR(addr.addressBytes)
	resultBytes = append(resultBytes, xor)

	str := hex.EncodeToString(resultBytes)
	_ = str
	encoded := B58encode(resultBytes)

	lenthPrefix := getLengthPrefix(len(addr.prefix))
	resultBytes = append([]byte(addr.prefix), lenthPrefix...)
	resultBytes = append(resultBytes, []byte(encoded)[:]...)

	return string(resultBytes)
}


func getXOR(src []byte) (byte) {
	var check byte
	for i := 0; i < len(src);i++  {
		check ^= src[i]
	}

	return check
}

func getLengthPrefix(length int) []byte {
	prefixs := []string{"", "a", "b", "c", "d", "e"};

	return []byte(prefixs[length])
}

func (addr *Address) string() string {

	addrStr := addr.getStringAddressByBytes()

	return addrStr
}