package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyImageSharePermissionRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ImageId              string
	AddAccount_1         string
	AddAccount_2         string
	AddAccount_3         string
	AddAccount_4         string
	AddAccount_5         string
	AddAccount_6         string
	AddAccount_7         string
	AddAccount_8         string
	AddAccount_9         string
	AddAccount_10        string
	RemoveAccount_1      string
	RemoveAccount_2      string
	RemoveAccount_3      string
	RemoveAccount_4      string
	RemoveAccount_5      string
	RemoveAccount_6      string
	RemoveAccount_7      string
	RemoveAccount_8      string
	RemoveAccount_9      string
	RemoveAccount_10     string
	OwnerAccount         string
}

func (r *ModifyImageSharePermissionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyImageSharePermissionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyImageSharePermissionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyImageSharePermissionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyImageSharePermissionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyImageSharePermissionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyImageSharePermissionRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *ModifyImageSharePermissionRequest) GetImageId() string {
	return r.ImageId
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_1(value string) {
	r.AddAccount_1 = value
	r.QueryParams.Set("AddAccount.1", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_1() string {
	return r.AddAccount_1
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_2(value string) {
	r.AddAccount_2 = value
	r.QueryParams.Set("AddAccount.2", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_2() string {
	return r.AddAccount_2
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_3(value string) {
	r.AddAccount_3 = value
	r.QueryParams.Set("AddAccount.3", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_3() string {
	return r.AddAccount_3
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_4(value string) {
	r.AddAccount_4 = value
	r.QueryParams.Set("AddAccount.4", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_4() string {
	return r.AddAccount_4
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_5(value string) {
	r.AddAccount_5 = value
	r.QueryParams.Set("AddAccount.5", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_5() string {
	return r.AddAccount_5
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_6(value string) {
	r.AddAccount_6 = value
	r.QueryParams.Set("AddAccount.6", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_6() string {
	return r.AddAccount_6
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_7(value string) {
	r.AddAccount_7 = value
	r.QueryParams.Set("AddAccount.7", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_7() string {
	return r.AddAccount_7
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_8(value string) {
	r.AddAccount_8 = value
	r.QueryParams.Set("AddAccount.8", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_8() string {
	return r.AddAccount_8
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_9(value string) {
	r.AddAccount_9 = value
	r.QueryParams.Set("AddAccount.9", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_9() string {
	return r.AddAccount_9
}
func (r *ModifyImageSharePermissionRequest) SetAddAccount_10(value string) {
	r.AddAccount_10 = value
	r.QueryParams.Set("AddAccount.10", value)
}
func (r *ModifyImageSharePermissionRequest) GetAddAccount_10() string {
	return r.AddAccount_10
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_1(value string) {
	r.RemoveAccount_1 = value
	r.QueryParams.Set("RemoveAccount.1", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_1() string {
	return r.RemoveAccount_1
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_2(value string) {
	r.RemoveAccount_2 = value
	r.QueryParams.Set("RemoveAccount.2", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_2() string {
	return r.RemoveAccount_2
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_3(value string) {
	r.RemoveAccount_3 = value
	r.QueryParams.Set("RemoveAccount.3", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_3() string {
	return r.RemoveAccount_3
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_4(value string) {
	r.RemoveAccount_4 = value
	r.QueryParams.Set("RemoveAccount.4", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_4() string {
	return r.RemoveAccount_4
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_5(value string) {
	r.RemoveAccount_5 = value
	r.QueryParams.Set("RemoveAccount.5", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_5() string {
	return r.RemoveAccount_5
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_6(value string) {
	r.RemoveAccount_6 = value
	r.QueryParams.Set("RemoveAccount.6", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_6() string {
	return r.RemoveAccount_6
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_7(value string) {
	r.RemoveAccount_7 = value
	r.QueryParams.Set("RemoveAccount.7", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_7() string {
	return r.RemoveAccount_7
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_8(value string) {
	r.RemoveAccount_8 = value
	r.QueryParams.Set("RemoveAccount.8", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_8() string {
	return r.RemoveAccount_8
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_9(value string) {
	r.RemoveAccount_9 = value
	r.QueryParams.Set("RemoveAccount.9", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_9() string {
	return r.RemoveAccount_9
}
func (r *ModifyImageSharePermissionRequest) SetRemoveAccount_10(value string) {
	r.RemoveAccount_10 = value
	r.QueryParams.Set("RemoveAccount.10", value)
}
func (r *ModifyImageSharePermissionRequest) GetRemoveAccount_10() string {
	return r.RemoveAccount_10
}
func (r *ModifyImageSharePermissionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyImageSharePermissionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyImageSharePermissionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyImageSharePermission")
	r.SetProduct(Product)
}

type ModifyImageSharePermissionResponse struct {
}

func ModifyImageSharePermission(req *ModifyImageSharePermissionRequest, accessId, accessSecret string) (*ModifyImageSharePermissionResponse, error) {
	var pResponse ModifyImageSharePermissionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
