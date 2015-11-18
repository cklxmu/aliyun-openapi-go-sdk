package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type WebAttackNumRequest struct {
	RpcRequest
}

func (r *WebAttackNumRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("WebAttackNum")
	r.SetProduct(Product)
}

type WebAttackNumResponse struct {
	WebAttackNum int `xml:"WebAttackNum" json:"WebAttackNum"`
}

func WebAttackNum(req *WebAttackNumRequest, accessId, accessSecret string) (*WebAttackNumResponse, error) {
	var pResponse WebAttackNumResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
