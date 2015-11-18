package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyVSwitchAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VSwitchId            string
	VSwitchName          string
	Description          string
	OwnerAccount         string
}

func (r *ModifyVSwitchAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyVSwitchAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyVSwitchAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyVSwitchAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyVSwitchAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyVSwitchAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyVSwitchAttributeRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *ModifyVSwitchAttributeRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *ModifyVSwitchAttributeRequest) SetVSwitchName(value string) {
	r.VSwitchName = value
	r.QueryParams.Set("VSwitchName", value)
}
func (r *ModifyVSwitchAttributeRequest) GetVSwitchName() string {
	return r.VSwitchName
}
func (r *ModifyVSwitchAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyVSwitchAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyVSwitchAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyVSwitchAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyVSwitchAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyVSwitchAttribute")
	r.SetProduct(Product)
}

type ModifyVSwitchAttributeResponse struct {
}

func ModifyVSwitchAttribute(req *ModifyVSwitchAttributeRequest, accessId, accessSecret string) (*ModifyVSwitchAttributeResponse, error) {
	var pResponse ModifyVSwitchAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
