package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDiskRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ZoneId               string
	SnapshotId           string
	DiskName             string
	Size                 int
	DiskCategory         string
	Description          string
	ClientToken          string
	OwnerAccount         string
}

func (r *CreateDiskRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateDiskRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateDiskRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateDiskRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateDiskRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateDiskRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateDiskRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateDiskRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateDiskRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *CreateDiskRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *CreateDiskRequest) SetDiskName(value string) {
	r.DiskName = value
	r.QueryParams.Set("DiskName", value)
}
func (r *CreateDiskRequest) GetDiskName() string {
	return r.DiskName
}
func (r *CreateDiskRequest) SetSize(value int) {
	r.Size = value
	r.QueryParams.Set("Size", strconv.Itoa(value))
}
func (r *CreateDiskRequest) GetSize() int {
	return r.Size
}
func (r *CreateDiskRequest) SetDiskCategory(value string) {
	r.DiskCategory = value
	r.QueryParams.Set("DiskCategory", value)
}
func (r *CreateDiskRequest) GetDiskCategory() string {
	return r.DiskCategory
}
func (r *CreateDiskRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateDiskRequest) GetDescription() string {
	return r.Description
}
func (r *CreateDiskRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateDiskRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateDiskRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateDiskRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateDiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDisk")
	r.SetProduct(Product)
}

type CreateDiskResponse struct {
	DiskId string `xml:"DiskId" json:"DiskId"`
}

func CreateDisk(req *CreateDiskRequest, accessId, accessSecret string) (*CreateDiskResponse, error) {
	var pResponse CreateDiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
