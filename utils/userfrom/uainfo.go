package userfrom

import (
	"regexp"
	"strings"
)

/*
  根据UA 获取 userFrom, device, authType
*/
func GetUserWalletInfo(ua string, wallets map[string]string) (userFrom, device, authType string) {
	// 添加用户登陆记录。
	if ua == "" {
		device = "web"
		authType = "QR"
	} else {
		device = getDevice(ua)
		if device == "android" || device == "iphone" {
			authType = "Dapp"
			if uaUserFrom := getUserFrom(ua, wallets); uaUserFrom != "" {
				userFrom = uaUserFrom
			}
		} else {
			authType = "Plug"
		}
	}
	return
}

/*
  根据特征匹配 终端设备
*/
func getDevice(uainfo string) (deviceType string) {
	devreg, _ := regexp.Compile("(?i:android|iphone)")
	devices := devreg.FindAllString(uainfo, -1)
	if len(devices) > 0 {
		deviceType = strings.ToLower(devices[0])
	} else {
		deviceType = "web"
	}
	return
}

/*
  根据字典表 匹配渠道信息
*/
func getUserFrom(uainfo string, wallets map[string]string) (userfrom string) {
Loop:
	for key, val := range wallets {

		devreg, _ := regexp.Compile("(?i:" + key + ")")
		devices := devreg.FindAllString(uainfo, -1)
		if len(devices) > 0 {
			userfrom = val
			break Loop
		}
	}
	return
}
