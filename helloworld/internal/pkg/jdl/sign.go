package jdl

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strings"
	"time"
)

// 构建带有参数的URL
func (c *JdLClient) buildSignURL(fullUrl, urlPath string, body string) (string, error) {
	// accessToken, err := c.getAccessToken(context.Background())
	// if err != nil {
	// 	c.log.Error("get access token err:", err)
	// 	return "", err
	// }
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("timeStamp: %s", timeStamp)

	signstr, err := generateSign(
		c.config.AppSecret,
		c.token.AccessToken,
		c.config.AppKey,
		urlPath,
		body,      //body
		timeStamp, //使用当前时间string
		c.config.V,
	)
	if err != nil {
		return "", err
	}

	log.Printf("c.config.AppSecret: %s", c.config.AppSecret)
	log.Printf("c.token.AccessToken: %s", c.token.AccessToken)
	log.Printf("c.config.AppKey: %s", c.config.AppKey)
	log.Printf("urlPath: %s", urlPath)
	log.Printf("body: %s", body)
	log.Printf("timeStamp: %s", timeStamp)
	log.Printf("c.config.V: %s", c.config.V)
	log.Printf("LOP-DN: %s", c.config.LOP_DN)
	log.Printf("signstr: %s", signstr)

	reqParams := map[string]string{
		"app_key":      c.config.AppKey,
		"access_token": c.token.AccessToken,
		"timestamp":    timeStamp,
		"v":            c.config.V,
		"LOP-DN":       c.config.LOP_DN,
		"sign":         signstr,
	}

	signUrl, err := getURLParams(fullUrl, reqParams) //返回签名后的url
	if err != nil {
		return "", err
	}

	return signUrl, nil
}

// 生成URL参数签名
func generateSign(appSecret, accessToken, appKey, method, body, timestamp, v string) (string, error) {
	// content := appSecret + "access_token" + accessToken + "app_key" + appKey + "method" + method + "param_json" + param_json + "timestamp" + timestamp + "v" + v + appSecret

	content := strings.Join([]string{
		appSecret,
		"access_token", accessToken,
		"app_key", appKey,
		"method", method,
		"param_json", body,
		"timestamp", timestamp,
		"v", v,
		appSecret,
	}, "")

	sign := md5.Sum([]byte(content))
	signStr := hex.EncodeToString(sign[:])

	return signStr, nil
}
