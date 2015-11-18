package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetDdosAutoRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
}

func (r *SetDdosAutoRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *SetDdosAutoRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *SetDdosAutoRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *SetDdosAutoRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *SetDdosAutoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetDdosAuto")
	r.SetProduct(Product)
}

type SetDdosAutoResponse struct {
}

func SetDdosAuto(req *SetDdosAutoRequest, accessId, accessSecret string) (*SetDdosAutoResponse, error) {
	var pResponse SetDdosAutoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
