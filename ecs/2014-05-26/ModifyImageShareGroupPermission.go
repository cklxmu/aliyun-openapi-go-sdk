package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyImageShareGroupPermissionRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ImageId              string
	AddGroup_1           string
	RemoveGroup_1        string
	OwnerAccount         string
}

func (r *ModifyImageShareGroupPermissionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyImageShareGroupPermissionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyImageShareGroupPermissionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyImageShareGroupPermissionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyImageShareGroupPermissionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyImageShareGroupPermissionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyImageShareGroupPermissionRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *ModifyImageShareGroupPermissionRequest) GetImageId() string {
	return r.ImageId
}
func (r *ModifyImageShareGroupPermissionRequest) SetAddGroup_1(value string) {
	r.AddGroup_1 = value
	r.QueryParams.Set("AddGroup.1", value)
}
func (r *ModifyImageShareGroupPermissionRequest) GetAddGroup_1() string {
	return r.AddGroup_1
}
func (r *ModifyImageShareGroupPermissionRequest) SetRemoveGroup_1(value string) {
	r.RemoveGroup_1 = value
	r.QueryParams.Set("RemoveGroup.1", value)
}
func (r *ModifyImageShareGroupPermissionRequest) GetRemoveGroup_1() string {
	return r.RemoveGroup_1
}
func (r *ModifyImageShareGroupPermissionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyImageShareGroupPermissionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyImageShareGroupPermissionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyImageShareGroupPermission")
	r.SetProduct(Product)
}

type ModifyImageShareGroupPermissionResponse struct {
}

func ModifyImageShareGroupPermission(req *ModifyImageShareGroupPermissionRequest, accessId, accessSecret string) (*ModifyImageShareGroupPermissionResponse, error) {
	var pResponse ModifyImageShareGroupPermissionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
