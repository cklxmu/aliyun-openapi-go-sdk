package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateSnapshotRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	SnapshotName         string
	Description          string
	ClientToken          string
	OwnerAccount         string
}

func (r *CreateSnapshotRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateSnapshotRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateSnapshotRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateSnapshotRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateSnapshotRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateSnapshotRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateSnapshotRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *CreateSnapshotRequest) GetDiskId() string {
	return r.DiskId
}
func (r *CreateSnapshotRequest) SetSnapshotName(value string) {
	r.SnapshotName = value
	r.QueryParams.Set("SnapshotName", value)
}
func (r *CreateSnapshotRequest) GetSnapshotName() string {
	return r.SnapshotName
}
func (r *CreateSnapshotRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateSnapshotRequest) GetDescription() string {
	return r.Description
}
func (r *CreateSnapshotRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateSnapshotRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateSnapshotRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateSnapshotRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateSnapshotRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateSnapshot")
	r.SetProduct(Product)
}

type CreateSnapshotResponse struct {
	SnapshotId string `xml:"SnapshotId" json:"SnapshotId"`
}

func CreateSnapshot(req *CreateSnapshotRequest, accessId, accessSecret string) (*CreateSnapshotResponse, error) {
	var pResponse CreateSnapshotResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
