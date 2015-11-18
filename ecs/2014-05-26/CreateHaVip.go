package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateHaVipRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	ClientToken          string
	VSwitchId            string
	IpAddress            string
	Description          string
}

func (r *CreateHaVipRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateHaVipRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateHaVipRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateHaVipRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateHaVipRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateHaVipRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateHaVipRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateHaVipRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateHaVipRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateHaVipRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateHaVipRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *CreateHaVipRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *CreateHaVipRequest) SetIpAddress(value string) {
	r.IpAddress = value
	r.QueryParams.Set("IpAddress", value)
}
func (r *CreateHaVipRequest) GetIpAddress() string {
	return r.IpAddress
}
func (r *CreateHaVipRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateHaVipRequest) GetDescription() string {
	return r.Description
}

func (r *CreateHaVipRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateHaVip")
	r.SetProduct(Product)
}

type CreateHaVipResponse struct {
	HaVipId string `xml:"HaVipId" json:"HaVipId"`
}

func CreateHaVip(req *CreateHaVipRequest, accessId, accessSecret string) (*CreateHaVipResponse, error) {
	var pResponse CreateHaVipResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
