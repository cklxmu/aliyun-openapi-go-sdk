package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeImageSharePermissionRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ImageId              string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
}

func (r *DescribeImageSharePermissionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeImageSharePermissionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeImageSharePermissionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeImageSharePermissionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeImageSharePermissionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeImageSharePermissionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeImageSharePermissionRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *DescribeImageSharePermissionRequest) GetImageId() string {
	return r.ImageId
}
func (r *DescribeImageSharePermissionRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeImageSharePermissionRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeImageSharePermissionRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeImageSharePermissionRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeImageSharePermissionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeImageSharePermissionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeImageSharePermissionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeImageSharePermission")
	r.SetProduct(Product)
}

type DescribeImageSharePermissionResponse struct {
	RegionId    string `xml:"RegionId" json:"RegionId"`
	TotalCount  int    `xml:"TotalCount" json:"TotalCount"`
	PageNumber  int    `xml:"PageNumber" json:"PageNumber"`
	PageSize    int    `xml:"PageSize" json:"PageSize"`
	ImageId     string `xml:"ImageId" json:"ImageId"`
	ShareGroups struct {
		ShareGroup []struct {
			Group string `xml:"Group" json:"Group"`
		} `xml:"ShareGroup" json:"ShareGroup"`
	} `xml:"ShareGroups" json:"ShareGroups"`
	Accounts struct {
		Account []struct {
			AliyunId string `xml:"AliyunId" json:"AliyunId"`
		} `xml:"Account" json:"Account"`
	} `xml:"Accounts" json:"Accounts"`
}

func DescribeImageSharePermission(req *DescribeImageSharePermissionRequest, accessId, accessSecret string) (*DescribeImageSharePermissionResponse, error) {
	var pResponse DescribeImageSharePermissionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
