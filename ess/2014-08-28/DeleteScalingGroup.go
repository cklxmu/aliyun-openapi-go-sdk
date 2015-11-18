package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteScalingGroupRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScalingGroupId       string
	ForceDelete          bool
	OwnerAccount         string
}

func (r *DeleteScalingGroupRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteScalingGroupRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteScalingGroupRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteScalingGroupRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteScalingGroupRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteScalingGroupRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteScalingGroupRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *DeleteScalingGroupRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *DeleteScalingGroupRequest) SetForceDelete(value bool) {
	r.ForceDelete = value
	r.QueryParams.Set("ForceDelete", strconv.FormatBool(value))
}
func (r *DeleteScalingGroupRequest) GetForceDelete() bool {
	return r.ForceDelete
}
func (r *DeleteScalingGroupRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteScalingGroupRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteScalingGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteScalingGroup")
	r.SetProduct(Product)
}

type DeleteScalingGroupResponse struct {
}

func DeleteScalingGroup(req *DeleteScalingGroupRequest, accessId, accessSecret string) (*DeleteScalingGroupResponse, error) {
	var pResponse DeleteScalingGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
