package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetJobDescriptionRequest struct {
	RoaRequest
	ResourceName string
}

func (r *GetJobDescriptionRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.QueryParams.Set("ResourceName", value)
}
func (r *GetJobDescriptionRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *GetJobDescriptionRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/jobs/ResourceName/description"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetJobDescriptionResponse struct {
}

func GetJobDescription(req *GetJobDescriptionRequest, accessId, accessSecret string) (*GetJobDescriptionResponse, error) {
	var pResponse GetJobDescriptionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
