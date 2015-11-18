package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RebootInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	ForceStop            bool
	OwnerAccount         string
}

func (r *RebootInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RebootInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RebootInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RebootInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RebootInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RebootInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RebootInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *RebootInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *RebootInstanceRequest) SetForceStop(value bool) {
	r.ForceStop = value
	r.QueryParams.Set("ForceStop", strconv.FormatBool(value))
}
func (r *RebootInstanceRequest) GetForceStop() bool {
	return r.ForceStop
}
func (r *RebootInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RebootInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RebootInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RebootInstance")
	r.SetProduct(Product)
}

type RebootInstanceResponse struct {
}

func RebootInstance(req *RebootInstanceRequest, accessId, accessSecret string) (*RebootInstanceResponse, error) {
	var pResponse RebootInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
