package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	RpcRequest
	LoadBalancerId     string
	ListenerPort       int
	XForwardedFor      string
	Scheduler          string
	StickySession      string
	StickySessionType  string
	CookieTimeout      int
	Cookie             string
	HealthCheck        string
	Domain             string
	URI                string
	HealthyThreshold   int
	UnhealthyThreshold int
	HealthCheckTimeout int
	Interval           int
	HostId             string
	OwnerAccount       string
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetXForwardedFor(value string) {
	r.XForwardedFor = value
	r.QueryParams.Set("XForwardedFor", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetXForwardedFor() string {
	return r.XForwardedFor
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySession(value string) {
	r.StickySession = value
	r.QueryParams.Set("StickySession", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetStickySession() string {
	return r.StickySession
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetStickySessionType(value string) {
	r.StickySessionType = value
	r.QueryParams.Set("StickySessionType", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetStickySessionType() string {
	return r.StickySessionType
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetCookieTimeout(value int) {
	r.CookieTimeout = value
	r.QueryParams.Set("CookieTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetCookieTimeout() int {
	return r.CookieTimeout
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetCookie(value string) {
	r.Cookie = value
	r.QueryParams.Set("Cookie", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetCookie() string {
	return r.Cookie
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheck(value string) {
	r.HealthCheck = value
	r.QueryParams.Set("HealthCheck", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheck() string {
	return r.HealthCheck
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetDomain(value string) {
	r.Domain = value
	r.QueryParams.Set("Domain", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetDomain() string {
	return r.Domain
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetURI(value string) {
	r.URI = value
	r.QueryParams.Set("URI", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetURI() string {
	return r.URI
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckTimeout(value int) {
	r.HealthCheckTimeout = value
	r.QueryParams.Set("HealthCheckTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckTimeout() int {
	return r.HealthCheckTimeout
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetInterval(value int) {
	r.Interval = value
	r.QueryParams.Set("Interval", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetInterval() int {
	return r.Interval
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHostId() string {
	return r.HostId
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerHTTPListenerAttribute")
	r.SetProduct(Product)
}

type SetLoadBalancerHTTPListenerAttributeResponse struct {
}

func SetLoadBalancerHTTPListenerAttribute(req *SetLoadBalancerHTTPListenerAttributeRequest, accessId, accessSecret string) (*SetLoadBalancerHTTPListenerAttributeResponse, error) {
	var pResponse SetLoadBalancerHTTPListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
