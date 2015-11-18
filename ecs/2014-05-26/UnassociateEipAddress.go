package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UnassociateEipAddressRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	AllocationId         string
	InstanceId           string
	OwnerAccount         string
	InstanceType         string
}

func (r *UnassociateEipAddressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *UnassociateEipAddressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *UnassociateEipAddressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *UnassociateEipAddressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *UnassociateEipAddressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *UnassociateEipAddressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *UnassociateEipAddressRequest) SetAllocationId(value string) {
	r.AllocationId = value
	r.QueryParams.Set("AllocationId", value)
}
func (r *UnassociateEipAddressRequest) GetAllocationId() string {
	return r.AllocationId
}
func (r *UnassociateEipAddressRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *UnassociateEipAddressRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *UnassociateEipAddressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *UnassociateEipAddressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *UnassociateEipAddressRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *UnassociateEipAddressRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *UnassociateEipAddressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UnassociateEipAddress")
	r.SetProduct(Product)
}

type UnassociateEipAddressResponse struct {
}

func UnassociateEipAddress(req *UnassociateEipAddressRequest, accessId, accessSecret string) (*UnassociateEipAddressResponse, error) {
	var pResponse UnassociateEipAddressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
