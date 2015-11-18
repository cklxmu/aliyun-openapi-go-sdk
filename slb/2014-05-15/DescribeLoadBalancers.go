package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancersRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ServerId             string
	LoadBalancerId       string
	AddressType          string
	InternetChargeType   string
	VpcId                string
	VSwitchId            string
	NetworkType          string
	Address              string
	SecurityStatus       string
	MasterZoneId         string
	SlaveZoneId          string
	OwnerAccount         string
}

func (r *DescribeLoadBalancersRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancersRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLoadBalancersRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLoadBalancersRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLoadBalancersRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancersRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeLoadBalancersRequest) SetServerId(value string) {
	r.ServerId = value
	r.QueryParams.Set("ServerId", value)
}
func (r *DescribeLoadBalancersRequest) GetServerId() string {
	return r.ServerId
}
func (r *DescribeLoadBalancersRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeLoadBalancersRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeLoadBalancersRequest) SetAddressType(value string) {
	r.AddressType = value
	r.QueryParams.Set("AddressType", value)
}
func (r *DescribeLoadBalancersRequest) GetAddressType() string {
	return r.AddressType
}
func (r *DescribeLoadBalancersRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *DescribeLoadBalancersRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *DescribeLoadBalancersRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DescribeLoadBalancersRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DescribeLoadBalancersRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *DescribeLoadBalancersRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *DescribeLoadBalancersRequest) SetNetworkType(value string) {
	r.NetworkType = value
	r.QueryParams.Set("NetworkType", value)
}
func (r *DescribeLoadBalancersRequest) GetNetworkType() string {
	return r.NetworkType
}
func (r *DescribeLoadBalancersRequest) SetAddress(value string) {
	r.Address = value
	r.QueryParams.Set("Address", value)
}
func (r *DescribeLoadBalancersRequest) GetAddress() string {
	return r.Address
}
func (r *DescribeLoadBalancersRequest) SetSecurityStatus(value string) {
	r.SecurityStatus = value
	r.QueryParams.Set("SecurityStatus", value)
}
func (r *DescribeLoadBalancersRequest) GetSecurityStatus() string {
	return r.SecurityStatus
}
func (r *DescribeLoadBalancersRequest) SetMasterZoneId(value string) {
	r.MasterZoneId = value
	r.QueryParams.Set("MasterZoneId", value)
}
func (r *DescribeLoadBalancersRequest) GetMasterZoneId() string {
	return r.MasterZoneId
}
func (r *DescribeLoadBalancersRequest) SetSlaveZoneId(value string) {
	r.SlaveZoneId = value
	r.QueryParams.Set("SlaveZoneId", value)
}
func (r *DescribeLoadBalancersRequest) GetSlaveZoneId() string {
	return r.SlaveZoneId
}
func (r *DescribeLoadBalancersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLoadBalancersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeLoadBalancersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLoadBalancers")
	r.SetProduct(Product)
}

type DescribeLoadBalancersResponse struct {
	LoadBalancers struct {
		LoadBalancer []struct {
			LoadBalancerId     string `xml:"LoadBalancerId" json:"LoadBalancerId"`
			LoadBalancerName   string `xml:"LoadBalancerName" json:"LoadBalancerName"`
			LoadBalancerStatus string `xml:"LoadBalancerStatus" json:"LoadBalancerStatus"`
			Address            string `xml:"Address" json:"Address"`
			AddressType        string `xml:"AddressType" json:"AddressType"`
			RegionId           string `xml:"RegionId" json:"RegionId"`
			RegionIdAlias      string `xml:"RegionIdAlias" json:"RegionIdAlias"`
			VSwitchId          string `xml:"VSwitchId" json:"VSwitchId"`
			VpcId              string `xml:"VpcId" json:"VpcId"`
			NetworkType        string `xml:"NetworkType" json:"NetworkType"`
			MasterZoneId       string `xml:"MasterZoneId" json:"MasterZoneId"`
			SlaveZoneId        string `xml:"SlaveZoneId" json:"SlaveZoneId"`
			MaxConnLimit       int    `xml:"MaxConnLimit" json:"MaxConnLimit"`
			SecurityStatus     string `xml:"SecurityStatus" json:"SecurityStatus"`
			InternetChargeType string `xml:"InternetChargeType" json:"InternetChargeType"`
			SysBandwidth       int    `xml:"SysBandwidth" json:"SysBandwidth"`
		} `xml:"LoadBalancer" json:"LoadBalancer"`
	} `xml:"LoadBalancers" json:"LoadBalancers"`
}

func DescribeLoadBalancers(req *DescribeLoadBalancersRequest, accessId, accessSecret string) (*DescribeLoadBalancersResponse, error) {
	var pResponse DescribeLoadBalancersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
