package jdl

import (
	"time"
)

type BaseResp struct {
	ErrCode     int    `json:"code"`
	Message     string `json:"msg"`
	Success     bool   `json:"success"`
	RequestId   string `json:"requestId"`
	RawResponse string `json:"-"`
}

type GetAccessTokenResp struct {
	AccessToken   string    `json:"accessToken"`
	AccessExpire  time.Time `json:"accessExpire"`
	RefreshToken  string    `json:"refreshToken"`
	RefreshExpire time.Time `json:"refreshExpire"`
	SellerId      string    `json:"sellerId"`
	// Code          string    `json:"code"`
	BaseResp
}

type GetRefreshTokenResp struct {
	AccessToken   string    `json:"accessToken"`
	AccessExpire  time.Time `json:"accessExpire"`
	RefreshToken  string    `json:"refreshToken"`
	RefreshExpire time.Time `json:"refreshExpire"`
	SellerId      string    `json:"sellerId"`
	BaseResp
}

// 下单前置校验响应
type PreCheckResp struct {
	BaseResp
}

// 下单响应
type PlaceOrderResp struct {
	Data PlaceOrderData `json:"data"`
	BaseResp
}

// 下单数据
type PlaceOrderData struct {
	OrderCode           string            `json:"orderId"`
	WaybillCode         string            `json:"waybillCode"`
	PickupPromiseTime   string            `json:"pickupPromiseTime"`
	DeliveryPromiseTime string            `json:"deliveryPromiseTime"`
	FaceSheetResponse   FaceSheetResp     `json:"faceSheetResponse"`
	ExtendFields        map[string]string `json:"extendFields"` //  key说明:verificationCodeCollect:签收码
}

// 面单数据
type FaceSheetResp struct {
	FaceSheetInfo FaceSheetInfo `json:"faceSheetInfo"`
}

type FaceSheetInfo struct {
	StartSiteName        string `json:"startSiteName"`
	EndSiteName          string `json:"endSiteName"`
	Road                 string `json:"road"`
	SourceSortCenterName string `json:"sourceSortCenterName"`
	TargetSortCenterName string `json:"targetSortCenterName"`
	Aging                int    `json:"aging"` // 时效
	AgingName            string `json:"agingName"`
	OrderId              string `json:"orderId"` // 商家订单号
}

// 修改订单响应
type UpdateOrderResp struct {
	Data UpdateOrdeData `json:"data"`
	BaseResp
}

type UpdateOrdeData struct {
	OrderCode   string `json:"orderCode"`
	WaybillCode string `json:"waybillCode"`
}

// 取消订单响应
type CancelOrderResp struct {
	Data CancelOrderData `json:"data"`
	BaseResp
}

type CancelOrderData struct {
	OrderCode   string `json:"orderCode"`
	WaybillCode string `json:"waybillCode"`
	ResultType  int    `json:"resultType"` // 0 取消成功；1 拦截成功 2 取消失败；3 拦截失败，4拦截中
}

// 获取订单状态响应
type GetOrderStatusResp struct {
	Data GetOrderStatusData `json:"data"`
	BaseResp
}

type GetOrderStatusData struct {
	Status     string `json:"status"`
	StatusDesc string `json:"statusDesc"`
}

// 获取订单轨迹响应
type GetOrderTrackResp struct {
	Data []GetOrderTrackData `json:"data"`
	BaseResp
}

type GetOrderTrackData struct {
	Trackings []Tracking `json:"trackings"`
}

type Tracking struct {
	Remark          string          `json:"remark"`
	OperatorName    string          `json:"operatorName"`
	OperationTime   string          `json:"operationTime"`
	OperationType   string          `json:"operationType"`
	OperationCOde   string          `json:"operationCode"`
	OperateSite     string          `json:"operateSite"`
	OperateLocation OperateLocation `json:"operateLocation"`
	Extend          TrackingExtend  `json:"extend"`
	NextLocation    NextLocation    `json:"nextLocation"`
}

type OperateLocation struct {
	RouteProvinceName string `json:"routeProvinceName"`
	RouteCityName     string `json:"routeCityName"`
	RouteDistrictName string `json:"routeDistrictName"`
	RouteAddress      string `json:"routeAddress"`
}

type TrackingExtend struct {
	OperatorPhone string `json:"operatorPhone"`
}

type NextLocation struct {
	NextProvinceName string `json:"nextProvinceName"`
	NextCityName     string `json:"nextCityName"`
	NextDistrictName string `json:"nextDistrictName"`
	NextAddress      string `json:"nextAddress"`
}
