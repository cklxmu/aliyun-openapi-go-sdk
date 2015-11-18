package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type StopJobRequest struct {
	RoaRequest
	ResourceName string
}

func (r *StopJobRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *StopJobRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *StopJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName?Action=Stop"
	r.SetMethod(PUT)
	r.SetProduct(Product)
}

type StopJobResponse struct {
}

func StopJob(req *StopJobRequest, accessId, accessSecret string) (*StopJobResponse, error) {
	var pResponse StopJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
