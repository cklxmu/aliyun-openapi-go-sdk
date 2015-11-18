package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CopyImageRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	DestinationImageName   string
	DestinationDescription string
	ImageId                string
	DestinationRegionId    string
	OwnerAccount           string
}

func (r *CopyImageRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CopyImageRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CopyImageRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CopyImageRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CopyImageRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CopyImageRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CopyImageRequest) SetDestinationImageName(value string) {
	r.DestinationImageName = value
	r.QueryParams.Set("DestinationImageName", value)
}
func (r *CopyImageRequest) GetDestinationImageName() string {
	return r.DestinationImageName
}
func (r *CopyImageRequest) SetDestinationDescription(value string) {
	r.DestinationDescription = value
	r.QueryParams.Set("DestinationDescription", value)
}
func (r *CopyImageRequest) GetDestinationDescription() string {
	return r.DestinationDescription
}
func (r *CopyImageRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *CopyImageRequest) GetImageId() string {
	return r.ImageId
}
func (r *CopyImageRequest) SetDestinationRegionId(value string) {
	r.DestinationRegionId = value
	r.QueryParams.Set("DestinationRegionId", value)
}
func (r *CopyImageRequest) GetDestinationRegionId() string {
	return r.DestinationRegionId
}
func (r *CopyImageRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CopyImageRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CopyImageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CopyImage")
	r.SetProduct(Product)
}

type CopyImageResponse struct {
	ImageId string `xml:"ImageId" json:"ImageId"`
}

func CopyImage(req *CopyImageRequest, accessId, accessSecret string) (*CopyImageResponse, error) {
	var pResponse CopyImageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
