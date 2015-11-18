package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteJobRequest struct {
	RoaRequest
	ResourceName string
}

func (r *DeleteJobRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *DeleteJobRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *DeleteJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName"
	r.SetMethod("DELETE")
	r.SetProduct(Product)
}

type DeleteJobResponse struct {
}

func DeleteJob(req *DeleteJobRequest, accessId, accessSecret string) (*DeleteJobResponse, error) {
	var pResponse DeleteJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
