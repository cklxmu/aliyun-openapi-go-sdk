package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AuthorizeSecurityGroupEgressRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	SecurityGroupId       string
	IpProtocol            string
	PortRange             string
	DestGroupId           string
	DestGroupOwnerAccount string
	DestCidrIp            string
	Policy                string
	Priority              string
	NicType               string
	ClientToken           string
	OwnerAccount          string
}

func (r *AuthorizeSecurityGroupEgressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AuthorizeSecurityGroupEgressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AuthorizeSecurityGroupEgressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AuthorizeSecurityGroupEgressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AuthorizeSecurityGroupEgressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(value string) {
	r.IpProtocol = value
	r.QueryParams.Set("IpProtocol", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetIpProtocol() string {
	return r.IpProtocol
}
func (r *AuthorizeSecurityGroupEgressRequest) SetPortRange(value string) {
	r.PortRange = value
	r.QueryParams.Set("PortRange", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetPortRange() string {
	return r.PortRange
}
func (r *AuthorizeSecurityGroupEgressRequest) SetDestGroupId(value string) {
	r.DestGroupId = value
	r.QueryParams.Set("DestGroupId", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetDestGroupId() string {
	return r.DestGroupId
}
func (r *AuthorizeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(value string) {
	r.DestGroupOwnerAccount = value
	r.QueryParams.Set("DestGroupOwnerAccount", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() string {
	return r.DestGroupOwnerAccount
}
func (r *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(value string) {
	r.DestCidrIp = value
	r.QueryParams.Set("DestCidrIp", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetDestCidrIp() string {
	return r.DestCidrIp
}
func (r *AuthorizeSecurityGroupEgressRequest) SetPolicy(value string) {
	r.Policy = value
	r.QueryParams.Set("Policy", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetPolicy() string {
	return r.Policy
}
func (r *AuthorizeSecurityGroupEgressRequest) SetPriority(value string) {
	r.Priority = value
	r.QueryParams.Set("Priority", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetPriority() string {
	return r.Priority
}
func (r *AuthorizeSecurityGroupEgressRequest) SetNicType(value string) {
	r.NicType = value
	r.QueryParams.Set("NicType", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetNicType() string {
	return r.NicType
}
func (r *AuthorizeSecurityGroupEgressRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *AuthorizeSecurityGroupEgressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AuthorizeSecurityGroupEgressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AuthorizeSecurityGroupEgressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AuthorizeSecurityGroupEgress")
	r.SetProduct(Product)
}

type AuthorizeSecurityGroupEgressResponse struct {
}

func AuthorizeSecurityGroupEgress(req *AuthorizeSecurityGroupEgressRequest, accessId, accessSecret string) (*AuthorizeSecurityGroupEgressResponse, error) {
	var pResponse AuthorizeSecurityGroupEgressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
