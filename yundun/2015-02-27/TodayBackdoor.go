package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type TodayBackdoorRequest struct {
	RpcRequest
}

func (r *TodayBackdoorRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TodayBackdoor")
	r.SetProduct(Product)
}

type TodayBackdoorResponse struct {
	Backdoor int `xml:"Backdoor" json:"Backdoor"`
}

func TodayBackdoor(req *TodayBackdoorRequest, accessId, accessSecret string) (*TodayBackdoorResponse, error) {
	var pResponse TodayBackdoorResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
