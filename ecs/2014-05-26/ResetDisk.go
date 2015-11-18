package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ResetDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	SnapshotId           string
	OwnerAccount         string
}

func (r *ResetDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ResetDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ResetDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ResetDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ResetDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ResetDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ResetDiskRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *ResetDiskRequest) GetDiskId() string {
	return r.DiskId
}
func (r *ResetDiskRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *ResetDiskRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *ResetDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ResetDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ResetDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ResetDisk")
	r.SetProduct(Product)
}

type ResetDiskResponse struct {
}

func ResetDisk(req *ResetDiskRequest, accessId, accessSecret string) (*ResetDiskResponse, error) {
	var pResponse ResetDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
