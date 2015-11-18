package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeRouteTablesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VRouterId            string
	RouteTableId         string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
}

func (r *DescribeRouteTablesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeRouteTablesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeRouteTablesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeRouteTablesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeRouteTablesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeRouteTablesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeRouteTablesRequest) SetVRouterId(value string) {
	r.VRouterId = value
	r.QueryParams.Set("VRouterId", value)
}
func (r *DescribeRouteTablesRequest) GetVRouterId() string {
	return r.VRouterId
}
func (r *DescribeRouteTablesRequest) SetRouteTableId(value string) {
	r.RouteTableId = value
	r.QueryParams.Set("RouteTableId", value)
}
func (r *DescribeRouteTablesRequest) GetRouteTableId() string {
	return r.RouteTableId
}
func (r *DescribeRouteTablesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeRouteTablesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeRouteTablesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeRouteTablesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeRouteTablesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeRouteTablesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeRouteTablesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeRouteTables")
	r.SetProduct(Product)
}

type DescribeRouteTablesResponse struct {
	TotalCount  int `xml:"TotalCount" json:"TotalCount"`
	PageNumber  int `xml:"PageNumber" json:"PageNumber"`
	PageSize    int `xml:"PageSize" json:"PageSize"`
	RouteTables struct {
		RouteTable []struct {
			VRouterId      string `xml:"VRouterId" json:"VRouterId"`
			RouteTableId   string `xml:"RouteTableId" json:"RouteTableId"`
			RouteTableType string `xml:"RouteTableType" json:"RouteTableType"`
			CreationTime   string `xml:"CreationTime" json:"CreationTime"`
			RouteEntrys    struct {
				RouteEntry []struct {
					RouteTableId         string `xml:"RouteTableId" json:"RouteTableId"`
					DestinationCidrBlock string `xml:"DestinationCidrBlock" json:"DestinationCidrBlock"`
					Type                 string `xml:"Type" json:"Type"`
					Status               string `xml:"Status" json:"Status"`
					InstanceId           string `xml:"InstanceId" json:"InstanceId"`
					NextHopType          string `xml:"NextHopType" json:"NextHopType"`
				} `xml:"RouteEntry" json:"RouteEntry"`
			} `xml:"RouteEntrys" json:"RouteEntrys"`
		} `xml:"RouteTable" json:"RouteTable"`
	} `xml:"RouteTables" json:"RouteTables"`
}

func DescribeRouteTables(req *DescribeRouteTablesRequest, accessId, accessSecret string) (*DescribeRouteTablesResponse, error) {
	var pResponse DescribeRouteTablesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
