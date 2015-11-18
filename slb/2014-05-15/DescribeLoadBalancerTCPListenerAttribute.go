package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
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
	ListenerPort              int    `xml:"ListenerPort" json:"ListenerPort"`
	BackendServerPort         int    `xml:"BackendServerPort" json:"BackendServerPort"`
	Status                    string `xml:"Status" json:"Status"`
	Bandwidth                 int    `xml:"Bandwidth" json:"Bandwidth"`
	Scheduler                 string `xml:"Scheduler" json:"Scheduler"`
	SynProxy                  string `xml:"SynProxy" json:"SynProxy"`
	PersistenceTimeout        int    `xml:"PersistenceTimeout" json:"PersistenceTimeout"`
	HealthCheck               string `xml:"HealthCheck" json:"HealthCheck"`
	HealthyThreshold          int    `xml:"HealthyThreshold" json:"HealthyThreshold"`
	UnhealthyThreshold        int    `xml:"UnhealthyThreshold" json:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `xml:"HealthCheckConnectTimeout" json:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int    `xml:"HealthCheckConnectPort" json:"HealthCheckConnectPort"`
	HealthCheckInterval       int    `xml:"HealthCheckInterval" json:"HealthCheckInterval"`
	HealthCheckHttpCode       string `xml:"HealthCheckHttpCode" json:"HealthCheckHttpCode"`
	HealthCheckDomain         string `xml:"HealthCheckDomain" json:"HealthCheckDomain"`
	HealthCheckURI            string `xml:"HealthCheckURI" json:"HealthCheckURI"`
	HealthCheckType           string `xml:"HealthCheckType" json:"HealthCheckType"`
	MaxConnLimit              int    `xml:"MaxConnLimit" json:"MaxConnLimit"`
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
