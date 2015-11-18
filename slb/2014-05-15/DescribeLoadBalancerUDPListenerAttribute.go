package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *DescribeLoadBalancerUDPListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLoadBalancerUDPListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeLoadBalancerUDPListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLoadBalancerUDPListenerAttribute")
	r.SetProduct(Product)
}

type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	ListenerPort              int    `xml:"ListenerPort" json:"ListenerPort"`
	BackendServerPort         int    `xml:"BackendServerPort" json:"BackendServerPort"`
	Status                    string `xml:"Status" json:"Status"`
	Bandwidth                 int    `xml:"Bandwidth" json:"Bandwidth"`
	Scheduler                 string `xml:"Scheduler" json:"Scheduler"`
	PersistenceTimeout        int    `xml:"PersistenceTimeout" json:"PersistenceTimeout"`
	HealthCheck               string `xml:"HealthCheck" json:"HealthCheck"`
	HealthyThreshold          int    `xml:"HealthyThreshold" json:"HealthyThreshold"`
	UnhealthyThreshold        int    `xml:"UnhealthyThreshold" json:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `xml:"HealthCheckConnectTimeout" json:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int    `xml:"HealthCheckConnectPort" json:"HealthCheckConnectPort"`
	HealthCheckInterval       int    `xml:"HealthCheckInterval" json:"HealthCheckInterval"`
	MaxConnLimit              int    `xml:"MaxConnLimit" json:"MaxConnLimit"`
}

func DescribeLoadBalancerUDPListenerAttribute(req *DescribeLoadBalancerUDPListenerAttributeRequest, accessId, accessSecret string) (*DescribeLoadBalancerUDPListenerAttributeResponse, error) {
	var pResponse DescribeLoadBalancerUDPListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
