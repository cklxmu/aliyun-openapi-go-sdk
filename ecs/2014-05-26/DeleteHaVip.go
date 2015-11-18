package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteHaVipRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ClientToken          string
	HaVipId              string
}

func (r *DeleteHaVipRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteHaVipRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteHaVipRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteHaVipRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteHaVipRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteHaVipRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteHaVipRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteHaVipRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DeleteHaVipRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DeleteHaVipRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DeleteHaVipRequest) SetHaVipId(value string) {
	r.HaVipId = value
	r.QueryParams.Set("HaVipId", value)
}
func (r *DeleteHaVipRequest) GetHaVipId() string {
	return r.HaVipId
}

func (r *DeleteHaVipRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteHaVip")
	r.SetProduct(Product)
}

type DeleteHaVipResponse struct {
}

func DeleteHaVip(req *DeleteHaVipRequest, accessId, accessSecret string) (*DeleteHaVipResponse, error) {
	var pResponse DeleteHaVipResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
