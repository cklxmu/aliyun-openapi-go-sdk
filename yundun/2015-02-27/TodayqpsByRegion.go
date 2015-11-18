package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type TodayqpsByRegionRequest struct {
	RpcRequest
}

func (r *TodayqpsByRegionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("TodayqpsByRegion")
	r.SetProduct(Product)
}

type TodayqpsByRegionResponse struct {
	Data struct {
		Region []struct {
			RegionId     string `xml:"RegionId" json:"RegionId"`
			RegionNumber int    `xml:"RegionNumber" json:"RegionNumber"`
			RegionFlow   int    `xml:"RegionFlow" json:"RegionFlow"`
		} `xml:"Region" json:"Region"`
	} `xml:"Data" json:"Data"`
}

func TodayqpsByRegion(req *TodayqpsByRegionRequest, accessId, accessSecret string) (*TodayqpsByRegionResponse, error) {
	var pResponse TodayqpsByRegionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
