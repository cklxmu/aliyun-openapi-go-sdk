package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeVRoutersRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VRouterId            string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
}

func (r *DescribeVRoutersRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeVRoutersRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeVRoutersRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeVRoutersRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeVRoutersRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeVRoutersRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeVRoutersRequest) SetVRouterId(value string) {
	r.VRouterId = value
	r.QueryParams.Set("VRouterId", value)
}
func (r *DescribeVRoutersRequest) GetVRouterId() string {
	return r.VRouterId
}
func (r *DescribeVRoutersRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeVRoutersRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeVRoutersRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeVRoutersRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeVRoutersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeVRoutersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeVRoutersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeVRouters")
	r.SetProduct(Product)
}

type DescribeVRoutersResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	VRouters   struct {
		VRouter []struct {
			RegionId      string `xml:"RegionId" json:"RegionId"`
			VpcId         string `xml:"VpcId" json:"VpcId"`
			VRouterName   string `xml:"VRouterName" json:"VRouterName"`
			Description   string `xml:"Description" json:"Description"`
			VRouterId     string `xml:"VRouterId" json:"VRouterId"`
			CreationTime  string `xml:"CreationTime" json:"CreationTime"`
			RouteTableIds struct {
				RouteTableId []string `xml:"RouteTableId" json:"RouteTableId"`
			} `xml:"RouteTableIds" json:"RouteTableIds"`
		} `xml:"VRouter" json:"VRouter"`
	} `xml:"VRouters" json:"VRouters"`
}

func DescribeVRouters(req *DescribeVRoutersRequest, accessId, accessSecret string) (*DescribeVRoutersResponse, error) {
	var pResponse DescribeVRoutersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
