package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateScalingGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingGroupName     string
	MinSize              int
	MaxSize              int
	DefaultCooldown      int
	LoadBalancerId       string
	DBInstanceId_1       string
	DBInstanceId_2       string
	DBInstanceId_3       string
	RemovalPolicy_1      string
	RemovalPolicy_2      string
	OwnerAccount         string
}

func (r *CreateScalingGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateScalingGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateScalingGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateScalingGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateScalingGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateScalingGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateScalingGroupRequest) SetScalingGroupName(value string) {
	r.ScalingGroupName = value
	r.QueryParams.Set("ScalingGroupName", value)
}
func (r *CreateScalingGroupRequest) GetScalingGroupName() string {
	return r.ScalingGroupName
}
func (r *CreateScalingGroupRequest) SetMinSize(value int) {
	r.MinSize = value
	r.QueryParams.Set("MinSize", strconv.Itoa(value))
}
func (r *CreateScalingGroupRequest) GetMinSize() int {
	return r.MinSize
}
func (r *CreateScalingGroupRequest) SetMaxSize(value int) {
	r.MaxSize = value
	r.QueryParams.Set("MaxSize", strconv.Itoa(value))
}
func (r *CreateScalingGroupRequest) GetMaxSize() int {
	return r.MaxSize
}
func (r *CreateScalingGroupRequest) SetDefaultCooldown(value int) {
	r.DefaultCooldown = value
	r.QueryParams.Set("DefaultCooldown", strconv.Itoa(value))
}
func (r *CreateScalingGroupRequest) GetDefaultCooldown() int {
	return r.DefaultCooldown
}
func (r *CreateScalingGroupRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *CreateScalingGroupRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *CreateScalingGroupRequest) SetDBInstanceId_1(value string) {
	r.DBInstanceId_1 = value
	r.QueryParams.Set("DBInstanceId.1", value)
}
func (r *CreateScalingGroupRequest) GetDBInstanceId_1() string {
	return r.DBInstanceId_1
}
func (r *CreateScalingGroupRequest) SetDBInstanceId_2(value string) {
	r.DBInstanceId_2 = value
	r.QueryParams.Set("DBInstanceId.2", value)
}
func (r *CreateScalingGroupRequest) GetDBInstanceId_2() string {
	return r.DBInstanceId_2
}
func (r *CreateScalingGroupRequest) SetDBInstanceId_3(value string) {
	r.DBInstanceId_3 = value
	r.QueryParams.Set("DBInstanceId.3", value)
}
func (r *CreateScalingGroupRequest) GetDBInstanceId_3() string {
	return r.DBInstanceId_3
}
func (r *CreateScalingGroupRequest) SetRemovalPolicy_1(value string) {
	r.RemovalPolicy_1 = value
	r.QueryParams.Set("RemovalPolicy.1", value)
}
func (r *CreateScalingGroupRequest) GetRemovalPolicy_1() string {
	return r.RemovalPolicy_1
}
func (r *CreateScalingGroupRequest) SetRemovalPolicy_2(value string) {
	r.RemovalPolicy_2 = value
	r.QueryParams.Set("RemovalPolicy.2", value)
}
func (r *CreateScalingGroupRequest) GetRemovalPolicy_2() string {
	return r.RemovalPolicy_2
}
func (r *CreateScalingGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateScalingGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateScalingGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateScalingGroup")
	r.SetProduct(Product)
}

type CreateScalingGroupResponse struct {
	ScalingGroupId string `xml:"ScalingGroupId" json:"ScalingGroupId"`
}

func CreateScalingGroup(req *CreateScalingGroupRequest, accessId, accessSecret string) (*CreateScalingGroupResponse, error) {
	var pResponse CreateScalingGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
