package jdl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

const (
	ErrCodeParseFailed = -999
)

func (c *JdLClient) Precheck(ctx context.Context, req *PreCheckReq) (*PreCheckResp, error) {
	reqParams := &PreCheckReq{
		SenderContact:   req.SenderContact,
		ReceiverContact: req.ReceiverContact,
		OrderOrigin:     req.OrderOrigin,
	}

	reqBody, err := json.Marshal([]PreCheckReq{*reqParams})
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}
	log.Print(string(reqBody))

	signURL, err := c.buildSignURL(c.path.PreCheck, PreCheckPath, string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &PreCheckResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err == nil { // 解析成功
		if result.ErrCode != 0 {
			c.log.Errorf("Precheck jd err: (%d) %s", result.ErrCode, result.Message)
			return nil, fmt.Errorf("jd err: %s", result.Message)
		}

		return result, err
	}

	result = &PreCheckResp{ // 解析失败
		BaseResp: BaseResp{
			ErrCode:     ErrCodeParseFailed,
			Message:     "response parse failed",
			Success:     false,
			RawResponse: string(respBodyBytes),
		},
	}

	return result, nil
}

func (c *JdLClient) GetWaybillCode(ctx context.Context, req string) (string, error) {
	return "", nil
}

func (c *JdLClient) PlaceOrder(ctx context.Context, req *PlaceOrderReq) (*PlaceOrderResp, error) {
	reqParams := &PlaceOrderReq{
		OrderID:           req.OrderID,
		SenderContact:     req.SenderContact,
		ReceiverContact:   req.ReceiverContact,
		OrderOrigin:       req.OrderOrigin,
		CustomerCode:      req.CustomerCode,
		ProductsReq:       req.ProductsReq,
		SettleType:        req.SettleType,
		Cargoes:           req.Cargoes,
		CommonChannelInfo: req.CommonChannelInfo,
		// PickupStartTime: req.PickupStartTime,
		// PickupEndTime:   req.PickupEndTime,
		// ExpectDeliveryStartTime: req.ExpectDeliveryStartTime,
		// ExpectDeliveryEndTime:   req.ExpectDeliveryEndTime,
	}

	reqBody, err := json.Marshal([]PlaceOrderReq{*reqParams})
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}
	log.Print(string(reqBody))

	signURL, err := c.buildSignURL(c.path.PlaceOrder, PlaceOrderPath, string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}
	log.Print(string(respBodyBytes))

	result := &PlaceOrderResp{}

	err = json.Unmarshal(respBodyBytes, result)
	log.Print(result)
	if err == nil {
		if result.ErrCode != 0 {
			c.log.Errorf("PlaceOrder jdl return err: (%d) %s", result.ErrCode, result.Message)
			return nil, fmt.Errorf("jdl return err: %s", result.Message)
		}

		return result, err
	}

	result = &PlaceOrderResp{ // 解析失败
		BaseResp: BaseResp{
			ErrCode:     ErrCodeParseFailed,
			Message:     "response parse failed",
			Success:     false,
			RawResponse: string(respBodyBytes),
		},
	}
	return result, nil
}

func (c *JdLClient) UpdateOrder(ctx context.Context, req *UpdateOrderReq) (*UpdateOrderResp, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	reqParams := &UpdateOrderReq{
		OrderID:         req.OrderID,
		WaybillCode:     req.WaybillCode,
		OrderCode:       req.OrderCode,
		SenderContact:   req.SenderContact,
		ReceiverContact: req.ReceiverContact,
		OrderOrigin:     req.OrderOrigin,
		ProductsReq:     req.ProductsReq,
	}

	reqBody, err := json.Marshal(reqParams)
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.UpdateOrder, UpdateOrderPath, string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &UpdateOrderResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err == nil {
		if result.ErrCode != 0 {
			c.log.Errorf("updateorder jdl return err: (%d) %s", result.ErrCode, result.Message)
			return nil, fmt.Errorf("jdl return err: %s", result.Message)
		}

		return result, err
	}

	result = &UpdateOrderResp{ // 解析失败
		BaseResp: BaseResp{
			ErrCode:     ErrCodeParseFailed,
			Message:     "response parse failed",
			Success:     false,
			RawResponse: string(respBodyBytes),
		},
	}

	return result, nil
}

// func (c *JdLClient) CancelOrder(ctx context.Context, req *CancelOrderReq) (*CancelOrderResp, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	reqParams := &CancelOrderReq{
// 		WaybillCode:        req.WaybillCode,
// 		OrderCode:          req.OrderCode,
// 		CustomerOrderId:    req.CustomerOrderId,
// 		OrderOrigin:        req.OrderOrigin,
// 		CustomerCode:       req.CustomerCode,
// 		CancelReason:       req.CancelReason,
// 		CancelReasonCode:   req.CancelReasonCode,
// 		CancelType:         req.CancelType,
// 		SubscribeIntercept: req.SubscribeIntercept,
// 	}

// 	respBodyBytes, err := c.doRequest(ctx, "POST", c.path.CancelOrder, CancelOrderPath, reqParams)
// 	if err != nil {
// 		c.log.Errorf("cancelorder doRequest err: %s", err)
// 		return nil, err
// 	}

// 	result := &CancelOrderResp{}

// 	err = json.Unmarshal(respBodyBytes, result)
// 	if err == nil {
// 		if result.ErrCode != 0 {
// 			c.log.Errorf("cancelorder jdl return err: (%d) %s", result.ErrCode, result.Message)
// 			return nil, fmt.Errorf("jdl return err: %s", result.Message)
// 		}

// 		return result, err
// 	}

// 	result = &CancelOrderResp{ // 解析失败
// 		BaseResp: BaseResp{
// 			ErrCode:     ErrCodeParseFailed,
// 			Message:     "response parse failed",
// 			Success:     false,
// 			RawResponse: string(respBodyBytes),
// 		},
// 	}

// 	return result, nil
// }

// func (c *JdLClient) GetOrderStatus(ctx context.Context, req *GetOrderStatusReq) (*GetOrderStatusResp, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	reqParams := &GetOrderStatusReq{
// 		WaybillCode:  req.WaybillCode,
// 		OrderCode:    req.OrderCode,
// 		OrderOrigin:  req.OrderOrigin,
// 		CustomerCode: req.CustomerCode,
// 	}

// 	resBodyBytes, err := c.doRequest(ctx, "POST", c.path.GetOrderStatus, GetOrderStatusPath, reqParams)
// 	if err != nil {
// 		c.log.Errorf("execute request err: %s", err)
// 		return nil, err
// 	}

// 	result := &GetOrderStatusResp{}

// 	err = json.Unmarshal(resBodyBytes, result)
// 	if err == nil {
// 		if result.ErrCode != 0 {
// 			c.log.Errorf("getorderstatus jdl return err: (%d) %s", result.ErrCode, result.Message)
// 			return nil, fmt.Errorf("jdl return err: %s", result.Message)
// 		}
// 	}

// 	result = &GetOrderStatusResp{ // 解析失败
// 		BaseResp: BaseResp{
// 			ErrCode:     ErrCodeParseFailed,
// 			Message:     "response parse failed",
// 			Success:     false,
// 			RawResponse: string(resBodyBytes),
// 		},
// 	}

// 	return result, nil
// }

// func (c *JdLClient) GetOrderTrack(ctx context.Context, req *GetOrderTrackReq) (*GetOrderTrackResp, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	reqParams := &GetOrderTrackReq{
// 		ReferenceNumber: req.ReferenceNumber,
// 		ReferenceType:   req.ReferenceType,
// 		Phone:           req.Phone,
// 		CustomerCode:    req.CustomerCode,
// 	}

// 	respBodyBytes, err := c.doRequest(ctx, "POST", c.path.GetOrderTrack, GetOrderTrackPath, reqParams)
// 	if err != nil {
// 		c.log.Errorf("GetOrderTrack doRequest err: %s", err)
// 		return nil, err
// 	}

// 	result := &GetOrderTrackResp{}

// 	err = json.Unmarshal(respBodyBytes, result)
// 	if err == nil {
// 		if result.ErrCode != 0 {
// 			c.log.Errorf("GetOrderTrack jd return err: %s", result.Message)
// 			return nil, fmt.Errorf("jd return err: %s", result.Message)
// 		}

// 		return result, nil
// 	}

// 	result = &GetOrderTrackResp{
// 		BaseResp: BaseResp{
// 			ErrCode:     ErrCodeParseFailed,
// 			Message:     "response parse failed",
// 			Success:     false,
// 			RawResponse: string(respBodyBytes),
// 		},
// 	}

// 	return result, nil
// }
