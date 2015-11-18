package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RemoveListenerWhiteListItemRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	SourceItems          string
	OwnerAccount         string
}

func (r *RemoveListenerWhiteListItemRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RemoveListenerWhiteListItemRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RemoveListenerWhiteListItemRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RemoveListenerWhiteListItemRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RemoveListenerWhiteListItemRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RemoveListenerWhiteListItemRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RemoveListenerWhiteListItemRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *RemoveListenerWhiteListItemRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *RemoveListenerWhiteListItemRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *RemoveListenerWhiteListItemRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *RemoveListenerWhiteListItemRequest) SetSourceItems(value string) {
	r.SourceItems = value
	r.QueryParams.Set("SourceItems", value)
}
func (r *RemoveListenerWhiteListItemRequest) GetSourceItems() string {
	return r.SourceItems
}
func (r *RemoveListenerWhiteListItemRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RemoveListenerWhiteListItemRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RemoveListenerWhiteListItemRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveListenerWhiteListItem")
	r.SetProduct(Product)
}

type RemoveListenerWhiteListItemResponse struct {
}

func RemoveListenerWhiteListItem(req *RemoveListenerWhiteListItemRequest, accessId, accessSecret string) (*RemoveListenerWhiteListItemResponse, error) {
	var pResponse RemoveListenerWhiteListItemResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
