package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopInstanceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	ForceStop            bool
	OwnerAccount         string
}

func (r *StopInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StopInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StopInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StopInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StopInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StopInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StopInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *StopInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *StopInstanceRequest) SetForceStop(value bool) {
	r.ForceStop = value
	r.QueryParams.Set("ForceStop", strconv.FormatBool(value))
}
func (r *StopInstanceRequest) GetForceStop() bool {
	return r.ForceStop
}
func (r *StopInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StopInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StopInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopInstance")
	r.SetProduct(Product)
}

type StopInstanceResponse struct {
}

func StopInstance(req *StopInstanceRequest, accessId, accessSecret string) (*StopInstanceResponse, error) {
	var pResponse StopInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
