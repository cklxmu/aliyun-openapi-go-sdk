package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyImageAttributeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ImageId              string
	ImageName            string
	Description          string
	OwnerAccount         string
}

func (r *ModifyImageAttributeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyImageAttributeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyImageAttributeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyImageAttributeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyImageAttributeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyImageAttributeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyImageAttributeRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *ModifyImageAttributeRequest) GetImageId() string {
	return r.ImageId
}
func (r *ModifyImageAttributeRequest) SetImageName(value string) {
	r.ImageName = value
	r.QueryParams.Set("ImageName", value)
}
func (r *ModifyImageAttributeRequest) GetImageName() string {
	return r.ImageName
}
func (r *ModifyImageAttributeRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *ModifyImageAttributeRequest) GetDescription() string {
	return r.Description
}
func (r *ModifyImageAttributeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyImageAttributeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyImageAttributeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyImageAttribute")
	r.SetProduct(Product)
}

type ModifyImageAttributeResponse struct {
}

func ModifyImageAttribute(req *ModifyImageAttributeRequest, accessId, accessSecret string) (*ModifyImageAttributeResponse, error) {
	var pResponse ModifyImageAttributeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
