package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateVpcRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	CidrBlock            string
	VpcName              string
	Description          string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
}

func (r *CreateVpcRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateVpcRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateVpcRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateVpcRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateVpcRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateVpcRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateVpcRequest) SetCidrBlock(value string) {
	r.CidrBlock = value
	r.QueryParams.Set("CidrBlock", value)
}
func (r *CreateVpcRequest) GetCidrBlock() string {
	return r.CidrBlock
}
func (r *CreateVpcRequest) SetVpcName(value string) {
	r.VpcName = value
	r.QueryParams.Set("VpcName", value)
}
func (r *CreateVpcRequest) GetVpcName() string {
	return r.VpcName
}
func (r *CreateVpcRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateVpcRequest) GetDescription() string {
	return r.Description
}
func (r *CreateVpcRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateVpcRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateVpcRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateVpcRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateVpcRequest) SetUserCidr(value string) {
	r.UserCidr = value
	r.QueryParams.Set("UserCidr", value)
}
func (r *CreateVpcRequest) GetUserCidr() string {
	return r.UserCidr
}

func (r *CreateVpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateVpc")
	r.SetProduct(Product)
}

type CreateVpcResponse struct {
	VpcId        string `xml:"VpcId" json:"VpcId"`
	VRouterId    string `xml:"VRouterId" json:"VRouterId"`
	RouteTableId string `xml:"RouteTableId" json:"RouteTableId"`
}

func CreateVpc(req *CreateVpcRequest, accessId, accessSecret string) (*CreateVpcResponse, error) {
	var pResponse CreateVpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
