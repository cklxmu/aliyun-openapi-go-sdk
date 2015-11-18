package pts

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetTasksRequest struct {
	RpcRequest
	Status string
}

func (r *GetTasksRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *GetTasksRequest) GetStatus() string {
	return r.Status
}

func (r *GetTasksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetTasks")
	r.SetProduct(Product)
}

type GetTasksResponse struct {
	Tasks string `xml:"Tasks" json:"Tasks"`
}

func GetTasks(req *GetTasksRequest, accessId, accessSecret string) (*GetTasksResponse, error) {
	var pResponse GetTasksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
