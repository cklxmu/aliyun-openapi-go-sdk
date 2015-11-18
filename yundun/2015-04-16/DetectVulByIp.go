package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type DetectVulByIpRequest struct {
	RpcRequest
	InstanceId string
	VulIp      string
}

func (r *DetectVulByIpRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DetectVulByIpRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DetectVulByIpRequest) SetVulIp(value string) {
	r.VulIp = value
	r.QueryParams.Set("VulIp", value)
}
func (r *DetectVulByIpRequest) GetVulIp() string {
	return r.VulIp
}

func (r *DetectVulByIpRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DetectVulByIp")
	r.SetProduct(Product)
}

type DetectVulByIpResponse struct {
}

func DetectVulByIp(req *DetectVulByIpRequest, accessId, accessSecret string) (*DetectVulByIpResponse, error) {
	var pResponse DetectVulByIpResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
