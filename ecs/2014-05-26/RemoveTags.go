package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RemoveTagsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ResourceType         string
	ResourceId           string
	Tag_1_Key            string
	Tag_2_Key            string
	Tag_3_Key            string
	Tag_4_Key            string
	Tag_5_Key            string
	Tag_1_Value          string
	Tag_2_Value          string
	Tag_3_Value          string
	Tag_4_Value          string
	Tag_5_Value          string
}

func (r *RemoveTagsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RemoveTagsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RemoveTagsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RemoveTagsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RemoveTagsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RemoveTagsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RemoveTagsRequest) SetResourceType(value string) {
	r.ResourceType = value
	r.QueryParams.Set("ResourceType", value)
}
func (r *RemoveTagsRequest) GetResourceType() string {
	return r.ResourceType
}
func (r *RemoveTagsRequest) SetResourceId(value string) {
	r.ResourceId = value
	r.QueryParams.Set("ResourceId", value)
}
func (r *RemoveTagsRequest) GetResourceId() string {
	return r.ResourceId
}
func (r *RemoveTagsRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *RemoveTagsRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *RemoveTagsRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *RemoveTagsRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *RemoveTagsRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *RemoveTagsRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *RemoveTagsRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *RemoveTagsRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *RemoveTagsRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *RemoveTagsRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *RemoveTagsRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *RemoveTagsRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *RemoveTagsRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *RemoveTagsRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *RemoveTagsRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *RemoveTagsRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *RemoveTagsRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *RemoveTagsRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *RemoveTagsRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *RemoveTagsRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *RemoveTagsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveTags")
	r.SetProduct(Product)
}

type RemoveTagsResponse struct {
}

func RemoveTags(req *RemoveTagsRequest, accessId, accessSecret string) (*RemoveTagsResponse, error) {
	var pResponse RemoveTagsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
