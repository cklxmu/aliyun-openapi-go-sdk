package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AssociateEipAddressRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	AllocationId         string
	InstanceId           string
	OwnerAccount         string
	InstanceType         string
}

func (r *AssociateEipAddressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AssociateEipAddressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AssociateEipAddressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AssociateEipAddressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AssociateEipAddressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AssociateEipAddressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AssociateEipAddressRequest) SetAllocationId(value string) {
	r.AllocationId = value
	r.QueryParams.Set("AllocationId", value)
}
func (r *AssociateEipAddressRequest) GetAllocationId() string {
	return r.AllocationId
}
func (r *AssociateEipAddressRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *AssociateEipAddressRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *AssociateEipAddressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AssociateEipAddressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *AssociateEipAddressRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *AssociateEipAddressRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *AssociateEipAddressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AssociateEipAddress")
	r.SetProduct(Product)
}

type AssociateEipAddressResponse struct {
}

func AssociateEipAddress(req *AssociateEipAddressRequest, accessId, accessSecret string) (*AssociateEipAddressResponse, error) {
	var pResponse AssociateEipAddressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
