package jdl

import (
	"net/url"
)

// func (c *JdLClient) doRequest(
// 	ctx context.Context,
// 	method, fullUrl, urlPath string,
// 	reqParams interface{},
// ) ([]byte, error) {
// 	reqBody, err := json.Marshal(reqParams)
// 	if err != nil {
// 		c.log.Errorf("json.Marshal(req) err: %s", err)
// 		return nil, err
// 	}

// 	signURL, err := c.buildSignURL(fullUrl, urlPath, string(reqBody))
// 	if err != nil {
// 		c.log.Errorf("build sign url err: %s", err)
// 		return nil, err
// 	}

// 	respBodyBytes, err := c.executeRequest(ctx, method, signURL, reqBody)
// 	if err != nil {
// 		c.log.Errorf("execute request err: %s", err)
// 		return nil, err
// 	}

// 	return respBodyBytes, nil
// }

func getURLParams(fullUrl string, params map[string]string) (string, error) {

	values := url.Values{}

	URL, err := url.Parse(fullUrl)
	if err != nil {
		return "", err
	}

	for k, v := range params {
		values.Set(k, v)
	}

	URL.RawQuery = values.Encode()

	return URL.String(), nil
}
