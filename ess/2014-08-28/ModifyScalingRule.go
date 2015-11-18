package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyScalingRuleRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingRuleId        string
	ScalingRuleName      string
	Cooldown             int
	AdjustmentType       string
	AdjustmentValue      int
	OwnerAccount         string
}

func (r *ModifyScalingRuleRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyScalingRuleRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyScalingRuleRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyScalingRuleRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyScalingRuleRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyScalingRuleRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyScalingRuleRequest) SetScalingRuleId(value string) {
	r.ScalingRuleId = value
	r.QueryParams.Set("ScalingRuleId", value)
}
func (r *ModifyScalingRuleRequest) GetScalingRuleId() string {
	return r.ScalingRuleId
}
func (r *ModifyScalingRuleRequest) SetScalingRuleName(value string) {
	r.ScalingRuleName = value
	r.QueryParams.Set("ScalingRuleName", value)
}
func (r *ModifyScalingRuleRequest) GetScalingRuleName() string {
	return r.ScalingRuleName
}
func (r *ModifyScalingRuleRequest) SetCooldown(value int) {
	r.Cooldown = value
	r.QueryParams.Set("Cooldown", strconv.Itoa(value))
}
func (r *ModifyScalingRuleRequest) GetCooldown() int {
	return r.Cooldown
}
func (r *ModifyScalingRuleRequest) SetAdjustmentType(value string) {
	r.AdjustmentType = value
	r.QueryParams.Set("AdjustmentType", value)
}
func (r *ModifyScalingRuleRequest) GetAdjustmentType() string {
	return r.AdjustmentType
}
func (r *ModifyScalingRuleRequest) SetAdjustmentValue(value int) {
	r.AdjustmentValue = value
	r.QueryParams.Set("AdjustmentValue", strconv.Itoa(value))
}
func (r *ModifyScalingRuleRequest) GetAdjustmentValue() int {
	return r.AdjustmentValue
}
func (r *ModifyScalingRuleRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyScalingRuleRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyScalingRuleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyScalingRule")
	r.SetProduct(Product)
}

type ModifyScalingRuleResponse struct {
}

func ModifyScalingRule(req *ModifyScalingRuleRequest, accessId, accessSecret string) (*ModifyScalingRuleResponse, error) {
	var pResponse ModifyScalingRuleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
