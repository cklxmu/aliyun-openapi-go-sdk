package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeBackendServersRequest struct {
	RpcRequest
	LoadBalancerId string
	ListenerPort   int
	HostId         string
	OwnerAccount   string
}

func (r *DescribeBackendServersRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *DescribeBackendServersRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *DescribeBackendServersRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *DescribeBackendServersRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *DescribeBackendServersRequest) SetHostId(value string) {
	r.HostId = value
	r.QueryParams.Set("HostId", value)
}
func (r *DescribeBackendServersRequest) GetHostId() string {
	return r.HostId
}
func (r *DescribeBackendServersRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeBackendServersRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeBackendServersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeBackendServers")
	r.SetProduct(Product)
}

type DescribeBackendServersResponse struct {
	Listeners struct {
		Listener []struct {
			ListenerPort   int `xml:"ListenerPort" json:"ListenerPort"`
			BackendServers struct {
				BackendServer []struct {
					ServerId           string `xml:"ServerId" json:"ServerId"`
					ServerHealthStatus string `xml:"ServerHealthStatus" json:"ServerHealthStatus"`
				} `xml:"BackendServer" json:"BackendServer"`
			} `xml:"BackendServers" json:"BackendServers"`
		} `xml:"Listener" json:"Listener"`
	} `xml:"Listeners" json:"Listeners"`
}

func DescribeBackendServers(req *DescribeBackendServersRequest, accessId, accessSecret string) (*DescribeBackendServersResponse, error) {
	var pResponse DescribeBackendServersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
