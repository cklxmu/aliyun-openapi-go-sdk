package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type TodayAllppsRequest struct {
	RpcRequest
}

func (r *TodayAllppsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TodayAllpps")
	r.SetProduct(Product)
}

type TodayAllppsResponse struct {
	Pps int `xml:"Pps" json:"Pps"`
}

func TodayAllpps(req *TodayAllppsRequest, accessId, accessSecret string) (*TodayAllppsResponse, error) {
	var pResponse TodayAllppsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
