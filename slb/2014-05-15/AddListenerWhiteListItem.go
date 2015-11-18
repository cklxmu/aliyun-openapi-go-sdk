package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AddListenerWhiteListItemRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	LoadBalancerId       string
	ListenerPort         int
	SourceItems          string
	OwnerAccount         string
}

func (r *AddListenerWhiteListItemRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AddListenerWhiteListItemRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AddListenerWhiteListItemRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AddListenerWhiteListItemRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AddListenerWhiteListItemRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AddListenerWhiteListItemRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AddListenerWhiteListItemRequest) SetLoadBalancerId(value string) {
	r.LoadBalancerId = value
	r.QueryParams.Set("LoadBalancerId", value)
}
func (r *AddListenerWhiteListItemRequest) GetLoadBalancerId() string {
	return r.LoadBalancerId
}
func (r *AddListenerWhiteListItemRequest) SetListenerPort(value int) {
	r.ListenerPort = value
	r.QueryParams.Set("ListenerPort", strconv.Itoa(value))
}
func (r *AddListenerWhiteListItemRequest) GetListenerPort() int {
	return r.ListenerPort
}
func (r *AddListenerWhiteListItemRequest) SetSourceItems(value string) {
	r.SourceItems = value
	r.QueryParams.Set("SourceItems", value)
}
func (r *AddListenerWhiteListItemRequest) GetSourceItems() string {
	return r.SourceItems
}
func (r *AddListenerWhiteListItemRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AddListenerWhiteListItemRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AddListenerWhiteListItemRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddListenerWhiteListItem")
	r.SetProduct(Product)
}

type AddListenerWhiteListItemResponse struct {
}

func AddListenerWhiteListItem(req *AddListenerWhiteListItemRequest, accessId, accessSecret string) (*AddListenerWhiteListItemResponse, error) {
	var pResponse AddListenerWhiteListItemResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
