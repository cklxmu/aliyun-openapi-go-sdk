package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	InstanceName         string
	NewPassword          string
}

func (r *ModifyInstanceAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyInstanceAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyInstanceAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyInstanceAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyInstanceAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifyInstanceAttributeRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceAttributeRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceAttributeRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *ModifyInstanceAttributeRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *ModifyInstanceAttributeRequest) SetNewPassword(value string) {
	r.NewPassword = value
	r.QueryParams.Set("NewPassword", value)
}
func (r *ModifyInstanceAttributeRequest) GetNewPassword() string {
	return r.NewPassword
}

func (r *ModifyInstanceAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyInstanceAttribute")
	r.SetProduct(Product)
}

type ModifyInstanceAttributeResponse struct {
}

func ModifyInstanceAttribute(req *ModifyInstanceAttributeRequest, accessId, accessSecret string) (*ModifyInstanceAttributeResponse, error) {
	var pResponse ModifyInstanceAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
