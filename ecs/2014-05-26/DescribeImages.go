package ecs

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeImagesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	Status               string
	ImageId              string
	ShowExpired          bool
	SnapshotId           string
	ImageName            string
	ImageOwnerAlias      string
	PageNumber           int
	PageSize             int
	OwnerAccount         string
	Filter_1_Key         string
	Filter_2_Key         string
	Filter_1_Value       string
	Filter_2_Value       string
	Usage                string
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

func (r *DescribeImagesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeImagesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeImagesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeImagesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeImagesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeImagesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeImagesRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeImagesRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeImagesRequest) SetImageId(value string) {
	r.ImageId = value
	r.QueryParams.Set("ImageId", value)
}
func (r *DescribeImagesRequest) GetImageId() string {
	return r.ImageId
}
func (r *DescribeImagesRequest) SetShowExpired(value bool) {
	r.ShowExpired = value
	r.QueryParams.Set("ShowExpired", strconv.FormatBool(value))
}
func (r *DescribeImagesRequest) GetShowExpired() bool {
	return r.ShowExpired
}
func (r *DescribeImagesRequest) SetSnapshotId(value string) {
	r.SnapshotId = value
	r.QueryParams.Set("SnapshotId", value)
}
func (r *DescribeImagesRequest) GetSnapshotId() string {
	return r.SnapshotId
}
func (r *DescribeImagesRequest) SetImageName(value string) {
	r.ImageName = value
	r.QueryParams.Set("ImageName", value)
}
func (r *DescribeImagesRequest) GetImageName() string {
	return r.ImageName
}
func (r *DescribeImagesRequest) SetImageOwnerAlias(value string) {
	r.ImageOwnerAlias = value
	r.QueryParams.Set("ImageOwnerAlias", value)
}
func (r *DescribeImagesRequest) GetImageOwnerAlias() string {
	return r.ImageOwnerAlias
}
func (r *DescribeImagesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeImagesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeImagesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeImagesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeImagesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeImagesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeImagesRequest) SetFilter_1_Key(value string) {
	r.Filter_1_Key = value
	r.QueryParams.Set("Filter.1.Key", value)
}
func (r *DescribeImagesRequest) GetFilter_1_Key() string {
	return r.Filter_1_Key
}
func (r *DescribeImagesRequest) SetFilter_2_Key(value string) {
	r.Filter_2_Key = value
	r.QueryParams.Set("Filter.2.Key", value)
}
func (r *DescribeImagesRequest) GetFilter_2_Key() string {
	return r.Filter_2_Key
}
func (r *DescribeImagesRequest) SetFilter_1_Value(value string) {
	r.Filter_1_Value = value
	r.QueryParams.Set("Filter.1.Value", value)
}
func (r *DescribeImagesRequest) GetFilter_1_Value() string {
	return r.Filter_1_Value
}
func (r *DescribeImagesRequest) SetFilter_2_Value(value string) {
	r.Filter_2_Value = value
	r.QueryParams.Set("Filter.2.Value", value)
}
func (r *DescribeImagesRequest) GetFilter_2_Value() string {
	return r.Filter_2_Value
}
func (r *DescribeImagesRequest) SetUsage(value string) {
	r.Usage = value
	r.QueryParams.Set("Usage", value)
}
func (r *DescribeImagesRequest) GetUsage() string {
	return r.Usage
}
func (r *DescribeImagesRequest) SetTag_1_Key(value string) {
	r.Tag_1_Key = value
	r.QueryParams.Set("Tag.1.Key", value)
}
func (r *DescribeImagesRequest) GetTag_1_Key() string {
	return r.Tag_1_Key
}
func (r *DescribeImagesRequest) SetTag_2_Key(value string) {
	r.Tag_2_Key = value
	r.QueryParams.Set("Tag.2.Key", value)
}
func (r *DescribeImagesRequest) GetTag_2_Key() string {
	return r.Tag_2_Key
}
func (r *DescribeImagesRequest) SetTag_3_Key(value string) {
	r.Tag_3_Key = value
	r.QueryParams.Set("Tag.3.Key", value)
}
func (r *DescribeImagesRequest) GetTag_3_Key() string {
	return r.Tag_3_Key
}
func (r *DescribeImagesRequest) SetTag_4_Key(value string) {
	r.Tag_4_Key = value
	r.QueryParams.Set("Tag.4.Key", value)
}
func (r *DescribeImagesRequest) GetTag_4_Key() string {
	return r.Tag_4_Key
}
func (r *DescribeImagesRequest) SetTag_5_Key(value string) {
	r.Tag_5_Key = value
	r.QueryParams.Set("Tag.5.Key", value)
}
func (r *DescribeImagesRequest) GetTag_5_Key() string {
	return r.Tag_5_Key
}
func (r *DescribeImagesRequest) SetTag_1_Value(value string) {
	r.Tag_1_Value = value
	r.QueryParams.Set("Tag.1.Value", value)
}
func (r *DescribeImagesRequest) GetTag_1_Value() string {
	return r.Tag_1_Value
}
func (r *DescribeImagesRequest) SetTag_2_Value(value string) {
	r.Tag_2_Value = value
	r.QueryParams.Set("Tag.2.Value", value)
}
func (r *DescribeImagesRequest) GetTag_2_Value() string {
	return r.Tag_2_Value
}
func (r *DescribeImagesRequest) SetTag_3_Value(value string) {
	r.Tag_3_Value = value
	r.QueryParams.Set("Tag.3.Value", value)
}
func (r *DescribeImagesRequest) GetTag_3_Value() string {
	return r.Tag_3_Value
}
func (r *DescribeImagesRequest) SetTag_4_Value(value string) {
	r.Tag_4_Value = value
	r.QueryParams.Set("Tag.4.Value", value)
}
func (r *DescribeImagesRequest) GetTag_4_Value() string {
	return r.Tag_4_Value
}
func (r *DescribeImagesRequest) SetTag_5_Value(value string) {
	r.Tag_5_Value = value
	r.QueryParams.Set("Tag.5.Value", value)
}
func (r *DescribeImagesRequest) GetTag_5_Value() string {
	return r.Tag_5_Value
}

func (r *DescribeImagesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeImages")
	r.SetProduct(Product)
}

type DescribeImagesResponse struct {
	RegionId   string `xml:"RegionId" json:"RegionId"`
	TotalCount int    `xml:"TotalCount" json:"TotalCount"`
	PageNumber int    `xml:"PageNumber" json:"PageNumber"`
	PageSize   int    `xml:"PageSize" json:"PageSize"`
	Images     struct {
		Image []struct {
			Progress           string `xml:"Progress" json:"Progress"`
			ImageId            string `xml:"ImageId" json:"ImageId"`
			ImageName          string `xml:"ImageName" json:"ImageName"`
			ImageVersion       string `xml:"ImageVersion" json:"ImageVersion"`
			Description        string `xml:"Description" json:"Description"`
			Size               int    `xml:"Size" json:"Size"`
			ImageOwnerAlias    string `xml:"ImageOwnerAlias" json:"ImageOwnerAlias"`
			OSName             string `xml:"OSName" json:"OSName"`
			Architecture       string `xml:"Architecture" json:"Architecture"`
			Status             string `xml:"Status" json:"Status"`
			ProductCode        string `xml:"ProductCode" json:"ProductCode"`
			IsSubscribed       bool   `xml:"IsSubscribed" json:"IsSubscribed"`
			CreationTime       string `xml:"CreationTime" json:"CreationTime"`
			IsSelfShared       string `xml:"IsSelfShared" json:"IsSelfShared"`
			OSType             string `xml:"OSType" json:"OSType"`
			Platform           string `xml:"Platform" json:"Platform"`
			Usage              string `xml:"Usage" json:"Usage"`
			IsCopied           bool   `xml:"IsCopied" json:"IsCopied"`
			DiskDeviceMappings struct {
				DiskDeviceMapping []struct {
					SnapshotId string `xml:"SnapshotId" json:"SnapshotId"`
					Size       string `xml:"Size" json:"Size"`
					Device     string `xml:"Device" json:"Device"`
				} `xml:"DiskDeviceMapping" json:"DiskDeviceMapping"`
			} `xml:"DiskDeviceMappings" json:"DiskDeviceMappings"`
			Tags struct {
				Tag []struct {
					TagKey   string `xml:"TagKey" json:"TagKey"`
					TagValue string `xml:"TagValue" json:"TagValue"`
				} `xml:"Tag" json:"Tag"`
			} `xml:"Tags" json:"Tags"`
		} `xml:"Image" json:"Image"`
	} `xml:"Images" json:"Images"`
}

func DescribeImages(req *DescribeImagesRequest, accessId, accessSecret string) (*DescribeImagesResponse, error) {
	var pResponse DescribeImagesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
