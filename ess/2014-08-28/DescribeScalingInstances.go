package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeScalingInstancesRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	ScalingGroupId         string
	ScalingConfigurationId string
	HealthStatus           string
	LifecycleState         string
	CreationType           string
	PageNumber             int
	PageSize               int
	InstanceId_1           string
	InstanceId_2           string
	InstanceId_3           string
	InstanceId_4           string
	InstanceId_5           string
	InstanceId_6           string
	InstanceId_7           string
	InstanceId_8           string
	InstanceId_9           string
	InstanceId_10          string
	InstanceId_11          string
	InstanceId_12          string
	InstanceId_13          string
	InstanceId_14          string
	InstanceId_15          string
	InstanceId_16          string
	InstanceId_17          string
	InstanceId_18          string
	InstanceId_19          string
	InstanceId_20          string
	OwnerAccount           string
}

func (r *DescribeScalingInstancesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingInstancesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeScalingInstancesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeScalingInstancesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeScalingInstancesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingInstancesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeScalingInstancesRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *DescribeScalingInstancesRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *DescribeScalingInstancesRequest) SetScalingConfigurationId(value string) {
	r.ScalingConfigurationId = value
	r.QueryParams.Set("ScalingConfigurationId", value)
}
func (r *DescribeScalingInstancesRequest) GetScalingConfigurationId() string {
	return r.ScalingConfigurationId
}
func (r *DescribeScalingInstancesRequest) SetHealthStatus(value string) {
	r.HealthStatus = value
	r.QueryParams.Set("HealthStatus", value)
}
func (r *DescribeScalingInstancesRequest) GetHealthStatus() string {
	return r.HealthStatus
}
func (r *DescribeScalingInstancesRequest) SetLifecycleState(value string) {
	r.LifecycleState = value
	r.QueryParams.Set("LifecycleState", value)
}
func (r *DescribeScalingInstancesRequest) GetLifecycleState() string {
	return r.LifecycleState
}
func (r *DescribeScalingInstancesRequest) SetCreationType(value string) {
	r.CreationType = value
	r.QueryParams.Set("CreationType", value)
}
func (r *DescribeScalingInstancesRequest) GetCreationType() string {
	return r.CreationType
}
func (r *DescribeScalingInstancesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeScalingInstancesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeScalingInstancesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeScalingInstancesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_1(value string) {
	r.InstanceId_1 = value
	r.QueryParams.Set("InstanceId.1", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_1() string {
	return r.InstanceId_1
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_2(value string) {
	r.InstanceId_2 = value
	r.QueryParams.Set("InstanceId.2", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_2() string {
	return r.InstanceId_2
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_3(value string) {
	r.InstanceId_3 = value
	r.QueryParams.Set("InstanceId.3", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_3() string {
	return r.InstanceId_3
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_4(value string) {
	r.InstanceId_4 = value
	r.QueryParams.Set("InstanceId.4", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_4() string {
	return r.InstanceId_4
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_5(value string) {
	r.InstanceId_5 = value
	r.QueryParams.Set("InstanceId.5", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_5() string {
	return r.InstanceId_5
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_6(value string) {
	r.InstanceId_6 = value
	r.QueryParams.Set("InstanceId.6", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_6() string {
	return r.InstanceId_6
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_7(value string) {
	r.InstanceId_7 = value
	r.QueryParams.Set("InstanceId.7", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_7() string {
	return r.InstanceId_7
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_8(value string) {
	r.InstanceId_8 = value
	r.QueryParams.Set("InstanceId.8", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_8() string {
	return r.InstanceId_8
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_9(value string) {
	r.InstanceId_9 = value
	r.QueryParams.Set("InstanceId.9", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_9() string {
	return r.InstanceId_9
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_10(value string) {
	r.InstanceId_10 = value
	r.QueryParams.Set("InstanceId.10", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_10() string {
	return r.InstanceId_10
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_11(value string) {
	r.InstanceId_11 = value
	r.QueryParams.Set("InstanceId.11", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_11() string {
	return r.InstanceId_11
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_12(value string) {
	r.InstanceId_12 = value
	r.QueryParams.Set("InstanceId.12", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_12() string {
	return r.InstanceId_12
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_13(value string) {
	r.InstanceId_13 = value
	r.QueryParams.Set("InstanceId.13", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_13() string {
	return r.InstanceId_13
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_14(value string) {
	r.InstanceId_14 = value
	r.QueryParams.Set("InstanceId.14", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_14() string {
	return r.InstanceId_14
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_15(value string) {
	r.InstanceId_15 = value
	r.QueryParams.Set("InstanceId.15", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_15() string {
	return r.InstanceId_15
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_16(value string) {
	r.InstanceId_16 = value
	r.QueryParams.Set("InstanceId.16", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_16() string {
	return r.InstanceId_16
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_17(value string) {
	r.InstanceId_17 = value
	r.QueryParams.Set("InstanceId.17", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_17() string {
	return r.InstanceId_17
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_18(value string) {
	r.InstanceId_18 = value
	r.QueryParams.Set("InstanceId.18", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_18() string {
	return r.InstanceId_18
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_19(value string) {
	r.InstanceId_19 = value
	r.QueryParams.Set("InstanceId.19", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_19() string {
	return r.InstanceId_19
}
func (r *DescribeScalingInstancesRequest) SetInstanceId_20(value string) {
	r.InstanceId_20 = value
	r.QueryParams.Set("InstanceId.20", value)
}
func (r *DescribeScalingInstancesRequest) GetInstanceId_20() string {
	return r.InstanceId_20
}
func (r *DescribeScalingInstancesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeScalingInstancesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeScalingInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeScalingInstances")
	r.SetProduct(Product)
}

type DescribeScalingInstancesResponse struct {
	TotalCount       int `xml:"TotalCount" json:"TotalCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageSize         int `xml:"PageSize" json:"PageSize"`
	ScalingInstances struct {
		ScalingInstance []struct {
			InstanceId             string `xml:"InstanceId" json:"InstanceId"`
			ScalingConfigurationId string `xml:"ScalingConfigurationId" json:"ScalingConfigurationId"`
			ScalingGroupId         string `xml:"ScalingGroupId" json:"ScalingGroupId"`
			HealthStatus           string `xml:"HealthStatus" json:"HealthStatus"`
			LifecycleState         string `xml:"LifecycleState" json:"LifecycleState"`
			CreationTime           string `xml:"CreationTime" json:"CreationTime"`
			CreationType           string `xml:"CreationType" json:"CreationType"`
		} `xml:"ScalingInstance" json:"ScalingInstance"`
	} `xml:"ScalingInstances" json:"ScalingInstances"`
}

func DescribeScalingInstances(req *DescribeScalingInstancesRequest, accessId, accessSecret string) (*DescribeScalingInstancesResponse, error) {
	var pResponse DescribeScalingInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
