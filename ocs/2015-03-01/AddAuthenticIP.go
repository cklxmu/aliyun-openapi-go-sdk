package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AddAuthenticIPRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	AuthenticIP          string
}

func (r *AddAuthenticIPRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AddAuthenticIPRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AddAuthenticIPRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AddAuthenticIPRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AddAuthenticIPRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AddAuthenticIPRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AddAuthenticIPRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AddAuthenticIPRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *AddAuthenticIPRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *AddAuthenticIPRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *AddAuthenticIPRequest) SetAuthenticIP(value string) {
	r.AuthenticIP = value
	r.QueryParams.Set("AuthenticIP", value)
}
func (r *AddAuthenticIPRequest) GetAuthenticIP() string {
	return r.AuthenticIP
}

func (r *AddAuthenticIPRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddAuthenticIP")
	r.SetProduct(Product)
}

type AddAuthenticIPResponse struct {
}

func AddAuthenticIP(req *AddAuthenticIPRequest, accessId, accessSecret string) (*AddAuthenticIPResponse, error) {
	var pResponse AddAuthenticIPResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
