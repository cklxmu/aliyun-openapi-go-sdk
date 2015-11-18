package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type QueryDdosConfigRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
}

func (r *QueryDdosConfigRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *QueryDdosConfigRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *QueryDdosConfigRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *QueryDdosConfigRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *QueryDdosConfigRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("QueryDdosConfig")
	r.SetProduct(Product)
}

type QueryDdosConfigResponse struct {
	Bps              int    `xml:"Bps" json:"Bps"`
	Pps              int    `xml:"Pps" json:"Pps"`
	Qps              int    `xml:"Qps" json:"Qps"`
	Sipconn          int    `xml:"Sipconn" json:"Sipconn"`
	Sipnew           int    `xml:"Sipnew" json:"Sipnew"`
	Layer7Config     bool   `xml:"Layer7Config" json:"Layer7Config"`
	FlowPosition     int    `xml:"FlowPosition" json:"FlowPosition"`
	QpsPosition      int    `xml:"QpsPosition" json:"QpsPosition"`
	StrategyPosition int    `xml:"StrategyPosition" json:"StrategyPosition"`
	Level            int    `xml:"Level" json:"Level"`
	HoleBps          string `xml:"HoleBps" json:"HoleBps"`
	ConfigType       string `xml:"ConfigType" json:"ConfigType"`
}

func QueryDdosConfig(req *QueryDdosConfigRequest, accessId, accessSecret string) (*QueryDdosConfigResponse, error) {
	var pResponse QueryDdosConfigResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
