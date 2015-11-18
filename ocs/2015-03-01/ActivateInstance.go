package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ActivateInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *ActivateInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ActivateInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ActivateInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ActivateInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ActivateInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ActivateInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ActivateInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ActivateInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ActivateInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ActivateInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *ActivateInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ActivateInstance")
	r.SetProduct(Product)
}

type ActivateInstanceResponse struct {
}

func ActivateInstance(req *ActivateInstanceRequest, accessId, accessSecret string) (*ActivateInstanceResponse, error) {
	var pResponse ActivateInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
