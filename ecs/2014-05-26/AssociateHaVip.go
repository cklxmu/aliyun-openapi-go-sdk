package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AssociateHaVipRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ClientToken          string
	HaVipId              string
	InstanceId           string
}

func (r *AssociateHaVipRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AssociateHaVipRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AssociateHaVipRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AssociateHaVipRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AssociateHaVipRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AssociateHaVipRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AssociateHaVipRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AssociateHaVipRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *AssociateHaVipRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *AssociateHaVipRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *AssociateHaVipRequest) SetHaVipId(value string) {
	r.HaVipId = value
	r.QueryParams.Set("HaVipId", value)
}
func (r *AssociateHaVipRequest) GetHaVipId() string {
	return r.HaVipId
}
func (r *AssociateHaVipRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *AssociateHaVipRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *AssociateHaVipRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AssociateHaVip")
	r.SetProduct(Product)
}

type AssociateHaVipResponse struct {
}

func AssociateHaVip(req *AssociateHaVipRequest, accessId, accessSecret string) (*AssociateHaVipResponse, error) {
	var pResponse AssociateHaVipResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
