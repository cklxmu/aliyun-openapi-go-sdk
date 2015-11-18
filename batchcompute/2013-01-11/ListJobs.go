package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListJobsRequest struct {
	RoaRequest
}

func (r *ListJobsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type ListJobsResponse struct {
}

func ListJobs(req *ListJobsRequest, accessId, accessSecret string) (*ListJobsResponse, error) {
	var pResponse ListJobsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
