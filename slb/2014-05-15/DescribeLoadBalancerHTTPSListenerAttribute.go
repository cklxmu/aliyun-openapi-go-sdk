package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeLoadBalancerHTTPSListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLoadBalancerHTTPSListenerAttribute")
	r.SetProduct(Product)
}

type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	ListenerPort           int    `xml:"ListenerPort" json:"ListenerPort"`
	BackendServerPort      int    `xml:"BackendServerPort" json:"BackendServerPort"`
	Bandwidth              int    `xml:"Bandwidth" json:"Bandwidth"`
	Status                 string `xml:"Status" json:"Status"`
	XForwardedFor          string `xml:"XForwardedFor" json:"XForwardedFor"`
	Scheduler              string `xml:"Scheduler" json:"Scheduler"`
	StickySession          string `xml:"StickySession" json:"StickySession"`
	StickySessionType      string `xml:"StickySessionType" json:"StickySessionType"`
	CookieTimeout          int    `xml:"CookieTimeout" json:"CookieTimeout"`
	Cookie                 string `xml:"Cookie" json:"Cookie"`
	HealthCheck            string `xml:"HealthCheck" json:"HealthCheck"`
	HealthCheckDomain      string `xml:"HealthCheckDomain" json:"HealthCheckDomain"`
	HealthCheckURI         string `xml:"HealthCheckURI" json:"HealthCheckURI"`
	HealthyThreshold       int    `xml:"HealthyThreshold" json:"HealthyThreshold"`
	UnhealthyThreshold     int    `xml:"UnhealthyThreshold" json:"UnhealthyThreshold"`
	HealthCheckTimeout     int    `xml:"HealthCheckTimeout" json:"HealthCheckTimeout"`
	HealthCheckInterval    int    `xml:"HealthCheckInterval" json:"HealthCheckInterval"`
	HealthCheckConnectPort int    `xml:"HealthCheckConnectPort" json:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string `xml:"HealthCheckHttpCode" json:"HealthCheckHttpCode"`
	ServerCertificateId    string `xml:"ServerCertificateId" json:"ServerCertificateId"`
	MaxConnLimit           int    `xml:"MaxConnLimit" json:"MaxConnLimit"`
}

func DescribeLoadBalancerHTTPSListenerAttribute(req *DescribeLoadBalancerHTTPSListenerAttributeRequest, accessId, accessSecret string) (*DescribeLoadBalancerHTTPSListenerAttributeResponse, error) {
	var pResponse DescribeLoadBalancerHTTPSListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
