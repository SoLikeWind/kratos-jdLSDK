package jdl

import "fmt"

type Environment string

const (
	EnvironmentRelease Environment = "release"
	EnvironmentDevelop Environment = "develop"
)

var BaseURLs = map[Environment]string{
	EnvironmentRelease: "https://oauth.jdl.com",
	EnvironmentDevelop: "https://uat-oauth.jdl.com",
}

const (
	GetAuthCodePath    = "/oauth/authorize"
	GetAccessTokenPath = "/oauth/token"
	RefreshTokenPath   = "/oauth/refresh"
	PreCheckPath       = "/ecap/v1/orders/precheck" // 下单前置校验
	// GetWaybillCodePath    = "/ecap/v1/orders/pregetwaybillcodes"   // 预售情况下提前生成快递订单号
	// PickUpOrderPath       = "/ecap/v1/deliverypickuporder/create"  // 送取同步下单
	PlaceOrderPath = "/ecap/v1/orders/create" // 普通下单
	// UnbindPickUpOrderPath = "/ecap/v1/deliverypickuporder/unbound" // 解绑送取同步下单
	UpdateOrderPath    = "/ecap/v1/orders/modify"     // 修改订单
	CancelOrderPath    = "/ecap/v1/orders/cancel"     // 取消订单
	GetOrderStatusPath = "/ecap/v1/orders/status/get" // 订单状态
	GetOrderTrackPath  = "/jd/tracking/query"         // 运单轨迹

	PrintPath    = "/cloud/print/render"       // 打印
	PullDataPath = "/PullDataService/pullData" // 获取打印数据

	TrackQueryPath = "/jd/tracking/query" // 运单轨迹
)

type Path struct {
	GetAuthCode    string
	GetAccessToken string
	RefreshToken   string
	PreCheck       string
	// GetWaybillCode string
	// PickUpOrder    string
	PlaceOrder     string
	UpdateOrder    string
	CancelOrder    string
	GetOrderStatus string
	GetOrderTrack  string
	Print          string
	PullData       string
	TrackQuery     string
}

func NewPath(env Environment) (*Path, error) {
	baseUrl, exists := BaseURLs[env]
	if !exists {
		return nil, fmt.Errorf("unsupported environment: %s", env)
	}

	return &Path{
		GetAuthCode:    baseUrl + GetAuthCodePath,
		GetAccessToken: baseUrl + GetAccessTokenPath,
		RefreshToken:   baseUrl + RefreshTokenPath,
		PreCheck:       baseUrl + PreCheckPath,
		// GetWaybillCode: baseUrl + GetWaybillCodePath,
		// PickUpOrder:    baseUrl + PickUpOrderPath,
		PlaceOrder:     baseUrl + PlaceOrderPath,
		UpdateOrder:    baseUrl + UpdateOrderPath,
		CancelOrder:    baseUrl + CancelOrderPath,
		GetOrderStatus: baseUrl + GetOrderStatusPath,
		GetOrderTrack:  baseUrl + GetOrderTrackPath,
		Print:          baseUrl + PrintPath,
		PullData:       baseUrl + PullDataPath,
		TrackQuery:     baseUrl + TrackQueryPath,
	}, nil
}
