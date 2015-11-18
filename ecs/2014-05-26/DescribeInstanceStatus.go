package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstanceStatusRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ZoneId               string
	ClusterId            string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
}

func (r *DescribeInstanceStatusRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceStatusRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstanceStatusRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstanceStatusRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstanceStatusRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstanceStatusRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstanceStatusRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribeInstanceStatusRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *DescribeInstanceStatusRequest) SetClusterId(value string) {
	r.ClusterId = value
	r.QueryParams.Set("ClusterId", value)
}
func (r *DescribeInstanceStatusRequest) GetClusterId() string {
	return r.ClusterId
}
func (r *DescribeInstanceStatusRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeInstanceStatusRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeInstanceStatusRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeInstanceStatusRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeInstanceStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstanceStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeInstanceStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstanceStatus")
	r.SetProduct(Product)
}

type DescribeInstanceStatusResponse struct {
	TotalCount       int `xml:"TotalCount" json:"TotalCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageSize         int `xml:"PageSize" json:"PageSize"`
	InstanceStatuses struct {
		InstanceStatus []struct {
			InstanceId string `xml:"InstanceId" json:"InstanceId"`
			Status     string `xml:"Status" json:"Status"`
		} `xml:"InstanceStatus" json:"InstanceStatus"`
	} `xml:"InstanceStatuses" json:"InstanceStatuses"`
}

func DescribeInstanceStatus(req *DescribeInstanceStatusRequest, accessId, accessSecret string) (*DescribeInstanceStatusResponse, error) {
	var pResponse DescribeInstanceStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
