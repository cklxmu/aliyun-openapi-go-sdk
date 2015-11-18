package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RenewInstanceRequest struct {
	RpcRequest
	OwnerId                     int
	ResourceOwnerAccount        string
	ResourceOwnerId             int
	OwnerAccount                string
	InstanceId                  string
	InstanceType                string
	InternetMaxBandwidthOut     int
	InternetChargeType          string
	Period                      int
	RebootTime                  string
	CovertDiskPortable_1_DiskId string
	CovertDiskPortable_2_DiskId string
	CovertDiskPortable_3_DiskId string
	CovertDiskPortable_4_DiskId string
}

func (r *RenewInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RenewInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RenewInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RenewInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RenewInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RenewInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *RenewInstanceRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *RenewInstanceRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *RenewInstanceRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *RenewInstanceRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *RenewInstanceRequest) SetInternetMaxBandwidthOut(value int) {
	r.InternetMaxBandwidthOut = value
	r.QueryParams.Set("InternetMaxBandwidthOut", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetInternetMaxBandwidthOut() int {
	return r.InternetMaxBandwidthOut
}
func (r *RenewInstanceRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *RenewInstanceRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *RenewInstanceRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *RenewInstanceRequest) GetPeriod() int {
	return r.Period
}
func (r *RenewInstanceRequest) SetRebootTime(value string) {
	r.RebootTime = value
	r.QueryParams.Set("RebootTime", value)
}
func (r *RenewInstanceRequest) GetRebootTime() string {
	return r.RebootTime
}
func (r *RenewInstanceRequest) SetCovertDiskPortable_1_DiskId(value string) {
	r.CovertDiskPortable_1_DiskId = value
	r.QueryParams.Set("CovertDiskPortable.1.DiskId", value)
}
func (r *RenewInstanceRequest) GetCovertDiskPortable_1_DiskId() string {
	return r.CovertDiskPortable_1_DiskId
}
func (r *RenewInstanceRequest) SetCovertDiskPortable_2_DiskId(value string) {
	r.CovertDiskPortable_2_DiskId = value
	r.QueryParams.Set("CovertDiskPortable.2.DiskId", value)
}
func (r *RenewInstanceRequest) GetCovertDiskPortable_2_DiskId() string {
	return r.CovertDiskPortable_2_DiskId
}
func (r *RenewInstanceRequest) SetCovertDiskPortable_3_DiskId(value string) {
	r.CovertDiskPortable_3_DiskId = value
	r.QueryParams.Set("CovertDiskPortable.3.DiskId", value)
}
func (r *RenewInstanceRequest) GetCovertDiskPortable_3_DiskId() string {
	return r.CovertDiskPortable_3_DiskId
}
func (r *RenewInstanceRequest) SetCovertDiskPortable_4_DiskId(value string) {
	r.CovertDiskPortable_4_DiskId = value
	r.QueryParams.Set("CovertDiskPortable.4.DiskId", value)
}
func (r *RenewInstanceRequest) GetCovertDiskPortable_4_DiskId() string {
	return r.CovertDiskPortable_4_DiskId
}

func (r *RenewInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RenewInstance")
	r.SetProduct(Product)
}

type RenewInstanceResponse struct {
}

func RenewInstance(req *RenewInstanceRequest, accessId, accessSecret string) (*RenewInstanceResponse, error) {
	var pResponse RenewInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
