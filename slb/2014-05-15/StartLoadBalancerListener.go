package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StartLoadBalancerListenerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *StartLoadBalancerListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StartLoadBalancerListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StartLoadBalancerListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StartLoadBalancerListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StartLoadBalancerListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StartLoadBalancerListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StartLoadBalancerListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *StartLoadBalancerListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *StartLoadBalancerListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *StartLoadBalancerListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *StartLoadBalancerListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StartLoadBalancerListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StartLoadBalancerListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StartLoadBalancerListener")
	r.SetProduct(Product)
}

type StartLoadBalancerListenerResponse struct {
}

func StartLoadBalancerListener(req *StartLoadBalancerListenerRequest, accessId, accessSecret string) (*StartLoadBalancerListenerResponse, error) {
	var pResponse StartLoadBalancerListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
