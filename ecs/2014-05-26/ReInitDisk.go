package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReInitDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DiskId               string
	OwnerAccount         string
}

func (r *ReInitDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ReInitDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ReInitDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ReInitDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ReInitDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ReInitDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ReInitDiskRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *ReInitDiskRequest) GetDiskId() string {
	return r.DiskId
}
func (r *ReInitDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ReInitDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ReInitDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReInitDisk")
	r.SetProduct(Product)
}

type ReInitDiskResponse struct {
}

func ReInitDisk(req *ReInitDiskRequest, accessId, accessSecret string) (*ReInitDiskResponse, error) {
	var pResponse ReInitDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
