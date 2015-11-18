package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerHTTPSListenerRequest struct {
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
	ServerCertificateId    string
	MaxConnLimit           int
	OwnerAccount           string
}

func (r *CreateLoadBalancerHTTPSListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetBackendServerPort(value int) {
	r.BackendServerPort = value
	r.QueryParams.Set("BackendServerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetBackendServerPort() int {
	return r.BackendServerPort
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetXForwardedFor(value string) {
	r.XForwardedFor = value
	r.QueryParams.Set("XForwardedFor", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetXForwardedFor() string {
	return r.XForwardedFor
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetStickySession(value string) {
	r.StickySession = value
	r.QueryParams.Set("StickySession", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetStickySession() string {
	return r.StickySession
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetStickySessionType(value string) {
	r.StickySessionType = value
	r.QueryParams.Set("StickySessionType", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetStickySessionType() string {
	return r.StickySessionType
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetCookieTimeout(value int) {
	r.CookieTimeout = value
	r.QueryParams.Set("CookieTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetCookieTimeout() int {
	return r.CookieTimeout
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetCookie(value string) {
	r.Cookie = value
	r.QueryParams.Set("Cookie", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetCookie() string {
	return r.Cookie
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheck(value string) {
	r.HealthCheck = value
	r.QueryParams.Set("HealthCheck", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheck() string {
	return r.HealthCheck
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckDomain(value string) {
	r.HealthCheckDomain = value
	r.QueryParams.Set("HealthCheckDomain", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckDomain() string {
	return r.HealthCheckDomain
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckURI(value string) {
	r.HealthCheckURI = value
	r.QueryParams.Set("HealthCheckURI", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckURI() string {
	return r.HealthCheckURI
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckTimeout(value int) {
	r.HealthCheckTimeout = value
	r.QueryParams.Set("HealthCheckTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckTimeout() int {
	return r.HealthCheckTimeout
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckInterval(value int) {
	r.HealthCheckInterval = value
	r.QueryParams.Set("HealthCheckInterval", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckInterval() int {
	return r.HealthCheckInterval
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetHealthCheckHttpCode(value string) {
	r.HealthCheckHttpCode = value
	r.QueryParams.Set("HealthCheckHttpCode", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetHealthCheckHttpCode() string {
	return r.HealthCheckHttpCode
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetServerCertificateId(value string) {
	r.ServerCertificateId = value
	r.QueryParams.Set("ServerCertificateId", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetServerCertificateId() string {
	return r.ServerCertificateId
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *CreateLoadBalancerHTTPSListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateLoadBalancerHTTPSListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateLoadBalancerHTTPSListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoadBalancerHTTPSListener")
	r.SetProduct(Product)
}

type CreateLoadBalancerHTTPSListenerResponse struct {
}

func CreateLoadBalancerHTTPSListener(req *CreateLoadBalancerHTTPSListenerRequest, accessId, accessSecret string) (*CreateLoadBalancerHTTPSListenerResponse, error) {
	var pResponse CreateLoadBalancerHTTPSListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
