package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReleaseEipAddressRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	AllocationId         string
	OwnerAccount         string
}

func (r *ReleaseEipAddressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ReleaseEipAddressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ReleaseEipAddressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ReleaseEipAddressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ReleaseEipAddressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ReleaseEipAddressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ReleaseEipAddressRequest) SetAllocationId(value string) {
	r.AllocationId = value
	r.QueryParams.Set("AllocationId", value)
}
func (r *ReleaseEipAddressRequest) GetAllocationId() string {
	return r.AllocationId
}
func (r *ReleaseEipAddressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ReleaseEipAddressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ReleaseEipAddressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReleaseEipAddress")
	r.SetProduct(Product)
}

type ReleaseEipAddressResponse struct {
}

func ReleaseEipAddress(req *ReleaseEipAddressRequest, accessId, accessSecret string) (*ReleaseEipAddressResponse, error) {
	var pResponse ReleaseEipAddressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
