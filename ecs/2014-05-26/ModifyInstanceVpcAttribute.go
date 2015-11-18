package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceVpcAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	VSwitchId            string
	PrivateIpAddress     string
	OwnerAccount         string
}

func (r *ModifyInstanceVpcAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceVpcAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyInstanceVpcAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyInstanceVpcAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyInstanceVpcAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceVpcAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyInstanceVpcAttributeRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceVpcAttributeRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceVpcAttributeRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *ModifyInstanceVpcAttributeRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *ModifyInstanceVpcAttributeRequest) SetPrivateIpAddress(value string) {
	r.PrivateIpAddress = value
	r.QueryParams.Set("PrivateIpAddress", value)
}
func (r *ModifyInstanceVpcAttributeRequest) GetPrivateIpAddress() string {
	return r.PrivateIpAddress
}
func (r *ModifyInstanceVpcAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceVpcAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyInstanceVpcAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyInstanceVpcAttribute")
	r.SetProduct(Product)
}

type ModifyInstanceVpcAttributeResponse struct {
}

func ModifyInstanceVpcAttribute(req *ModifyInstanceVpcAttributeRequest, accessId, accessSecret string) (*ModifyInstanceVpcAttributeResponse, error) {
	var pResponse ModifyInstanceVpcAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
