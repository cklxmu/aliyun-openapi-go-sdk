package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerHTTPListenerRequest struct {
	RpcRequest
	LoadBalancerId     string
	ListenerPort       int
	BackendServerPort  int
	ListenerStatus     string
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

func (r *CreateLoadBalancerHTTPListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetBackendServerPort(value int) {
	r.BackendServerPort = value
	r.QueryParams.Set("BackendServerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetBackendServerPort() int {
	return r.BackendServerPort
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetListenerStatus(value string) {
	r.ListenerStatus = value
	r.QueryParams.Set("ListenerStatus", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetListenerStatus() string {
	return r.ListenerStatus
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetXForwardedFor(value string) {
	r.XForwardedFor = value
	r.QueryParams.Set("XForwardedFor", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetXForwardedFor() string {
	return r.XForwardedFor
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetStickySession(value string) {
	r.StickySession = value
	r.QueryParams.Set("StickySession", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetStickySession() string {
	return r.StickySession
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetStickySessionType(value string) {
	r.StickySessionType = value
	r.QueryParams.Set("StickySessionType", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetStickySessionType() string {
	return r.StickySessionType
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetCookieTimeout(value int) {
	r.CookieTimeout = value
	r.QueryParams.Set("CookieTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetCookieTimeout() int {
	return r.CookieTimeout
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetCookie(value string) {
	r.Cookie = value
	r.QueryParams.Set("Cookie", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetCookie() string {
	return r.Cookie
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheck(value string) {
	r.HealthCheck = value
	r.QueryParams.Set("HealthCheck", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheck() string {
	return r.HealthCheck
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetDomain(value string) {
	r.Domain = value
	r.QueryParams.Set("Domain", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetDomain() string {
	return r.Domain
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetURI(value string) {
	r.URI = value
	r.QueryParams.Set("URI", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetURI() string {
	return r.URI
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckTimeout(value int) {
	r.HealthCheckTimeout = value
	r.QueryParams.Set("HealthCheckTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckTimeout() int {
	return r.HealthCheckTimeout
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetInterval(value int) {
	r.Interval = value
	r.QueryParams.Set("Interval", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetInterval() int {
	return r.Interval
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHostId() string {
	return r.HostId
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateLoadBalancerHTTPListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoadBalancerHTTPListener")
	r.SetProduct(Product)
}

type CreateLoadBalancerHTTPListenerResponse struct {
}

func CreateLoadBalancerHTTPListener(req *CreateLoadBalancerHTTPListenerRequest, accessId, accessSecret string) (*CreateLoadBalancerHTTPListenerResponse, error) {
	var pResponse CreateLoadBalancerHTTPListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
