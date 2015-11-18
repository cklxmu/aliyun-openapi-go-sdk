package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyEipAddressAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	AllocationId         string
	Bandwidth            string
	OwnerAccount         string
}

func (r *ModifyEipAddressAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyEipAddressAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyEipAddressAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyEipAddressAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyEipAddressAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyEipAddressAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyEipAddressAttributeRequest) SetAllocationId(value string) {
	r.AllocationId = value
	r.QueryParams.Set("AllocationId", value)
}
func (r *ModifyEipAddressAttributeRequest) GetAllocationId() string {
	return r.AllocationId
}
func (r *ModifyEipAddressAttributeRequest) SetBandwidth(value string) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", value)
}
func (r *ModifyEipAddressAttributeRequest) GetBandwidth() string {
	return r.Bandwidth
}
func (r *ModifyEipAddressAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyEipAddressAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyEipAddressAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyEipAddressAttribute")
	r.SetProduct(Product)
}

type ModifyEipAddressAttributeResponse struct {
}

func ModifyEipAddressAttribute(req *ModifyEipAddressAttributeRequest, accessId, accessSecret string) (*ModifyEipAddressAttributeResponse, error) {
	var pResponse ModifyEipAddressAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
