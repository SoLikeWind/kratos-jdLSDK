package jdl

import (
	"errors"
)

// 下单前置校验请求
type PreCheckReq struct {
	SenderContact   PlaceContact `json:"senderContact"`
	ReceiverContact PlaceContact `json:"receiverContact"`
	OrderOrigin     int          `json:"orderOrigin"` // 下单来源B2C--1
	CustomerCode    string       `json:"customerCode"`
}

// PlaceOrderReq 下单请求
type PlaceOrderReq struct {
	OrderID           string            `json:"orderId"`
	SenderContact     PlaceContact      `json:"senderContact"`
	ReceiverContact   PlaceContact      `json:"receiverContact"`
	OrderOrigin       int               `json:"orderOrigin"`  // 订单来源 1-b2c
	CustomerCode      string            `json:"customerCode"` // 客户编码
	ProductsReq       ProductInfo       `json:"productsReq"`
	SettleType        int               `json:"settleType"` // 结算方式 1-寄付；2-到付；3-月结；
	Cargoes           []CargoInfo       `json:"cargoes"`    // 货物信息
	CommonChannelInfo CommonChannelInfo `json:"commonChannelInfo"`
	// PickupStartTime int64        `json:"pickupStartTime"` // 期望揽收开始时间
	// PickupEndTime   int64        `json:"pickupEndTime"`   // 期望揽收结束时间
	// ExpectDeliveryStartTime int64        `json:"expectDeliveryStartTime"` // 期望配送开始时间
	// ExpectDeliveryEndTime   int64        `json:"expectDeliveryEndTime"`   // 期望配送结束时间
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

// // 货物信息
// type CargoInfo struct {
// 	Object Object `json:"object"`
// }

type CargoInfo struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"` // 件数
	Weight   int    `json:"weight"`   // 重量，单位：kg
	Volume   int    `json:"volume"`   // 体积，单位：cm³，必须>0
}

type CommonChannelInfo struct {
	ChannelCode string `json:"channelCode"`
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
	ProductsReq     *ProductInfo   `json:"products,omitempty"`
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
	OrderOrigin        string  `json:"orderOrigin"`               // 订单来源 1-b2c
	CustomerCode       string  `json:"customerCode"`              // 客户编码
	CancelReason       string  `json:"cancelReason"`
	CancelReasonCode   string  `json:"cancelReasonCode"`             // 1 客户取消 2 超时未支付
	CancelType         int     `json:"cancelType"`                   // 1
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
}

func (r *GetOrderTrackReq) Validate() error {
	if r.Phone == nil && r.CustomerCode == nil {
		return errors.New("phone or customerCode is required")
	}

	return nil
}
