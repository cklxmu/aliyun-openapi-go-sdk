package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetLoadBanancerListenerStatusRequest struct {
	RpcRequest
	LoadBalancerId string
	ListenerPort   int
	ListenerStatus string
	HostId         string
	OwnerAccount   string
}

func (r *SetLoadBanancerListenerStatusRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetLoadBanancerListenerStatusRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetLoadBanancerListenerStatusRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetLoadBanancerListenerStatusRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetLoadBanancerListenerStatusRequest) SetListenerStatus(value string) {
	r.ListenerStatus = value
	r.QueryParams.Set("ListenerStatus", value)
}
func (r *SetLoadBanancerListenerStatusRequest) GetListenerStatus() string {
	return r.ListenerStatus
}
func (r *SetLoadBanancerListenerStatusRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *SetLoadBanancerListenerStatusRequest) GetHostId() string {
	return r.HostId
}
func (r *SetLoadBanancerListenerStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetLoadBanancerListenerStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetLoadBanancerListenerStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetLoadBanancerListenerStatus")
	r.SetProduct(Product)
}

type SetLoadBanancerListenerStatusResponse struct {
}

func SetLoadBanancerListenerStatus(req *SetLoadBanancerListenerStatusRequest, accessId, accessSecret string) (*SetLoadBanancerListenerStatusResponse, error) {
	var pResponse SetLoadBanancerListenerStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
