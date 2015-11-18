package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DetachDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	DiskId               string
	OwnerAccount         string
}

func (r *DetachDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DetachDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DetachDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DetachDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DetachDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DetachDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DetachDiskRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DetachDiskRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DetachDiskRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *DetachDiskRequest) GetDiskId() string {
	return r.DiskId
}
func (r *DetachDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DetachDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DetachDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DetachDisk")
	r.SetProduct(Product)
}

type DetachDiskResponse struct {
}

func DetachDisk(req *DetachDiskRequest, accessId, accessSecret string) (*DetachDiskResponse, error) {
	var pResponse DetachDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
