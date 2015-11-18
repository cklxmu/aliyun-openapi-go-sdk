package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type ServiceStatusRequest struct {
	RpcRequest
	InstanceId string
}

func (r *ServiceStatusRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ServiceStatusRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *ServiceStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ServiceStatus")
	r.SetProduct(Product)
}

type ServiceStatusResponse struct {
	PortScan bool `xml:"PortScan" json:"PortScan"`
	VulScan  bool `xml:"VulScan" json:"VulScan"`
}

func ServiceStatus(req *ServiceStatusRequest, accessId, accessSecret string) (*ServiceStatusResponse, error) {
	var pResponse ServiceStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
