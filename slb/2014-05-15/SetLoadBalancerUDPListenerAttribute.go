package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerUDPListenerAttributeRequest struct {
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
	MaxConnLimit              int
	OwnerAccount              string
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetPersistenceTimeout(value int) {
	r.PersistenceTimeout = value
	r.QueryParams.Set("PersistenceTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetPersistenceTimeout() int {
	return r.PersistenceTimeout
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectTimeout(value int) {
	r.HealthCheckConnectTimeout = value
	r.QueryParams.Set("HealthCheckConnectTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckConnectTimeout() int {
	return r.HealthCheckConnectTimeout
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetHealthCheckInterval(value int) {
	r.HealthCheckInterval = value
	r.QueryParams.Set("HealthCheckInterval", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetHealthCheckInterval() int {
	return r.HealthCheckInterval
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerUDPListenerAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerUDPListenerAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerUDPListenerAttribute")
	r.SetProduct(Product)
}

type SetLoadBalancerUDPListenerAttributeResponse struct {
}

func SetLoadBalancerUDPListenerAttribute(req *SetLoadBalancerUDPListenerAttributeRequest, accessId, accessSecret string) (*SetLoadBalancerUDPListenerAttributeResponse, error) {
	var pResponse SetLoadBalancerUDPListenerAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
