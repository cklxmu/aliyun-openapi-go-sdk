package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	OwnerAccount         string
}

func (r *DeleteDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteDiskRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *DeleteDiskRequest) GetDiskId() string {
	return r.DiskId
}
func (r *DeleteDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteDisk")
	r.SetProduct(Product)
}

type DeleteDiskResponse struct {
}

func DeleteDisk(req *DeleteDiskRequest, accessId, accessSecret string) (*DeleteDiskResponse, error) {
	var pResponse DeleteDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
