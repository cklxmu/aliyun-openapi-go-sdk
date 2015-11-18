package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyScalingGroupRequest struct {
	RpcRequest
	OwnerId                      int
	ResourceOwnerAccount         string
	ResourceOwnerId              int
	ScalingGroupId               string
	ScalingGroupName             string
	MinSize                      int
	MaxSize                      int
	DefaultCooldown              int
	RemovalPolicy_1              string
	RemovalPolicy_2              string
	ActiveScalingConfigurationId string
	OwnerAccount                 string
}

func (r *ModifyScalingGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyScalingGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyScalingGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyScalingGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyScalingGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyScalingGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyScalingGroupRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *ModifyScalingGroupRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *ModifyScalingGroupRequest) SetScalingGroupName(value string) {
	r.ScalingGroupName = value
	r.QueryParams.Set("ScalingGroupName", value)
}
func (r *ModifyScalingGroupRequest) GetScalingGroupName() string {
	return r.ScalingGroupName
}
func (r *ModifyScalingGroupRequest) SetMinSize(value int) {
	r.MinSize = value
	r.QueryParams.Set("MinSize", strconv.Itoa(value))
}
func (r *ModifyScalingGroupRequest) GetMinSize() int {
	return r.MinSize
}
func (r *ModifyScalingGroupRequest) SetMaxSize(value int) {
	r.MaxSize = value
	r.QueryParams.Set("MaxSize", strconv.Itoa(value))
}
func (r *ModifyScalingGroupRequest) GetMaxSize() int {
	return r.MaxSize
}
func (r *ModifyScalingGroupRequest) SetDefaultCooldown(value int) {
	r.DefaultCooldown = value
	r.QueryParams.Set("DefaultCooldown", strconv.Itoa(value))
}
func (r *ModifyScalingGroupRequest) GetDefaultCooldown() int {
	return r.DefaultCooldown
}
func (r *ModifyScalingGroupRequest) SetRemovalPolicy_1(value string) {
	r.RemovalPolicy_1 = value
	r.QueryParams.Set("RemovalPolicy.1", value)
}
func (r *ModifyScalingGroupRequest) GetRemovalPolicy_1() string {
	return r.RemovalPolicy_1
}
func (r *ModifyScalingGroupRequest) SetRemovalPolicy_2(value string) {
	r.RemovalPolicy_2 = value
	r.QueryParams.Set("RemovalPolicy.2", value)
}
func (r *ModifyScalingGroupRequest) GetRemovalPolicy_2() string {
	return r.RemovalPolicy_2
}
func (r *ModifyScalingGroupRequest) SetActiveScalingConfigurationId(value string) {
	r.ActiveScalingConfigurationId = value
	r.QueryParams.Set("ActiveScalingConfigurationId", value)
}
func (r *ModifyScalingGroupRequest) GetActiveScalingConfigurationId() string {
	return r.ActiveScalingConfigurationId
}
func (r *ModifyScalingGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyScalingGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyScalingGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyScalingGroup")
	r.SetProduct(Product)
}

type ModifyScalingGroupResponse struct {
}

func ModifyScalingGroup(req *ModifyScalingGroupRequest, accessId, accessSecret string) (*ModifyScalingGroupResponse, error) {
	var pResponse ModifyScalingGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
