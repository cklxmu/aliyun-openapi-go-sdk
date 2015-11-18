package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSecurityGroupsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VpcId                string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
	SecurityGroupIds     string
}

func (r *DescribeSecurityGroupsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSecurityGroupsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSecurityGroupsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSecurityGroupsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSecurityGroupsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSecurityGroupsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSecurityGroupsRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DescribeSecurityGroupsRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DescribeSecurityGroupsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSecurityGroupsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSecurityGroupsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSecurityGroupsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSecurityGroupsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSecurityGroupsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeSecurityGroupsRequest) SetSecurityGroupIds(value string) {
	r.SecurityGroupIds = value
	r.QueryParams.Set("SecurityGroupIds", value)
}
func (r *DescribeSecurityGroupsRequest) GetSecurityGroupIds() string {
	return r.SecurityGroupIds
}

func (r *DescribeSecurityGroupsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSecurityGroups")
	r.SetProduct(Product)
}

type DescribeSecurityGroupsResponse struct {
	RegionId       string `xml:"RegionId" json:"RegionId"`
	TotalCount     int    `xml:"TotalCount" json:"TotalCount"`
	PageNumber     int    `xml:"PageNumber" json:"PageNumber"`
	PageSize       int    `xml:"PageSize" json:"PageSize"`
	SecurityGroups struct {
		SecurityGroup []struct {
			SecurityGroupId   string `xml:"SecurityGroupId" json:"SecurityGroupId"`
			Description       string `xml:"Description" json:"Description"`
			SecurityGroupName string `xml:"SecurityGroupName" json:"SecurityGroupName"`
			VpcId             string `xml:"VpcId" json:"VpcId"`
			CreationTime      string `xml:"CreationTime" json:"CreationTime"`
		} `xml:"SecurityGroup" json:"SecurityGroup"`
	} `xml:"SecurityGroups" json:"SecurityGroups"`
}

func DescribeSecurityGroups(req *DescribeSecurityGroupsRequest, accessId, accessSecret string) (*DescribeSecurityGroupsResponse, error) {
	var pResponse DescribeSecurityGroupsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
