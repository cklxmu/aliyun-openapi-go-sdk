package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDiskAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	DiskName             string
	Description          string
	DeleteWithInstance   bool
	DeleteAutoSnapshot   bool
	EnableAutoSnapshot   bool
	OwnerAccount         string
}

func (r *ModifyDiskAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDiskAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDiskAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDiskAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDiskAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDiskAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDiskAttributeRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *ModifyDiskAttributeRequest) GetDiskId() string {
	return r.DiskId
}
func (r *ModifyDiskAttributeRequest) SetDiskName(value string) {
	r.DiskName = value
	r.QueryParams.Set("DiskName", value)
}
func (r *ModifyDiskAttributeRequest) GetDiskName() string {
	return r.DiskName
}
func (r *ModifyDiskAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyDiskAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyDiskAttributeRequest) SetDeleteWithInstance(value bool) {
	r.DeleteWithInstance = value
	r.QueryParams.Set("DeleteWithInstance", strconv.FormatBool(value))
}
func (r *ModifyDiskAttributeRequest) GetDeleteWithInstance() bool {
	return r.DeleteWithInstance
}
func (r *ModifyDiskAttributeRequest) SetDeleteAutoSnapshot(value bool) {
	r.DeleteAutoSnapshot = value
	r.QueryParams.Set("DeleteAutoSnapshot", strconv.FormatBool(value))
}
func (r *ModifyDiskAttributeRequest) GetDeleteAutoSnapshot() bool {
	return r.DeleteAutoSnapshot
}
func (r *ModifyDiskAttributeRequest) SetEnableAutoSnapshot(value bool) {
	r.EnableAutoSnapshot = value
	r.QueryParams.Set("EnableAutoSnapshot", strconv.FormatBool(value))
}
func (r *ModifyDiskAttributeRequest) GetEnableAutoSnapshot() bool {
	return r.EnableAutoSnapshot
}
func (r *ModifyDiskAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDiskAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDiskAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDiskAttribute")
	r.SetProduct(Product)
}

type ModifyDiskAttributeResponse struct {
}

func ModifyDiskAttribute(req *ModifyDiskAttributeRequest, accessId, accessSecret string) (*ModifyDiskAttributeResponse, error) {
	var pResponse ModifyDiskAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
