package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteVpcRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VpcId                string
	OwnerAccount         string
}

func (r *DeleteVpcRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteVpcRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteVpcRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteVpcRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteVpcRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteVpcRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteVpcRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DeleteVpcRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DeleteVpcRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteVpcRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteVpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteVpc")
	r.SetProduct(Product)
}

type DeleteVpcResponse struct {
}

func DeleteVpc(req *DeleteVpcRequest, accessId, accessSecret string) (*DeleteVpcResponse, error) {
	var pResponse DeleteVpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
