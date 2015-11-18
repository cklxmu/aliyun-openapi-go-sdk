package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyHaVipAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ClientToken          string
	HaVipId              string
	Description          string
}

func (r *ModifyHaVipAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyHaVipAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyHaVipAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyHaVipAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyHaVipAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyHaVipAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyHaVipAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyHaVipAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifyHaVipAttributeRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifyHaVipAttributeRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifyHaVipAttributeRequest) SetHaVipId(value string) {
	r.HaVipId = value
	r.QueryParams.Set("HaVipId", value)
}
func (r *ModifyHaVipAttributeRequest) GetHaVipId() string {
	return r.HaVipId
}
func (r *ModifyHaVipAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyHaVipAttributeRequest) GetDescription() string {
	return r.Description
}

func (r *ModifyHaVipAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyHaVipAttribute")
	r.SetProduct(Product)
}

type ModifyHaVipAttributeResponse struct {
}

func ModifyHaVipAttribute(req *ModifyHaVipAttributeRequest, accessId, accessSecret string) (*ModifyHaVipAttributeResponse, error) {
	var pResponse ModifyHaVipAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
