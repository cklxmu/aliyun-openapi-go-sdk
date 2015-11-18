package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBalancerListenerStatusRequest struct {
	RpcRequest
	LoadBalancerId string
	ListenerPort   int
	ListenerStatus string
	HostId         string
	OwnerAccount   string
}

func (r *SetLoadBalancerListenerStatusRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBalancerListenerStatusRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBalancerListenerStatusRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetLoadBalancerListenerStatusRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetLoadBalancerListenerStatusRequest) SetListenerStatus(value string) {
	r.ListenerStatus = value
	r.QueryParams.Set("ListenerStatus", value)
}
func (r *SetLoadBalancerListenerStatusRequest) GetListenerStatus() string {
	return r.ListenerStatus
}
func (r *SetLoadBalancerListenerStatusRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *SetLoadBalancerListenerStatusRequest) GetHostId() string {
	return r.HostId
}
func (r *SetLoadBalancerListenerStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBalancerListenerStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBalancerListenerStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBalancerListenerStatus")
	r.SetProduct(Product)
}

type SetLoadBalancerListenerStatusResponse struct {
}

func SetLoadBalancerListenerStatus(req *SetLoadBalancerListenerStatusRequest, accessId, accessSecret string) (*SetLoadBalancerListenerStatusResponse, error) {
	var pResponse SetLoadBalancerListenerStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
