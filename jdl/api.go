package jdl

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	ErrCodeParseFailed = -999
)

func (c *Client) Precheck(ctx context.Context, req *PreCheckReq) (*PreCheckResp, error) {
	customerCode := req.CustomerCode
	if customerCode == "" {
		customerCode = c.config.CustomerCode
	}
	reqParams := []PreCheckReq{
		{
			SenderContact:   req.SenderContact,
			ReceiverContact: req.ReceiverContact,
			OrderOrigin:     req.OrderOrigin,
			CustomerCode:    customerCode,
			ProductsReq:     req.ProductsReq,
			Cargoes:         req.Cargoes,
		},
	}

	reqBody, err := json.Marshal(reqParams)
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.PreCheck, PreCheckPath, "ECAP", string(reqBody))
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
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.ErrorResponse != nil {
		c.log.Errorf("Precheck api error:(%d) %s", result.ErrorResponse.Code, result.ErrorResponse.ZhDesc)
		return nil, fmt.Errorf("jdl api error: %s", result.ErrorResponse.ZhDesc)
	}

	if result.ErrCode != 0 {
		c.log.Errorf("Precheck jdl return err: (%d) %s", result.ErrCode, result.Message)
		return nil, fmt.Errorf("jdl return err: %s", result.Message)
	}

	return result, nil
}

func (c *Client) GetWaybillCode(ctx context.Context, req string) (string, error) {
	return "", nil
}

func (c *Client) PlaceOrder(ctx context.Context, req *PlaceOrderReq) (*PlaceOrderResp, error) {
	customerCode := req.CustomerCode
	if customerCode == "" {
		customerCode = c.config.CustomerCode
	}
	reqParams := &PlaceOrderReq{
		OrderID:                 req.OrderID,
		SenderContact:           req.SenderContact,
		ReceiverContact:         req.ReceiverContact,
		OrderOrigin:             req.OrderOrigin,
		CustomerCode:            customerCode,
		ProductsReq:             req.ProductsReq,
		SettleType:              req.SettleType,
		Cargoes:                 req.Cargoes,
		CommonChannelInfo:       req.CommonChannelInfo,
		PickupStartTime:         req.PickupStartTime,
		PickupEndTime:           req.PickupEndTime,
		ExpectDeliveryStartTime: req.ExpectDeliveryStartTime,
		ExpectDeliveryEndTime:   req.ExpectDeliveryEndTime,
	}

	reqBody, err := json.Marshal([]PlaceOrderReq{*reqParams})
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.PlaceOrder, PlaceOrderPath, "ECAP", string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &PlaceOrderResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.ErrorResponse != nil {
		c.log.Errorf("PlaceOrder api error:(%d) %s", result.ErrorResponse.Code, result.ErrorResponse.ZhDesc)
		return nil, fmt.Errorf("jdl api error: %s", result.ErrorResponse.ZhDesc)
	}

	if result.ErrCode != 0 {
		c.log.Errorf("PlaceOrder jdl return err: (%d) %s", result.ErrCode, result.Message)
		return nil, fmt.Errorf("jdl return err: %s", result.Message)
	}
	return result, nil
}

func (c *Client) UpdateOrder(ctx context.Context, req *UpdateOrderReq) (*UpdateOrderResp, error) {
	customerCode := req.CustomerCode
	if customerCode == "" {
		customerCode = c.config.CustomerCode
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	reqParams := &UpdateOrderReq{
		OrderID:         req.OrderID,
		WaybillCode:     req.WaybillCode,
		OrderCode:       req.OrderCode,
		CustomerCode:    customerCode,
		OrderOrigin:     req.OrderOrigin,
		SenderContact:   req.SenderContact,
		ReceiverContact: req.ReceiverContact,
		ProductsReq:     req.ProductsReq,
		PickupStartTime: req.PickupStartTime,
		PickupEndTime:   req.PickupEndTime,
	}

	reqBody, err := json.Marshal([]UpdateOrderReq{*reqParams})
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.UpdateOrder, UpdateOrderPath, "ECAP", string(reqBody))
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
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.ErrorResponse != nil {
		c.log.Errorf("UpdateOrder api error:(%d) %s", result.ErrorResponse.Code, result.ErrorResponse.ZhDesc)
		return nil, fmt.Errorf("jdl api error: %s", result.ErrorResponse.ZhDesc)
	}

	if result.ErrCode != 0 {
		c.log.Errorf("UpdateOrder jdl return err: (%d) %s", result.ErrCode, result.Message)
		return nil, fmt.Errorf("jdl return err: %s", result.Message)
	}

	return result, nil
}

func (c *Client) CancelOrder(ctx context.Context, req *CancelOrderReq) (*CancelOrderResp, error) {
	customerCode := req.CustomerCode
	if customerCode == "" {
		customerCode = c.config.CustomerCode
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	reqParams := &CancelOrderReq{
		WaybillCode:        req.WaybillCode,
		OrderCode:          req.OrderCode,
		CustomerOrderId:    req.CustomerOrderId,
		OrderOrigin:        req.OrderOrigin,
		CustomerCode:       customerCode,
		CancelReason:       req.CancelReason,
		CancelReasonCode:   req.CancelReasonCode,
		CancelType:         req.CancelType,
		SubscribeIntercept: req.SubscribeIntercept,
	}

	respBodyBytes, err := c.doRequest(ctx, "POST", c.path.CancelOrder, CancelOrderPath, "ECAP", []CancelOrderReq{*reqParams})
	if err != nil {
		c.log.Errorf("cancelorder doRequest err: %s", err)
		return nil, err
	}

	result := &CancelOrderResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.ErrorResponse != nil {
		c.log.Errorf("CancelOrder api error:(%d) %s", result.ErrorResponse.Code, result.ErrorResponse.ZhDesc)
		return nil, fmt.Errorf("jdl api error: %s", result.ErrorResponse.ZhDesc)
	}

	if result.ErrCode != 0 {
		c.log.Errorf("CancelOrder jdl return err: (%d) %s", result.ErrCode, result.Message)
		return nil, fmt.Errorf("jdl return err: %s", result.Message)
	}

	return result, nil
}

func (c *Client) GetOrderStatus(ctx context.Context, req *GetOrderStatusReq) (*GetOrderStatusResp, error) {
	customerCode := req.CustomerCode
	if customerCode == "" {
		customerCode = c.config.CustomerCode
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	reqParams := &GetOrderStatusReq{
		WaybillCode:  req.WaybillCode,
		OrderCode:    req.OrderCode,
		OrderOrigin:  req.OrderOrigin,
		CustomerCode: customerCode,
	}

	respBodyBytes, err := c.doRequest(ctx, "POST", c.path.GetOrderStatus, GetOrderStatusPath, "ECAP", []GetOrderStatusReq{*reqParams})
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &GetOrderStatusResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.ErrorResponse != nil {
		c.log.Errorf("GetOrderStatus api error:(%d) %s", result.ErrorResponse.Code, result.ErrorResponse.ZhDesc)
		return nil, fmt.Errorf("jdl api error: %s", result.ErrorResponse.ZhDesc)
	}

	if result.ErrCode != 0 {
		c.log.Errorf("GetOrderStatus jdl return err: (%d) %s", result.ErrCode, result.Message)
		return nil, fmt.Errorf("jdl return err: %s", result.Message)
	}

	return result, nil
}

func (c *Client) GetOrderTrack(ctx context.Context, req *GetOrderTrackReq) (*GetOrderTrackResp, error) {
	customerCode := req.CustomerCode
	if customerCode == nil {
		customerCode = &c.config.CustomerCode
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	reqParams := &GetOrderTrackReq{
		ReferenceNumber: req.ReferenceNumber,
		ReferenceType:   req.ReferenceType,
		Phone:           req.Phone,
		CustomerCode:    customerCode,
		Scope:           req.Scope,
	}

	respBodyBytes, err := c.doRequest(ctx, "POST", c.path.GetOrderTrack, GetOrderTrackPath, "ECAP", []GetOrderTrackReq{*reqParams})
	if err != nil {
		c.log.Errorf("GetOrderTrack doRequest err: %s", err)
		return nil, err
	}

	result := &GetOrderTrackResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.ErrorResponse != nil {
		c.log.Errorf("GetOrderTrack api error:(%d) %s", result.ErrorResponse.Code, result.ErrorResponse.ZhDesc)
		return nil, fmt.Errorf("jdl api error: %s", result.ErrorResponse.ZhDesc)
	}

	// 检查业务级别错误
	if result.ErrCode != 0 {
		c.log.Errorf("GetOrderTrack jdl return err: (%d) %s", result.ErrCode, result.Message)
		return nil, fmt.Errorf("jdl return err: %s", result.Message)
	}

	return result, nil
}

func (c *Client) Print(ctx context.Context, req []*PrintReq) (*PrintResp, error) {
	for _, r := range req {
		if r.CustomerCode == "" {
			r.CustomerCode = c.config.CustomerCode
		}
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.Print, PrintPath, "jdcloudprint", string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &PrintResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.Code != 1000 && result.Code != 0 {
		c.log.Errorf("Print api error:(%d) %s", result.Code, result.Msg)
		return nil, fmt.Errorf("jdl api error: %s", result.Msg)
	}

	return result, nil
}

func (c *Client) PullData(ctx context.Context, req []*PullDataReq) (*PullDataResp, error) {
	for _, r := range req {
		if r.CpCode == "JD" {
			if r.Parameters == nil {
				r.Parameters = make(map[string]string)
			}
			if r.Parameters["ewCustomerCode"] == "" {
				r.Parameters["ewCustomerCode"] = c.config.CustomerCode
			}
		}
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.PullData, PullDataPath, "jdcloudprint", string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &PullDataResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.Code != "1" {
		c.log.Errorf("PullData api error:(%s) %s", result.Code, result.Message)
		return nil, fmt.Errorf("jdl api error: %s", result.Message)
	}

	return result, nil
}

func (c *Client) TrackQuery(ctx context.Context, req []*TrackQueryReq) (*TrackQueryResp, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		c.log.Errorf("json.Marshal(req) err: %s", err)
		return nil, err
	}

	signURL, err := c.buildSignURL(c.path.TrackQuery, TrackQueryPath, "Tracking_JD", string(reqBody))
	if err != nil {
		c.log.Errorf("build sign url err: %s", err)
		return nil, err
	}

	respBodyBytes, err := c.executeRequest(ctx, "POST", signURL, reqBody)
	if err != nil {
		c.log.Errorf("execute request err: %s", err)
		return nil, err
	}

	result := &TrackQueryResp{}

	err = json.Unmarshal(respBodyBytes, result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal body err: %w", err)
	}

	// 检查API级别错误
	if result.Code != 1000 {
		c.log.Errorf("TrackQuery api error:(%d) %s", result.Code, result.Msg)
		return nil, fmt.Errorf("jdl api error: %s", result.Msg)
	}

	return result, nil
}
