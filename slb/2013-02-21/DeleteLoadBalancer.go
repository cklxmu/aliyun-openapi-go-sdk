package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteLoadBalancerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	OwnerAccount         string
}

func (r *DeleteLoadBalancerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteLoadBalancerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteLoadBalancerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteLoadBalancerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteLoadBalancerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteLoadBalancerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteLoadBalancerRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DeleteLoadBalancerRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DeleteLoadBalancerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteLoadBalancerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteLoadBalancerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteLoadBalancer")
	r.SetProduct(Product)
}

type DeleteLoadBalancerResponse struct {
}

func DeleteLoadBalancer(req *DeleteLoadBalancerRequest, accessId, accessSecret string) (*DeleteLoadBalancerResponse, error) {
	var pResponse DeleteLoadBalancerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
