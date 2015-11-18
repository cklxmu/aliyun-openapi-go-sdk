package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeHealthStatusRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	OwnerAccount         string
}

func (r *DescribeHealthStatusRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeHealthStatusRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeHealthStatusRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeHealthStatusRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeHealthStatusRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeHealthStatusRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeHealthStatusRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeHealthStatusRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeHealthStatusRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeHealthStatusRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeHealthStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeHealthStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeHealthStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeHealthStatus")
	r.SetProduct(Product)
}

type DescribeHealthStatusResponse struct {
	BackendServers struct {
		BackendServer []struct {
			ServerId           string `xml:"ServerId" json:"ServerId"`
			ServerHealthStatus string `xml:"ServerHealthStatus" json:"ServerHealthStatus"`
		} `xml:"BackendServer" json:"BackendServer"`
	} `xml:"BackendServers" json:"BackendServers"`
}

func DescribeHealthStatus(req *DescribeHealthStatusRequest, accessId, accessSecret string) (*DescribeHealthStatusResponse, error) {
	var pResponse DescribeHealthStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
