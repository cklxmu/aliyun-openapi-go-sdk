package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type OpenVulScanRequest struct {
	RpcRequest
	InstanceId string
}

func (r *OpenVulScanRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *OpenVulScanRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *OpenVulScanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OpenVulScan")
	r.SetProduct(Product)
}

type OpenVulScanResponse struct {
}

func OpenVulScan(req *OpenVulScanRequest, accessId, accessSecret string) (*OpenVulScanResponse, error) {
	var pResponse OpenVulScanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
