package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyVRouterAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VRouterId            string
	VRouterName          string
	Description          string
	OwnerAccount         string
}

func (r *ModifyVRouterAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyVRouterAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyVRouterAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyVRouterAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyVRouterAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyVRouterAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyVRouterAttributeRequest) SetVRouterId(value string) {
	r.VRouterId = value
	r.QueryParams.Set("VRouterId", value)
}
func (r *ModifyVRouterAttributeRequest) GetVRouterId() string {
	return r.VRouterId
}
func (r *ModifyVRouterAttributeRequest) SetVRouterName(value string) {
	r.VRouterName = value
	r.QueryParams.Set("VRouterName", value)
}
func (r *ModifyVRouterAttributeRequest) GetVRouterName() string {
	return r.VRouterName
}
func (r *ModifyVRouterAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyVRouterAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyVRouterAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyVRouterAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyVRouterAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyVRouterAttribute")
	r.SetProduct(Product)
}

type ModifyVRouterAttributeResponse struct {
}

func ModifyVRouterAttribute(req *ModifyVRouterAttributeRequest, accessId, accessSecret string) (*ModifyVRouterAttributeResponse, error) {
	var pResponse ModifyVRouterAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
