package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AttachDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	DiskId               string
	Device               string
	DeleteWithInstance   bool
	OwnerAccount         string
}

func (r *AttachDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AttachDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AttachDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AttachDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AttachDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AttachDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AttachDiskRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *AttachDiskRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *AttachDiskRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *AttachDiskRequest) GetDiskId() string {
	return r.DiskId
}
func (r *AttachDiskRequest) SetDevice(value string) {
	r.Device = value
	r.QueryParams.Set("Device", value)
}
func (r *AttachDiskRequest) GetDevice() string {
	return r.Device
}
func (r *AttachDiskRequest) SetDeleteWithInstance(value bool) {
	r.DeleteWithInstance = value
	r.QueryParams.Set("DeleteWithInstance", strconv.FormatBool(value))
}
func (r *AttachDiskRequest) GetDeleteWithInstance() bool {
	return r.DeleteWithInstance
}
func (r *AttachDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *AttachDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *AttachDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AttachDisk")
	r.SetProduct(Product)
}

type AttachDiskResponse struct {
}

func AttachDisk(req *AttachDiskRequest, accessId, accessSecret string) (*AttachDiskResponse, error) {
	var pResponse AttachDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
