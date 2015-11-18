package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	RpcRequest
	LoadBalancerId string
	ListenerPort   int
	HostId         string
	OwnerAccount   string
}

func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetHostId() string {
	return r.HostId
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeLoadBalancerTCPListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLoadBalancerTCPListenerAttribute")
	r.SetProduct(Product)
}

type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	ListenerPort       int    `xml:"ListenerPort" json:"ListenerPort"`
	BackendServerPort  int    `xml:"BackendServerPort" json:"BackendServerPort"`
	Status             string `xml:"Status" json:"Status"`
	Scheduler          string `xml:"Scheduler" json:"Scheduler"`
	PersistenceTimeout int    `xml:"PersistenceTimeout" json:"PersistenceTimeout"`
	HealthCheck        string `xml:"HealthCheck" json:"HealthCheck"`
	HealthyThreshold   int    `xml:"HealthyThreshold" json:"HealthyThreshold"`
	UnhealthyThreshold int    `xml:"UnhealthyThreshold" json:"UnhealthyThreshold"`
	ConnectTimeout     int    `xml:"ConnectTimeout" json:"ConnectTimeout"`
	ConnectPort        int    `xml:"ConnectPort" json:"ConnectPort"`
	Interval           int    `xml:"Interval" json:"Interval"`
}

func DescribeLoadBalancerTCPListenerAttribute(req *DescribeLoadBalancerTCPListenerAttributeRequest, accessId, accessSecret string) (*DescribeLoadBalancerTCPListenerAttributeResponse, error) {
	var pResponse DescribeLoadBalancerTCPListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
