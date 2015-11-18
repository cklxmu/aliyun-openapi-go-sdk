package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyScheduledTaskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ScheduledTaskId      string
	ScheduledTaskName    string
	Description          string
	ScheduledAction      string
	RecurrenceEndTime    string
	LaunchTime           string
	RecurrenceType       string
	RecurrenceValue      string
	TaskEnabled          bool
	LaunchExpirationTime int
	OwnerAccount         string
}

func (r *ModifyScheduledTaskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyScheduledTaskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyScheduledTaskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyScheduledTaskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyScheduledTaskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyScheduledTaskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyScheduledTaskRequest) SetScheduledTaskId(value string) {
	r.ScheduledTaskId = value
	r.QueryParams.Set("ScheduledTaskId", value)
}
func (r *ModifyScheduledTaskRequest) GetScheduledTaskId() string {
	return r.ScheduledTaskId
}
func (r *ModifyScheduledTaskRequest) SetScheduledTaskName(value string) {
	r.ScheduledTaskName = value
	r.QueryParams.Set("ScheduledTaskName", value)
}
func (r *ModifyScheduledTaskRequest) GetScheduledTaskName() string {
	return r.ScheduledTaskName
}
func (r *ModifyScheduledTaskRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyScheduledTaskRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyScheduledTaskRequest) SetScheduledAction(value string) {
	r.ScheduledAction = value
	r.QueryParams.Set("ScheduledAction", value)
}
func (r *ModifyScheduledTaskRequest) GetScheduledAction() string {
	return r.ScheduledAction
}
func (r *ModifyScheduledTaskRequest) SetRecurrenceEndTime(value string) {
	r.RecurrenceEndTime = value
	r.QueryParams.Set("RecurrenceEndTime", value)
}
func (r *ModifyScheduledTaskRequest) GetRecurrenceEndTime() string {
	return r.RecurrenceEndTime
}
func (r *ModifyScheduledTaskRequest) SetLaunchTime(value string) {
	r.LaunchTime = value
	r.QueryParams.Set("LaunchTime", value)
}
func (r *ModifyScheduledTaskRequest) GetLaunchTime() string {
	return r.LaunchTime
}
func (r *ModifyScheduledTaskRequest) SetRecurrenceType(value string) {
	r.RecurrenceType = value
	r.QueryParams.Set("RecurrenceType", value)
}
func (r *ModifyScheduledTaskRequest) GetRecurrenceType() string {
	return r.RecurrenceType
}
func (r *ModifyScheduledTaskRequest) SetRecurrenceValue(value string) {
	r.RecurrenceValue = value
	r.QueryParams.Set("RecurrenceValue", value)
}
func (r *ModifyScheduledTaskRequest) GetRecurrenceValue() string {
	return r.RecurrenceValue
}
func (r *ModifyScheduledTaskRequest) SetTaskEnabled(value bool) {
	r.TaskEnabled = value
	r.QueryParams.Set("TaskEnabled", strconv.FormatBool(value))
}
func (r *ModifyScheduledTaskRequest) GetTaskEnabled() bool {
	return r.TaskEnabled
}
func (r *ModifyScheduledTaskRequest) SetLaunchExpirationTime(value int) {
	r.LaunchExpirationTime = value
	r.QueryParams.Set("LaunchExpirationTime", strconv.Itoa(value))
}
func (r *ModifyScheduledTaskRequest) GetLaunchExpirationTime() int {
	return r.LaunchExpirationTime
}
func (r *ModifyScheduledTaskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyScheduledTaskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyScheduledTaskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyScheduledTask")
	r.SetProduct(Product)
}

type ModifyScheduledTaskResponse struct {
}

func ModifyScheduledTask(req *ModifyScheduledTaskRequest, accessId, accessSecret string) (*ModifyScheduledTaskResponse, error) {
	var pResponse ModifyScheduledTaskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
