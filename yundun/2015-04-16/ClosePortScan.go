package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type ClosePortScanRequest struct {
	RpcRequest
	InstanceId string
}

func (r *ClosePortScanRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ClosePortScanRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *ClosePortScanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ClosePortScan")
	r.SetProduct(Product)
}

type ClosePortScanResponse struct {
}

func ClosePortScan(req *ClosePortScanRequest, accessId, accessSecret string) (*ClosePortScanResponse, error) {
	var pResponse ClosePortScanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
