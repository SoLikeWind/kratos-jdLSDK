package jdl

import (
	"time"
)

type Option func(*Config)

type Config struct {
	AppKey       string
	AppSecret    string
	AccessToken  string
	Env          string
	V            string
	CustomerCode string // 客户编码
	// TokenStore  store.TokenStore
	Timeout time.Duration
}

func NewConfig(options ...Option) *Config {
	c := &Config{
		Env: "develop",
		V:   "2.0", // 默认2.0
	}
	for _, opt := range options {
		opt(c)
	}

	return c
}

func WithAppKey(appKey string) Option {
	return func(c *Config) {
		c.AppKey = appKey
	}
}

func WithAppSecret(appSecret string) Option {
	return func(c *Config) {
		c.AppSecret = appSecret
	}
}

func WithAccessToken(accessToken string) Option {
	return func(c *Config) {
		c.AccessToken = accessToken
	}
}

func WithEnv(env string) Option {
	return func(c *Config) {
		c.Env = env
	}
}

func WithV(V string) Option {
	return func(c *Config) {
		c.V = V
	}
}

func WithCustomerCode(customerCode string) Option {
	return func(c *Config) {
		c.CustomerCode = customerCode
	}
}

// func WithTokenStore(tokenStore store.TokenStore) Option {
// 	return func(c *Config) {
// 		c.TokenStore = tokenStore
// 	}
// }

func WithTimeout(timeout int64) Option {
	return func(c *Config) {
		c.Timeout = time.Duration(timeout) * time.Second
	}
}
