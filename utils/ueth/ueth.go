package ueth

import (
	"strings"
)

const (
	AddrByteLen = 20
)

func HasHexPrefix(str string) bool {
	return strings.HasPrefix(str, "0x")
}

// Strip hex prefix if there is one.
func StripHexPrefix(hex string) string {
	return strings.TrimPrefix(hex, "0x")
}

func ChainNameById(chainId int64) string {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md
	switch chainId {
	case 1:
		return "Mainnet"
	case 2:
		return "Morden"
	case 3:
		return "Ropsten"
	case 4:
		return "Rinkeby"
	case 30:
		return "Rootstock Mainnet"
	case 31:
		return "Rootstock Testnet"
	case 42:
		return "Kovan"
	case 61:
		return "Ethereum Classic Mainnet"
	case 62:
		return "Ethereum Classic Testnet"
	case 1337:
		return "Geth Private Default"
	default:
		return "Unknown"
	}
}
