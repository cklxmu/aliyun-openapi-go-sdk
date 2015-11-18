package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type PostJobRequest struct {
	RoaRequest
}

func (r *PostJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs"
	r.SetMethod("POST")
	r.SetProduct(Product)
}

type PostJobResponse struct {
}

func PostJob(req *PostJobRequest, accessId, accessSecret string) (*PostJobResponse, error) {
	var pResponse PostJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
