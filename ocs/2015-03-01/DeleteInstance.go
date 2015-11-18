package ocs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	InstanceId           string
}

func (r *DeleteInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DeleteInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DeleteInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}

func (r *DeleteInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteInstance")
	r.SetProduct(Product)
}

type DeleteInstanceResponse struct {
}

func DeleteInstance(req *DeleteInstanceRequest, accessId, accessSecret string) (*DeleteInstanceResponse, error) {
	var pResponse DeleteInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
