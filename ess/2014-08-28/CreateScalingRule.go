package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateScalingRuleRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingGroupId       string
	ScalingRuleName      string
	Cooldown             int
	AdjustmentType       string
	AdjustmentValue      int
	OwnerAccount         string
}

func (r *CreateScalingRuleRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateScalingRuleRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateScalingRuleRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateScalingRuleRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateScalingRuleRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateScalingRuleRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateScalingRuleRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *CreateScalingRuleRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *CreateScalingRuleRequest) SetScalingRuleName(value string) {
	r.ScalingRuleName = value
	r.QueryParams.Set("ScalingRuleName", value)
}
func (r *CreateScalingRuleRequest) GetScalingRuleName() string {
	return r.ScalingRuleName
}
func (r *CreateScalingRuleRequest) SetCooldown(value int) {
	r.Cooldown = value
	r.QueryParams.Set("Cooldown", strconv.Itoa(value))
}
func (r *CreateScalingRuleRequest) GetCooldown() int {
	return r.Cooldown
}
func (r *CreateScalingRuleRequest) SetAdjustmentType(value string) {
	r.AdjustmentType = value
	r.QueryParams.Set("AdjustmentType", value)
}
func (r *CreateScalingRuleRequest) GetAdjustmentType() string {
	return r.AdjustmentType
}
func (r *CreateScalingRuleRequest) SetAdjustmentValue(value int) {
	r.AdjustmentValue = value
	r.QueryParams.Set("AdjustmentValue", strconv.Itoa(value))
}
func (r *CreateScalingRuleRequest) GetAdjustmentValue() int {
	return r.AdjustmentValue
}
func (r *CreateScalingRuleRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateScalingRuleRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateScalingRuleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateScalingRule")
	r.SetProduct(Product)
}

type CreateScalingRuleResponse struct {
	ScalingRuleId  string `xml:"ScalingRuleId" json:"ScalingRuleId"`
	ScalingRuleAri string `xml:"ScalingRuleAri" json:"ScalingRuleAri"`
}

func CreateScalingRule(req *CreateScalingRuleRequest, accessId, accessSecret string) (*CreateScalingRuleResponse, error) {
	var pResponse CreateScalingRuleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
