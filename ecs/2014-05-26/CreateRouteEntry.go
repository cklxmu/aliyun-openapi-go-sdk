package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateRouteEntryRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	RouteTableId         string
	DestinationCidrBlock string
	NextHopId            string
	ClientToken          string
	NextHopType          string
	OwnerAccount         string
}

func (r *CreateRouteEntryRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateRouteEntryRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateRouteEntryRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateRouteEntryRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateRouteEntryRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateRouteEntryRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateRouteEntryRequest) SetRouteTableId(value string) {
	r.RouteTableId = value
	r.QueryParams.Set("RouteTableId", value)
}
func (r *CreateRouteEntryRequest) GetRouteTableId() string {
	return r.RouteTableId
}
func (r *CreateRouteEntryRequest) SetDestinationCidrBlock(value string) {
	r.DestinationCidrBlock = value
	r.QueryParams.Set("DestinationCidrBlock", value)
}
func (r *CreateRouteEntryRequest) GetDestinationCidrBlock() string {
	return r.DestinationCidrBlock
}
func (r *CreateRouteEntryRequest) SetNextHopId(value string) {
	r.NextHopId = value
	r.QueryParams.Set("NextHopId", value)
}
func (r *CreateRouteEntryRequest) GetNextHopId() string {
	return r.NextHopId
}
func (r *CreateRouteEntryRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateRouteEntryRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateRouteEntryRequest) SetNextHopType(value string) {
	r.NextHopType = value
	r.QueryParams.Set("NextHopType", value)
}
func (r *CreateRouteEntryRequest) GetNextHopType() string {
	return r.NextHopType
}
func (r *CreateRouteEntryRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateRouteEntryRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateRouteEntryRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateRouteEntry")
	r.SetProduct(Product)
}

type CreateRouteEntryResponse struct {
}

func CreateRouteEntry(req *CreateRouteEntryRequest, accessId, accessSecret string) (*CreateRouteEntryResponse, error) {
	var pResponse CreateRouteEntryResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
