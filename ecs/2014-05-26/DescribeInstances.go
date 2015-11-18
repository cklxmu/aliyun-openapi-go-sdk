package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeInstancesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	VpcId                string
	VSwitchId            string
	ZoneId               string
	InstanceNetworkType  string
	SecurityGroupId      string
	InstanceIds          string
	PageNumber           int
	PageSize             int
	InnerIpAddresses     string
	PrivateIpAddresses   string
	PublicIpAddresses    string
	OwnerAccount         string
	InstanceChargeType   string
	InternetChargeType   string
	InstanceName         string
	ImageId              string
	Status               string
	LockReason           string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_3_Key         string
	Filter_4_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
	Filter_3_Value       string
	Filter_4_Value       string
	DeviceAvailable      bool
	IoOptimized          bool
	Tag_1_Key            string
	Tag_2_Key            string
	Tag_3_Key            string
	Tag_4_Key            string
	Tag_5_Key            string
	Tag_1_Value          string
	Tag_2_Value          string
	Tag_3_Value          string
	Tag_4_Value          string
	Tag_5_Value          string
}

func (r *DescribeInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeInstancesRequest) SetVpcId(value string) {
	r.VpcId = value
	r.QueryParams.Set("VpcId", value)
}
func (r *DescribeInstancesRequest) GetVpcId() string {
	return r.VpcId
}
func (r *DescribeInstancesRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *DescribeInstancesRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *DescribeInstancesRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribeInstancesRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *DescribeInstancesRequest) SetInstanceNetworkType(value string) {
	r.InstanceNetworkType = value
	r.QueryParams.Set("InstanceNetworkType", value)
}
func (r *DescribeInstancesRequest) GetInstanceNetworkType() string {
	return r.InstanceNetworkType
}
func (r *DescribeInstancesRequest) SetSecurityGroupId(value string) {
	r.SecurityGroupId = value
	r.QueryParams.Set("SecurityGroupId", value)
}
func (r *DescribeInstancesRequest) GetSecurityGroupId() string {
	return r.SecurityGroupId
}
func (r *DescribeInstancesRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *DescribeInstancesRequest) GetInstanceIds() string {
	return r.InstanceIds
}
func (r *DescribeInstancesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeInstancesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeInstancesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeInstancesRequest) SetInnerIpAddresses(value string) {
	r.InnerIpAddresses = value
	r.QueryParams.Set("InnerIpAddresses", value)
}
func (r *DescribeInstancesRequest) GetInnerIpAddresses() string {
	return r.InnerIpAddresses
}
func (r *DescribeInstancesRequest) SetPrivateIpAddresses(value string) {
	r.PrivateIpAddresses = value
	r.QueryParams.Set("PrivateIpAddresses", value)
}
func (r *DescribeInstancesRequest) GetPrivateIpAddresses() string {
	return r.PrivateIpAddresses
}
func (r *DescribeInstancesRequest) SetPublicIpAddresses(value string) {
	r.PublicIpAddresses = value
	r.QueryParams.Set("PublicIpAddresses", value)
}
func (r *DescribeInstancesRequest) GetPublicIpAddresses() string {
	return r.PublicIpAddresses
}
func (r *DescribeInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeInstancesRequest) SetInstanceChargeType(value string) {
	r.InstanceChargeType = value
	r.QueryParams.Set("InstanceChargeType", value)
}
func (r *DescribeInstancesRequest) GetInstanceChargeType() string {
	return r.InstanceChargeType
}
func (r *DescribeInstancesRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *DescribeInstancesRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}
func (r *DescribeInstancesRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *DescribeInstancesRequest) GetInstanceName() string {
	return r.InstanceName
}
func (r *DescribeInstancesRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *DescribeInstancesRequest) GetImageId() string {
	return r.ImageId
}
func (r *DescribeInstancesRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeInstancesRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeInstancesRequest) SetLockReason(value string) {
	r.LockReason = value
	r.QueryParams.Set("LockReason", value)
}
func (r *DescribeInstancesRequest) GetLockReason() string {
	return r.LockReason
}
func (r *DescribeInstancesRequest) SetFilter_1_Key(value string) {
	r.Filter_1_Key = value
	r.QueryParams.Set("Filter.1.Key", value)
}
func (r *DescribeInstancesRequest) GetFilter_1_Key() string {
	return r.Filter_1_Key
}
func (r *DescribeInstancesRequest) SetFilter_2_Key(value string) {
	r.Filter_2_Key = value
	r.QueryParams.Set("Filter.2.Key", value)
}
func (r *DescribeInstancesRequest) GetFilter_2_Key() string {
	return r.Filter_2_Key
}
func (r *DescribeInstancesRequest) SetFilter_3_Key(value string) {
	r.Filter_3_Key = value
	r.QueryParams.Set("Filter.3.Key", value)
}
func (r *DescribeInstancesRequest) GetFilter_3_Key() string {
	return r.Filter_3_Key
}
func (r *DescribeInstancesRequest) SetFilter_4_Key(value string) {
	r.Filter_4_Key = value
	r.QueryParams.Set("Filter.4.Key", value)
}
func (r *DescribeInstancesRequest) GetFilter_4_Key() string {
	return r.Filter_4_Key
}
func (r *DescribeInstancesRequest) SetFilter_1_Value(value string) {
	r.Filter_1_Value = value
	r.QueryParams.Set("Filter.1.Value", value)
}
func (r *DescribeInstancesRequest) GetFilter_1_Value() string {
	return r.Filter_1_Value
}
func (r *DescribeInstancesRequest) SetFilter_2_Value(value string) {
	r.Filter_2_Value = value
	r.QueryParams.Set("Filter.2.Value", value)
}
func (r *DescribeInstancesRequest) GetFilter_2_Value() string {
	return r.Filter_2_Value
}
func (r *DescribeInstancesRequest) SetFilter_3_Value(value string) {
	r.Filter_3_Value = value
	r.QueryParams.Set("Filter.3.Value", value)
}
func (r *DescribeInstancesRequest) GetFilter_3_Value() string {
	return r.Filter_3_Value
}
func (r *DescribeInstancesRequest) SetFilter_4_Value(value string) {
	r.Filter_4_Value = value
	r.QueryParams.Set("Filter.4.Value", value)
}
func (r *DescribeInstancesRequest) GetFilter_4_Value() string {
	return r.Filter_4_Value
}
func (r *DescribeInstancesRequest) SetDeviceAvailable(value bool) {
	r.DeviceAvailable = value
	r.QueryParams.Set("DeviceAvailable", strconv.FormatBool(value))
}
func (r *DescribeInstancesRequest) GetDeviceAvailable() bool {
	return r.DeviceAvailable
}
func (r *DescribeInstancesRequest) SetIoOptimized(value bool) {
	r.IoOptimized = value
	r.QueryParams.Set("IoOptimized", strconv.FormatBool(value))
}
func (r *DescribeInstancesRequest) GetIoOptimized() bool {
	return r.IoOptimized
}
func (r *DescribeInstancesRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *DescribeInstancesRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *DescribeInstancesRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *DescribeInstancesRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *DescribeInstancesRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *DescribeInstancesRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *DescribeInstancesRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *DescribeInstancesRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *DescribeInstancesRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *DescribeInstancesRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *DescribeInstancesRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *DescribeInstancesRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *DescribeInstancesRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *DescribeInstancesRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *DescribeInstancesRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *DescribeInstancesRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *DescribeInstancesRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *DescribeInstancesRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *DescribeInstancesRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *DescribeInstancesRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *DescribeInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeInstances")
	r.SetProduct(Product)
}

type DescribeInstancesResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Instances  struct {
		Instance []struct {
			InstanceId              string `xml:"InstanceId" json:"InstanceId"`
			InstanceName            string `xml:"InstanceName" json:"InstanceName"`
			Description             string `xml:"Description" json:"Description"`
			ImageId                 string `xml:"ImageId" json:"ImageId"`
			RegionId                string `xml:"RegionId" json:"RegionId"`
			ZoneId                  string `xml:"ZoneId" json:"ZoneId"`
			ClusterId               string `xml:"ClusterId" json:"ClusterId"`
			InstanceType            string `xml:"InstanceType" json:"InstanceType"`
			HostName                string `xml:"HostName" json:"HostName"`
			Status                  string `xml:"Status" json:"Status"`
			SerialNumber            string `xml:"SerialNumber" json:"SerialNumber"`
			InternetChargeType      string `xml:"InternetChargeType" json:"InternetChargeType"`
			InternetMaxBandwidthIn  int    `xml:"InternetMaxBandwidthIn" json:"InternetMaxBandwidthIn"`
			InternetMaxBandwidthOut int    `xml:"InternetMaxBandwidthOut" json:"InternetMaxBandwidthOut"`
			VlanId                  string `xml:"VlanId" json:"VlanId"`
			CreationTime            string `xml:"CreationTime" json:"CreationTime"`
			InstanceNetworkType     string `xml:"InstanceNetworkType" json:"InstanceNetworkType"`
			InstanceChargeType      string `xml:"InstanceChargeType" json:"InstanceChargeType"`
			ExpiredTime             string `xml:"ExpiredTime" json:"ExpiredTime"`
			IoOptimized             bool   `xml:"IoOptimized" json:"IoOptimized"`
			DeviceAvailable         bool   `xml:"DeviceAvailable" json:"DeviceAvailable"`
			SecurityGroupIds        struct {
				SecurityGroupId []string `xml:"SecurityGroupId" json:"SecurityGroupId"`
			} `xml:"SecurityGroupIds" json:"SecurityGroupIds"`
			PublicIpAddress struct {
				IpAddress []string `xml:"IpAddress" json:"IpAddress"`
			} `xml:"PublicIpAddress" json:"PublicIpAddress"`
			InnerIpAddress struct {
				IpAddress []string `xml:"IpAddress" json:"IpAddress"`
			} `xml:"InnerIpAddress" json:"InnerIpAddress"`
			OperationLocks struct {
				LockReason []struct {
					LockReason string `xml:"LockReason" json:"LockReason"`
				} `xml:"LockReason" json:"LockReason"`
			} `xml:"OperationLocks" json:"OperationLocks"`
			Tags struct {
				Tag []struct {
					TagKey   string `xml:"TagKey" json:"TagKey"`
					TagValue string `xml:"TagValue" json:"TagValue"`
				} `xml:"Tag" json:"Tag"`
			} `xml:"Tags" json:"Tags"`
			VpcAttributes struct {
				VpcId            string `xml:"VpcId" json:"VpcId"`
				VSwitchId        string `xml:"VSwitchId" json:"VSwitchId"`
				NatIpAddress     string `xml:"NatIpAddress" json:"NatIpAddress"`
				PrivateIpAddress struct {
					IpAddress []string `xml:"IpAddress" json:"IpAddress"`
				} `xml:"PrivateIpAddress" json:"PrivateIpAddress"`
			} `xml:"VpcAttributes" json:"VpcAttributes"`
			EipAddress struct {
				AllocationId       string `xml:"AllocationId" json:"AllocationId"`
				IpAddress          string `xml:"IpAddress" json:"IpAddress"`
				Bandwidth          int    `xml:"Bandwidth" json:"Bandwidth"`
				InternetChargeType string `xml:"InternetChargeType" json:"InternetChargeType"`
			} `xml:"EipAddress" json:"EipAddress"`
		} `xml:"Instance" json:"Instance"`
	} `xml:"Instances" json:"Instances"`
}

func DescribeInstances(req *DescribeInstancesRequest, accessId, accessSecret string) (*DescribeInstancesResponse, error) {
	var pResponse DescribeInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
