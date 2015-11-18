package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyVpcAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VpcId                string
	Description          string
	VpcName              string
	OwnerAccount         string
	UserCidr             string
}

func (r *ModifyVpcAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyVpcAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyVpcAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyVpcAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyVpcAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyVpcAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyVpcAttributeRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *ModifyVpcAttributeRequest) GetVpcId() string {
	return r.VpcId
}
func (r *ModifyVpcAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyVpcAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyVpcAttributeRequest) SetVpcName(value string) {
	r.VpcName = value
	r.QueryParams.Set("VpcName", value)
}
func (r *ModifyVpcAttributeRequest) GetVpcName() string {
	return r.VpcName
}
func (r *ModifyVpcAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyVpcAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifyVpcAttributeRequest) SetUserCidr(value string) {
	r.UserCidr = value
	r.QueryParams.Set("UserCidr", value)
}
func (r *ModifyVpcAttributeRequest) GetUserCidr() string {
	return r.UserCidr
}

func (r *ModifyVpcAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyVpcAttribute")
	r.SetProduct(Product)
}

type ModifyVpcAttributeResponse struct {
}

func ModifyVpcAttribute(req *ModifyVpcAttributeRequest, accessId, accessSecret string) (*ModifyVpcAttributeResponse, error) {
	var pResponse ModifyVpcAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
