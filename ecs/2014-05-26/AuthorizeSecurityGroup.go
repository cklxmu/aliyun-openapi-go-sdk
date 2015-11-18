package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AuthorizeSecurityGroupRequest struct {
	RpcRequest
	OwnerId                 int
	ResourceOwnerAccount    string
	ResourceOwnerId         int
	SecurityGroupId         string
	IpProtocol              string
	PortRange               string
	SourceGroupId           string
	SourceGroupOwnerAccount string
	SourceCidrIp            string
	Policy                  string
	Priority                string
	NicType                 string
	ClientToken             string
	OwnerAccount            string
}

func (r *AuthorizeSecurityGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AuthorizeSecurityGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AuthorizeSecurityGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AuthorizeSecurityGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AuthorizeSecurityGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AuthorizeSecurityGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AuthorizeSecurityGroupRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *AuthorizeSecurityGroupRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *AuthorizeSecurityGroupRequest) SetIpProtocol(value string) {
	r.IpProtocol = value
	r.QueryParams.Set("IpProtocol", value)
}
func (r *AuthorizeSecurityGroupRequest) GetIpProtocol() string {
	return r.IpProtocol
}
func (r *AuthorizeSecurityGroupRequest) SetPortRange(value string) {
	r.PortRange = value
	r.QueryParams.Set("PortRange", value)
}
func (r *AuthorizeSecurityGroupRequest) GetPortRange() string {
	return r.PortRange
}
func (r *AuthorizeSecurityGroupRequest) SetSourceGroupId(value string) {
	r.SourceGroupId = value
	r.QueryParams.Set("SourceGroupId", value)
}
func (r *AuthorizeSecurityGroupRequest) GetSourceGroupId() string {
	return r.SourceGroupId
}
func (r *AuthorizeSecurityGroupRequest) SetSourceGroupOwnerAccount(value string) {
	r.SourceGroupOwnerAccount = value
	r.QueryParams.Set("SourceGroupOwnerAccount", value)
}
func (r *AuthorizeSecurityGroupRequest) GetSourceGroupOwnerAccount() string {
	return r.SourceGroupOwnerAccount
}
func (r *AuthorizeSecurityGroupRequest) SetSourceCidrIp(value string) {
	r.SourceCidrIp = value
	r.QueryParams.Set("SourceCidrIp", value)
}
func (r *AuthorizeSecurityGroupRequest) GetSourceCidrIp() string {
	return r.SourceCidrIp
}
func (r *AuthorizeSecurityGroupRequest) SetPolicy(value string) {
	r.Policy = value
	r.QueryParams.Set("Policy", value)
}
func (r *AuthorizeSecurityGroupRequest) GetPolicy() string {
	return r.Policy
}
func (r *AuthorizeSecurityGroupRequest) SetPriority(value string) {
	r.Priority = value
	r.QueryParams.Set("Priority", value)
}
func (r *AuthorizeSecurityGroupRequest) GetPriority() string {
	return r.Priority
}
func (r *AuthorizeSecurityGroupRequest) SetNicType(value string) {
	r.NicType = value
	r.QueryParams.Set("NicType", value)
}
func (r *AuthorizeSecurityGroupRequest) GetNicType() string {
	return r.NicType
}
func (r *AuthorizeSecurityGroupRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *AuthorizeSecurityGroupRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *AuthorizeSecurityGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AuthorizeSecurityGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AuthorizeSecurityGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AuthorizeSecurityGroup")
	r.SetProduct(Product)
}

type AuthorizeSecurityGroupResponse struct {
}

func AuthorizeSecurityGroup(req *AuthorizeSecurityGroupRequest, accessId, accessSecret string) (*AuthorizeSecurityGroupResponse, error) {
	var pResponse AuthorizeSecurityGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
