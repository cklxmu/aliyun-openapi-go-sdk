package slb

import (
	. "aliyun-openapi-go-sdk/core"
)

type AddBackendServersRequest struct {
	RpcRequest
	LoadBalancerId string
	BackendServers string
	HostId         string
	OwnerAccount   string
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
func (r *AddBackendServersRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *AddBackendServersRequest) GetHostId() string {
	return r.HostId
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
			Weight   int    `xml:"Weight" json:"Weight"`
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
