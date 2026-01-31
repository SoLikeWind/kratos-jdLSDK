package jdl

import (
	"errors"
)

// 下单前置校验请求
type PreCheckReq struct {
	SenderContact   *PlaceContact `json:"senderContact"`
	ReceiverContact *PlaceContact `json:"receiverContact"`
	OrderOrigin     int           `json:"orderOrigin"`  // 下单来源B2C--1
	CustomerCode    string        `json:"customerCode"` // 客户编码
	ProductsReq     *ProductInfo  `json:"productsReq"`
	Cargoes         []*CargoInfo  `json:"cargoes"`
}

// PlaceOrderReq 下单请求
type PlaceOrderReq struct {
	OrderID                 string             `json:"orderId"`
	SenderContact           *PlaceContact      `json:"senderContact"`
	ReceiverContact         *PlaceContact      `json:"receiverContact"`
	OrderOrigin             int                `json:"orderOrigin"`  // 订单来源 1-b2c
	CustomerCode            string             `json:"customerCode"` // 客户编码
	ProductsReq             *ProductInfo       `json:"productsReq"`
	SettleType              int                `json:"settleType"`                        // 结算方式 1-寄付；2-到付；3-月结；
	Cargoes                 []*CargoInfo       `json:"cargoes"`                           // 货物信息
	CommonChannelInfo       *CommonChannelInfo `json:"commonChannelInfo"`                 // 渠道信息
	PickupStartTime         *int64             `json:"pickupStartTime,omitempty"`         // 期望揽收开始时间（可选），毫秒
	PickupEndTime           *int64             `json:"pickupEndTime,omitempty"`           // 期望揽收结束时间（可选），毫秒
	ExpectDeliveryStartTime *int64             `json:"expectDeliveryStartTime,omitempty"` // 期望配送开始时间（可选），毫秒
	ExpectDeliveryEndTime   *int64             `json:"expectDeliveryEndTime,omitempty"`   // 期望配送结束时间（可选），毫秒
}

// 寄/收件人
type PlaceContact struct {
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	FullAddress string `json:"fullAddress"`
}

// 产品信息
type ProductInfo struct {
	ProductCode string `json:"productCode"` //电商标快	ed-m-0059
}

//	type CargoInfo struct {
//		Object string `json:"object"`
//	}
//

// 货物信息
type CargoInfo struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"` // 件数
	Weight   float64 `json:"weight"`   // 重量，单位：kg
	Volume   float64 `json:"volume"`   // 体积，单位：cm³，必须>0
}

// 渠道信息
type CommonChannelInfo struct {
	ChannelCode string `json:"channelCode"` // 0030001-其他平台
}

// 修改订单请求(参数结构同下单接口一致)
type UpdateOrderReq struct {
	OrderID         *string        `json:"orderId,omitempty"`
	WaybillCode     *string        `json:"waybillCode,omitempty"` // jdl运单号(同jdl订单号必填其一)
	OrderCode       *string        `json:"orderCode,omitempty"`   // jdl订单号
	SenderContact   *UpdateContact `json:"senderContact,omitempty"`
	ReceiverContact *UpdateContact `json:"receiverContact,omitempty"`
	OrderOrigin     string         `json:"orderOrigin"`  // 订单来源 1-b2c
	CustomerCode    string         `json:"customerCode"` // 客户编码(1-b2c-必填)
	ProductsReq     *ProductInfo   `json:"productsReq,omitempty"`
	PickupStartTime *int64         `json:"pickupStartTime,omitempty"` //期望揽收开始时间,需大于下单时间小于预约揽收结束时间；毫秒级
	PickupEndTime   *int64         `json:"pickupEndTime,omitempty"`
}

// 修改收/寄件人
type UpdateContact struct {
	Name        *string `json:"name,omitempty"`
	Mobile      *string `json:"mobile,omitempty"`
	FullAddress *string `json:"fullAddress,omitempty"`
}

func (r *UpdateOrderReq) Validate() error {
	if r.WaybillCode == nil && r.OrderCode == nil {
		return errors.New("waybillCode or orderId is required")
	}

	return nil
}

// 取消订单请求
type CancelOrderReq struct {
	WaybillCode        *string `json:"waybillCode,omitempty"`     // jdl运单号(同jdl订单号必填其一)
	OrderCode          *string `json:"orderCode,omitempty"`       // jdl订单号
	CustomerOrderId    *string `json:"customerOrderId,omitempty"` // 商家订单号
	OrderOrigin        int64   `json:"orderOrigin"`               // 订单来源 1-b2c
	CustomerCode       string  `json:"customerCode"`              // 客户编码
	CancelReason       string  `json:"cancelReason"`
	CancelReasonCode   string  `json:"cancelReasonCode"`             // 1 客户取消 2 超时未支付
	CancelType         int64   `json:"cancelType"`                   // 1
	SubscribeIntercept *string `json:"subscribeIntercept,omitempty"` // 1
}

func (r *CancelOrderReq) Validate() error {
	if r.WaybillCode == nil && r.OrderCode == nil {
		return errors.New("waybillCode or orderId is required")
	}

	return nil
}

type GetOrderStatusReq struct {
	WaybillCode  *string `json:"waybillCode,omitempty"`
	OrderCode    *string `json:"orderCode,omitempty"`
	OrderOrigin  string  `json:"orderOrigin"`
	CustomerCode string  `json:"customerCode"`
}

func (r *GetOrderStatusReq) Validate() error {
	if r.WaybillCode == nil && r.OrderCode == nil {
		return errors.New("waybillCode or orderId is required")
	}

	return nil
}

type GetOrderTrackReq struct {
	ReferenceNumber string  `json:"referenceNumber"`
	ReferenceType   string  `json:"referenceType"` // 单号类型：物流运单	20000 商家订单号 40000
	Phone           *string `json:"phone,omitempty"`
	CustomerCode    *string `json:"customerCode,omitempty"`
	Scope           *string `json:"scope,omitempty"`
}

func (r *GetOrderTrackReq) Validate() error {
	if r.Phone == nil && r.CustomerCode == nil {
		return errors.New("phone or customerCode is required")
	}

	return nil
}

type PrintReq struct {
	CustomerCode string          `json:"customerCode"`
	TemplateCode string          `json:"templateCode"`
	Operator     string          `json:"operator"`
	TaskID       string          `json:"taskId"`
	PrintData    []*PrintData    `json:"printData"`
	OutputConfig []*OutputConfig `json:"outputConfig"`
}

type PrintData struct {
	OrderNumber   string `json:"orderNumber"`
	CarrierCode   string `json:"carrierCode"`
	BillCodeValue string `json:"billCodeValue"`
	BillCodeType  string `json:"billCodeType"`
	Scene         int    `json:"scene"`
}

type OutputConfig struct {
	DataFormat int `json:"dataFormat"`
	OutputType int `json:"outputType"`
	FileFormat int `json:"fileFormat"`
}

type PrintResp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result []struct {
		DataFormat int64  `json:"dataFormat"`
		FileFormat int64  `json:"fileFormat"`
		URL        string `json:"url"`
		Base64     string `json:"base64"`
	} `json:"result"`
}

type PullDataReq struct {
	CpCode       string                 `json:"cpCode"`
	ObjectID     string                 `json:"objectId"`
	WayBillInfos []*PullDataWayBillInfo `json:"wayBillInfos"`
	Parameters   map[string]string      `json:"parameters"`
}

type PullDataWayBillInfo struct {
	JdWayBillCode string `json:"jdWayBillCode"`
	PopFlag       int    `json:"popFlag"`
}

type TrackQueryReq struct {
	ReferenceNumber  string `json:"referenceNumber"`
	ReferenceType    int    `json:"referenceType"`
	Phone            string `json:"phone"`
	CustomerCode     string `json:"customerCode"`
	BusinessUnitCode string `json:"businessUnitCode"`
	Scope            string `json:"scope"`
}
