package pts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopTaskRequest struct {
	RpcRequest
	TaskId int
	Type   string
	Msg    string
}

func (r *StopTaskRequest) SetTaskId(value int) {
	r.TaskId = value
	r.QueryParams.Set("TaskId", strconv.Itoa(value))
}
func (r *StopTaskRequest) GetTaskId() int {
	return r.TaskId
}
func (r *StopTaskRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("Type", value)
}
func (r *StopTaskRequest) GetType() string {
	return r.Type
}
func (r *StopTaskRequest) SetMsg(value string) {
	r.Msg = value
	r.QueryParams.Set("Msg", value)
}
func (r *StopTaskRequest) GetMsg() string {
	return r.Msg
}

func (r *StopTaskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopTask")
	r.SetProduct(Product)
}

type StopTaskResponse struct {
}

func StopTask(req *StopTaskRequest, accessId, accessSecret string) (*StopTaskResponse, error) {
	var pResponse StopTaskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
