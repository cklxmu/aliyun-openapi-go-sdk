package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UnassociateHaVipRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ClientToken          string
	HaVipId              string
	InstanceId           string
	Force                string
}

func (r *UnassociateHaVipRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *UnassociateHaVipRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *UnassociateHaVipRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *UnassociateHaVipRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *UnassociateHaVipRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *UnassociateHaVipRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *UnassociateHaVipRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *UnassociateHaVipRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *UnassociateHaVipRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *UnassociateHaVipRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *UnassociateHaVipRequest) SetHaVipId(value string) {
	r.HaVipId = value
	r.QueryParams.Set("HaVipId", value)
}
func (r *UnassociateHaVipRequest) GetHaVipId() string {
	return r.HaVipId
}
func (r *UnassociateHaVipRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *UnassociateHaVipRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *UnassociateHaVipRequest) SetForce(value string) {
	r.Force = value
	r.QueryParams.Set("Force", value)
}
func (r *UnassociateHaVipRequest) GetForce() string {
	return r.Force
}

func (r *UnassociateHaVipRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UnassociateHaVip")
	r.SetProduct(Product)
}

type UnassociateHaVipResponse struct {
}

func UnassociateHaVip(req *UnassociateHaVipRequest, accessId, accessSecret string) (*UnassociateHaVipResponse, error) {
	var pResponse UnassociateHaVipResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
