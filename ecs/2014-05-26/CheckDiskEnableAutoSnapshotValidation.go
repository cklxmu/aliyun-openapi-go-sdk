package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CheckDiskEnableAutoSnapshotValidationRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	OwnerAccount         string
	DiskIds              string
}

func (r *CheckDiskEnableAutoSnapshotValidationRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) SetDiskIds(value string) {
	r.DiskIds = value
	r.QueryParams.Set("DiskIds", value)
}
func (r *CheckDiskEnableAutoSnapshotValidationRequest) GetDiskIds() string {
	return r.DiskIds
}

func (r *CheckDiskEnableAutoSnapshotValidationRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CheckDiskEnableAutoSnapshotValidation")
	r.SetProduct(Product)
}

type CheckDiskEnableAutoSnapshotValidationResponse struct {
	IsPermitted            string `xml:"IsPermitted" json:"IsPermitted"`
	AutoSnapshotOccupation int    `xml:"AutoSnapshotOccupation" json:"AutoSnapshotOccupation"`
}

func CheckDiskEnableAutoSnapshotValidation(req *CheckDiskEnableAutoSnapshotValidationRequest, accessId, accessSecret string) (*CheckDiskEnableAutoSnapshotValidationResponse, error) {
	var pResponse CheckDiskEnableAutoSnapshotValidationResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
