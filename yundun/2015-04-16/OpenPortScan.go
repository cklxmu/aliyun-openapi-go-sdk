package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type OpenPortScanRequest struct {
	RpcRequest
	InstanceId string
}

func (r *OpenPortScanRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *OpenPortScanRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *OpenPortScanRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OpenPortScan")
	r.SetProduct(Product)
}

type OpenPortScanResponse struct {
}

func OpenPortScan(req *OpenPortScanRequest, accessId, accessSecret string) (*OpenPortScanResponse, error) {
	var pResponse OpenPortScanResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
