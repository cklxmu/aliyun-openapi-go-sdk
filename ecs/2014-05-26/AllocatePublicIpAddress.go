package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AllocatePublicIpAddressRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	IpAddress            string
	VlanId               string
	OwnerAccount         string
}

func (r *AllocatePublicIpAddressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AllocatePublicIpAddressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AllocatePublicIpAddressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AllocatePublicIpAddressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AllocatePublicIpAddressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AllocatePublicIpAddressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AllocatePublicIpAddressRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *AllocatePublicIpAddressRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *AllocatePublicIpAddressRequest) SetIpAddress(value string) {
	r.IpAddress = value
	r.QueryParams.Set("IpAddress", value)
}
func (r *AllocatePublicIpAddressRequest) GetIpAddress() string {
	return r.IpAddress
}
func (r *AllocatePublicIpAddressRequest) SetVlanId(value string) {
	r.VlanId = value
	r.QueryParams.Set("VlanId", value)
}
func (r *AllocatePublicIpAddressRequest) GetVlanId() string {
	return r.VlanId
}
func (r *AllocatePublicIpAddressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AllocatePublicIpAddressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AllocatePublicIpAddressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AllocatePublicIpAddress")
	r.SetProduct(Product)
}

type AllocatePublicIpAddressResponse struct {
	IpAddress string `xml:"IpAddress" json:"IpAddress"`
}

func AllocatePublicIpAddress(req *AllocatePublicIpAddressRequest, accessId, accessSecret string) (*AllocatePublicIpAddressResponse, error) {
	var pResponse AllocatePublicIpAddressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
