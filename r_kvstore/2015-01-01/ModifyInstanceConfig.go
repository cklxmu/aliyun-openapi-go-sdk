package r_kvstore

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyInstanceConfigRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
	Config               string
}

func (r *ModifyInstanceConfigRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceConfigRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyInstanceConfigRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyInstanceConfigRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyInstanceConfigRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyInstanceConfigRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyInstanceConfigRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyInstanceConfigRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ModifyInstanceConfigRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ModifyInstanceConfigRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ModifyInstanceConfigRequest) SetConfig(value string) {
	r.Config = value
	r.QueryParams.Set("Config", value)
}
func (r *ModifyInstanceConfigRequest) GetConfig() string {
	return r.Config
}

func (r *ModifyInstanceConfigRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyInstanceConfig")
	r.SetProduct(Product)
}

type ModifyInstanceConfigResponse struct {
}

func ModifyInstanceConfig(req *ModifyInstanceConfigRequest, accessId, accessSecret string) (*ModifyInstanceConfigResponse, error) {
	var pResponse ModifyInstanceConfigResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
