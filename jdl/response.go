package jdl

import (
	"time"
)

type BaseResp struct {
	ErrCode       int        `json:"code"`
	Message       string     `json:"msg"`
	Success       bool       `json:"success"`
	RequestId     string     `json:"requestId"`
	Mills         int        `json:"mills,omitempty"`          // 请求耗时（毫秒）
	ErrorResponse *ErrorResp `json:"error_response,omitempty"` // 错误响应
}

// ErrorResp 京东物流错误响应格式
type ErrorResp struct {
	Code   int    `json:"code"`
	ZhDesc string `json:"zh_desc"`
	EnDesc string `json:"en_desc"`
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
	Data PreCheckData `json:"data"`
}

type PreCheckData struct {
	PickupSliceTimes []struct {
		DateKey          string `json:"dateKey"`
		PickupSliceTimes []struct {
			EndTime   string `json:"endTime"`
			StartTime string `json:"startTime"`
		} `json:"pickupSliceTimes"`
		TimeInterval int `json:"timeInterval"`
	} `json:"pickupSliceTimes"`
	ProductInfos []struct {
		AddedProducts []struct {
			FeeType         string  `json:"feeType,omitempty"`
			FeeTypeName     string  `json:"feeTypeName,omitempty"`
			FreightPre      float64 `json:"freightPre,omitempty"`
			FreightStandard float64 `json:"freightStandard,omitempty"`
			ProductCode     string  `json:"productCode"`
			ProductName     string  `json:"productName"`
		} `json:"addedProducts"`
		DeliveryPromiseTime int64   `json:"deliveryPromiseTime"`
		FeeType             string  `json:"feeType"`
		FeeTypeName         string  `json:"feeTypeName"`
		FreightPre          float64 `json:"freightPre"`
		FreightStandard     float64 `json:"freightStandard"`
		ProductCode         string  `json:"productCode"`
		ProductName         string  `json:"productName"`
	} `json:"productInfos"`
	ShipmentInfo struct {
		EndStationName   string `json:"endStationName"`
		EndStationNo     int    `json:"endStationNo"`
		StartStationName string `json:"startStationName"`
		StartStationNo   int    `json:"startStationNo"`
	} `json:"shipmentInfo"`
	CommonFeeInfoResponse struct {
		CalWeight                 float64 `json:"calWeight"`
		Currency                  string  `json:"currency"`
		ExpDate                   int64   `json:"expDate"`
		FeeDetailInfoResponseList []struct {
			Amount      float64 `json:"amount"`
			CostName    string  `json:"costName"`
			CostNo      string  `json:"costNo"`
			Currency    string  `json:"currency"`
			PreAmount   float64 `json:"preAmount"`
			ProductCode string  `json:"productCode"`
			ProductName string  `json:"productName"`
		} `json:"feeDetailInfoResponseList"`
		TotalAmount    float64 `json:"totalAmount"`
		TotalPreAmount float64 `json:"totalPreAmount"`
	} `json:"commonFeeInfoResponse"`
	TotalFreightPre      float64 `json:"totalFreightPre"`
	TotalFreightStandard float64 `json:"totalFreightStandard"`
}

// 下单响应
type PlaceOrderResp struct {
	Data PlaceOrderData `json:"data"`
	BaseResp
}

// 下单数据
type PlaceOrderData struct {
	OrderCode             string             `json:"orderCode"`                       // 京东订单号
	WaybillCode           string             `json:"waybillCode"`                     // 运单号
	PickupPromiseTime     int64              `json:"pickupPromiseTime"`               // 承诺揽收时间（时间戳）
	DeliveryPromiseTime   int64              `json:"deliveryPromiseTime"`             // 承诺送达时间（时间戳）
	FreightPre            float64            `json:"freightPre"`                      // 预估运费
	NeedRetry             bool               `json:"needRetry"`                       // 是否需要重试
	FaceSheetResponse     *FaceSheetResp     `json:"faceSheetResponse,omitempty"`     // 面单信息
	ExtendFields          map[string]string  `json:"extendFields,omitempty"`          // 扩展字段
	CommonFeeInfoResponse *CommonFeeInfoResp `json:"commonFeeInfoResponse,omitempty"` // 费用信息
	PresortResult         *PresortResult     `json:"presortResult,omitempty"`         // 预分拣结果
	ProductsResponse      *ProductsResponse  `json:"productsResponse,omitempty"`      // 产品响应
}

// 面单响应
type FaceSheetResp struct {
	FaceSheetInfo FaceSheetInfo `json:"faceSheetInfo"`
}

// 面单信息
type FaceSheetInfo struct {
	StartSiteId             int    `json:"startSiteId"`             // 始发站点ID
	StartSiteName           string `json:"startSiteName"`           // 始发站点名称
	EndSiteId               int    `json:"endSiteId"`               // 目的站点ID
	EndSiteName             string `json:"endSiteName"`             // 目的站点名称
	Road                    string `json:"road"`                    // 道口
	DeliveryTransferPointId int    `json:"deliveryTransferPointId"` // 配送中转点ID
}

// 费用信息响应
type CommonFeeInfoResp struct {
	CalWeight                 float64         `json:"calWeight"`                 // 计费重量
	Currency                  string          `json:"currency"`                  // 币种
	ExpDate                   int64           `json:"expDate"`                   // 过期时间（时间戳）
	TotalAmount               float64         `json:"totalAmount"`               // 总费用
	TotalPreAmount            float64         `json:"totalPreAmount"`            // 预估总费用
	FeeDetailInfoResponseList []FeeDetailInfo `json:"feeDetailInfoResponseList"` // 费用明细列表
}

// 费用明细信息
type FeeDetailInfo struct {
	Amount      float64 `json:"amount"`      // 费用金额
	PreAmount   float64 `json:"preAmount"`   // 预估费用金额
	CostName    string  `json:"costName"`    // 费用名称
	CostNo      string  `json:"costNo"`      // 费用编号
	Currency    string  `json:"currency"`    // 币种
	ProductCode string  `json:"productCode"` // 产品编码
	ProductName string  `json:"productName"` // 产品名称
}

// 预分拣结果
type PresortResult struct {
	StartStationNo                string `json:"startStationNo"`                // 始发站点编号
	StartStationName              string `json:"startStationName"`              // 始发站点名称
	StartStationCode              string `json:"startStationCode"`              // 始发站点代码
	StartStationType              string `json:"startStationType"`              // 始发站点类型
	StartRoadArea                 string `json:"startRoadArea"`                 // 始发道口区域
	StartDeliveryType             string `json:"startDeliveryType"`             // 始发配送类型
	StartAoiCode                  string `json:"startAoiCode"`                  // 始发AOI编码
	StartStationPresortResultType string `json:"startStationPresortResultType"` // 始发站点预分拣结果类型
	EndStationNo                  string `json:"endStationNo"`                  // 目的站点编号
	EndStationName                string `json:"endStationName"`                // 目的站点名称
	EndStationCode                string `json:"endStationCode"`                // 目的站点代码
	EndStationType                string `json:"endStationType"`                // 目的站点类型
	EndRoadArea                   string `json:"endRoadArea"`                   // 目的道口区域
	EndDeliveryType               string `json:"endDeliveryType"`               // 目的配送类型
	EndAoiCode                    string `json:"endAoiCode"`                    // 目的AOI编码
	EndStationPresortResultType   string `json:"endStationPresortResultType"`   // 目的站点预分拣结果类型
	EndTransferStationNo          string `json:"endTransferStationNo"`          // 目的中转站编号
}

// 产品响应
type ProductsResponse struct {
	ProductCode       string            `json:"productCode"`       // 产品编码
	ProductModifyFlag int               `json:"productModifyFlag"` // 产品修改标识
	ProductAttrs      map[string]string `json:"productAttrs"`      // 产品属性
	AddedProducts     []AddedProduct    `json:"addedProducts"`     // 附加产品列表
}

// 附加产品
type AddedProduct struct {
	ProductCode       string            `json:"productCode"`       // 产品编码
	ProductModifyFlag int               `json:"productModifyFlag"` // 产品修改标识
	ProductAttrs      map[string]string `json:"productAttrs"`      // 产品属性
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

type PullDataResp struct {
	Code          string `json:"code"`
	Message       string `json:"message"`
	ObjectID      string `json:"objectId"`
	PrePrintDatas []struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		PerPrintData string `json:"perPrintData"`
		WayBillNo    string `json:"wayBillNo"`
	} `json:"prePrintDatas"`
}

type TrackQueryResp struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		ReferenceNumber string  `json:"referenceNumber"`
		ReferenceType   float64 `json:"referenceType"`
		Trackings       []struct {
			OperationCode   string `json:"operationCode"`
			Remark          string `json:"remark"`
			OperationType   string `json:"operationType"`
			OperatorName    string `json:"operatorName"`
			OperationTime   string `json:"operationTime"`
			WaybillNo       string `json:"waybillNo"`
			OperateSiteID   string `json:"operateSiteId"`
			OperateSite     string `json:"operateSite"`
			OperateLocation struct {
				RouteProvinceName string `json:"routeProvinceName"`
				RouteAddress      string `json:"routeAddress"`
				RouteCityName     string `json:"routeCityName"`
				RouteDistrictName string `json:"routeDistrictName"`
				Lon               string `json:"lon"`
				Lat               string `json:"lat"`
			} `json:"operateLocation"`
			Extend       map[string]string `json:"extend"`
			NextLocation struct {
				Lat              string `json:"lat"`
				Lon              string `json:"lon"`
				NextAddress      string `json:"nextAddress"`
				NextCityName     string `json:"nextCityName"`
				NextDistrictName string `json:"nextDistrictName"`
				NextProvinceName string `json:"nextProvinceName"`
				NextStreetName   string `json:"nextStreetName"`
			} `json:"nextLocation"`
		} `json:"trackings"`
	} `json:"data"`
}
