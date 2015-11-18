package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceVncPasswdRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	VncPassword          string
	OwnerAccount         string
}

func (r *ModifyInstanceVncPasswdRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceVncPasswdRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyInstanceVncPasswdRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyInstanceVncPasswdRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyInstanceVncPasswdRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceVncPasswdRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyInstanceVncPasswdRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceVncPasswdRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceVncPasswdRequest) SetVncPassword(value string) {
	r.VncPassword = value
	r.QueryParams.Set("VncPassword", value)
}
func (r *ModifyInstanceVncPasswdRequest) GetVncPassword() string {
	return r.VncPassword
}
func (r *ModifyInstanceVncPasswdRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceVncPasswdRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyInstanceVncPasswdRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyInstanceVncPasswd")
	r.SetProduct(Product)
}

type ModifyInstanceVncPasswdResponse struct {
}

func ModifyInstanceVncPasswd(req *ModifyInstanceVncPasswdRequest, accessId, accessSecret string) (*ModifyInstanceVncPasswdResponse, error) {
	var pResponse ModifyInstanceVncPasswdResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
