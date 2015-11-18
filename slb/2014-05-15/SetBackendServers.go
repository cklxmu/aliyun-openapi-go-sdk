package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetBackendServersRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	BackendServers       string
	OwnerAccount         string
}

func (r *SetBackendServersRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetBackendServersRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetBackendServersRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetBackendServersRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetBackendServersRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetBackendServersRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetBackendServersRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *SetBackendServersRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *SetBackendServersRequest) SetBackendServers(value string) {
	r.BackendServers = value
	r.QueryParams.Set("BackendServers", value)
}
func (r *SetBackendServersRequest) GetBackendServers() string {
	return r.BackendServers
}
func (r *SetBackendServersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetBackendServersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetBackendServersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetBackendServers")
	r.SetProduct(Product)
}

type SetBackendServersResponse struct {
	LoadBalancerId string `xml:"LoadBalancerId" json:"LoadBalancerId"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `xml:"ServerId" json:"ServerId"`
			Weight   string `xml:"Weight" json:"Weight"`
		} `xml:"BackendServer" json:"BackendServer"`
	} `xml:"BackendServers" json:"BackendServers"`
}

func SetBackendServers(req *SetBackendServersRequest, accessId, accessSecret string) (*SetBackendServersResponse, error) {
	var pResponse SetBackendServersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
