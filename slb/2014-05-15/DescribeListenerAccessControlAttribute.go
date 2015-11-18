package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeListenerAccessControlAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *DescribeListenerAccessControlAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeListenerAccessControlAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeListenerAccessControlAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeListenerAccessControlAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeListenerAccessControlAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeListenerAccessControlAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeListenerAccessControlAttributeRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeListenerAccessControlAttributeRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeListenerAccessControlAttributeRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeListenerAccessControlAttributeRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeListenerAccessControlAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeListenerAccessControlAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeListenerAccessControlAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeListenerAccessControlAttribute")
	r.SetProduct(Product)
}

type DescribeListenerAccessControlAttributeResponse struct {
	AccessControlStatus string `xml:"AccessControlStatus" json:"AccessControlStatus"`
	SourceItems         string `xml:"SourceItems" json:"SourceItems"`
}

func DescribeListenerAccessControlAttribute(req *DescribeListenerAccessControlAttributeRequest, accessId, accessSecret string) (*DescribeListenerAccessControlAttributeResponse, error) {
	var pResponse DescribeListenerAccessControlAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
