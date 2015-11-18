package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type StartJobRequest struct {
	RoaRequest
	ResourceName string
}

func (r *StartJobRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *StartJobRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *StartJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName?Action=Start"
	r.SetMethod(PUT)
	r.SetProduct(Product)
}

type StartJobResponse struct {
}

func StartJob(req *StartJobRequest, accessId, accessSecret string) (*StartJobResponse, error) {
	var pResponse StartJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
