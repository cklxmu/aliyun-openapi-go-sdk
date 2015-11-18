package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type JoinSecurityGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SecurityGroupId      string
	InstanceId           string
	OwnerAccount         string
}

func (r *JoinSecurityGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *JoinSecurityGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *JoinSecurityGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *JoinSecurityGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *JoinSecurityGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *JoinSecurityGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *JoinSecurityGroupRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *JoinSecurityGroupRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *JoinSecurityGroupRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *JoinSecurityGroupRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *JoinSecurityGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *JoinSecurityGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *JoinSecurityGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("JoinSecurityGroup")
	r.SetProduct(Product)
}

type JoinSecurityGroupResponse struct {
}

func JoinSecurityGroup(req *JoinSecurityGroupRequest, accessId, accessSecret string) (*JoinSecurityGroupResponse, error) {
	var pResponse JoinSecurityGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
