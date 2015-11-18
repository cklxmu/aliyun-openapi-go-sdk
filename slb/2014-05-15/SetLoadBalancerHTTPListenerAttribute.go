package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerHTTPListenerAttributeRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	LoadBalancerId         string
	ListenerPort           int
	Bandwidth              int
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
	HealthCheckInterval    int
	HealthCheckConnectPort int
	HealthCheckHttpCode    string
	MaxConnLimit           int
	OwnerAccount           string
}

func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
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
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetBandwidth() int {
	return r.Bandwidth
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
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckDomain(value string) {
	r.HealthCheckDomain = value
	r.QueryParams.Set("HealthCheckDomain", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckDomain() string {
	return r.HealthCheckDomain
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckURI(value string) {
	r.HealthCheckURI = value
	r.QueryParams.Set("HealthCheckURI", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckURI() string {
	return r.HealthCheckURI
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
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckInterval(value int) {
	r.HealthCheckInterval = value
	r.QueryParams.Set("HealthCheckInterval", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckInterval() int {
	return r.HealthCheckInterval
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetHealthCheckHttpCode(value string) {
	r.HealthCheckHttpCode = value
	r.QueryParams.Set("HealthCheckHttpCode", value)
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetHealthCheckHttpCode() string {
	return r.HealthCheckHttpCode
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPListenerAttributeRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
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
