package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteSecurityGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SecurityGroupId      string
	OwnerAccount         string
}

func (r *DeleteSecurityGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteSecurityGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteSecurityGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteSecurityGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteSecurityGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteSecurityGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteSecurityGroupRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *DeleteSecurityGroupRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *DeleteSecurityGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteSecurityGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteSecurityGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteSecurityGroup")
	r.SetProduct(Product)
}

type DeleteSecurityGroupResponse struct {
}

func DeleteSecurityGroup(req *DeleteSecurityGroupRequest, accessId, accessSecret string) (*DeleteSecurityGroupResponse, error) {
	var pResponse DeleteSecurityGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
