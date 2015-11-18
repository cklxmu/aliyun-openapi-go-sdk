package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifySecurityIpsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	SecurityIps          string
}

func (r *ModifySecurityIpsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifySecurityIpsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifySecurityIpsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifySecurityIpsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifySecurityIpsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifySecurityIpsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifySecurityIpsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifySecurityIpsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifySecurityIpsRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifySecurityIpsRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifySecurityIpsRequest) SetSecurityIps(value string) {
	r.SecurityIps = value
	r.QueryParams.Set("SecurityIps", value)
}
func (r *ModifySecurityIpsRequest) GetSecurityIps() string {
	return r.SecurityIps
}

func (r *ModifySecurityIpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifySecurityIps")
	r.SetProduct(Product)
}

type ModifySecurityIpsResponse struct {
}

func ModifySecurityIps(req *ModifySecurityIpsRequest, accessId, accessSecret string) (*ModifySecurityIpsResponse, error) {
	var pResponse ModifySecurityIpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
