package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateImageRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	SnapshotId           string
	ImageName            string
	ImageVersion         string
	Description          string
	ClientToken          string
	OwnerAccount         string
}

func (r *CreateImageRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateImageRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateImageRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateImageRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateImageRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateImageRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateImageRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *CreateImageRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *CreateImageRequest) SetImageName(value string) {
	r.ImageName = value
	r.QueryParams.Set("ImageName", value)
}
func (r *CreateImageRequest) GetImageName() string {
	return r.ImageName
}
func (r *CreateImageRequest) SetImageVersion(value string) {
	r.ImageVersion = value
	r.QueryParams.Set("ImageVersion", value)
}
func (r *CreateImageRequest) GetImageVersion() string {
	return r.ImageVersion
}
func (r *CreateImageRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateImageRequest) GetDescription() string {
	return r.Description
}
func (r *CreateImageRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateImageRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateImageRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateImageRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateImageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateImage")
	r.SetProduct(Product)
}

type CreateImageResponse struct {
	ImageId string `xml:"ImageId" json:"ImageId"`
}

func CreateImage(req *CreateImageRequest, accessId, accessSecret string) (*CreateImageResponse, error) {
	var pResponse CreateImageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
