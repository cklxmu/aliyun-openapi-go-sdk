package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type CancelDDLTaskRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	TaskId         string
}

func (r *CancelDDLTaskRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *CancelDDLTaskRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *CancelDDLTaskRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *CancelDDLTaskRequest) GetDbName() string {
	return r.DbName
}
func (r *CancelDDLTaskRequest) SetTaskId(value string) {
	r.TaskId = value
	r.QueryParams.Set("TaskId", value)
}
func (r *CancelDDLTaskRequest) GetTaskId() string {
	return r.TaskId
}

func (r *CancelDDLTaskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CancelDDLTask")
	r.SetProduct(Product)
}

type CancelDDLTaskResponse struct {
}

func CancelDDLTask(req *CancelDDLTaskRequest, accessId, accessSecret string) (*CancelDDLTaskResponse, error) {
	var pResponse CancelDDLTaskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
