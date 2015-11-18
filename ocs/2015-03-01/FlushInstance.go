package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type FlushInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *FlushInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *FlushInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *FlushInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *FlushInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *FlushInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *FlushInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *FlushInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *FlushInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *FlushInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *FlushInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *FlushInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("FlushInstance")
	r.SetProduct(Product)
}

type FlushInstanceResponse struct {
}

func FlushInstance(req *FlushInstanceRequest, accessId, accessSecret string) (*FlushInstanceResponse, error) {
	var pResponse FlushInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
