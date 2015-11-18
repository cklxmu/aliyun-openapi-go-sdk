package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeDisksRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ZoneId               string
	DiskIds              string
	InstanceId           string
	DiskType             string
	Category             string
	Status               string
	SnapshotId           string
	Portable             bool
	DeleteWithInstance   bool
	DeleteAutoSnapshot   bool
	PageNumber           int
	PageSize             int
	OwnerAccount         string
	DiskName             string
	EnableAutoSnapshot   bool
	DiskChargeType       string
	LockReason           string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
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

func (r *DescribeDisksRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeDisksRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeDisksRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeDisksRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeDisksRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeDisksRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeDisksRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *DescribeDisksRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *DescribeDisksRequest) SetDiskIds(value string) {
	r.DiskIds = value
	r.QueryParams.Set("DiskIds", value)
}
func (r *DescribeDisksRequest) GetDiskIds() string {
	return r.DiskIds
}
func (r *DescribeDisksRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeDisksRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeDisksRequest) SetDiskType(value string) {
	r.DiskType = value
	r.QueryParams.Set("DiskType", value)
}
func (r *DescribeDisksRequest) GetDiskType() string {
	return r.DiskType
}
func (r *DescribeDisksRequest) SetCategory(value string) {
	r.Category = value
	r.QueryParams.Set("Category", value)
}
func (r *DescribeDisksRequest) GetCategory() string {
	return r.Category
}
func (r *DescribeDisksRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeDisksRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeDisksRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *DescribeDisksRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *DescribeDisksRequest) SetPortable(value bool) {
	r.Portable = value
	r.QueryParams.Set("Portable", strconv.FormatBool(value))
}
func (r *DescribeDisksRequest) GetPortable() bool {
	return r.Portable
}
func (r *DescribeDisksRequest) SetDeleteWithInstance(value bool) {
	r.DeleteWithInstance = value
	r.QueryParams.Set("DeleteWithInstance", strconv.FormatBool(value))
}
func (r *DescribeDisksRequest) GetDeleteWithInstance() bool {
	return r.DeleteWithInstance
}
func (r *DescribeDisksRequest) SetDeleteAutoSnapshot(value bool) {
	r.DeleteAutoSnapshot = value
	r.QueryParams.Set("DeleteAutoSnapshot", strconv.FormatBool(value))
}
func (r *DescribeDisksRequest) GetDeleteAutoSnapshot() bool {
	return r.DeleteAutoSnapshot
}
func (r *DescribeDisksRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeDisksRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeDisksRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeDisksRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeDisksRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeDisksRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeDisksRequest) SetDiskName(value string) {
	r.DiskName = value
	r.QueryParams.Set("DiskName", value)
}
func (r *DescribeDisksRequest) GetDiskName() string {
	return r.DiskName
}
func (r *DescribeDisksRequest) SetEnableAutoSnapshot(value bool) {
	r.EnableAutoSnapshot = value
	r.QueryParams.Set("EnableAutoSnapshot", strconv.FormatBool(value))
}
func (r *DescribeDisksRequest) GetEnableAutoSnapshot() bool {
	return r.EnableAutoSnapshot
}
func (r *DescribeDisksRequest) SetDiskChargeType(value string) {
	r.DiskChargeType = value
	r.QueryParams.Set("DiskChargeType", value)
}
func (r *DescribeDisksRequest) GetDiskChargeType() string {
	return r.DiskChargeType
}
func (r *DescribeDisksRequest) SetLockReason(value string) {
	r.LockReason = value
	r.QueryParams.Set("LockReason", value)
}
func (r *DescribeDisksRequest) GetLockReason() string {
	return r.LockReason
}
func (r *DescribeDisksRequest) SetFilter_1_Key(value string) {
	r.Filter_1_Key = value
	r.QueryParams.Set("Filter.1.Key", value)
}
func (r *DescribeDisksRequest) GetFilter_1_Key() string {
	return r.Filter_1_Key
}
func (r *DescribeDisksRequest) SetFilter_2_Key(value string) {
	r.Filter_2_Key = value
	r.QueryParams.Set("Filter.2.Key", value)
}
func (r *DescribeDisksRequest) GetFilter_2_Key() string {
	return r.Filter_2_Key
}
func (r *DescribeDisksRequest) SetFilter_1_Value(value string) {
	r.Filter_1_Value = value
	r.QueryParams.Set("Filter.1.Value", value)
}
func (r *DescribeDisksRequest) GetFilter_1_Value() string {
	return r.Filter_1_Value
}
func (r *DescribeDisksRequest) SetFilter_2_Value(value string) {
	r.Filter_2_Value = value
	r.QueryParams.Set("Filter.2.Value", value)
}
func (r *DescribeDisksRequest) GetFilter_2_Value() string {
	return r.Filter_2_Value
}
func (r *DescribeDisksRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *DescribeDisksRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *DescribeDisksRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *DescribeDisksRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *DescribeDisksRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *DescribeDisksRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *DescribeDisksRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *DescribeDisksRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *DescribeDisksRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *DescribeDisksRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *DescribeDisksRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *DescribeDisksRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *DescribeDisksRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *DescribeDisksRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *DescribeDisksRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *DescribeDisksRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *DescribeDisksRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *DescribeDisksRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *DescribeDisksRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *DescribeDisksRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *DescribeDisksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDisks")
	r.SetProduct(Product)
}

type DescribeDisksResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Disks      struct {
		Disk []struct {
			DiskId             string `xml:"DiskId" json:"DiskId"`
			RegionId           string `xml:"RegionId" json:"RegionId"`
			ZoneId             string `xml:"ZoneId" json:"ZoneId"`
			DiskName           string `xml:"DiskName" json:"DiskName"`
			Description        string `xml:"Description" json:"Description"`
			Type               string `xml:"Type" json:"Type"`
			Category           string `xml:"Category" json:"Category"`
			Size               int    `xml:"Size" json:"Size"`
			ImageId            string `xml:"ImageId" json:"ImageId"`
			SourceSnapshotId   string `xml:"SourceSnapshotId" json:"SourceSnapshotId"`
			ProductCode        string `xml:"ProductCode" json:"ProductCode"`
			Portable           bool   `xml:"Portable" json:"Portable"`
			Status             string `xml:"Status" json:"Status"`
			InstanceId         string `xml:"InstanceId" json:"InstanceId"`
			Device             string `xml:"Device" json:"Device"`
			DeleteWithInstance bool   `xml:"DeleteWithInstance" json:"DeleteWithInstance"`
			DeleteAutoSnapshot bool   `xml:"DeleteAutoSnapshot" json:"DeleteAutoSnapshot"`
			EnableAutoSnapshot bool   `xml:"EnableAutoSnapshot" json:"EnableAutoSnapshot"`
			CreationTime       string `xml:"CreationTime" json:"CreationTime"`
			AttachedTime       string `xml:"AttachedTime" json:"AttachedTime"`
			DetachedTime       string `xml:"DetachedTime" json:"DetachedTime"`
			DiskChargeType     string `xml:"DiskChargeType" json:"DiskChargeType"`
			ExpiredTime        string `xml:"ExpiredTime" json:"ExpiredTime"`
			OperationLocks     struct {
				OperationLock []struct {
					LockReason string `xml:"LockReason" json:"LockReason"`
				} `xml:"OperationLock" json:"OperationLock"`
			} `xml:"OperationLocks" json:"OperationLocks"`
			Tags struct {
				Tag []struct {
					TagKey   string `xml:"TagKey" json:"TagKey"`
					TagValue string `xml:"TagValue" json:"TagValue"`
				} `xml:"Tag" json:"Tag"`
			} `xml:"Tags" json:"Tags"`
		} `xml:"Disk" json:"Disk"`
	} `xml:"Disks" json:"Disks"`
}

func DescribeDisks(req *DescribeDisksRequest, accessId, accessSecret string) (*DescribeDisksResponse, error) {
	var pResponse DescribeDisksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
