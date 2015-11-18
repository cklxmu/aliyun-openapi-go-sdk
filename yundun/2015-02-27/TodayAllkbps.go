package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type TodayAllkbpsRequest struct {
	RpcRequest
}

func (r *TodayAllkbpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TodayAllkbps")
	r.SetProduct(Product)
}

type TodayAllkbpsResponse struct {
	Kbps int `xml:"Kbps" json:"Kbps"`
}

func TodayAllkbps(req *TodayAllkbpsRequest, accessId, accessSecret string) (*TodayAllkbpsResponse, error) {
	var pResponse TodayAllkbpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
