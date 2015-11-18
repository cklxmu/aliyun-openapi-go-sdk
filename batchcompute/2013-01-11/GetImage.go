package batchcompute

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetImageRequest struct {
	RoaRequest
	ResourceName string
}

func (r *GetImageRequest) SetResourceName(value string) {
	r.ResourceName = value
	r.PathParams.Set("ResourceName", value)
}
func (r *GetImageRequest) GetResourceName() string {
	return r.ResourceName
}

func (r *GetImageRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/images/ResourceName"
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetImageResponse struct {
}

func GetImage(req *GetImageRequest, accessId, accessSecret string) (*GetImageResponse, error) {
	var pResponse GetImageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
