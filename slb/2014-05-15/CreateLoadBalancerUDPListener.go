package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerUDPListenerRequest struct {
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
	MaxConnLimit              int
	OwnerAccount              string
}

func (r *CreateLoadBalancerUDPListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateLoadBalancerUDPListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateLoadBalancerUDPListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateLoadBalancerUDPListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateLoadBalancerUDPListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *CreateLoadBalancerUDPListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *CreateLoadBalancerUDPListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *CreateLoadBalancerUDPListenerRequest) SetBackendServerPort(value int) {
	r.BackendServerPort = value
	r.QueryParams.Set("BackendServerPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetBackendServerPort() int {
	return r.BackendServerPort
}
func (r *CreateLoadBalancerUDPListenerRequest) SetBandwidth(value int) {
	r.Bandwidth = value
	r.QueryParams.Set("Bandwidth", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetBandwidth() int {
	return r.Bandwidth
}
func (r *CreateLoadBalancerUDPListenerRequest) SetScheduler(value string) {
	r.Scheduler = value
	r.QueryParams.Set("Scheduler", value)
}
func (r *CreateLoadBalancerUDPListenerRequest) GetScheduler() string {
	return r.Scheduler
}
func (r *CreateLoadBalancerUDPListenerRequest) SetPersistenceTimeout(value int) {
	r.PersistenceTimeout = value
	r.QueryParams.Set("PersistenceTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetPersistenceTimeout() int {
	return r.PersistenceTimeout
}
func (r *CreateLoadBalancerUDPListenerRequest) SetHealthyThreshold(value int) {
	r.HealthyThreshold = value
	r.QueryParams.Set("HealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetHealthyThreshold() int {
	return r.HealthyThreshold
}
func (r *CreateLoadBalancerUDPListenerRequest) SetUnhealthyThreshold(value int) {
	r.UnhealthyThreshold = value
	r.QueryParams.Set("UnhealthyThreshold", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetUnhealthyThreshold() int {
	return r.UnhealthyThreshold
}
func (r *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectTimeout(value int) {
	r.HealthCheckConnectTimeout = value
	r.QueryParams.Set("HealthCheckConnectTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetHealthCheckConnectTimeout() int {
	return r.HealthCheckConnectTimeout
}
func (r *CreateLoadBalancerUDPListenerRequest) SetHealthCheckConnectPort(value int) {
	r.HealthCheckConnectPort = value
	r.QueryParams.Set("HealthCheckConnectPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetHealthCheckConnectPort() int {
	return r.HealthCheckConnectPort
}
func (r *CreateLoadBalancerUDPListenerRequest) SethealthCheckInterval(value int) {
	r.healthCheckInterval = value
	r.QueryParams.Set("healthCheckInterval", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GethealthCheckInterval() int {
	return r.healthCheckInterval
}
func (r *CreateLoadBalancerUDPListenerRequest) SetMaxConnLimit(value int) {
	r.MaxConnLimit = value
	r.QueryParams.Set("MaxConnLimit", strconv.Itoa(value))
}
func (r *CreateLoadBalancerUDPListenerRequest) GetMaxConnLimit() int {
	return r.MaxConnLimit
}
func (r *CreateLoadBalancerUDPListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateLoadBalancerUDPListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateLoadBalancerUDPListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoadBalancerUDPListener")
	r.SetProduct(Product)
}

type CreateLoadBalancerUDPListenerResponse struct {
}

func CreateLoadBalancerUDPListener(req *CreateLoadBalancerUDPListenerRequest, accessId, accessSecret string) (*CreateLoadBalancerUDPListenerResponse, error) {
	var pResponse CreateLoadBalancerUDPListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
