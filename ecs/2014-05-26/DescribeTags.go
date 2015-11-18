package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeTagsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageSize             int
	PageNumber           int
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

func (r *DescribeTagsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeTagsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeTagsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeTagsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeTagsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeTagsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeTagsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeTagsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeTagsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeTagsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeTagsRequest) SetResourceType(value string) {
	r.ResourceType = value
	r.QueryParams.Set("ResourceType", value)
}
func (r *DescribeTagsRequest) GetResourceType() string {
	return r.ResourceType
}
func (r *DescribeTagsRequest) SetResourceId(value string) {
	r.ResourceId = value
	r.QueryParams.Set("ResourceId", value)
}
func (r *DescribeTagsRequest) GetResourceId() string {
	return r.ResourceId
}
func (r *DescribeTagsRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *DescribeTagsRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *DescribeTagsRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *DescribeTagsRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *DescribeTagsRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *DescribeTagsRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *DescribeTagsRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *DescribeTagsRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *DescribeTagsRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *DescribeTagsRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *DescribeTagsRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *DescribeTagsRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *DescribeTagsRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *DescribeTagsRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *DescribeTagsRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *DescribeTagsRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *DescribeTagsRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *DescribeTagsRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *DescribeTagsRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *DescribeTagsRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *DescribeTagsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeTags")
	r.SetProduct(Product)
}

type DescribeTagsResponse struct {
	PageSize   int `xml:"PageSize" json:"PageSize"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	Tags       struct {
		Tag []struct {
			TagKey   string `xml:"TagKey" json:"TagKey"`
			TagValue string `xml:"TagValue" json:"TagValue"`
		} `xml:"Tag" json:"Tag"`
	} `xml:"Tags" json:"Tags"`
}

func DescribeTags(req *DescribeTagsRequest, accessId, accessSecret string) (*DescribeTagsResponse, error) {
	var pResponse DescribeTagsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
