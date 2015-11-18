package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type ReleaseJobRequest struct {
	RoaRequest
	ResourceName string
}

func (r *ReleaseJobRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *ReleaseJobRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *ReleaseJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName"
	r.SetMethod("DELETE")
	r.SetProduct(Product)
}

type ReleaseJobResponse struct {
}

func ReleaseJob(req *ReleaseJobRequest, accessId, accessSecret string) (*ReleaseJobResponse, error) {
	var pResponse ReleaseJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
