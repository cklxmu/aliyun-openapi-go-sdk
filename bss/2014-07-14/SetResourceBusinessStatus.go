package bss

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetResourceBusinessStatusRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ResourceType         string
	ResourceId           string
	BusinessStatus       string
	OwnerAccount         string
}

func (r *SetResourceBusinessStatusRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetResourceBusinessStatusRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetResourceBusinessStatusRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetResourceBusinessStatusRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetResourceBusinessStatusRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetResourceBusinessStatusRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetResourceBusinessStatusRequest) SetResourceType(value string) {
	r.ResourceType = value
	r.QueryParams.Set("ResourceType", value)
}
func (r *SetResourceBusinessStatusRequest) GetResourceType() string {
	return r.ResourceType
}
func (r *SetResourceBusinessStatusRequest) SetResourceId(value string) {
	r.ResourceId = value
	r.QueryParams.Set("ResourceId", value)
}
func (r *SetResourceBusinessStatusRequest) GetResourceId() string {
	return r.ResourceId
}
func (r *SetResourceBusinessStatusRequest) SetBusinessStatus(value string) {
	r.BusinessStatus = value
	r.QueryParams.Set("BusinessStatus", value)
}
func (r *SetResourceBusinessStatusRequest) GetBusinessStatus() string {
	return r.BusinessStatus
}
func (r *SetResourceBusinessStatusRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetResourceBusinessStatusRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetResourceBusinessStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetResourceBusinessStatus")
	r.SetProduct(Product)
}

type SetResourceBusinessStatusResponse struct {
}

func SetResourceBusinessStatus(req *SetResourceBusinessStatusRequest, accessId, accessSecret string) (*SetResourceBusinessStatusResponse, error) {
	var pResponse SetResourceBusinessStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
