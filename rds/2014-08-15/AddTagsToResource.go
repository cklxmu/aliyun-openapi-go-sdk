package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AddTagsToResourceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	proxyId              string
	DBInstanceId         string
	Tag_1_key            string
	Tag_2_key            string
	Tag_3_key            string
	Tag_4_key            string
	Tag_5_key            string
	Tag_1_value          string
	Tag_2_value          string
	Tag_3_value          string
	Tag_4_value          string
	Tag_5_value          string
	OwnerAccount         string
}

func (r *AddTagsToResourceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AddTagsToResourceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AddTagsToResourceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AddTagsToResourceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AddTagsToResourceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AddTagsToResourceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AddTagsToResourceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *AddTagsToResourceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *AddTagsToResourceRequest) SetproxyId(value string) {
	r.proxyId = value
	r.QueryParams.Set("proxyId", value)
}
func (r *AddTagsToResourceRequest) GetproxyId() string {
	return r.proxyId
}
func (r *AddTagsToResourceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *AddTagsToResourceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *AddTagsToResourceRequest) SetTag_1_key(value string) {
	r.Tag_1_key = value
	r.QueryParams.Set("Tag.1.key", value)
}
func (r *AddTagsToResourceRequest) GetTag_1_key() string {
	return r.Tag_1_key
}
func (r *AddTagsToResourceRequest) SetTag_2_key(value string) {
	r.Tag_2_key = value
	r.QueryParams.Set("Tag.2.key", value)
}
func (r *AddTagsToResourceRequest) GetTag_2_key() string {
	return r.Tag_2_key
}
func (r *AddTagsToResourceRequest) SetTag_3_key(value string) {
	r.Tag_3_key = value
	r.QueryParams.Set("Tag.3.key", value)
}
func (r *AddTagsToResourceRequest) GetTag_3_key() string {
	return r.Tag_3_key
}
func (r *AddTagsToResourceRequest) SetTag_4_key(value string) {
	r.Tag_4_key = value
	r.QueryParams.Set("Tag.4.key", value)
}
func (r *AddTagsToResourceRequest) GetTag_4_key() string {
	return r.Tag_4_key
}
func (r *AddTagsToResourceRequest) SetTag_5_key(value string) {
	r.Tag_5_key = value
	r.QueryParams.Set("Tag.5.key", value)
}
func (r *AddTagsToResourceRequest) GetTag_5_key() string {
	return r.Tag_5_key
}
func (r *AddTagsToResourceRequest) SetTag_1_value(value string) {
	r.Tag_1_value = value
	r.QueryParams.Set("Tag.1.value", value)
}
func (r *AddTagsToResourceRequest) GetTag_1_value() string {
	return r.Tag_1_value
}
func (r *AddTagsToResourceRequest) SetTag_2_value(value string) {
	r.Tag_2_value = value
	r.QueryParams.Set("Tag.2.value", value)
}
func (r *AddTagsToResourceRequest) GetTag_2_value() string {
	return r.Tag_2_value
}
func (r *AddTagsToResourceRequest) SetTag_3_value(value string) {
	r.Tag_3_value = value
	r.QueryParams.Set("Tag.3.value", value)
}
func (r *AddTagsToResourceRequest) GetTag_3_value() string {
	return r.Tag_3_value
}
func (r *AddTagsToResourceRequest) SetTag_4_value(value string) {
	r.Tag_4_value = value
	r.QueryParams.Set("Tag.4.value", value)
}
func (r *AddTagsToResourceRequest) GetTag_4_value() string {
	return r.Tag_4_value
}
func (r *AddTagsToResourceRequest) SetTag_5_value(value string) {
	r.Tag_5_value = value
	r.QueryParams.Set("Tag.5.value", value)
}
func (r *AddTagsToResourceRequest) GetTag_5_value() string {
	return r.Tag_5_value
}
func (r *AddTagsToResourceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AddTagsToResourceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AddTagsToResourceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddTagsToResource")
	r.SetProduct(Product)
}

type AddTagsToResourceResponse struct {
}

func AddTagsToResource(req *AddTagsToResourceRequest, accessId, accessSecret string) (*AddTagsToResourceResponse, error) {
	var pResponse AddTagsToResourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
