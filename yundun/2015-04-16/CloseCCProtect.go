package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type CloseCCProtectRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
}

func (r *CloseCCProtectRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *CloseCCProtectRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *CloseCCProtectRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *CloseCCProtectRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *CloseCCProtectRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CloseCCProtect")
	r.SetProduct(Product)
}

type CloseCCProtectResponse struct {
}

func CloseCCProtect(req *CloseCCProtectRequest, accessId, accessSecret string) (*CloseCCProtectResponse, error) {
	var pResponse CloseCCProtectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
