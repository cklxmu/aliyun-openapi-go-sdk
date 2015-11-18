package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StartInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	OwnerAccount         string
}

func (r *StartInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StartInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StartInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StartInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StartInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StartInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StartInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *StartInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *StartInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StartInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StartInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StartInstance")
	r.SetProduct(Product)
}

type StartInstanceResponse struct {
}

func StartInstance(req *StartInstanceRequest, accessId, accessSecret string) (*StartInstanceResponse, error) {
	var pResponse StartInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
