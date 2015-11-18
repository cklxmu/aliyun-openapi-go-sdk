package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReplaceAuthenticIPRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	OldAuthenticIP       string
	NewAuthenticIP       string
}

func (r *ReplaceAuthenticIPRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ReplaceAuthenticIPRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ReplaceAuthenticIPRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ReplaceAuthenticIPRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ReplaceAuthenticIPRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ReplaceAuthenticIPRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ReplaceAuthenticIPRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ReplaceAuthenticIPRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ReplaceAuthenticIPRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ReplaceAuthenticIPRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ReplaceAuthenticIPRequest) SetOldAuthenticIP(value string) {
	r.OldAuthenticIP = value
	r.QueryParams.Set("OldAuthenticIP", value)
}
func (r *ReplaceAuthenticIPRequest) GetOldAuthenticIP() string {
	return r.OldAuthenticIP
}
func (r *ReplaceAuthenticIPRequest) SetNewAuthenticIP(value string) {
	r.NewAuthenticIP = value
	r.QueryParams.Set("NewAuthenticIP", value)
}
func (r *ReplaceAuthenticIPRequest) GetNewAuthenticIP() string {
	return r.NewAuthenticIP
}

func (r *ReplaceAuthenticIPRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReplaceAuthenticIP")
	r.SetProduct(Product)
}

type ReplaceAuthenticIPResponse struct {
}

func ReplaceAuthenticIP(req *ReplaceAuthenticIPRequest, accessId, accessSecret string) (*ReplaceAuthenticIPResponse, error) {
	var pResponse ReplaceAuthenticIPResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
