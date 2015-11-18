package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSnapshotsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InstanceId           string
	DiskId               string
	SnapshotIds          string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
	SnapshotName         string
	Status               string
	SnapshotType         string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
	Usage                string
	SourceDiskType       string
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

func (r *DescribeSnapshotsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSnapshotsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSnapshotsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSnapshotsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSnapshotsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSnapshotsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSnapshotsRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DescribeSnapshotsRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DescribeSnapshotsRequest) SetDiskId(value string) {
	r.DiskId = value
	r.QueryParams.Set("DiskId", value)
}
func (r *DescribeSnapshotsRequest) GetDiskId() string {
	return r.DiskId
}
func (r *DescribeSnapshotsRequest) SetSnapshotIds(value string) {
	r.SnapshotIds = value
	r.QueryParams.Set("SnapshotIds", value)
}
func (r *DescribeSnapshotsRequest) GetSnapshotIds() string {
	return r.SnapshotIds
}
func (r *DescribeSnapshotsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSnapshotsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSnapshotsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSnapshotsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSnapshotsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSnapshotsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeSnapshotsRequest) SetSnapshotName(value string) {
	r.SnapshotName = value
	r.QueryParams.Set("SnapshotName", value)
}
func (r *DescribeSnapshotsRequest) GetSnapshotName() string {
	return r.SnapshotName
}
func (r *DescribeSnapshotsRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeSnapshotsRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeSnapshotsRequest) SetSnapshotType(value string) {
	r.SnapshotType = value
	r.QueryParams.Set("SnapshotType", value)
}
func (r *DescribeSnapshotsRequest) GetSnapshotType() string {
	return r.SnapshotType
}
func (r *DescribeSnapshotsRequest) SetFilter_1_Key(value string) {
	r.Filter_1_Key = value
	r.QueryParams.Set("Filter.1.Key", value)
}
func (r *DescribeSnapshotsRequest) GetFilter_1_Key() string {
	return r.Filter_1_Key
}
func (r *DescribeSnapshotsRequest) SetFilter_2_Key(value string) {
	r.Filter_2_Key = value
	r.QueryParams.Set("Filter.2.Key", value)
}
func (r *DescribeSnapshotsRequest) GetFilter_2_Key() string {
	return r.Filter_2_Key
}
func (r *DescribeSnapshotsRequest) SetFilter_1_Value(value string) {
	r.Filter_1_Value = value
	r.QueryParams.Set("Filter.1.Value", value)
}
func (r *DescribeSnapshotsRequest) GetFilter_1_Value() string {
	return r.Filter_1_Value
}
func (r *DescribeSnapshotsRequest) SetFilter_2_Value(value string) {
	r.Filter_2_Value = value
	r.QueryParams.Set("Filter.2.Value", value)
}
func (r *DescribeSnapshotsRequest) GetFilter_2_Value() string {
	return r.Filter_2_Value
}
func (r *DescribeSnapshotsRequest) SetUsage(value string) {
	r.Usage = value
	r.QueryParams.Set("Usage", value)
}
func (r *DescribeSnapshotsRequest) GetUsage() string {
	return r.Usage
}
func (r *DescribeSnapshotsRequest) SetSourceDiskType(value string) {
	r.SourceDiskType = value
	r.QueryParams.Set("SourceDiskType", value)
}
func (r *DescribeSnapshotsRequest) GetSourceDiskType() string {
	return r.SourceDiskType
}
func (r *DescribeSnapshotsRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *DescribeSnapshotsRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *DescribeSnapshotsRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *DescribeSnapshotsRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *DescribeSnapshotsRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *DescribeSnapshotsRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *DescribeSnapshotsRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *DescribeSnapshotsRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *DescribeSnapshotsRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *DescribeSnapshotsRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *DescribeSnapshotsRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *DescribeSnapshotsRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *DescribeSnapshotsRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *DescribeSnapshotsRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *DescribeSnapshotsRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *DescribeSnapshotsRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *DescribeSnapshotsRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *DescribeSnapshotsRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *DescribeSnapshotsRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *DescribeSnapshotsRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *DescribeSnapshotsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSnapshots")
	r.SetProduct(Product)
}

type DescribeSnapshotsResponse struct {
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	Snapshots  struct {
		Snapshot []struct {
			SnapshotId     string `xml:"SnapshotId" json:"SnapshotId"`
			SnapshotName   string `xml:"SnapshotName" json:"SnapshotName"`
			Progress       string `xml:"Progress" json:"Progress"`
			ProductCode    string `xml:"ProductCode" json:"ProductCode"`
			SourceDiskId   string `xml:"SourceDiskId" json:"SourceDiskId"`
			SourceDiskType string `xml:"SourceDiskType" json:"SourceDiskType"`
			SourceDiskSize string `xml:"SourceDiskSize" json:"SourceDiskSize"`
			Description    string `xml:"Description" json:"Description"`
			CreationTime   string `xml:"CreationTime" json:"CreationTime"`
			Status         string `xml:"Status" json:"Status"`
			Usage          string `xml:"Usage" json:"Usage"`
			Tags           struct {
				Tag []struct {
					TagKey   string `xml:"TagKey" json:"TagKey"`
					TagValue string `xml:"TagValue" json:"TagValue"`
				} `xml:"Tag" json:"Tag"`
			} `xml:"Tags" json:"Tags"`
		} `xml:"Snapshot" json:"Snapshot"`
	} `xml:"Snapshots" json:"Snapshots"`
}

func DescribeSnapshots(req *DescribeSnapshotsRequest, accessId, accessSecret string) (*DescribeSnapshotsResponse, error) {
	var pResponse DescribeSnapshotsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
