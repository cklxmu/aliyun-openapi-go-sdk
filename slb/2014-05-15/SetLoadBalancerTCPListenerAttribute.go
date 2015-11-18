package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerTCPListenerAttributeRequest struct {
	RpcRequest
	OwnerId                   int
	ResourceOwnerAccount      string
	ResourceOwnerId           int
	LoadBalancerId            string
	ListenerPort              int
	Bandwidth                 int
	Scheduler                 string
	PersistenceTimeout        int
	HealthyThreshold          int
	UnhealthyThreshold        int
	HealthCheckConnectTimeout int
	HealthCheckConnectPort    int
	HealthCheckInterval       int
	HealthCheckDomain         string
	HealthCheckURI            string
	HealthCheckHttpCode       string
	HealthCheckType           string
	SynProxy                  string
	MaxConnLimit              int
	OwnerAccount              string
}

func (r *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetPersistenceTimeout(value int) {
	r.PersistenceTimeout = value
	r.QueryParams.Set("PersistenceTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetPersistenceTimeout() int {
	return r.PersistenceTimeout
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectTimeout(value int) {
	r.HealthCheckConnectTimeout = value
	r.QueryParams.Set("HealthCheckConnectTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckConnectTimeout() int {
	return r.HealthCheckConnectTimeout
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckInterval(value int) {
	r.HealthCheckInterval = value
	r.QueryParams.Set("HealthCheckInterval", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckInterval() int {
	return r.HealthCheckInterval
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckDomain(value string) {
	r.HealthCheckDomain = value
	r.QueryParams.Set("HealthCheckDomain", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckDomain() string {
	return r.HealthCheckDomain
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckURI(value string) {
	r.HealthCheckURI = value
	r.QueryParams.Set("HealthCheckURI", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckURI() string {
	return r.HealthCheckURI
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckHttpCode(value string) {
	r.HealthCheckHttpCode = value
	r.QueryParams.Set("HealthCheckHttpCode", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckHttpCode() string {
	return r.HealthCheckHttpCode
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheckType(value string) {
	r.HealthCheckType = value
	r.QueryParams.Set("HealthCheckType", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheckType() string {
	return r.HealthCheckType
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetSynProxy(value string) {
	r.SynProxy = value
	r.QueryParams.Set("SynProxy", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetSynProxy() string {
	return r.SynProxy
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerTCPListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerTCPListenerAttribute")
	r.SetProduct(Product)
}

type SetLoadBalancerTCPListenerAttributeResponse struct {
}

func SetLoadBalancerTCPListenerAttribute(req *SetLoadBalancerTCPListenerAttributeRequest, accessId, accessSecret string) (*SetLoadBalancerTCPListenerAttributeResponse, error) {
	var pResponse SetLoadBalancerTCPListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
