package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateVSwitchRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ZoneId               string
	CidrBlock            string
	VpcId                string
	VSwitchName          string
	Description          string
	ClientToken          string
	OwnerAccount         string
}

func (r *CreateVSwitchRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateVSwitchRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateVSwitchRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateVSwitchRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateVSwitchRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateVSwitchRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateVSwitchRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateVSwitchRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateVSwitchRequest) SetCidrBlock(value string) {
	r.CidrBlock = value
	r.QueryParams.Set("CidrBlock", value)
}
func (r *CreateVSwitchRequest) GetCidrBlock() string {
	return r.CidrBlock
}
func (r *CreateVSwitchRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *CreateVSwitchRequest) GetVpcId() string {
	return r.VpcId
}
func (r *CreateVSwitchRequest) SetVSwitchName(value string) {
	r.VSwitchName = value
	r.QueryParams.Set("VSwitchName", value)
}
func (r *CreateVSwitchRequest) GetVSwitchName() string {
	return r.VSwitchName
}
func (r *CreateVSwitchRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateVSwitchRequest) GetDescription() string {
	return r.Description
}
func (r *CreateVSwitchRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateVSwitchRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateVSwitchRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateVSwitchRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateVSwitchRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateVSwitch")
	r.SetProduct(Product)
}

type CreateVSwitchResponse struct {
	VSwitchId string `xml:"VSwitchId" json:"VSwitchId"`
}

func CreateVSwitch(req *CreateVSwitchRequest, accessId, accessSecret string) (*CreateVSwitchResponse, error) {
	var pResponse CreateVSwitchResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
