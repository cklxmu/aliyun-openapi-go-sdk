package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	RpcRequest
	LoadBalancerId string
	ListenerPort   int
	HostId         string
	OwnerAccount   string
}

func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) GetHostId() string {
	return r.HostId
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeLoadBalancerHTTPListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeLoadBalancerHTTPListenerAttribute")
	r.SetProduct(Product)
}

type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	ListenerPort         int    `xml:"ListenerPort" json:"ListenerPort"`
	BackendServerPort    int    `xml:"BackendServerPort" json:"BackendServerPort"`
	Status               string `xml:"Status" json:"Status"`
	XForwardedFor        string `xml:"XForwardedFor" json:"XForwardedFor"`
	Scheduler            string `xml:"Scheduler" json:"Scheduler"`
	StickySession        string `xml:"StickySession" json:"StickySession"`
	StickySessionapiType string `xml:"StickySessionapiType" json:"StickySessionapiType"`
	CookieTimeout        int    `xml:"CookieTimeout" json:"CookieTimeout"`
	Cookie               string `xml:"Cookie" json:"Cookie"`
	HealthCheck          string `xml:"HealthCheck" json:"HealthCheck"`
	Domain               string `xml:"Domain" json:"Domain"`
	URI                  string `xml:"URI" json:"URI"`
	HealthyThreshold     int    `xml:"HealthyThreshold" json:"HealthyThreshold"`
	UnhealthyThreshold   int    `xml:"UnhealthyThreshold" json:"UnhealthyThreshold"`
	HealthCheckTimeout   int    `xml:"HealthCheckTimeout" json:"HealthCheckTimeout"`
	Interval             int    `xml:"Interval" json:"Interval"`
}

func DescribeLoadBalancerHTTPListenerAttribute(req *DescribeLoadBalancerHTTPListenerAttributeRequest, accessId, accessSecret string) (*DescribeLoadBalancerHTTPListenerAttributeResponse, error) {
	var pResponse DescribeLoadBalancerHTTPListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
