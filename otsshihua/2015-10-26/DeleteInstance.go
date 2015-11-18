package otsshihua

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteInstanceRequest struct {
	RpcRequest
	InstanceName string
}

func (r *DeleteInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *DeleteInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}

func (r *DeleteInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteInstance")
	r.SetMethod("POST")
	r.SetProduct(Product)
}

type DeleteInstanceResponse struct {
}

func DeleteInstance(req *DeleteInstanceRequest, accessId, accessSecret string) (*DeleteInstanceResponse, error) {
	var pResponse DeleteInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
