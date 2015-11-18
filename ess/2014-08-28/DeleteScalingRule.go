package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteScalingRuleRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingRuleId        string
	OwnerAccount         string
}

func (r *DeleteScalingRuleRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteScalingRuleRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteScalingRuleRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteScalingRuleRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteScalingRuleRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteScalingRuleRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteScalingRuleRequest) SetScalingRuleId(value string) {
	r.ScalingRuleId = value
	r.QueryParams.Set("ScalingRuleId", value)
}
func (r *DeleteScalingRuleRequest) GetScalingRuleId() string {
	return r.ScalingRuleId
}
func (r *DeleteScalingRuleRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteScalingRuleRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteScalingRuleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteScalingRule")
	r.SetProduct(Product)
}

type DeleteScalingRuleResponse struct {
}

func DeleteScalingRule(req *DeleteScalingRuleRequest, accessId, accessSecret string) (*DeleteScalingRuleResponse, error) {
	var pResponse DeleteScalingRuleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
