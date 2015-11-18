package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeResourceByTagsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageSize             int
	PageNumber           int
	ResourceType         string
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

func (r *DescribeResourceByTagsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeResourceByTagsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeResourceByTagsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeResourceByTagsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeResourceByTagsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeResourceByTagsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeResourceByTagsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeResourceByTagsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeResourceByTagsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeResourceByTagsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeResourceByTagsRequest) SetResourceType(value string) {
	r.ResourceType = value
	r.QueryParams.Set("ResourceType", value)
}
func (r *DescribeResourceByTagsRequest) GetResourceType() string {
	return r.ResourceType
}
func (r *DescribeResourceByTagsRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *DescribeResourceByTagsRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *DescribeResourceByTagsRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *DescribeResourceByTagsRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *DescribeResourceByTagsRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *DescribeResourceByTagsRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *DescribeResourceByTagsRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *DescribeResourceByTagsRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *DescribeResourceByTagsRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *DescribeResourceByTagsRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *DescribeResourceByTagsRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *DescribeResourceByTagsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeResourceByTags")
	r.SetProduct(Product)
}

type DescribeResourceByTagsResponse struct {
	PageSize   int `xml:"PageSize" json:"PageSize"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	Resources  struct {
		Resource []struct {
			ResourceId   string `xml:"ResourceId" json:"ResourceId"`
			ResourceType string `xml:"ResourceType" json:"ResourceType"`
			RegionId     string `xml:"RegionId" json:"RegionId"`
		} `xml:"Resource" json:"Resource"`
	} `xml:"Resources" json:"Resources"`
}

func DescribeResourceByTags(req *DescribeResourceByTagsRequest, accessId, accessSecret string) (*DescribeResourceByTagsResponse, error) {
	var pResponse DescribeResourceByTagsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
