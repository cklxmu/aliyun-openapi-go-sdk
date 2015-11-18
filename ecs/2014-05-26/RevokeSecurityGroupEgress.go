package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RevokeSecurityGroupEgressRequest struct {
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
	NicType               string
	OwnerAccount          string
}

func (r *RevokeSecurityGroupEgressRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RevokeSecurityGroupEgressRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RevokeSecurityGroupEgressRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RevokeSecurityGroupEgressRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RevokeSecurityGroupEgressRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RevokeSecurityGroupEgressRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *RevokeSecurityGroupEgressRequest) SetIpProtocol(value string) {
	r.IpProtocol = value
	r.QueryParams.Set("IpProtocol", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetIpProtocol() string {
	return r.IpProtocol
}
func (r *RevokeSecurityGroupEgressRequest) SetPortRange(value string) {
	r.PortRange = value
	r.QueryParams.Set("PortRange", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetPortRange() string {
	return r.PortRange
}
func (r *RevokeSecurityGroupEgressRequest) SetDestGroupId(value string) {
	r.DestGroupId = value
	r.QueryParams.Set("DestGroupId", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetDestGroupId() string {
	return r.DestGroupId
}
func (r *RevokeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(value string) {
	r.DestGroupOwnerAccount = value
	r.QueryParams.Set("DestGroupOwnerAccount", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() string {
	return r.DestGroupOwnerAccount
}
func (r *RevokeSecurityGroupEgressRequest) SetDestCidrIp(value string) {
	r.DestCidrIp = value
	r.QueryParams.Set("DestCidrIp", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetDestCidrIp() string {
	return r.DestCidrIp
}
func (r *RevokeSecurityGroupEgressRequest) SetPolicy(value string) {
	r.Policy = value
	r.QueryParams.Set("Policy", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetPolicy() string {
	return r.Policy
}
func (r *RevokeSecurityGroupEgressRequest) SetNicType(value string) {
	r.NicType = value
	r.QueryParams.Set("NicType", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetNicType() string {
	return r.NicType
}
func (r *RevokeSecurityGroupEgressRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RevokeSecurityGroupEgressRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RevokeSecurityGroupEgressRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RevokeSecurityGroupEgress")
	r.SetProduct(Product)
}

type RevokeSecurityGroupEgressResponse struct {
}

func RevokeSecurityGroupEgress(req *RevokeSecurityGroupEgressRequest, accessId, accessSecret string) (*RevokeSecurityGroupEgressResponse, error) {
	var pResponse RevokeSecurityGroupEgressResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
