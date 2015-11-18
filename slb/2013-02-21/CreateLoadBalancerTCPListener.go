package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerTCPListenerRequest struct {
	RpcRequest
	LoadBalancerId     string
	ListenerPort       int
	BackendServerPort  int
	ListenerStatus     string
	Scheduler          string
	PersistenceTimeout int
	HealthCheck        string
	ConnectTimeout     int
	ConnectPort        int
	Interval           int
	HostId             string
	OwnerAccount       string
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
func (r *CreateLoadBalancerTCPListenerRequest) SetListenerStatus(value string) {
	r.ListenerStatus = value
	r.QueryParams.Set("ListenerStatus", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetListenerStatus() string {
	return r.ListenerStatus
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
func (r *CreateLoadBalancerTCPListenerRequest) SetHealthCheck(value string) {
	r.HealthCheck = value
	r.QueryParams.Set("HealthCheck", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHealthCheck() string {
	return r.HealthCheck
}
func (r *CreateLoadBalancerTCPListenerRequest) SetConnectTimeout(value int) {
	r.ConnectTimeout = value
	r.QueryParams.Set("ConnectTimeout", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetConnectTimeout() int {
	return r.ConnectTimeout
}
func (r *CreateLoadBalancerTCPListenerRequest) SetConnectPort(value int) {
	r.ConnectPort = value
	r.QueryParams.Set("ConnectPort", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetConnectPort() int {
	return r.ConnectPort
}
func (r *CreateLoadBalancerTCPListenerRequest) SetInterval(value int) {
	r.Interval = value
	r.QueryParams.Set("Interval", strconv.Itoa(value))
}
func (r *CreateLoadBalancerTCPListenerRequest) GetInterval() int {
	return r.Interval
}
func (r *CreateLoadBalancerTCPListenerRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *CreateLoadBalancerTCPListenerRequest) GetHostId() string {
	return r.HostId
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
