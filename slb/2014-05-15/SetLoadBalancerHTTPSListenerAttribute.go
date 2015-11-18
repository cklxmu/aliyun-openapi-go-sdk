package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerHTTPSListenerAttributeRequest struct {
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
	ServerCertificateId    string
	MaxConnLimit           int
	OwnerAccount           string
}

func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetXForwardedFor(value string) {
	r.XForwardedFor = value
	r.QueryParams.Set("XForwardedFor", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetXForwardedFor() string {
	return r.XForwardedFor
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetStickySession(value string) {
	r.StickySession = value
	r.QueryParams.Set("StickySession", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetStickySession() string {
	return r.StickySession
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetStickySessionType(value string) {
	r.StickySessionType = value
	r.QueryParams.Set("StickySessionType", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetStickySessionType() string {
	return r.StickySessionType
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetCookieTimeout(value int) {
	r.CookieTimeout = value
	r.QueryParams.Set("CookieTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetCookieTimeout() int {
	return r.CookieTimeout
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetCookie(value string) {
	r.Cookie = value
	r.QueryParams.Set("Cookie", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetCookie() string {
	return r.Cookie
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheck(value string) {
	r.HealthCheck = value
	r.QueryParams.Set("HealthCheck", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheck() string {
	return r.HealthCheck
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckDomain(value string) {
	r.HealthCheckDomain = value
	r.QueryParams.Set("HealthCheckDomain", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckDomain() string {
	return r.HealthCheckDomain
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckURI(value string) {
	r.HealthCheckURI = value
	r.QueryParams.Set("HealthCheckURI", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckURI() string {
	return r.HealthCheckURI
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckTimeout(value int) {
	r.HealthCheckTimeout = value
	r.QueryParams.Set("HealthCheckTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckTimeout() int {
	return r.HealthCheckTimeout
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckInterval(value int) {
	r.HealthCheckInterval = value
	r.QueryParams.Set("HealthCheckInterval", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckInterval() int {
	return r.HealthCheckInterval
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetHealthCheckHttpCode(value string) {
	r.HealthCheckHttpCode = value
	r.QueryParams.Set("HealthCheckHttpCode", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetHealthCheckHttpCode() string {
	return r.HealthCheckHttpCode
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetServerCertificateId(value string) {
	r.ServerCertificateId = value
	r.QueryParams.Set("ServerCertificateId", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetServerCertificateId() string {
	return r.ServerCertificateId
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerHTTPSListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerHTTPSListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerHTTPSListenerAttribute")
	r.SetProduct(Product)
}

type SetLoadBalancerHTTPSListenerAttributeResponse struct {
}

func SetLoadBalancerHTTPSListenerAttribute(req *SetLoadBalancerHTTPSListenerAttributeRequest, accessId, accessSecret string) (*SetLoadBalancerHTTPSListenerAttributeResponse, error) {
	var pResponse SetLoadBalancerHTTPSListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
