package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopLoadBalancerListenerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *StopLoadBalancerListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StopLoadBalancerListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StopLoadBalancerListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StopLoadBalancerListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StopLoadBalancerListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StopLoadBalancerListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StopLoadBalancerListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *StopLoadBalancerListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *StopLoadBalancerListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *StopLoadBalancerListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *StopLoadBalancerListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StopLoadBalancerListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StopLoadBalancerListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopLoadBalancerListener")
	r.SetProduct(Product)
}

type StopLoadBalancerListenerResponse struct {
}

func StopLoadBalancerListener(req *StopLoadBalancerListenerRequest, accessId, accessSecret string) (*StopLoadBalancerListenerResponse, error) {
	var pResponse StopLoadBalancerListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
