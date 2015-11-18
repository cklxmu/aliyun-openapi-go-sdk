package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateInstanceRequest struct {
	RpcRequest
	OwnerId                       int
	ResourceOwnerAccount          string
	ResourceOwnerId               int
	ImageId                       string
	InstanceType                  string
	SecurityGroupId               string
	InstanceName                  string
	InternetChargeType            string
	InternetMaxBandwidthIn        int
	InternetMaxBandwidthOut       int
	HostName                      string
	Password                      string
	ZoneId                        string
	ClusterId                     string
	ClientToken                   string
	VlanId                        string
	InnerIpAddress                string
	SystemDisk_Category           string
	SystemDisk_DiskName           string
	SystemDisk_Description        string
	DataDisk_1_Size               int
	DataDisk_1_Category           string
	DataDisk_1_SnapshotId         string
	DataDisk_1_DiskName           string
	DataDisk_1_Description        string
	DataDisk_1_Device             string
	DataDisk_1_DeleteWithInstance bool
	DataDisk_2_Size               int
	DataDisk_2_Category           string
	DataDisk_2_SnapshotId         string
	DataDisk_2_DiskName           string
	DataDisk_2_Description        string
	DataDisk_2_Device             string
	DataDisk_2_DeleteWithInstance bool
	DataDisk_3_Size               int
	DataDisk_3_Category           string
	DataDisk_3_SnapshotId         string
	DataDisk_3_DiskName           string
	DataDisk_3_Description        string
	DataDisk_3_Device             string
	DataDisk_3_DeleteWithInstance bool
	DataDisk_4_Size               int
	DataDisk_4_Category           string
	DataDisk_4_SnapshotId         string
	DataDisk_4_DiskName           string
	DataDisk_4_Description        string
	DataDisk_4_Device             string
	DataDisk_4_DeleteWithInstance bool
	NodeControllerId              string
	Description                   string
	VSwitchId                     string
	PrivateIpAddress              string
	IoOptimized                   string
	OwnerAccount                  string
	UseAdditionalService          bool
	InstanceChargeType            string
	Period                        int
}

func (r *CreateInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateInstanceRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *CreateInstanceRequest) GetImageId() string {
	return r.ImageId
}
func (r *CreateInstanceRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *CreateInstanceRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *CreateInstanceRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *CreateInstanceRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *CreateInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *CreateInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *CreateInstanceRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *CreateInstanceRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *CreateInstanceRequest) SetInternetMaxBandwidthIn(value int) {
	r.InternetMaxBandwidthIn = value
	r.QueryParams.Set("InternetMaxBandwidthIn", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetInternetMaxBandwidthIn() int {
	return r.InternetMaxBandwidthIn
}
func (r *CreateInstanceRequest) SetInternetMaxBandwidthOut(value int) {
	r.InternetMaxBandwidthOut = value
	r.QueryParams.Set("InternetMaxBandwidthOut", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetInternetMaxBandwidthOut() int {
	return r.InternetMaxBandwidthOut
}
func (r *CreateInstanceRequest) SetHostName(value string) {
	r.HostName = value
	r.QueryParams.Set("HostName", value)
}
func (r *CreateInstanceRequest) GetHostName() string {
	return r.HostName
}
func (r *CreateInstanceRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *CreateInstanceRequest) GetPassword() string {
	return r.Password
}
func (r *CreateInstanceRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateInstanceRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateInstanceRequest) SetClusterId(value string) {
	r.ClusterId = value
	r.QueryParams.Set("ClusterId", value)
}
func (r *CreateInstanceRequest) GetClusterId() string {
	return r.ClusterId
}
func (r *CreateInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateInstanceRequest) SetVlanId(value string) {
	r.VlanId = value
	r.QueryParams.Set("VlanId", value)
}
func (r *CreateInstanceRequest) GetVlanId() string {
	return r.VlanId
}
func (r *CreateInstanceRequest) SetInnerIpAddress(value string) {
	r.InnerIpAddress = value
	r.QueryParams.Set("InnerIpAddress", value)
}
func (r *CreateInstanceRequest) GetInnerIpAddress() string {
	return r.InnerIpAddress
}
func (r *CreateInstanceRequest) SetSystemDisk_Category(value string) {
	r.SystemDisk_Category = value
	r.QueryParams.Set("SystemDisk.Category", value)
}
func (r *CreateInstanceRequest) GetSystemDisk_Category() string {
	return r.SystemDisk_Category
}
func (r *CreateInstanceRequest) SetSystemDisk_DiskName(value string) {
	r.SystemDisk_DiskName = value
	r.QueryParams.Set("SystemDisk.DiskName", value)
}
func (r *CreateInstanceRequest) GetSystemDisk_DiskName() string {
	return r.SystemDisk_DiskName
}
func (r *CreateInstanceRequest) SetSystemDisk_Description(value string) {
	r.SystemDisk_Description = value
	r.QueryParams.Set("SystemDisk.Description", value)
}
func (r *CreateInstanceRequest) GetSystemDisk_Description() string {
	return r.SystemDisk_Description
}
func (r *CreateInstanceRequest) SetDataDisk_1_Size(value int) {
	r.DataDisk_1_Size = value
	r.QueryParams.Set("DataDisk.1.Size", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetDataDisk_1_Size() int {
	return r.DataDisk_1_Size
}
func (r *CreateInstanceRequest) SetDataDisk_1_Category(value string) {
	r.DataDisk_1_Category = value
	r.QueryParams.Set("DataDisk.1.Category", value)
}
func (r *CreateInstanceRequest) GetDataDisk_1_Category() string {
	return r.DataDisk_1_Category
}
func (r *CreateInstanceRequest) SetDataDisk_1_SnapshotId(value string) {
	r.DataDisk_1_SnapshotId = value
	r.QueryParams.Set("DataDisk.1.SnapshotId", value)
}
func (r *CreateInstanceRequest) GetDataDisk_1_SnapshotId() string {
	return r.DataDisk_1_SnapshotId
}
func (r *CreateInstanceRequest) SetDataDisk_1_DiskName(value string) {
	r.DataDisk_1_DiskName = value
	r.QueryParams.Set("DataDisk.1.DiskName", value)
}
func (r *CreateInstanceRequest) GetDataDisk_1_DiskName() string {
	return r.DataDisk_1_DiskName
}
func (r *CreateInstanceRequest) SetDataDisk_1_Description(value string) {
	r.DataDisk_1_Description = value
	r.QueryParams.Set("DataDisk.1.Description", value)
}
func (r *CreateInstanceRequest) GetDataDisk_1_Description() string {
	return r.DataDisk_1_Description
}
func (r *CreateInstanceRequest) SetDataDisk_1_Device(value string) {
	r.DataDisk_1_Device = value
	r.QueryParams.Set("DataDisk.1.Device", value)
}
func (r *CreateInstanceRequest) GetDataDisk_1_Device() string {
	return r.DataDisk_1_Device
}
func (r *CreateInstanceRequest) SetDataDisk_1_DeleteWithInstance(value bool) {
	r.DataDisk_1_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.1.DeleteWithInstance", strconv.FormatBool(value))
}
func (r *CreateInstanceRequest) GetDataDisk_1_DeleteWithInstance() bool {
	return r.DataDisk_1_DeleteWithInstance
}
func (r *CreateInstanceRequest) SetDataDisk_2_Size(value int) {
	r.DataDisk_2_Size = value
	r.QueryParams.Set("DataDisk.2.Size", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetDataDisk_2_Size() int {
	return r.DataDisk_2_Size
}
func (r *CreateInstanceRequest) SetDataDisk_2_Category(value string) {
	r.DataDisk_2_Category = value
	r.QueryParams.Set("DataDisk.2.Category", value)
}
func (r *CreateInstanceRequest) GetDataDisk_2_Category() string {
	return r.DataDisk_2_Category
}
func (r *CreateInstanceRequest) SetDataDisk_2_SnapshotId(value string) {
	r.DataDisk_2_SnapshotId = value
	r.QueryParams.Set("DataDisk.2.SnapshotId", value)
}
func (r *CreateInstanceRequest) GetDataDisk_2_SnapshotId() string {
	return r.DataDisk_2_SnapshotId
}
func (r *CreateInstanceRequest) SetDataDisk_2_DiskName(value string) {
	r.DataDisk_2_DiskName = value
	r.QueryParams.Set("DataDisk.2.DiskName", value)
}
func (r *CreateInstanceRequest) GetDataDisk_2_DiskName() string {
	return r.DataDisk_2_DiskName
}
func (r *CreateInstanceRequest) SetDataDisk_2_Description(value string) {
	r.DataDisk_2_Description = value
	r.QueryParams.Set("DataDisk.2.Description", value)
}
func (r *CreateInstanceRequest) GetDataDisk_2_Description() string {
	return r.DataDisk_2_Description
}
func (r *CreateInstanceRequest) SetDataDisk_2_Device(value string) {
	r.DataDisk_2_Device = value
	r.QueryParams.Set("DataDisk.2.Device", value)
}
func (r *CreateInstanceRequest) GetDataDisk_2_Device() string {
	return r.DataDisk_2_Device
}
func (r *CreateInstanceRequest) SetDataDisk_2_DeleteWithInstance(value bool) {
	r.DataDisk_2_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.2.DeleteWithInstance", strconv.FormatBool(value))
}
func (r *CreateInstanceRequest) GetDataDisk_2_DeleteWithInstance() bool {
	return r.DataDisk_2_DeleteWithInstance
}
func (r *CreateInstanceRequest) SetDataDisk_3_Size(value int) {
	r.DataDisk_3_Size = value
	r.QueryParams.Set("DataDisk.3.Size", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetDataDisk_3_Size() int {
	return r.DataDisk_3_Size
}
func (r *CreateInstanceRequest) SetDataDisk_3_Category(value string) {
	r.DataDisk_3_Category = value
	r.QueryParams.Set("DataDisk.3.Category", value)
}
func (r *CreateInstanceRequest) GetDataDisk_3_Category() string {
	return r.DataDisk_3_Category
}
func (r *CreateInstanceRequest) SetDataDisk_3_SnapshotId(value string) {
	r.DataDisk_3_SnapshotId = value
	r.QueryParams.Set("DataDisk.3.SnapshotId", value)
}
func (r *CreateInstanceRequest) GetDataDisk_3_SnapshotId() string {
	return r.DataDisk_3_SnapshotId
}
func (r *CreateInstanceRequest) SetDataDisk_3_DiskName(value string) {
	r.DataDisk_3_DiskName = value
	r.QueryParams.Set("DataDisk.3.DiskName", value)
}
func (r *CreateInstanceRequest) GetDataDisk_3_DiskName() string {
	return r.DataDisk_3_DiskName
}
func (r *CreateInstanceRequest) SetDataDisk_3_Description(value string) {
	r.DataDisk_3_Description = value
	r.QueryParams.Set("DataDisk.3.Description", value)
}
func (r *CreateInstanceRequest) GetDataDisk_3_Description() string {
	return r.DataDisk_3_Description
}
func (r *CreateInstanceRequest) SetDataDisk_3_Device(value string) {
	r.DataDisk_3_Device = value
	r.QueryParams.Set("DataDisk.3.Device", value)
}
func (r *CreateInstanceRequest) GetDataDisk_3_Device() string {
	return r.DataDisk_3_Device
}
func (r *CreateInstanceRequest) SetDataDisk_3_DeleteWithInstance(value bool) {
	r.DataDisk_3_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.3.DeleteWithInstance", strconv.FormatBool(value))
}
func (r *CreateInstanceRequest) GetDataDisk_3_DeleteWithInstance() bool {
	return r.DataDisk_3_DeleteWithInstance
}
func (r *CreateInstanceRequest) SetDataDisk_4_Size(value int) {
	r.DataDisk_4_Size = value
	r.QueryParams.Set("DataDisk.4.Size", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetDataDisk_4_Size() int {
	return r.DataDisk_4_Size
}
func (r *CreateInstanceRequest) SetDataDisk_4_Category(value string) {
	r.DataDisk_4_Category = value
	r.QueryParams.Set("DataDisk.4.Category", value)
}
func (r *CreateInstanceRequest) GetDataDisk_4_Category() string {
	return r.DataDisk_4_Category
}
func (r *CreateInstanceRequest) SetDataDisk_4_SnapshotId(value string) {
	r.DataDisk_4_SnapshotId = value
	r.QueryParams.Set("DataDisk.4.SnapshotId", value)
}
func (r *CreateInstanceRequest) GetDataDisk_4_SnapshotId() string {
	return r.DataDisk_4_SnapshotId
}
func (r *CreateInstanceRequest) SetDataDisk_4_DiskName(value string) {
	r.DataDisk_4_DiskName = value
	r.QueryParams.Set("DataDisk.4.DiskName", value)
}
func (r *CreateInstanceRequest) GetDataDisk_4_DiskName() string {
	return r.DataDisk_4_DiskName
}
func (r *CreateInstanceRequest) SetDataDisk_4_Description(value string) {
	r.DataDisk_4_Description = value
	r.QueryParams.Set("DataDisk.4.Description", value)
}
func (r *CreateInstanceRequest) GetDataDisk_4_Description() string {
	return r.DataDisk_4_Description
}
func (r *CreateInstanceRequest) SetDataDisk_4_Device(value string) {
	r.DataDisk_4_Device = value
	r.QueryParams.Set("DataDisk.4.Device", value)
}
func (r *CreateInstanceRequest) GetDataDisk_4_Device() string {
	return r.DataDisk_4_Device
}
func (r *CreateInstanceRequest) SetDataDisk_4_DeleteWithInstance(value bool) {
	r.DataDisk_4_DeleteWithInstance = value
	r.QueryParams.Set("DataDisk.4.DeleteWithInstance", strconv.FormatBool(value))
}
func (r *CreateInstanceRequest) GetDataDisk_4_DeleteWithInstance() bool {
	return r.DataDisk_4_DeleteWithInstance
}
func (r *CreateInstanceRequest) SetNodeControllerId(value string) {
	r.NodeControllerId = value
	r.QueryParams.Set("NodeControllerId", value)
}
func (r *CreateInstanceRequest) GetNodeControllerId() string {
	return r.NodeControllerId
}
func (r *CreateInstanceRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreateInstanceRequest) GetDescription() string {
	return r.Description
}
func (r *CreateInstanceRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *CreateInstanceRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *CreateInstanceRequest) SetPrivateIpAddress(value string) {
	r.PrivateIpAddress = value
	r.QueryParams.Set("PrivateIpAddress", value)
}
func (r *CreateInstanceRequest) GetPrivateIpAddress() string {
	return r.PrivateIpAddress
}
func (r *CreateInstanceRequest) SetIoOptimized(value string) {
	r.IoOptimized = value
	r.QueryParams.Set("IoOptimized", value)
}
func (r *CreateInstanceRequest) GetIoOptimized() string {
	return r.IoOptimized
}
func (r *CreateInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *CreateInstanceRequest) SetUseAdditionalService(value bool) {
	r.UseAdditionalService = value
	r.QueryParams.Set("UseAdditionalService", strconv.FormatBool(value))
}
func (r *CreateInstanceRequest) GetUseAdditionalService() bool {
	return r.UseAdditionalService
}
func (r *CreateInstanceRequest) SetInstanceChargeType(value string) {
	r.InstanceChargeType = value
	r.QueryParams.Set("InstanceChargeType", value)
}
func (r *CreateInstanceRequest) GetInstanceChargeType() string {
	return r.InstanceChargeType
}
func (r *CreateInstanceRequest) SetPeriod(value int) {
	r.Period = value
	r.QueryParams.Set("Period", strconv.Itoa(value))
}
func (r *CreateInstanceRequest) GetPeriod() int {
	return r.Period
}

func (r *CreateInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateInstance")
	r.SetProduct(Product)
}

type CreateInstanceResponse struct {
	InstanceId string `xml:"InstanceId" json:"InstanceId"`
}

func CreateInstance(req *CreateInstanceRequest, accessId, accessSecret string) (*CreateInstanceResponse, error) {
	var pResponse CreateInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
