package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RemoveBackendServersRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	BackendServers       string
	OwnerAccount         string
}

func (r *RemoveBackendServersRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RemoveBackendServersRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RemoveBackendServersRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RemoveBackendServersRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RemoveBackendServersRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RemoveBackendServersRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
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
