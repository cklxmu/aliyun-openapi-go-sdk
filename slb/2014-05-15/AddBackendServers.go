package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AddBackendServersRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	BackendServers       string
	OwnerAccount         string
}

func (r *AddBackendServersRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AddBackendServersRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AddBackendServersRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AddBackendServersRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AddBackendServersRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AddBackendServersRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AddBackendServersRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *AddBackendServersRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *AddBackendServersRequest) SetBackendServers(value string) {
	r.BackendServers = value
	r.QueryParams.Set("BackendServers", value)
}
func (r *AddBackendServersRequest) GetBackendServers() string {
	return r.BackendServers
}
func (r *AddBackendServersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AddBackendServersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AddBackendServersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddBackendServers")
	r.SetProduct(Product)
}

type AddBackendServersResponse struct {
	LoadBalancerId string `xml:"LoadBalancerId" json:"LoadBalancerId"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `xml:"ServerId" json:"ServerId"`
			Weight   string `xml:"Weight" json:"Weight"`
		} `xml:"BackendServer" json:"BackendServer"`
	} `xml:"BackendServers" json:"BackendServers"`
}

func AddBackendServers(req *AddBackendServersRequest, accessId, accessSecret string) (*AddBackendServersResponse, error) {
	var pResponse AddBackendServersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
