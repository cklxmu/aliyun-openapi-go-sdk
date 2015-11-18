package otsfinance

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type InsertInstanceRequest struct {
	RpcRequest
	InstanceName  string
	ClusterName   string
	WriteCapacity int
	ReadCapacity  int
	EntityQuota   int
	Description   string
}

func (r *InsertInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *InsertInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *InsertInstanceRequest) SetClusterName(value string) {
	r.ClusterName = value
	r.QueryParams.Set("ClusterName", value)
}
func (r *InsertInstanceRequest) GetClusterName() string {
	return r.ClusterName
}
func (r *InsertInstanceRequest) SetWriteCapacity(value int) {
	r.WriteCapacity = value
	r.QueryParams.Set("WriteCapacity", strconv.Itoa(value))
}
func (r *InsertInstanceRequest) GetWriteCapacity() int {
	return r.WriteCapacity
}
func (r *InsertInstanceRequest) SetReadCapacity(value int) {
	r.ReadCapacity = value
	r.QueryParams.Set("ReadCapacity", strconv.Itoa(value))
}
func (r *InsertInstanceRequest) GetReadCapacity() int {
	return r.ReadCapacity
}
func (r *InsertInstanceRequest) SetEntityQuota(value int) {
	r.EntityQuota = value
	r.QueryParams.Set("EntityQuota", strconv.Itoa(value))
}
func (r *InsertInstanceRequest) GetEntityQuota() int {
	return r.EntityQuota
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
