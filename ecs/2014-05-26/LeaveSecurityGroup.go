package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type LeaveSecurityGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SecurityGroupId      string
	InstanceId           string
	OwnerAccount         string
}

func (r *LeaveSecurityGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *LeaveSecurityGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *LeaveSecurityGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *LeaveSecurityGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *LeaveSecurityGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *LeaveSecurityGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *LeaveSecurityGroupRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *LeaveSecurityGroupRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *LeaveSecurityGroupRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *LeaveSecurityGroupRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *LeaveSecurityGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *LeaveSecurityGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *LeaveSecurityGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("LeaveSecurityGroup")
	r.SetProduct(Product)
}

type LeaveSecurityGroupResponse struct {
}

func LeaveSecurityGroup(req *LeaveSecurityGroupRequest, accessId, accessSecret string) (*LeaveSecurityGroupResponse, error) {
	var pResponse LeaveSecurityGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
