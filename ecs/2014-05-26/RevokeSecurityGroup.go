package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RevokeSecurityGroupRequest struct {
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
	NicType                 string
	OwnerAccount            string
}

func (r *RevokeSecurityGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RevokeSecurityGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RevokeSecurityGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RevokeSecurityGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RevokeSecurityGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RevokeSecurityGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RevokeSecurityGroupRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *RevokeSecurityGroupRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *RevokeSecurityGroupRequest) SetIpProtocol(value string) {
	r.IpProtocol = value
	r.QueryParams.Set("IpProtocol", value)
}
func (r *RevokeSecurityGroupRequest) GetIpProtocol() string {
	return r.IpProtocol
}
func (r *RevokeSecurityGroupRequest) SetPortRange(value string) {
	r.PortRange = value
	r.QueryParams.Set("PortRange", value)
}
func (r *RevokeSecurityGroupRequest) GetPortRange() string {
	return r.PortRange
}
func (r *RevokeSecurityGroupRequest) SetSourceGroupId(value string) {
	r.SourceGroupId = value
	r.QueryParams.Set("SourceGroupId", value)
}
func (r *RevokeSecurityGroupRequest) GetSourceGroupId() string {
	return r.SourceGroupId
}
func (r *RevokeSecurityGroupRequest) SetSourceGroupOwnerAccount(value string) {
	r.SourceGroupOwnerAccount = value
	r.QueryParams.Set("SourceGroupOwnerAccount", value)
}
func (r *RevokeSecurityGroupRequest) GetSourceGroupOwnerAccount() string {
	return r.SourceGroupOwnerAccount
}
func (r *RevokeSecurityGroupRequest) SetSourceCidrIp(value string) {
	r.SourceCidrIp = value
	r.QueryParams.Set("SourceCidrIp", value)
}
func (r *RevokeSecurityGroupRequest) GetSourceCidrIp() string {
	return r.SourceCidrIp
}
func (r *RevokeSecurityGroupRequest) SetPolicy(value string) {
	r.Policy = value
	r.QueryParams.Set("Policy", value)
}
func (r *RevokeSecurityGroupRequest) GetPolicy() string {
	return r.Policy
}
func (r *RevokeSecurityGroupRequest) SetNicType(value string) {
	r.NicType = value
	r.QueryParams.Set("NicType", value)
}
func (r *RevokeSecurityGroupRequest) GetNicType() string {
	return r.NicType
}
func (r *RevokeSecurityGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RevokeSecurityGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RevokeSecurityGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RevokeSecurityGroup")
	r.SetProduct(Product)
}

type RevokeSecurityGroupResponse struct {
}

func RevokeSecurityGroup(req *RevokeSecurityGroupRequest, accessId, accessSecret string) (*RevokeSecurityGroupResponse, error) {
	var pResponse RevokeSecurityGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
