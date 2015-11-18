package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateSecurityGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	Description          string
	ClientToken          string
	SecurityGroupName    string
	VpcId                string
	OwnerAccount         string
}

func (r *CreateSecurityGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateSecurityGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateSecurityGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateSecurityGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateSecurityGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateSecurityGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateSecurityGroupRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateSecurityGroupRequest) GetDescription() string {
	return r.Description
}
func (r *CreateSecurityGroupRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateSecurityGroupRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateSecurityGroupRequest) SetSecurityGroupName(value string) {
	r.SecurityGroupName = value
	r.QueryParams.Set("SecurityGroupName", value)
}
func (r *CreateSecurityGroupRequest) GetSecurityGroupName() string {
	return r.SecurityGroupName
}
func (r *CreateSecurityGroupRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *CreateSecurityGroupRequest) GetVpcId() string {
	return r.VpcId
}
func (r *CreateSecurityGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateSecurityGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateSecurityGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateSecurityGroup")
	r.SetProduct(Product)
}

type CreateSecurityGroupResponse struct {
	SecurityGroupId string `xml:"SecurityGroupId" json:"SecurityGroupId"`
}

func CreateSecurityGroup(req *CreateSecurityGroupRequest, accessId, accessSecret string) (*CreateSecurityGroupResponse, error) {
	var pResponse CreateSecurityGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
