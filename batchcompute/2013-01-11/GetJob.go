package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetJobRequest struct {
	RoaRequest
	ResourceName string
}

func (r *GetJobRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *GetJobRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *GetJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetJobResponse struct {
}

func GetJob(req *GetJobRequest, accessId, accessSecret string) (*GetJobResponse, error) {
	var pResponse GetJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
