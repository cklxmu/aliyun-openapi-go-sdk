package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ExecuteScalingRuleRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingRuleAri       string
	ClientToken          string
	OwnerAccount         string
}

func (r *ExecuteScalingRuleRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ExecuteScalingRuleRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ExecuteScalingRuleRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ExecuteScalingRuleRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ExecuteScalingRuleRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ExecuteScalingRuleRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ExecuteScalingRuleRequest) SetScalingRuleAri(value string) {
	r.ScalingRuleAri = value
	r.QueryParams.Set("ScalingRuleAri", value)
}
func (r *ExecuteScalingRuleRequest) GetScalingRuleAri() string {
	return r.ScalingRuleAri
}
func (r *ExecuteScalingRuleRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ExecuteScalingRuleRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ExecuteScalingRuleRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ExecuteScalingRuleRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ExecuteScalingRuleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ExecuteScalingRule")
	r.SetProduct(Product)
}

type ExecuteScalingRuleResponse struct {
	ScalingActivityId string `xml:"ScalingActivityId" json:"ScalingActivityId"`
}

func ExecuteScalingRule(req *ExecuteScalingRuleRequest, accessId, accessSecret string) (*ExecuteScalingRuleResponse, error) {
	var pResponse ExecuteScalingRuleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
