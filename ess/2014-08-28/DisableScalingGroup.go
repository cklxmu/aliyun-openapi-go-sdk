package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DisableScalingGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingGroupId       string
	OwnerAccount         string
}

func (r *DisableScalingGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DisableScalingGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DisableScalingGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DisableScalingGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DisableScalingGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DisableScalingGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DisableScalingGroupRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *DisableScalingGroupRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *DisableScalingGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DisableScalingGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DisableScalingGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DisableScalingGroup")
	r.SetProduct(Product)
}

type DisableScalingGroupResponse struct {
}

func DisableScalingGroup(req *DisableScalingGroupRequest, accessId, accessSecret string) (*DisableScalingGroupResponse, error) {
	var pResponse DisableScalingGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
