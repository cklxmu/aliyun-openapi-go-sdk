package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type PutJobRequest struct {
	RoaRequest
	ResourceName string
}

func (r *PutJobRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *PutJobRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *PutJobRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName/Priority"
	r.SetMethod(PUT)
	r.SetProduct(Product)
}

type PutJobResponse struct {
}

func PutJob(req *PutJobRequest, accessId, accessSecret string) (*PutJobResponse, error) {
	var pResponse PutJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
