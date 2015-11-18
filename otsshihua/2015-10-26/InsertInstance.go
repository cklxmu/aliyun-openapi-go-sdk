package otsshihua

import (
	. "aliyun-openapi-go-sdk/core"
)

type InsertInstanceRequest struct {
	RpcRequest
	InstanceName string
	ClusterType  string
	Description  string
}

func (r *InsertInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *InsertInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *InsertInstanceRequest) SetClusterType(value string) {
	r.ClusterType = value
	r.QueryParams.Set("ClusterType", value)
}
func (r *InsertInstanceRequest) GetClusterType() string {
	return r.ClusterType
}
func (r *InsertInstanceRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *InsertInstanceRequest) GetDescription() string {
	return r.Description
}

func (r *InsertInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("InsertInstance")
	r.SetMethod("POST")
	r.SetProduct(Product)
}

type InsertInstanceResponse struct {
}

func InsertInstance(req *InsertInstanceRequest, accessId, accessSecret string) (*InsertInstanceResponse, error) {
	var pResponse InsertInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
