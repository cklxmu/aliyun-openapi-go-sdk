package otsfinance

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UpdateInstanceRequest struct {
	RpcRequest
	InstanceName  string
	WriteCapacity int
	ReadCapacity  int
	EntityQuota   int
	Description   string
}

func (r *UpdateInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *UpdateInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *UpdateInstanceRequest) SetWriteCapacity(value int) {
	r.WriteCapacity = value
	r.QueryParams.Set("WriteCapacity", strconv.Itoa(value))
}
func (r *UpdateInstanceRequest) GetWriteCapacity() int {
	return r.WriteCapacity
}
func (r *UpdateInstanceRequest) SetReadCapacity(value int) {
	r.ReadCapacity = value
	r.QueryParams.Set("ReadCapacity", strconv.Itoa(value))
}
func (r *UpdateInstanceRequest) GetReadCapacity() int {
	return r.ReadCapacity
}
func (r *UpdateInstanceRequest) SetEntityQuota(value int) {
	r.EntityQuota = value
	r.QueryParams.Set("EntityQuota", strconv.Itoa(value))
}
func (r *UpdateInstanceRequest) GetEntityQuota() int {
	return r.EntityQuota
}
func (r *UpdateInstanceRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *UpdateInstanceRequest) GetDescription() string {
	return r.Description
}

func (r *UpdateInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateInstance")
	r.SetMethod("POST")
	r.SetProduct(Product)
}

type UpdateInstanceResponse struct {
}

func UpdateInstance(req *UpdateInstanceRequest, accessId, accessSecret string) (*UpdateInstanceResponse, error) {
	var pResponse UpdateInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
