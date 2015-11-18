package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListImagesRequest struct {
	RoaRequest
}

func (r *ListImagesRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/images"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type ListImagesResponse struct {
}

func ListImages(req *ListImagesRequest, accessId, accessSecret string) (*ListImagesResponse, error) {
	var pResponse ListImagesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
