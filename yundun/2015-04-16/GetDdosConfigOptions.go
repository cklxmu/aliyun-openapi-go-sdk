package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetDdosConfigOptionsRequest struct {
	RpcRequest
}

func (r *GetDdosConfigOptionsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetDdosConfigOptions")
	r.SetProduct(Product)
}

type GetDdosConfigOptionsResponse struct {
	QpsOptions1 struct {
		value []string `xml:"value" json:"value"`
	} `xml:"QpsOptions1" json:"QpsOptions1"`
	QpsOptions2 struct {
		value []string `xml:"value" json:"value"`
	} `xml:"QpsOptions2" json:"QpsOptions2"`
	RequestThresholdOptions1 struct {
		RequestThresholdOption []struct {
			Bps int `xml:"Bps" json:"Bps"`
			Pps int `xml:"Pps" json:"Pps"`
		} `xml:"RequestThresholdOption" json:"RequestThresholdOption"`
	} `xml:"RequestThresholdOptions1" json:"RequestThresholdOptions1"`
	RequestThresholdOptions2 struct {
		RequestThresholdOption []struct {
			Bps int `xml:"Bps" json:"Bps"`
			Pps int `xml:"Pps" json:"Pps"`
		} `xml:"RequestThresholdOption" json:"RequestThresholdOption"`
	} `xml:"RequestThresholdOptions2" json:"RequestThresholdOptions2"`
	ConnectionThresholdOptions struct {
		ConnectionThresholdOption []struct {
			Sipconn int `xml:"Sipconn" json:"Sipconn"`
			Sipnew  int `xml:"Sipnew" json:"Sipnew"`
		} `xml:"ConnectionThresholdOption" json:"ConnectionThresholdOption"`
	} `xml:"ConnectionThresholdOptions" json:"ConnectionThresholdOptions"`
}

func GetDdosConfigOptions(req *GetDdosConfigOptionsRequest, accessId, accessSecret string) (*GetDdosConfigOptionsResponse, error) {
	var pResponse GetDdosConfigOptionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
