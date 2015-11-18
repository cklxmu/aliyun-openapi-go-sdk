package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetListenerAccessControlStatusRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	AccessControlStatus  string
	OwnerAccount         string
}

func (r *SetListenerAccessControlStatusRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetListenerAccessControlStatusRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetListenerAccessControlStatusRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetListenerAccessControlStatusRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetListenerAccessControlStatusRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetListenerAccessControlStatusRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetListenerAccessControlStatusRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetListenerAccessControlStatusRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetListenerAccessControlStatusRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *SetListenerAccessControlStatusRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *SetListenerAccessControlStatusRequest) SetAccessControlStatus(value string) {
	r.AccessControlStatus = value
	r.QueryParams.Set("AccessControlStatus", value)
}
func (r *SetListenerAccessControlStatusRequest) GetAccessControlStatus() string {
	return r.AccessControlStatus
}
func (r *SetListenerAccessControlStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetListenerAccessControlStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetListenerAccessControlStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetListenerAccessControlStatus")
	r.SetProduct(Product)
}

type SetListenerAccessControlStatusResponse struct {
}

func SetListenerAccessControlStatus(req *SetListenerAccessControlStatusRequest, accessId, accessSecret string) (*SetListenerAccessControlStatusResponse, error) {
	var pResponse SetListenerAccessControlStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
