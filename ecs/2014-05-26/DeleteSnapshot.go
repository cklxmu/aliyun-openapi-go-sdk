package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteSnapshotRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SnapshotId           string
	OwnerAccount         string
}

func (r *DeleteSnapshotRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteSnapshotRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteSnapshotRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteSnapshotRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteSnapshotRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteSnapshotRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteSnapshotRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *DeleteSnapshotRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *DeleteSnapshotRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteSnapshotRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteSnapshotRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteSnapshot")
	r.SetProduct(Product)
}

type DeleteSnapshotResponse struct {
}

func DeleteSnapshot(req *DeleteSnapshotRequest, accessId, accessSecret string) (*DeleteSnapshotResponse, error) {
	var pResponse DeleteSnapshotResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
