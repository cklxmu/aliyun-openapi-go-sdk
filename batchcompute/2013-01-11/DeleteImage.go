package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteImageRequest struct {
	RoaRequest
	ResourceName string
}

func (r *DeleteImageRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *DeleteImageRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *DeleteImageRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/images/ResourceName"
	r.SetMethod("DELETE")
	r.SetProduct(Product)
}

type DeleteImageResponse struct {
}

func DeleteImage(req *DeleteImageRequest, accessId, accessSecret string) (*DeleteImageResponse, error) {
	var pResponse DeleteImageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
