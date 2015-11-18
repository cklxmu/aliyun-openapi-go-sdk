package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifySecurityGroupAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SecurityGroupId      string
	Description          string
	SecurityGroupName    string
	OwnerAccount         string
}

func (r *ModifySecurityGroupAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifySecurityGroupAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifySecurityGroupAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifySecurityGroupAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifySecurityGroupAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifySecurityGroupAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifySecurityGroupAttributeRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *ModifySecurityGroupAttributeRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *ModifySecurityGroupAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifySecurityGroupAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifySecurityGroupAttributeRequest) SetSecurityGroupName(value string) {
	r.SecurityGroupName = value
	r.QueryParams.Set("SecurityGroupName", value)
}
func (r *ModifySecurityGroupAttributeRequest) GetSecurityGroupName() string {
	return r.SecurityGroupName
}
func (r *ModifySecurityGroupAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifySecurityGroupAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifySecurityGroupAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifySecurityGroupAttribute")
	r.SetProduct(Product)
}

type ModifySecurityGroupAttributeResponse struct {
}

func ModifySecurityGroupAttribute(req *ModifySecurityGroupAttributeRequest, accessId, accessSecret string) (*ModifySecurityGroupAttributeResponse, error) {
	var pResponse ModifySecurityGroupAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
