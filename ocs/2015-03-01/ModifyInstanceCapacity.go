package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceCapacityRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	Capacity             int
}

func (r *ModifyInstanceCapacityRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceCapacityRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyInstanceCapacityRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyInstanceCapacityRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyInstanceCapacityRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceCapacityRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyInstanceCapacityRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceCapacityRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifyInstanceCapacityRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceCapacityRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceCapacityRequest) SetCapacity(value int) {
	r.Capacity = value
	r.QueryParams.Set("Capacity", strconv.Itoa(value))
}
func (r *ModifyInstanceCapacityRequest) GetCapacity() int {
	return r.Capacity
}

func (r *ModifyInstanceCapacityRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyInstanceCapacity")
	r.SetProduct(Product)
}

type ModifyInstanceCapacityResponse struct {
}

func ModifyInstanceCapacity(req *ModifyInstanceCapacityRequest, accessId, accessSecret string) (*ModifyInstanceCapacityResponse, error) {
	var pResponse ModifyInstanceCapacityResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
