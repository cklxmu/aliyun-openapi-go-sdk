package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteVSwitchRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VSwitchId            string
	OwnerAccount         string
}

func (r *DeleteVSwitchRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteVSwitchRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteVSwitchRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteVSwitchRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteVSwitchRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteVSwitchRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteVSwitchRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *DeleteVSwitchRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *DeleteVSwitchRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteVSwitchRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteVSwitchRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteVSwitch")
	r.SetProduct(Product)
}

type DeleteVSwitchResponse struct {
}

func DeleteVSwitch(req *DeleteVSwitchRequest, accessId, accessSecret string) (*DeleteVSwitchResponse, error) {
	var pResponse DeleteVSwitchResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
