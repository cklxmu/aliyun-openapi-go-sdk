package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateScheduledTaskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
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

func (r *CreateScheduledTaskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateScheduledTaskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateScheduledTaskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateScheduledTaskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateScheduledTaskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateScheduledTaskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateScheduledTaskRequest) SetScheduledTaskName(value string) {
	r.ScheduledTaskName = value
	r.QueryParams.Set("ScheduledTaskName", value)
}
func (r *CreateScheduledTaskRequest) GetScheduledTaskName() string {
	return r.ScheduledTaskName
}
func (r *CreateScheduledTaskRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateScheduledTaskRequest) GetDescription() string {
	return r.Description
}
func (r *CreateScheduledTaskRequest) SetScheduledAction(value string) {
	r.ScheduledAction = value
	r.QueryParams.Set("ScheduledAction", value)
}
func (r *CreateScheduledTaskRequest) GetScheduledAction() string {
	return r.ScheduledAction
}
func (r *CreateScheduledTaskRequest) SetRecurrenceEndTime(value string) {
	r.RecurrenceEndTime = value
	r.QueryParams.Set("RecurrenceEndTime", value)
}
func (r *CreateScheduledTaskRequest) GetRecurrenceEndTime() string {
	return r.RecurrenceEndTime
}
func (r *CreateScheduledTaskRequest) SetLaunchTime(value string) {
	r.LaunchTime = value
	r.QueryParams.Set("LaunchTime", value)
}
func (r *CreateScheduledTaskRequest) GetLaunchTime() string {
	return r.LaunchTime
}
func (r *CreateScheduledTaskRequest) SetRecurrenceType(value string) {
	r.RecurrenceType = value
	r.QueryParams.Set("RecurrenceType", value)
}
func (r *CreateScheduledTaskRequest) GetRecurrenceType() string {
	return r.RecurrenceType
}
func (r *CreateScheduledTaskRequest) SetRecurrenceValue(value string) {
	r.RecurrenceValue = value
	r.QueryParams.Set("RecurrenceValue", value)
}
func (r *CreateScheduledTaskRequest) GetRecurrenceValue() string {
	return r.RecurrenceValue
}
func (r *CreateScheduledTaskRequest) SetTaskEnabled(value bool) {
	r.TaskEnabled = value
	r.QueryParams.Set("TaskEnabled", strconv.FormatBool(value))
}
func (r *CreateScheduledTaskRequest) GetTaskEnabled() bool {
	return r.TaskEnabled
}
func (r *CreateScheduledTaskRequest) SetLaunchExpirationTime(value int) {
	r.LaunchExpirationTime = value
	r.QueryParams.Set("LaunchExpirationTime", strconv.Itoa(value))
}
func (r *CreateScheduledTaskRequest) GetLaunchExpirationTime() int {
	return r.LaunchExpirationTime
}
func (r *CreateScheduledTaskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateScheduledTaskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateScheduledTaskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateScheduledTask")
	r.SetProduct(Product)
}

type CreateScheduledTaskResponse struct {
	ScheduledTaskId string `xml:"ScheduledTaskId" json:"ScheduledTaskId"`
}

func CreateScheduledTask(req *CreateScheduledTaskRequest, accessId, accessSecret string) (*CreateScheduledTaskResponse, error) {
	var pResponse CreateScheduledTaskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
