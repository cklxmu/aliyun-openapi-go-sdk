package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifySnapshotAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	SnapshotId           string
	SnapshotName         string
	Description          string
}

func (r *ModifySnapshotAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifySnapshotAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifySnapshotAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifySnapshotAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifySnapshotAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifySnapshotAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifySnapshotAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifySnapshotAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifySnapshotAttributeRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *ModifySnapshotAttributeRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *ModifySnapshotAttributeRequest) SetSnapshotName(value string) {
	r.SnapshotName = value
	r.QueryParams.Set("SnapshotName", value)
}
func (r *ModifySnapshotAttributeRequest) GetSnapshotName() string {
	return r.SnapshotName
}
func (r *ModifySnapshotAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifySnapshotAttributeRequest) GetDescription() string {
	return r.Description
}

func (r *ModifySnapshotAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifySnapshotAttribute")
	r.SetProduct(Product)
}

type ModifySnapshotAttributeResponse struct {
}

func ModifySnapshotAttribute(req *ModifySnapshotAttributeRequest, accessId, accessSecret string) (*ModifySnapshotAttributeResponse, error) {
	var pResponse ModifySnapshotAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
