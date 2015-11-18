package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CheckAutoSnapshotPolicyRequest struct {
	RpcRequest
	OwnerId                           int
	ResourceOwnerAccount              string
	ResourceOwnerId                   int
	OwnerAccount                      string
	SystemDiskPolicyEnabled           bool
	SystemDiskPolicyTimePeriod        int
	SystemDiskPolicyRetentionDays     int
	SystemDiskPolicyRetentionLastWeek bool
	DataDiskPolicyEnabled             bool
	DataDiskPolicyTimePeriod          int
	DataDiskPolicyRetentionDays       int
	DataDiskPolicyRetentionLastWeek   bool
}

func (r *CheckAutoSnapshotPolicyRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CheckAutoSnapshotPolicyRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CheckAutoSnapshotPolicyRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CheckAutoSnapshotPolicyRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CheckAutoSnapshotPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CheckAutoSnapshotPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CheckAutoSnapshotPolicyRequest) SetSystemDiskPolicyEnabled(value bool) {
	r.SystemDiskPolicyEnabled = value
	r.QueryParams.Set("SystemDiskPolicyEnabled", strconv.FormatBool(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetSystemDiskPolicyEnabled() bool {
	return r.SystemDiskPolicyEnabled
}
func (r *CheckAutoSnapshotPolicyRequest) SetSystemDiskPolicyTimePeriod(value int) {
	r.SystemDiskPolicyTimePeriod = value
	r.QueryParams.Set("SystemDiskPolicyTimePeriod", strconv.Itoa(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetSystemDiskPolicyTimePeriod() int {
	return r.SystemDiskPolicyTimePeriod
}
func (r *CheckAutoSnapshotPolicyRequest) SetSystemDiskPolicyRetentionDays(value int) {
	r.SystemDiskPolicyRetentionDays = value
	r.QueryParams.Set("SystemDiskPolicyRetentionDays", strconv.Itoa(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetSystemDiskPolicyRetentionDays() int {
	return r.SystemDiskPolicyRetentionDays
}
func (r *CheckAutoSnapshotPolicyRequest) SetSystemDiskPolicyRetentionLastWeek(value bool) {
	r.SystemDiskPolicyRetentionLastWeek = value
	r.QueryParams.Set("SystemDiskPolicyRetentionLastWeek", strconv.FormatBool(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetSystemDiskPolicyRetentionLastWeek() bool {
	return r.SystemDiskPolicyRetentionLastWeek
}
func (r *CheckAutoSnapshotPolicyRequest) SetDataDiskPolicyEnabled(value bool) {
	r.DataDiskPolicyEnabled = value
	r.QueryParams.Set("DataDiskPolicyEnabled", strconv.FormatBool(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetDataDiskPolicyEnabled() bool {
	return r.DataDiskPolicyEnabled
}
func (r *CheckAutoSnapshotPolicyRequest) SetDataDiskPolicyTimePeriod(value int) {
	r.DataDiskPolicyTimePeriod = value
	r.QueryParams.Set("DataDiskPolicyTimePeriod", strconv.Itoa(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetDataDiskPolicyTimePeriod() int {
	return r.DataDiskPolicyTimePeriod
}
func (r *CheckAutoSnapshotPolicyRequest) SetDataDiskPolicyRetentionDays(value int) {
	r.DataDiskPolicyRetentionDays = value
	r.QueryParams.Set("DataDiskPolicyRetentionDays", strconv.Itoa(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetDataDiskPolicyRetentionDays() int {
	return r.DataDiskPolicyRetentionDays
}
func (r *CheckAutoSnapshotPolicyRequest) SetDataDiskPolicyRetentionLastWeek(value bool) {
	r.DataDiskPolicyRetentionLastWeek = value
	r.QueryParams.Set("DataDiskPolicyRetentionLastWeek", strconv.FormatBool(value))
}
func (r *CheckAutoSnapshotPolicyRequest) GetDataDiskPolicyRetentionLastWeek() bool {
	return r.DataDiskPolicyRetentionLastWeek
}

func (r *CheckAutoSnapshotPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CheckAutoSnapshotPolicy")
	r.SetProduct(Product)
}

type CheckAutoSnapshotPolicyResponse struct {
	AutoSnapshotOccupation int    `xml:"AutoSnapshotOccupation" json:"AutoSnapshotOccupation"`
	IsPermittedModify      string `xml:"IsPermittedModify" json:"IsPermittedModify"`
}

func CheckAutoSnapshotPolicy(req *CheckAutoSnapshotPolicyRequest, accessId, accessSecret string) (*CheckAutoSnapshotPolicyResponse, error) {
	var pResponse CheckAutoSnapshotPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
