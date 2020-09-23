package uhttp

import (
	"github.com/gisvr/deif-common/utils/ulog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
 * 发送POST请求并返回收到的结果。
 */
func DoPost(urlStr string, params map[string]string) string {
	msg := url.Values{}
	if params != nil && len(params) != 0 {
		for key, value := range params {
			msg.Add(key, value)
		}
	}

	var result = ""
	msgReader := strings.NewReader(msg.Encode())
	req, err := http.NewRequest("POST", urlStr, msgReader)
	if err != nil {
		ulog.Errorf("err occur when create POST request. url: %v, params: %v", urlStr,
			msg.Encode(), err)
		return result
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	var httpClient = http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		ulog.Errorf("err occur when client do request. url: %v, params: %v",
			urlStr, msg.Encode(), err)
		return result
	}

	if resp.StatusCode != 200 {
		ulog.Errorf("DoPost return: %d, url: %v, params: %v", resp.StatusCode, urlStr, msg.Encode())
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ulog.Errorf("err occur when read resp's body. url: %v", urlStr, err)
		return result
	}
	return string(bs)
}

/**
 * 使用body方式发送POST请求并返回收到的结果。
 */
func DoPostBody(urlStr, bodyMsg string) string {
	var result = ""
	if len(bodyMsg) == 0 {
		return result
	}

	msgReader := strings.NewReader(bodyMsg)
	req, err := http.NewRequest("POST", urlStr, msgReader)
	if err != nil {
		ulog.Errorf("err occur when create POST request. url: %v, params: %v", urlStr,
			bodyMsg, err)
		return result
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	var httpClient = http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		ulog.Errorf("err occur when client do request. url: %v, params: %v",
			urlStr, bodyMsg, err)
		return result
	}

	if resp.StatusCode != 200 {
		ulog.Errorf("DoPost return: %d, url: %v, params: %v", resp.StatusCode, urlStr, bodyMsg)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ulog.Errorf("err occur when read resp's body. url: %v", urlStr, err)
		return result
	}
	return string(bs)
}

/**
 * 发送GET请求并返回收到的结果。
 */
func DoGet(urlStr string) string {
	var result = ""
	resp, err := http.Get(urlStr)
	if err != nil {
		ulog.Errorf("err occur when use http.Get. url: %v", urlStr, err)
		return result
	}

	if resp.StatusCode != 200 {
		ulog.Errorf("DoGet return: %d, url: %v", resp.StatusCode, urlStr)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ulog.Errorf("err occur when read resp's body. url: %v", urlStr, err)
		return result
	}
	return string(bs)
}
