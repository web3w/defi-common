package robot

import (
	"encoding/json"
	"github.com/gisvr/defi-common/utils/uhttp"
)

/**
 * 机器人信息返回结果。
 */
type NotifyResult struct {
	ErrorCode int    `json:"errorcode"`
	ErrMsg    string `json:"errmsg"`
}

func (res *NotifyResult) GetCode() int {
	return res.ErrorCode
}

func (res *NotifyResult) GetErrMsg() string {
	return res.ErrMsg
}

/**
 * 根据传入的信息组件通知内容字符串。有问题的情况下会返回空字符串。
 * 不带提醒。
 */
func GetMsgStr(msg string) string {
	var params = make(map[string]interface{})
	params["msgtype"] = "text"
	params["text"] = map[string]string{
		"content": msg,
	}

	result, _ := json.Marshal(params)
	return string(result)
}

/**
 * 统合信息发送处理。
 */
func SendRobotNotify(robotUrl, msg string) *NotifyResult {
	data := uhttp.DoPostBody(robotUrl, GetMsgStr(msg))
	if len(data) == 0 {
		return nil
	}

	var result = NotifyResult{}
	_ = json.Unmarshal([]byte(data), &result)
	return &result
}
