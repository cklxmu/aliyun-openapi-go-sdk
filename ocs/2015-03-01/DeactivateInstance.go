package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeactivateInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *DeactivateInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeactivateInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeactivateInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeactivateInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeactivateInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeactivateInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeactivateInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeactivateInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DeactivateInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DeactivateInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *DeactivateInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeactivateInstance")
	r.SetProduct(Product)
}

type DeactivateInstanceResponse struct {
}

func DeactivateInstance(req *DeactivateInstanceRequest, accessId, accessSecret string) (*DeactivateInstanceResponse, error) {
	var pResponse DeactivateInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
