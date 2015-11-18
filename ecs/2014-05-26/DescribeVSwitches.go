package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeVSwitchesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VpcId                string
	VSwitchId            string
	ZoneId               string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
}

func (r *DescribeVSwitchesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeVSwitchesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeVSwitchesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeVSwitchesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeVSwitchesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeVSwitchesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeVSwitchesRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DescribeVSwitchesRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DescribeVSwitchesRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *DescribeVSwitchesRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *DescribeVSwitchesRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribeVSwitchesRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *DescribeVSwitchesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeVSwitchesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeVSwitchesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeVSwitchesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeVSwitchesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeVSwitchesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeVSwitchesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeVSwitches")
	r.SetProduct(Product)
}

type DescribeVSwitchesResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	VSwitches  struct {
		VSwitch []struct {
			VSwitchId               string `xml:"VSwitchId" json:"VSwitchId"`
			VpcId                   string `xml:"VpcId" json:"VpcId"`
			Status                  string `xml:"Status" json:"Status"`
			CidrBlock               string `xml:"CidrBlock" json:"CidrBlock"`
			ZoneId                  string `xml:"ZoneId" json:"ZoneId"`
			AvailableIpAddressCount int    `xml:"AvailableIpAddressCount" json:"AvailableIpAddressCount"`
			Description             string `xml:"Description" json:"Description"`
			VSwitchName             string `xml:"VSwitchName" json:"VSwitchName"`
			CreationTime            string `xml:"CreationTime" json:"CreationTime"`
		} `xml:"VSwitch" json:"VSwitch"`
	} `xml:"VSwitches" json:"VSwitches"`
}

func DescribeVSwitches(req *DescribeVSwitchesRequest, accessId, accessSecret string) (*DescribeVSwitchesResponse, error) {
	var pResponse DescribeVSwitchesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
