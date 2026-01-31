package jdl

import (
	"context"
	"encoding/json"
	"net/url"
)

func (c *Client) doRequest(
	ctx context.Context,
	method, fullUrl, urlPath, domain string,
	reqParams any,
) ([]byte, error) {
	reqBody, err := json.Marshal(reqParams)
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(fullUrl, urlPath, domain, string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, method, signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	return respBodyBytes, nil
}

func getURLParams(fullUrl string, params map[string]string) (string, error) {
	URL, err := url.Parse(fullUrl)
	if err != nil {
		return "", err
	}

	query := URL.Query()
	for k, v := range params {
		query.Set(k, v)
	}

	URL.RawQuery = query.Encode() // 将编码后的 query 设置回 URL

	return URL.String(), nil
}
