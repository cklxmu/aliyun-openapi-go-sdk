package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteRouteEntryRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	RouteTableId         string
	DestinationCidrBlock string
	NextHopId            string
	OwnerAccount         string
}

func (r *DeleteRouteEntryRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteRouteEntryRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteRouteEntryRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteRouteEntryRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteRouteEntryRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteRouteEntryRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteRouteEntryRequest) SetRouteTableId(value string) {
	r.RouteTableId = value
	r.QueryParams.Set("RouteTableId", value)
}
func (r *DeleteRouteEntryRequest) GetRouteTableId() string {
	return r.RouteTableId
}
func (r *DeleteRouteEntryRequest) SetDestinationCidrBlock(value string) {
	r.DestinationCidrBlock = value
	r.QueryParams.Set("DestinationCidrBlock", value)
}
func (r *DeleteRouteEntryRequest) GetDestinationCidrBlock() string {
	return r.DestinationCidrBlock
}
func (r *DeleteRouteEntryRequest) SetNextHopId(value string) {
	r.NextHopId = value
	r.QueryParams.Set("NextHopId", value)
}
func (r *DeleteRouteEntryRequest) GetNextHopId() string {
	return r.NextHopId
}
func (r *DeleteRouteEntryRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteRouteEntryRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteRouteEntryRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteRouteEntry")
	r.SetProduct(Product)
}

type DeleteRouteEntryResponse struct {
}

func DeleteRouteEntry(req *DeleteRouteEntryRequest, accessId, accessSecret string) (*DeleteRouteEntryResponse, error) {
	var pResponse DeleteRouteEntryResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
