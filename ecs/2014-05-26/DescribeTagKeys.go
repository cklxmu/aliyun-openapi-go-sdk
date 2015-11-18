package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeTagKeysRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageSize             int
	PageNumber           int
	ResourceType         string
	ResourceId           string
}

func (r *DescribeTagKeysRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeTagKeysRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeTagKeysRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeTagKeysRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeTagKeysRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeTagKeysRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeTagKeysRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeTagKeysRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeTagKeysRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeTagKeysRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeTagKeysRequest) SetResourceType(value string) {
	r.ResourceType = value
	r.QueryParams.Set("ResourceType", value)
}
func (r *DescribeTagKeysRequest) GetResourceType() string {
	return r.ResourceType
}
func (r *DescribeTagKeysRequest) SetResourceId(value string) {
	r.ResourceId = value
	r.QueryParams.Set("ResourceId", value)
}
func (r *DescribeTagKeysRequest) GetResourceId() string {
	return r.ResourceId
}

func (r *DescribeTagKeysRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeTagKeys")
	r.SetProduct(Product)
}

type DescribeTagKeysResponse struct {
	PageSize   int `xml:"PageSize" json:"PageSize"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	TagKeys    struct {
		TagKey []string `xml:"TagKey" json:"TagKey"`
	} `xml:"TagKeys" json:"TagKeys"`
}

func DescribeTagKeys(req *DescribeTagKeysRequest, accessId, accessSecret string) (*DescribeTagKeysResponse, error) {
	var pResponse DescribeTagKeysResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
