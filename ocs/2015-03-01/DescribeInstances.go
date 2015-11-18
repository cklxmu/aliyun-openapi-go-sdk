package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceIds          string
	InstanceStatus       string
	PageNumber           int
	PageSize             int
	NetworkType          string
	VpcId                string
	VSwitchId            string
	PrivateIpAddresses   string
}

func (r *DescribeInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeInstancesRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *DescribeInstancesRequest) GetInstanceIds() string {
	return r.InstanceIds
}
func (r *DescribeInstancesRequest) SetInstanceStatus(value string) {
	r.InstanceStatus = value
	r.QueryParams.Set("InstanceStatus", value)
}
func (r *DescribeInstancesRequest) GetInstanceStatus() string {
	return r.InstanceStatus
}
func (r *DescribeInstancesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeInstancesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeInstancesRequest) SetNetworkType(value string) {
	r.NetworkType = value
	r.QueryParams.Set("NetworkType", value)
}
func (r *DescribeInstancesRequest) GetNetworkType() string {
	return r.NetworkType
}
func (r *DescribeInstancesRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DescribeInstancesRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DescribeInstancesRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *DescribeInstancesRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *DescribeInstancesRequest) SetPrivateIpAddresses(value string) {
	r.PrivateIpAddresses = value
	r.QueryParams.Set("PrivateIpAddresses", value)
}
func (r *DescribeInstancesRequest) GetPrivateIpAddresses() string {
	return r.PrivateIpAddresses
}

func (r *DescribeInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstances")
	r.SetProduct(Product)
}

type DescribeInstancesResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Instances  struct {
		OcsInstance []struct {
			InstanceId       string `xml:"InstanceId" json:"InstanceId"`
			InstanceName     string `xml:"InstanceName" json:"InstanceName"`
			ConnectionDomain string `xml:"ConnectionDomain" json:"ConnectionDomain"`
			Port             int    `xml:"Port" json:"Port"`
			UserName         string `xml:"UserName" json:"UserName"`
			InstanceStatus   string `xml:"InstanceStatus" json:"InstanceStatus"`
			Capacity         int    `xml:"Capacity" json:"Capacity"`
			QPS              int    `xml:"QPS" json:"QPS"`
			Bandwidth        int    `xml:"Bandwidth" json:"Bandwidth"`
			Connections      int    `xml:"Connections" json:"Connections"`
			RegionId         string `xml:"RegionId" json:"RegionId"`
			ZoneId           string `xml:"ZoneId" json:"ZoneId"`
			NetworkType      string `xml:"NetworkType" json:"NetworkType"`
			VpcId            string `xml:"VpcId" json:"VpcId"`
			VSwitchId        string `xml:"VSwitchId" json:"VSwitchId"`
			PrivateIpAddress string `xml:"PrivateIpAddress" json:"PrivateIpAddress"`
			CreationTime     string `xml:"CreationTime" json:"CreationTime"`
		} `xml:"OcsInstance" json:"OcsInstance"`
	} `xml:"Instances" json:"Instances"`
}

func DescribeInstances(req *DescribeInstancesRequest, accessId, accessSecret string) (*DescribeInstancesResponse, error) {
	var pResponse DescribeInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
