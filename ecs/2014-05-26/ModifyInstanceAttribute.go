package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	Password             string
	HostName             string
	InstanceName         string
	Description          string
	OwnerAccount         string
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
func (r *ModifyInstanceAttributeRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceAttributeRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceAttributeRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *ModifyInstanceAttributeRequest) GetPassword() string {
	return r.Password
}
func (r *ModifyInstanceAttributeRequest) SetHostName(value string) {
	r.HostName = value
	r.QueryParams.Set("HostName", value)
}
func (r *ModifyInstanceAttributeRequest) GetHostName() string {
	return r.HostName
}
func (r *ModifyInstanceAttributeRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *ModifyInstanceAttributeRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *ModifyInstanceAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyInstanceAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyInstanceAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
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
