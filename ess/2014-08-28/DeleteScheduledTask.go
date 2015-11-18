package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteScheduledTaskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScheduledTaskId      string
	OwnerAccount         string
}

func (r *DeleteScheduledTaskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteScheduledTaskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteScheduledTaskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteScheduledTaskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteScheduledTaskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteScheduledTaskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteScheduledTaskRequest) SetScheduledTaskId(value string) {
	r.ScheduledTaskId = value
	r.QueryParams.Set("ScheduledTaskId", value)
}
func (r *DeleteScheduledTaskRequest) GetScheduledTaskId() string {
	return r.ScheduledTaskId
}
func (r *DeleteScheduledTaskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteScheduledTaskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteScheduledTaskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteScheduledTask")
	r.SetProduct(Product)
}

type DeleteScheduledTaskResponse struct {
}

func DeleteScheduledTask(req *DeleteScheduledTaskRequest, accessId, accessSecret string) (*DeleteScheduledTaskResponse, error) {
	var pResponse DeleteScheduledTaskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
