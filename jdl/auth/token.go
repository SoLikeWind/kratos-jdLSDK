package auth

import (
	"time"

	"codeup.aliyun.com/sitao/zhinvyun/vega/internal/gateway/pkg/jdl/utils"
)

var CST = time.FixedZone("CST", 8*3600) // 东八区，北京时间

type Token struct {
	AccessToken   string
	AccessExpire  time.Time
	RefreshToken  string
	RefreshExpire time.Time
}

func (t *Token) GetAccessToken() (string, utils.TokenStatus) {
	if t.AccessExpire.IsZero() {
		return "", utils.TokenMemoryNil
	}

	now := time.Now().In(CST)
	tenDays := 10 * 24 * time.Hour

	if now.Before(t.AccessExpire) && t.AccessExpire.Sub(now) < tenDays {
		return t.AccessToken, utils.TokenValidLtTen
	} else if now.Before(t.AccessExpire) && t.AccessExpire.Sub(now) > tenDays {
		return t.AccessToken, utils.TokenValidGtTen
	}
	return "", utils.TokenMemoryNil
}

func (t *Token) GetRefreshToken() string {
	if time.Now().Before(t.RefreshExpire) {
		return t.RefreshToken
	}
	return ""
}

// true为小于十天且未过期
func (t *Token) IsLtTenDaysAndValid() bool {
	if t.AccessExpire.IsZero() {
		return false
	}

	now := time.Now().In(CST)
	tenDays := 10 * 24 * time.Hour

	return now.Before(t.AccessExpire) && t.AccessExpire.Sub(now) < tenDays
}

func (t *Token) IsAccessExpired() bool {
	return time.Now().After(t.AccessExpire)
}

func (t *Token) IsRefreshExpired() bool {
	return time.Now().After(t.RefreshExpire)
}
