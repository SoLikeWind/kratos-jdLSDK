package jdl

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"codeup.aliyun.com/sitao/zhinvyun/vega/internal/gateway/pkg/jdl/auth"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var (
	ProviderSet = wire.NewSet(NewClientManager)
	CST         = time.FixedZone("CST", 8*60*60)
)

type Client struct {
	config     *Config
	log        *log.Helper
	httpClient *http.Client
	token      *auth.Token // 保留，扩展用
	path       *Path
}

func NewClient(config *Config, logger log.Logger) (*Client, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "jd"))

	env := Environment(config.Env)
	paths, err := NewPath(env)
	if err != nil {
		return nil, nil, err
	}

	c := &Client{
		config: config,
		log:    log,
		httpClient: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   config.Timeout,
					KeepAlive: 0,
				}).Dial,
				DisableKeepAlives: true,
			},
		},
		token: &auth.Token{
			// AccessToken: config.AccessToken,
			// AccessExpire: time.Now().In(CST).Add(-1),
		},
		path: paths,
	}

	return c, func() {}, nil
}

// executeRequest 执行HTTP请求的通用方法
func (c *Client) executeRequest(ctx context.Context, method, signURL string, body []byte) ([]byte, error) {
	httpReq, err := http.NewRequestWithContext(ctx, method, signURL, bytes.NewBuffer(body))
	if err != nil {
		c.log.Errorf("create request err: %s", err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		c.log.Errorf("request err: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.log.Errorf("read body err: %s", err)
		return nil, err
	}

	return respBodyBytes, nil
}

// func (c *Client) getMemAccessToken(ctx context.Context) (string, error) {
// 	accessToken, tokenStatus := c.token.GetAccessToken()
// 	if tokenStatus == utils.TokenValidGtTen {
// 		return accessToken, nil
// 	}
// 	if tokenStatus == utils.TokenValidLtTen {
// 		if c.token.GetRefreshToken() != "" {
// 			c.refreshToken(ctx,
// 				c.config.AppSecret,
// 				c.config.AppKey,
// 				c.token.GetRefreshToken(),
// 			)
// 		}
// 	}

// 	return "", errors.New("memory token is expired")
// }

// func (c *Client) getAccessToken(ctx context.Context) (string, error) {
// 	accessToken, err := c.getMemAccessToken(ctx)

// 	if accessToken == "" {
// 		// 获取accesstoken到内存
// 		_, tokenStatus, err := c.ensureToken(ctx)
// 		if err != nil {
// 			return "", err
// 		}

// 		if tokenStatus == utils.TokenValid {
// 			accessToken, err = c.getMemAccessToken(ctx)
// 			if err != nil {
// 				return "", err
// 			}
// 			return accessToken, nil
// 		}

// 		if tokenStatus == utils.TokenMemoryNil {
// 			c.log.Infof("token is nil, get new token by code", tokenStatus, utils.TokenMemoryNil)
// 			accessToken, err = c.getTokenByCode(ctx)
// 			if err != nil {
// 				return "", err
// 			}

// 			return accessToken, nil
// 		}

// 	}

// 	return "", err
// }

// func (c *Client) getTokenByCode(ctx context.Context) (string, error) {
// 	code := os.Getenv("CODE")
// 	if code == "" {
// 		return "", fmt.Errorf("code is empty")
// 	}

// 	reqParams := map[string]string{
// 		"code":          code,
// 		"client_id":     c.config.AppKey,
// 		"client_secret": c.config.AppSecret,
// 	}

// 	reqBody, err := json.Marshal(reqParams)
// 	if err != nil {
// 		c.log.Errorf("marshal req err: %s", err)
// 		return "", err
// 	}

// 	respBodyBytes, err := c.executeRequest(ctx, http.MethodPost, c.path.GetAccessToken, reqBody)
// 	if err != nil {
// 		c.log.Errorf("execute request err: %s", err)
// 		return "", err
// 	}

// 	result := &GetAccessTokenResp{}
// 	err = json.Unmarshal(respBodyBytes, result) // 反序列化收到的json数据
// 	if err != nil {
// 		c.log.Errorf("unmarshal body err: %s", err)
// 		return "", err
// 	}

// 	if result.ErrCode != 0 {
// 		c.log.Errorf("jd return err: %s", result.Message)
// 		return "", fmt.Errorf("jd return err: %s", result.Message)
// 	}

// 	c.token.AccessToken = result.AccessToken
// 	c.token.AccessExpire = result.AccessExpire

// 	token := &auth.Token{
// 		AccessToken:   result.AccessToken,
// 		AccessExpire:  result.AccessExpire,
// 		RefreshToken:  result.RefreshToken,
// 		RefreshExpire: result.RefreshExpire,
// 	}

// 	if err := c.config.TokenStore.SetToken(ctx, token); err != nil { // 缓存token
// 		c.log.Errorf("set token err: %s", err)
// 		return "", err
// 	}

// 	return result.AccessToken, nil
// }

// func (c *Client) refreshToken(ctx context.Context, appSercet string, app_key string, refresh_token string) (*GetRefreshTokenResp, error) {
// 	timestamp := time.Now().In(CST).Format("2006-01-02 15:04:05")

// 	// 计算签名
// 	content := appSercet + "app_key" + app_key + "refresh_token" + refresh_token + "timestamp" + timestamp + appSercet

// 	sign := md5.Sum([]byte(content))
// 	signStr := hex.EncodeToString(sign[:])

// 	reqParams := map[string]string{
// 		"app_key":       c.config.AppKey,
// 		"refresh_token": c.token.RefreshToken,
// 		"timestamp":     timestamp,
// 		"sign":          signStr, //签名,小写十六进制
// 	}

// 	reqBody, err := json.Marshal(reqParams)
// 	if err != nil {
// 		c.log.Errorf("marshal req err: %s", err)
// 		return nil, err
// 	}

// 	signURL, err := getURLParams(c.path.RefreshToken, reqParams)
// 	if err != nil {
// 		return nil, err
// 	}

// 	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
// 	if err != nil {
// 		c.log.Errorf("execute request err: %s", err)
// 		return nil, err
// 	}

// 	result := &GetRefreshTokenResp{}
// 	err = json.Unmarshal(respBodyBytes, result)
// 	if err != nil {
// 		c.log.Errorf("unmarshal err: %s", err)
// 		return nil, err
// 	}

// 	if result.ErrCode != 0 {
// 		c.log.Errorf("jd return err: %s", result.Message)
// 		return nil, fmt.Errorf("jd return err: %s", result.Message)
// 	}

// 	c.token.AccessToken = result.AccessToken
// 	c.token.AccessExpire = result.AccessExpire
// 	c.token.RefreshToken = result.RefreshToken
// 	c.token.RefreshExpire = result.RefreshExpire

// 	token := &auth.Token{
// 		AccessToken:   result.AccessToken,
// 		AccessExpire:  result.AccessExpire,
// 		RefreshToken:  result.RefreshToken,
// 		RefreshExpire: result.RefreshExpire,
// 	}

// 	if err := c.config.TokenStore.SetToken(ctx, token); err != nil {
// 		c.log.Errorf("set token err: %s", err.Error())
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (c *Client) ensureToken(ctx context.Context) (*auth.Token, utils.TokenStatus, error) {

// 	if c.config.TokenStore != nil { // 配置了缓存
// 		token, err := c.config.TokenStore.GetToken(ctx)
// 		if err != nil {
// 			return nil, utils.TokenFailed, fmt.Errorf("get token failed: %w", err)
// 		}

// 		if token == nil {
// 			return nil, utils.TokenRedisNil, nil
// 		}

// 		c.token.AccessToken = token.AccessToken // 存入缓存
// 		c.token.AccessExpire = token.AccessExpire
// 		c.token.RefreshToken = token.RefreshToken
// 		c.token.RefreshExpire = token.RefreshExpire

// 		return token, utils.TokenValid, nil
// 	}
// 	return nil, utils.TokenFailed, fmt.Errorf("未配置缓存")
// }

// func (c *JdClient) BuildAuthURL(ctx context.Context) (string, error) {
// 	reqParams := map[string]string{
// 		"client_id":     c.config.ClientId,
// 		"redirect_uri":  c.config.RedirectURI,
// 		"response_type": "code", // 固定为 code
// 		// "state":         "",
// 	}

// 	authUrl, err := getURLParams(jdPath.getAccessToken, reqParams)
// 	if err != nil {
// 		c.log.Errorf("get url params err: %s", err)
// 		return "", err
// 	}

// 	return authUrl, err //重定向至xxx/xxx？code=xxx
// }
