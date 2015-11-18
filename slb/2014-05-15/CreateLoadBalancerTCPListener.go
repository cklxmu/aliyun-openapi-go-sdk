package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerTCPListenerRequest struct {
	RpcRequest
	OwnerId                   int
	ResourceOwnerAccount      string
	ResourceOwnerId           int
	LoadBalancerId            string
	ListenerPort              int
	BackendServerPort         int
	Bandwidth                 int
	Scheduler                 string
	PersistenceTimeout        int
	HealthyThreshold          int
	UnhealthyThreshold        int
	HealthCheckConnectTimeout int
	HealthCheckConnectPort    int
	healthCheckInterval       int
	HealthCheckDomain         string
	HealthCheckURI            string
	HealthCheckHttpCode       string
	HealthCheckType           string
	MaxConnLimit              int
	OwnerAccount              string
}

func (r *CreateLoadBalancerTCPListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateLoadBalancerTCPListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *CreateLoadBalancerTCPListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *CreateLoadBalancerTCPListenerRequest) SetBackendServerPort(value int) {
	r.BackendServerPort = value
	r.QueryParams.Set("BackendServerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetBackendServerPort() int {
	return r.BackendServerPort
}
func (r *CreateLoadBalancerTCPListenerRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *CreateLoadBalancerTCPListenerRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *CreateLoadBalancerTCPListenerRequest) SetPersistenceTimeout(value int) {
	r.PersistenceTimeout = value
	r.QueryParams.Set("PersistenceTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetPersistenceTimeout() int {
	return r.PersistenceTimeout
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *CreateLoadBalancerTCPListenerRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectTimeout(value int) {
	r.HealthCheckConnectTimeout = value
	r.QueryParams.Set("HealthCheckConnectTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheckConnectTimeout() int {
	return r.HealthCheckConnectTimeout
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *CreateLoadBalancerTCPListenerRequest) SethealthCheckInterval(value int) {
	r.healthCheckInterval = value
	r.QueryParams.Set("healthCheckInterval", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GethealthCheckInterval() int {
	return r.healthCheckInterval
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheckDomain(value string) {
	r.HealthCheckDomain = value
	r.QueryParams.Set("HealthCheckDomain", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheckDomain() string {
	return r.HealthCheckDomain
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheckURI(value string) {
	r.HealthCheckURI = value
	r.QueryParams.Set("HealthCheckURI", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheckURI() string {
	return r.HealthCheckURI
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheckHttpCode(value string) {
	r.HealthCheckHttpCode = value
	r.QueryParams.Set("HealthCheckHttpCode", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheckHttpCode() string {
	return r.HealthCheckHttpCode
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheckType(value string) {
	r.HealthCheckType = value
	r.QueryParams.Set("HealthCheckType", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheckType() string {
	return r.HealthCheckType
}
func (r *CreateLoadBalancerTCPListenerRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *CreateLoadBalancerTCPListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateLoadBalancerTCPListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoadBalancerTCPListener")
	r.SetProduct(Product)
}

type CreateLoadBalancerTCPListenerResponse struct {
}

func CreateLoadBalancerTCPListener(req *CreateLoadBalancerTCPListenerRequest, accessId, accessSecret string) (*CreateLoadBalancerTCPListenerResponse, error) {
	var pResponse CreateLoadBalancerTCPListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
