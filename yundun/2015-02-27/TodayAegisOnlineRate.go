package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type TodayAegisOnlineRateRequest struct {
	RpcRequest
}

func (r *TodayAegisOnlineRateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TodayAegisOnlineRate")
	r.SetProduct(Product)
}

type TodayAegisOnlineRateResponse struct {
	Rate int `xml:"Rate" json:"Rate"`
}

func TodayAegisOnlineRate(req *TodayAegisOnlineRateRequest, accessId, accessSecret string) (*TodayAegisOnlineRateResponse, error) {
	var pResponse TodayAegisOnlineRateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
