package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteImageRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ImageId              string
	OwnerAccount         string
}

func (r *DeleteImageRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteImageRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteImageRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteImageRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteImageRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteImageRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteImageRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *DeleteImageRequest) GetImageId() string {
	return r.ImageId
}
func (r *DeleteImageRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteImageRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteImageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteImage")
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
