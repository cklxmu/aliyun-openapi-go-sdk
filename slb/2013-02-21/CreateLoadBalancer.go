package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoadBalancerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	IsPublicAddress      string
	Address              string
	ClientToken          string
	LoadBalancerName     string
	LoadBalancerMode     string
	OwnerAccount         string
}

func (r *CreateLoadBalancerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateLoadBalancerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateLoadBalancerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateLoadBalancerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateLoadBalancerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateLoadBalancerRequest) SetIsPublicAddress(value string) {
	r.IsPublicAddress = value
	r.QueryParams.Set("IsPublicAddress", value)
}
func (r *CreateLoadBalancerRequest) GetIsPublicAddress() string {
	return r.IsPublicAddress
}
func (r *CreateLoadBalancerRequest) SetAddress(value string) {
	r.Address = value
	r.QueryParams.Set("Address", value)
}
func (r *CreateLoadBalancerRequest) GetAddress() string {
	return r.Address
}
func (r *CreateLoadBalancerRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateLoadBalancerRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateLoadBalancerRequest) SetLoadBalancerName(value string) {
	r.LoadBalancerName = value
	r.QueryParams.Set("LoadBalancerName", value)
}
func (r *CreateLoadBalancerRequest) GetLoadBalancerName() string {
	return r.LoadBalancerName
}
func (r *CreateLoadBalancerRequest) SetLoadBalancerMode(value string) {
	r.LoadBalancerMode = value
	r.QueryParams.Set("LoadBalancerMode", value)
}
func (r *CreateLoadBalancerRequest) GetLoadBalancerMode() string {
	return r.LoadBalancerMode
}
func (r *CreateLoadBalancerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateLoadBalancerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateLoadBalancerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoadBalancer")
	r.SetProduct(Product)
}

type CreateLoadBalancerResponse struct {
	Address          string `xml:"Address" json:"Address"`
	LoadBalancerId   string `xml:"LoadBalancerId" json:"LoadBalancerId"`
	LoadBalancerName string `xml:"LoadBalancerName" json:"LoadBalancerName"`
}

func CreateLoadBalancer(req *CreateLoadBalancerRequest, accessId, accessSecret string) (*CreateLoadBalancerResponse, error) {
	var pResponse CreateLoadBalancerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
