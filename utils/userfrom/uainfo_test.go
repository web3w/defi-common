package userfrom

import (
	"fmt"
	"testing"
)

func TestGetUserWalletInfo(t *testing.T) {
	testList := map[string]string{
		"TP":  "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) TokenPocket_iOS",
		"TP2": "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) TokenPocket_iOS",
		"TP3": "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) TokenPocket_iOS",
	}

	for key, val := range testList {
		GetUserWalletInfo(testList[key], testList)
		fmt.Println(val)
	}

}
