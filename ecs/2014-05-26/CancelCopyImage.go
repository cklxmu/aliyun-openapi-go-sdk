package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CancelCopyImageRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ImageId              string
	OwnerAccount         string
}

func (r *CancelCopyImageRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CancelCopyImageRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CancelCopyImageRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CancelCopyImageRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CancelCopyImageRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CancelCopyImageRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CancelCopyImageRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *CancelCopyImageRequest) GetImageId() string {
	return r.ImageId
}
func (r *CancelCopyImageRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CancelCopyImageRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CancelCopyImageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CancelCopyImage")
	r.SetProduct(Product)
}

type CancelCopyImageResponse struct {
}

func CancelCopyImage(req *CancelCopyImageRequest, accessId, accessSecret string) (*CancelCopyImageResponse, error) {
	var pResponse CancelCopyImageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
