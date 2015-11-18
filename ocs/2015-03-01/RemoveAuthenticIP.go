package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RemoveAuthenticIPRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	AuthenticIP          string
}

func (r *RemoveAuthenticIPRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RemoveAuthenticIPRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RemoveAuthenticIPRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RemoveAuthenticIPRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RemoveAuthenticIPRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RemoveAuthenticIPRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RemoveAuthenticIPRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RemoveAuthenticIPRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *RemoveAuthenticIPRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *RemoveAuthenticIPRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *RemoveAuthenticIPRequest) SetAuthenticIP(value string) {
	r.AuthenticIP = value
	r.QueryParams.Set("AuthenticIP", value)
}
func (r *RemoveAuthenticIPRequest) GetAuthenticIP() string {
	return r.AuthenticIP
}

func (r *RemoveAuthenticIPRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveAuthenticIP")
	r.SetProduct(Product)
}

type RemoveAuthenticIPResponse struct {
}

func RemoveAuthenticIP(req *RemoveAuthenticIPRequest, accessId, accessSecret string) (*RemoveAuthenticIPResponse, error) {
	var pResponse RemoveAuthenticIPResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
