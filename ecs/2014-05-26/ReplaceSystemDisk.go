package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReplaceSystemDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	ImageId              string
	ClientToken          string
	OwnerAccount         string
	UseAdditionalService bool
}

func (r *ReplaceSystemDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ReplaceSystemDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ReplaceSystemDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ReplaceSystemDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ReplaceSystemDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ReplaceSystemDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ReplaceSystemDiskRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ReplaceSystemDiskRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ReplaceSystemDiskRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *ReplaceSystemDiskRequest) GetImageId() string {
	return r.ImageId
}
func (r *ReplaceSystemDiskRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ReplaceSystemDiskRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ReplaceSystemDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ReplaceSystemDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *ReplaceSystemDiskRequest) SetUseAdditionalService(value bool) {
	r.UseAdditionalService = value
	r.QueryParams.Set("UseAdditionalService", strconv.FormatBool(value))
}
func (r *ReplaceSystemDiskRequest) GetUseAdditionalService() bool {
	return r.UseAdditionalService
}

func (r *ReplaceSystemDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReplaceSystemDisk")
	r.SetProduct(Product)
}

type ReplaceSystemDiskResponse struct {
	DiskId string `xml:"DiskId" json:"DiskId"`
}

func ReplaceSystemDisk(req *ReplaceSystemDiskRequest, accessId, accessSecret string) (*ReplaceSystemDiskResponse, error) {
	var pResponse ReplaceSystemDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
