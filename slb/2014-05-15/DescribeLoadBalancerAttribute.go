package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancerAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	OwnerAccount         string
}

func (r *DescribeLoadBalancerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLoadBalancerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLoadBalancerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLoadBalancerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeLoadBalancerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeLoadBalancerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeLoadBalancerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLoadBalancerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeLoadBalancerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLoadBalancerAttribute")
	r.SetProduct(Product)
}

type DescribeLoadBalancerAttributeResponse struct {
	LoadBalancerId     string `xml:"LoadBalancerId" json:"LoadBalancerId"`
	LoadBalancerName   string `xml:"LoadBalancerName" json:"LoadBalancerName"`
	LoadBalancerStatus string `xml:"LoadBalancerStatus" json:"LoadBalancerStatus"`
	RegionId           string `xml:"RegionId" json:"RegionId"`
	RegionIdAlias      string `xml:"RegionIdAlias" json:"RegionIdAlias"`
	Address            string `xml:"Address" json:"Address"`
	AddressType        string `xml:"AddressType" json:"AddressType"`
	VpcId              string `xml:"VpcId" json:"VpcId"`
	VSwitchId          string `xml:"VSwitchId" json:"VSwitchId"`
	NetworkType        string `xml:"NetworkType" json:"NetworkType"`
	InternetChargeType string `xml:"InternetChargeType" json:"InternetChargeType"`
	Bandwidth          int    `xml:"Bandwidth" json:"Bandwidth"`
	SysBandwidth       int    `xml:"SysBandwidth" json:"SysBandwidth"`
	CreateTime         string `xml:"CreateTime" json:"CreateTime"`
	CreateTimeStamp    int    `xml:"CreateTimeStamp" json:"CreateTimeStamp"`
	MasterZoneId       string `xml:"MasterZoneId" json:"MasterZoneId"`
	SlaveZoneId        string `xml:"SlaveZoneId" json:"SlaveZoneId"`
	MaxConnLimit       int    `xml:"MaxConnLimit" json:"MaxConnLimit"`
	SecurityStatus     string `xml:"SecurityStatus" json:"SecurityStatus"`
	ListenerPorts      struct {
		ListenerPort []string `xml:"ListenerPort" json:"ListenerPort"`
	} `xml:"ListenerPorts" json:"ListenerPorts"`
	ListenerPortsAndProtocal struct {
		ListenerPortAndProtocal []struct {
			ListenerPort     int    `xml:"ListenerPort" json:"ListenerPort"`
			ListenerProtocal string `xml:"ListenerProtocal" json:"ListenerProtocal"`
		} `xml:"ListenerPortAndProtocal" json:"ListenerPortAndProtocal"`
	} `xml:"ListenerPortsAndProtocal" json:"ListenerPortsAndProtocal"`
	ListenerPortsAndProtocol struct {
		ListenerPortAndProtocol []struct {
			ListenerPort     int    `xml:"ListenerPort" json:"ListenerPort"`
			ListenerProtocol string `xml:"ListenerProtocol" json:"ListenerProtocol"`
		} `xml:"ListenerPortAndProtocol" json:"ListenerPortAndProtocol"`
	} `xml:"ListenerPortsAndProtocol" json:"ListenerPortsAndProtocol"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `xml:"ServerId" json:"ServerId"`
			Weight   int    `xml:"Weight" json:"Weight"`
		} `xml:"BackendServer" json:"BackendServer"`
	} `xml:"BackendServers" json:"BackendServers"`
}

func DescribeLoadBalancerAttribute(req *DescribeLoadBalancerAttributeRequest, accessId, accessSecret string) (*DescribeLoadBalancerAttributeResponse, error) {
	var pResponse DescribeLoadBalancerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
