package jdl

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"
)

// 构建带有参数的URL
func (c *Client) buildSignURL(fullUrl string, urlPath string, domain string, body string) (string, error) {
	// accessToken, err := c.getAccessToken(context.Background())
	timeStamp := time.Now().Format("2006-01-02 15:04:05")

	signstr, err := generateSign(
		c.config.AppSecret,
		c.config.AccessToken,
		c.config.AppKey,
		urlPath,
		body,      //body
		timeStamp, //使用当前时间string
		c.config.V,
	)
	if err != nil {
		return "", err
	}

	reqParams := map[string]string{
		"access_token": c.config.AccessToken,
		"app_key":      c.config.AppKey,
		"timestamp":    timeStamp,
		"v":            c.config.V,
		"LOP-DN":       domain,
		"sign":         signstr,
	}

	signUrl, err := getURLParams(fullUrl, reqParams) //返回签名后的url
	if err != nil {
		return "", err
	}

	return signUrl, nil
}

// 生成URL参数签名
func generateSign(appSecret, accessToken, appKey, path, body, timestamp, v string) (string, error) {
	// content := appSecret + "access_token" + accessToken + "app_key" + appKey + "method" + method + "param_json" + param_json + "timestamp" + timestamp + "v" + v + appSecret

	content := strings.Join([]string{
		appSecret,
		"access_token", accessToken,
		"app_key", appKey,
		"method", path,
		"param_json", body,
		"timestamp", timestamp,
		"v", v,
		appSecret,
	}, "")

	sign := md5.Sum([]byte(content))
	signStr := hex.EncodeToString(sign[:])

	return signStr, nil
}
