package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeScalingConfigurationsRequest struct {
	RpcRequest
	OwnerId                     int
	ResourceOwnerAccount        string
	ResourceOwnerId             int
	PageNumber                  int
	PageSize                    int
	ScalingGroupId              string
	ScalingConfigurationId_1    string
	ScalingConfigurationId_2    string
	ScalingConfigurationId_3    string
	ScalingConfigurationId_4    string
	ScalingConfigurationId_5    string
	ScalingConfigurationId_6    string
	ScalingConfigurationId_7    string
	ScalingConfigurationId_8    string
	ScalingConfigurationId_9    string
	ScalingConfigurationId_10   string
	ScalingConfigurationName_1  string
	ScalingConfigurationName_2  string
	ScalingConfigurationName_3  string
	ScalingConfigurationName_4  string
	ScalingConfigurationName_5  string
	ScalingConfigurationName_6  string
	ScalingConfigurationName_7  string
	ScalingConfigurationName_8  string
	ScalingConfigurationName_9  string
	ScalingConfigurationName_10 string
	OwnerAccount                string
}

func (r *DescribeScalingConfigurationsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingConfigurationsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeScalingConfigurationsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeScalingConfigurationsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeScalingConfigurationsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingConfigurationsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeScalingConfigurationsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeScalingConfigurationsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeScalingConfigurationsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeScalingConfigurationsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeScalingConfigurationsRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_1(value string) {
	r.ScalingConfigurationId_1 = value
	r.QueryParams.Set("ScalingConfigurationId.1", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_1() string {
	return r.ScalingConfigurationId_1
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_2(value string) {
	r.ScalingConfigurationId_2 = value
	r.QueryParams.Set("ScalingConfigurationId.2", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_2() string {
	return r.ScalingConfigurationId_2
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_3(value string) {
	r.ScalingConfigurationId_3 = value
	r.QueryParams.Set("ScalingConfigurationId.3", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_3() string {
	return r.ScalingConfigurationId_3
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_4(value string) {
	r.ScalingConfigurationId_4 = value
	r.QueryParams.Set("ScalingConfigurationId.4", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_4() string {
	return r.ScalingConfigurationId_4
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_5(value string) {
	r.ScalingConfigurationId_5 = value
	r.QueryParams.Set("ScalingConfigurationId.5", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_5() string {
	return r.ScalingConfigurationId_5
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_6(value string) {
	r.ScalingConfigurationId_6 = value
	r.QueryParams.Set("ScalingConfigurationId.6", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_6() string {
	return r.ScalingConfigurationId_6
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_7(value string) {
	r.ScalingConfigurationId_7 = value
	r.QueryParams.Set("ScalingConfigurationId.7", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_7() string {
	return r.ScalingConfigurationId_7
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_8(value string) {
	r.ScalingConfigurationId_8 = value
	r.QueryParams.Set("ScalingConfigurationId.8", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_8() string {
	return r.ScalingConfigurationId_8
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_9(value string) {
	r.ScalingConfigurationId_9 = value
	r.QueryParams.Set("ScalingConfigurationId.9", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_9() string {
	return r.ScalingConfigurationId_9
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationId_10(value string) {
	r.ScalingConfigurationId_10 = value
	r.QueryParams.Set("ScalingConfigurationId.10", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationId_10() string {
	return r.ScalingConfigurationId_10
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_1(value string) {
	r.ScalingConfigurationName_1 = value
	r.QueryParams.Set("ScalingConfigurationName.1", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_1() string {
	return r.ScalingConfigurationName_1
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_2(value string) {
	r.ScalingConfigurationName_2 = value
	r.QueryParams.Set("ScalingConfigurationName.2", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_2() string {
	return r.ScalingConfigurationName_2
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_3(value string) {
	r.ScalingConfigurationName_3 = value
	r.QueryParams.Set("ScalingConfigurationName.3", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_3() string {
	return r.ScalingConfigurationName_3
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_4(value string) {
	r.ScalingConfigurationName_4 = value
	r.QueryParams.Set("ScalingConfigurationName.4", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_4() string {
	return r.ScalingConfigurationName_4
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_5(value string) {
	r.ScalingConfigurationName_5 = value
	r.QueryParams.Set("ScalingConfigurationName.5", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_5() string {
	return r.ScalingConfigurationName_5
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_6(value string) {
	r.ScalingConfigurationName_6 = value
	r.QueryParams.Set("ScalingConfigurationName.6", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_6() string {
	return r.ScalingConfigurationName_6
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_7(value string) {
	r.ScalingConfigurationName_7 = value
	r.QueryParams.Set("ScalingConfigurationName.7", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_7() string {
	return r.ScalingConfigurationName_7
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_8(value string) {
	r.ScalingConfigurationName_8 = value
	r.QueryParams.Set("ScalingConfigurationName.8", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_8() string {
	return r.ScalingConfigurationName_8
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_9(value string) {
	r.ScalingConfigurationName_9 = value
	r.QueryParams.Set("ScalingConfigurationName.9", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_9() string {
	return r.ScalingConfigurationName_9
}
func (r *DescribeScalingConfigurationsRequest) SetScalingConfigurationName_10(value string) {
	r.ScalingConfigurationName_10 = value
	r.QueryParams.Set("ScalingConfigurationName.10", value)
}
func (r *DescribeScalingConfigurationsRequest) GetScalingConfigurationName_10() string {
	return r.ScalingConfigurationName_10
}
func (r *DescribeScalingConfigurationsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeScalingConfigurationsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeScalingConfigurationsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeScalingConfigurations")
	r.SetProduct(Product)
}

type DescribeScalingConfigurationsResponse struct {
	TotalCount            int `xml:"TotalCount" json:"TotalCount"`
	PageNumber            int `xml:"PageNumber" json:"PageNumber"`
	PageSize              int `xml:"PageSize" json:"PageSize"`
	ScalingConfigurations struct {
		ScalingConfiguration []struct {
			ScalingConfigurationId   string `xml:"ScalingConfigurationId" json:"ScalingConfigurationId"`
			ScalingConfigurationName string `xml:"ScalingConfigurationName" json:"ScalingConfigurationName"`
			ScalingGroupId           string `xml:"ScalingGroupId" json:"ScalingGroupId"`
			ImageId                  string `xml:"ImageId" json:"ImageId"`
			InstanceType             string `xml:"InstanceType" json:"InstanceType"`
			SecurityGroupId          string `xml:"SecurityGroupId" json:"SecurityGroupId"`
			InternetChargeType       string `xml:"InternetChargeType" json:"InternetChargeType"`
			InternetMaxBandwidthIn   int    `xml:"InternetMaxBandwidthIn" json:"InternetMaxBandwidthIn"`
			InternetMaxBandwidthOut  int    `xml:"InternetMaxBandwidthOut" json:"InternetMaxBandwidthOut"`
			SystemDiskCategory       string `xml:"SystemDiskCategory" json:"SystemDiskCategory"`
			LifecycleState           string `xml:"LifecycleState" json:"LifecycleState"`
			CreationTime             string `xml:"CreationTime" json:"CreationTime"`
			DataDisks                struct {
				DataDisk []struct {
					Size       int    `xml:"Size" json:"Size"`
					Category   string `xml:"Category" json:"Category"`
					SnapshotId string `xml:"SnapshotId" json:"SnapshotId"`
					Device     string `xml:"Device" json:"Device"`
				} `xml:"DataDisk" json:"DataDisk"`
			} `xml:"DataDisks" json:"DataDisks"`
		} `xml:"ScalingConfiguration" json:"ScalingConfiguration"`
	} `xml:"ScalingConfigurations" json:"ScalingConfigurations"`
}

func DescribeScalingConfigurations(req *DescribeScalingConfigurationsRequest, accessId, accessSecret string) (*DescribeScalingConfigurationsResponse, error) {
	var pResponse DescribeScalingConfigurationsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
