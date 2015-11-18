package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type OpenCCProtectRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
}

func (r *OpenCCProtectRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *OpenCCProtectRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *OpenCCProtectRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *OpenCCProtectRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *OpenCCProtectRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OpenCCProtect")
	r.SetProduct(Product)
}

type OpenCCProtectResponse struct {
}

func OpenCCProtect(req *OpenCCProtectRequest, accessId, accessSecret string) (*OpenCCProtectResponse, error) {
	var pResponse OpenCCProtectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
