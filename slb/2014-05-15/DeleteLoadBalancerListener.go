package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteLoadBalancerListenerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *DeleteLoadBalancerListenerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteLoadBalancerListenerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteLoadBalancerListenerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteLoadBalancerListenerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteLoadBalancerListenerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteLoadBalancerListenerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteLoadBalancerListenerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DeleteLoadBalancerListenerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DeleteLoadBalancerListenerRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DeleteLoadBalancerListenerRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DeleteLoadBalancerListenerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteLoadBalancerListenerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteLoadBalancerListenerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteLoadBalancerListener")
	r.SetProduct(Product)
}

type DeleteLoadBalancerListenerResponse struct {
}

func DeleteLoadBalancerListener(req *DeleteLoadBalancerListenerRequest, accessId, accessSecret string) (*DeleteLoadBalancerListenerResponse, error) {
	var pResponse DeleteLoadBalancerListenerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
