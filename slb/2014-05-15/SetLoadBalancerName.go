package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerNameRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	LoadBalancerName     string
	OwnerAccount         string
}

func (r *SetLoadBalancerNameRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerNameRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetLoadBalancerNameRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetLoadBalancerNameRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetLoadBalancerNameRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerNameRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetLoadBalancerNameRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerNameRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerNameRequest) SetLoadBalancerName(value string) {
	r.LoadBalancerName = value
	r.QueryParams.Set("LoadBalancerName", value)
}
func (r *SetLoadBalancerNameRequest) GetLoadBalancerName() string {
	return r.LoadBalancerName
}
func (r *SetLoadBalancerNameRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerNameRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerNameRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerName")
	r.SetProduct(Product)
}

type SetLoadBalancerNameResponse struct {
}

func SetLoadBalancerName(req *SetLoadBalancerNameRequest, accessId, accessSecret string) (*SetLoadBalancerNameResponse, error) {
	var pResponse SetLoadBalancerNameResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
