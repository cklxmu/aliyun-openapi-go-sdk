package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ResizeDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	NewSize              int
	ClientToken          string
	OwnerAccount         string
}

func (r *ResizeDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ResizeDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ResizeDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ResizeDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ResizeDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ResizeDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ResizeDiskRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *ResizeDiskRequest) GetDiskId() string {
	return r.DiskId
}
func (r *ResizeDiskRequest) SetNewSize(value int) {
	r.NewSize = value
	r.QueryParams.Set("NewSize", strconv.Itoa(value))
}
func (r *ResizeDiskRequest) GetNewSize() int {
	return r.NewSize
}
func (r *ResizeDiskRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ResizeDiskRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ResizeDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ResizeDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ResizeDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ResizeDisk")
	r.SetProduct(Product)
}

type ResizeDiskResponse struct {
}

func ResizeDisk(req *ResizeDiskRequest, accessId, accessSecret string) (*ResizeDiskResponse, error) {
	var pResponse ResizeDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
