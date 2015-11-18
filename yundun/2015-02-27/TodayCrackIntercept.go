package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type TodayCrackInterceptRequest struct {
	RpcRequest
}

func (r *TodayCrackInterceptRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TodayCrackIntercept")
	r.SetProduct(Product)
}

type TodayCrackInterceptResponse struct {
	InterceptNum int `xml:"InterceptNum" json:"InterceptNum"`
}

func TodayCrackIntercept(req *TodayCrackInterceptRequest, accessId, accessSecret string) (*TodayCrackInterceptResponse, error) {
	var pResponse TodayCrackInterceptResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
