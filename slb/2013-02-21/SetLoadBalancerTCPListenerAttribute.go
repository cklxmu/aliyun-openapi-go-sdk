package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerTCPListenerAttributeRequest struct {
	RpcRequest
	LoadBalancerId     string
	ListenerPort       int
	Scheduler          string
	PersistenceTimeout int
	HealthCheck        string
	HealthyThreshold   int
	UnhealthyThreshold int
	ConnectTimeout     int
	ConnectPort        int
	Interval           int
	HostId             string
	OwnerAccount       string
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
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHealthCheck(value string) {
	r.HealthCheck = value
	r.QueryParams.Set("HealthCheck", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHealthCheck() string {
	return r.HealthCheck
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
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetConnectTimeout(value int) {
	r.ConnectTimeout = value
	r.QueryParams.Set("ConnectTimeout", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetConnectTimeout() int {
	return r.ConnectTimeout
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetConnectPort(value int) {
	r.ConnectPort = value
	r.QueryParams.Set("ConnectPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetConnectPort() int {
	return r.ConnectPort
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetInterval(value int) {
	r.Interval = value
	r.QueryParams.Set("Interval", strconv.Itoa(value))
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetInterval() int {
	return r.Interval
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *SetLoadBalancerTCPListenerAttributeRequest) GetHostId() string {
	return r.HostId
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
