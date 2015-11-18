package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyAutoSnapshotPolicyRequest struct {
	RpcRequest
	OwnerId                           int
	ResourceOwnerAccount              string
	ResourceOwnerId                   int
	SystemDiskPolicyEnabled           bool
	SystemDiskPolicyTimePeriod        int
	SystemDiskPolicyRetentionDays     int
	SystemDiskPolicyRetentionLastWeek bool
	DataDiskPolicyEnabled             bool
	DataDiskPolicyTimePeriod          int
	DataDiskPolicyRetentionDays       int
	DataDiskPolicyRetentionLastWeek   bool
	OwnerAccount                      string
}

func (r *ModifyAutoSnapshotPolicyRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyAutoSnapshotPolicyRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyAutoSnapshotPolicyRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyAutoSnapshotPolicyRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyEnabled(value bool) {
	r.SystemDiskPolicyEnabled = value
	r.QueryParams.Set("SystemDiskPolicyEnabled", strconv.FormatBool(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyEnabled() bool {
	return r.SystemDiskPolicyEnabled
}
func (r *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyTimePeriod(value int) {
	r.SystemDiskPolicyTimePeriod = value
	r.QueryParams.Set("SystemDiskPolicyTimePeriod", strconv.Itoa(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyTimePeriod() int {
	return r.SystemDiskPolicyTimePeriod
}
func (r *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyRetentionDays(value int) {
	r.SystemDiskPolicyRetentionDays = value
	r.QueryParams.Set("SystemDiskPolicyRetentionDays", strconv.Itoa(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyRetentionDays() int {
	return r.SystemDiskPolicyRetentionDays
}
func (r *ModifyAutoSnapshotPolicyRequest) SetSystemDiskPolicyRetentionLastWeek(value bool) {
	r.SystemDiskPolicyRetentionLastWeek = value
	r.QueryParams.Set("SystemDiskPolicyRetentionLastWeek", strconv.FormatBool(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetSystemDiskPolicyRetentionLastWeek() bool {
	return r.SystemDiskPolicyRetentionLastWeek
}
func (r *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyEnabled(value bool) {
	r.DataDiskPolicyEnabled = value
	r.QueryParams.Set("DataDiskPolicyEnabled", strconv.FormatBool(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyEnabled() bool {
	return r.DataDiskPolicyEnabled
}
func (r *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyTimePeriod(value int) {
	r.DataDiskPolicyTimePeriod = value
	r.QueryParams.Set("DataDiskPolicyTimePeriod", strconv.Itoa(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyTimePeriod() int {
	return r.DataDiskPolicyTimePeriod
}
func (r *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyRetentionDays(value int) {
	r.DataDiskPolicyRetentionDays = value
	r.QueryParams.Set("DataDiskPolicyRetentionDays", strconv.Itoa(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyRetentionDays() int {
	return r.DataDiskPolicyRetentionDays
}
func (r *ModifyAutoSnapshotPolicyRequest) SetDataDiskPolicyRetentionLastWeek(value bool) {
	r.DataDiskPolicyRetentionLastWeek = value
	r.QueryParams.Set("DataDiskPolicyRetentionLastWeek", strconv.FormatBool(value))
}
func (r *ModifyAutoSnapshotPolicyRequest) GetDataDiskPolicyRetentionLastWeek() bool {
	return r.DataDiskPolicyRetentionLastWeek
}
func (r *ModifyAutoSnapshotPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyAutoSnapshotPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyAutoSnapshotPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyAutoSnapshotPolicy")
	r.SetProduct(Product)
}

type ModifyAutoSnapshotPolicyResponse struct {
}

func ModifyAutoSnapshotPolicy(req *ModifyAutoSnapshotPolicyRequest, accessId, accessSecret string) (*ModifyAutoSnapshotPolicyResponse, error) {
	var pResponse ModifyAutoSnapshotPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
