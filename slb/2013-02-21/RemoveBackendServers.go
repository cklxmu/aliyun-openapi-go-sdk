package slb

import (
	. "aliyun-openapi-go-sdk/core"
)

type RemoveBackendServersRequest struct {
	RpcRequest
	LoadBalancerId string
	BackendServers string
	HostId         string
	OwnerAccount   string
}

func (r *RemoveBackendServersRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *RemoveBackendServersRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *RemoveBackendServersRequest) SetBackendServers(value string) {
	r.BackendServers = value
	r.QueryParams.Set("BackendServers", value)
}
func (r *RemoveBackendServersRequest) GetBackendServers() string {
	return r.BackendServers
}
func (r *RemoveBackendServersRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *RemoveBackendServersRequest) GetHostId() string {
	return r.HostId
}
func (r *RemoveBackendServersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RemoveBackendServersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RemoveBackendServersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveBackendServers")
	r.SetProduct(Product)
}

type RemoveBackendServersResponse struct {
	LoadBalancerId string `xml:"LoadBalancerId" json:"LoadBalancerId"`
	BackendServers struct {
		BackendServer []struct {
			ServerId string `xml:"ServerId" json:"ServerId"`
			Weight   int    `xml:"Weight" json:"Weight"`
		} `xml:"BackendServer" json:"BackendServer"`
	} `xml:"BackendServers" json:"BackendServers"`
}

func RemoveBackendServers(req *RemoveBackendServersRequest, accessId, accessSecret string) (*RemoveBackendServersResponse, error) {
	var pResponse RemoveBackendServersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
