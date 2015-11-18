package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateScalingConfigurationRequest struct {
	RpcRequest
	OwnerId                       int
	ResourceOwnerAccount          string
	ResourceOwnerId               int
	ScalingGroupId                string
	ImageId                       string
	InstanceType                  string
	SecurityGroupId               string
	InternetChargeType            string
	InternetMaxBandwidthIn        int
	InternetMaxBandwidthOut       int
	SystemDisk_Category           string
	ScalingConfigurationName      string
	DataDisk_1_Size               int
	DataDisk_2_Size               int
	DataDisk_3_Size               int
	DataDisk_4_Size               int
	DataDisk_1_Category           string
	DataDisk_2_Category           string
	DataDisk_3_Category           string
	DataDisk_4_Category           string
	DataDisk_1_SnapshotId         string
	DataDisk_2_SnapshotId         string
	DataDisk_3_SnapshotId         string
	DataDisk_4_SnapshotId         string
	DataDisk_1_Device             string
	DataDisk_2_Device             string
	DataDisk_3_Device             string
	DataDisk_4_Device             string
	DataDisk_1_DeleteWithInstance string
	DataDisk_2_DeleteWithInstance string
	DataDisk_3_DeleteWithInstance string
	DataDisk_4_DeleteWithInstance string
	OwnerAccount                  string
}

func (r *CreateScalingConfigurationRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateScalingConfigurationRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateScalingConfigurationRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateScalingConfigurationRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateScalingConfigurationRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *CreateScalingConfigurationRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *CreateScalingConfigurationRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *CreateScalingConfigurationRequest) GetImageId() string {
	return r.ImageId
}
func (r *CreateScalingConfigurationRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *CreateScalingConfigurationRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *CreateScalingConfigurationRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *CreateScalingConfigurationRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *CreateScalingConfigurationRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *CreateScalingConfigurationRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *CreateScalingConfigurationRequest) SetInternetMaxBandwidthIn(value int) {
	r.InternetMaxBandwidthIn = value
	r.QueryParams.Set("InternetMaxBandwidthIn", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetInternetMaxBandwidthIn() int {
	return r.InternetMaxBandwidthIn
}
func (r *CreateScalingConfigurationRequest) SetInternetMaxBandwidthOut(value int) {
	r.InternetMaxBandwidthOut = value
	r.QueryParams.Set("InternetMaxBandwidthOut", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetInternetMaxBandwidthOut() int {
	return r.InternetMaxBandwidthOut
}
func (r *CreateScalingConfigurationRequest) SetSystemDisk_Category(value string) {
	r.SystemDisk_Category = value
	r.QueryParams.Set("SystemDisk.Category", value)
}
func (r *CreateScalingConfigurationRequest) GetSystemDisk_Category() string {
	return r.SystemDisk_Category
}
func (r *CreateScalingConfigurationRequest) SetScalingConfigurationName(value string) {
	r.ScalingConfigurationName = value
	r.QueryParams.Set("ScalingConfigurationName", value)
}
func (r *CreateScalingConfigurationRequest) GetScalingConfigurationName() string {
	return r.ScalingConfigurationName
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_1_Size(value int) {
	r.DataDisk_1_Size = value
	r.QueryParams.Set("DataDisk.1.Size", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_1_Size() int {
	return r.DataDisk_1_Size
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_2_Size(value int) {
	r.DataDisk_2_Size = value
	r.QueryParams.Set("DataDisk.2.Size", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_2_Size() int {
	return r.DataDisk_2_Size
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_3_Size(value int) {
	r.DataDisk_3_Size = value
	r.QueryParams.Set("DataDisk.3.Size", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_3_Size() int {
	return r.DataDisk_3_Size
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_4_Size(value int) {
	r.DataDisk_4_Size = value
	r.QueryParams.Set("DataDisk.4.Size", strconv.Itoa(value))
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_4_Size() int {
	return r.DataDisk_4_Size
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_1_Category(value string) {
	r.DataDisk_1_Category = value
	r.QueryParams.Set("DataDisk.1.Category", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_1_Category() string {
	return r.DataDisk_1_Category
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_2_Category(value string) {
	r.DataDisk_2_Category = value
	r.QueryParams.Set("DataDisk.2.Category", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_2_Category() string {
	return r.DataDisk_2_Category
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_3_Category(value string) {
	r.DataDisk_3_Category = value
	r.QueryParams.Set("DataDisk.3.Category", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_3_Category() string {
	return r.DataDisk_3_Category
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_4_Category(value string) {
	r.DataDisk_4_Category = value
	r.QueryParams.Set("DataDisk.4.Category", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_4_Category() string {
	return r.DataDisk_4_Category
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_1_SnapshotId(value string) {
	r.DataDisk_1_SnapshotId = value
	r.QueryParams.Set("DataDisk.1.SnapshotId", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_1_SnapshotId() string {
	return r.DataDisk_1_SnapshotId
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_2_SnapshotId(value string) {
	r.DataDisk_2_SnapshotId = value
	r.QueryParams.Set("DataDisk.2.SnapshotId", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_2_SnapshotId() string {
	return r.DataDisk_2_SnapshotId
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_3_SnapshotId(value string) {
	r.DataDisk_3_SnapshotId = value
	r.QueryParams.Set("DataDisk.3.SnapshotId", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_3_SnapshotId() string {
	return r.DataDisk_3_SnapshotId
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_4_SnapshotId(value string) {
	r.DataDisk_4_SnapshotId = value
	r.QueryParams.Set("DataDisk.4.SnapshotId", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_4_SnapshotId() string {
	return r.DataDisk_4_SnapshotId
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_1_Device(value string) {
	r.DataDisk_1_Device = value
	r.QueryParams.Set("DataDisk.1.Device", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_1_Device() string {
	return r.DataDisk_1_Device
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_2_Device(value string) {
	r.DataDisk_2_Device = value
	r.QueryParams.Set("DataDisk.2.Device", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_2_Device() string {
	return r.DataDisk_2_Device
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_3_Device(value string) {
	r.DataDisk_3_Device = value
	r.QueryParams.Set("DataDisk.3.Device", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_3_Device() string {
	return r.DataDisk_3_Device
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_4_Device(value string) {
	r.DataDisk_4_Device = value
	r.QueryParams.Set("DataDisk.4.Device", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_4_Device() string {
	return r.DataDisk_4_Device
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_1_DeleteWithInstance(value string) {
	r.DataDisk_1_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.1.DeleteWithInstance", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_1_DeleteWithInstance() string {
	return r.DataDisk_1_DeleteWithInstance
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_2_DeleteWithInstance(value string) {
	r.DataDisk_2_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.2.DeleteWithInstance", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_2_DeleteWithInstance() string {
	return r.DataDisk_2_DeleteWithInstance
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_3_DeleteWithInstance(value string) {
	r.DataDisk_3_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.3.DeleteWithInstance", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_3_DeleteWithInstance() string {
	return r.DataDisk_3_DeleteWithInstance
}
func (r *CreateScalingConfigurationRequest) SetDataDisk_4_DeleteWithInstance(value string) {
	r.DataDisk_4_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.4.DeleteWithInstance", value)
}
func (r *CreateScalingConfigurationRequest) GetDataDisk_4_DeleteWithInstance() string {
	return r.DataDisk_4_DeleteWithInstance
}
func (r *CreateScalingConfigurationRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateScalingConfigurationRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateScalingConfigurationRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateScalingConfiguration")
	r.SetProduct(Product)
}

type CreateScalingConfigurationResponse struct {
	ScalingConfigurationId string `xml:"ScalingConfigurationId" json:"ScalingConfigurationId"`
}

func CreateScalingConfiguration(req *CreateScalingConfigurationRequest, accessId, accessSecret string) (*CreateScalingConfigurationResponse, error) {
	var pResponse CreateScalingConfigurationResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
