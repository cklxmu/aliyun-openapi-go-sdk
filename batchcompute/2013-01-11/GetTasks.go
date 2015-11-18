package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetTasksRequest struct {
	RoaRequest
	ResourceName string
}

func (r *GetTasksRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *GetTasksRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *GetTasksRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName/tasks"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetTasksResponse struct {
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
