package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerHTTPListenerRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	LoadBalancerId         string
	Bandwidth              int
	ListenerPort           int
	BackendServerPort      int
	XForwardedFor          string
	Scheduler              string
	StickySession          string
	StickySessionType      string
	CookieTimeout          int
	Cookie                 string
	HealthCheck            string
	HealthCheckDomain      string
	HealthCheckURI         string
	HealthyThreshold       int
	UnhealthyThreshold     int
	HealthCheckTimeout     int
	HealthCheckConnectPort int
	HealthCheckInterval    int
	HealthCheckHttpCode    string
	MaxConnLimit           int
	OwnerAccount           string
}

func (r *CreateLoadBalancerHTTPListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetBandwidth() int {
	return r.Bandwidth
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
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckDomain(value string) {
	r.HealthCheckDomain = value
	r.QueryParams.Set("HealthCheckDomain", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckDomain() string {
	return r.HealthCheckDomain
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckURI(value string) {
	r.HealthCheckURI = value
	r.QueryParams.Set("HealthCheckURI", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckURI() string {
	return r.HealthCheckURI
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
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckInterval(value int) {
	r.HealthCheckInterval = value
	r.QueryParams.Set("HealthCheckInterval", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckInterval() int {
	return r.HealthCheckInterval
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetHealthCheckHttpCode(value string) {
	r.HealthCheckHttpCode = value
	r.QueryParams.Set("HealthCheckHttpCode", value)
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetHealthCheckHttpCode() string {
	return r.HealthCheckHttpCode
}
func (r *CreateLoadBalancerHTTPListenerRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPListenerRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
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
