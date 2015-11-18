package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type CloseVulScanRequest struct {
	RpcRequest
	InstanceId string
}

func (r *CloseVulScanRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *CloseVulScanRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *CloseVulScanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CloseVulScan")
	r.SetProduct(Product)
}

type CloseVulScanResponse struct {
}

func CloseVulScan(req *CloseVulScanRequest, accessId, accessSecret string) (*CloseVulScanResponse, error) {
	var pResponse CloseVulScanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
