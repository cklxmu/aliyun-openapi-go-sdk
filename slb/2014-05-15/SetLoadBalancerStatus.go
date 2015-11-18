package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerStatusRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	LoadBalancerStatus   string
	OwnerAccount         string
}

func (r *SetLoadBalancerStatusRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerStatusRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetLoadBalancerStatusRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetLoadBalancerStatusRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetLoadBalancerStatusRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetLoadBalancerStatusRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetLoadBalancerStatusRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerStatusRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerStatusRequest) SetLoadBalancerStatus(value string) {
	r.LoadBalancerStatus = value
	r.QueryParams.Set("LoadBalancerStatus", value)
}
func (r *SetLoadBalancerStatusRequest) GetLoadBalancerStatus() string {
	return r.LoadBalancerStatus
}
func (r *SetLoadBalancerStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerStatus")
	r.SetProduct(Product)
}

type SetLoadBalancerStatusResponse struct {
}

func SetLoadBalancerStatus(req *SetLoadBalancerStatusRequest, accessId, accessSecret string) (*SetLoadBalancerStatusResponse, error) {
	var pResponse SetLoadBalancerStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
