package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeScalingGroupsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageNumber           int
	PageSize             int
	ScalingGroupId_1     string
	ScalingGroupId_2     string
	ScalingGroupId_3     string
	ScalingGroupId_4     string
	ScalingGroupId_5     string
	ScalingGroupId_6     string
	ScalingGroupId_7     string
	ScalingGroupId_8     string
	ScalingGroupId_9     string
	ScalingGroupId_10    string
	ScalingGroupId_12    string
	ScalingGroupId_13    string
	ScalingGroupId_14    string
	ScalingGroupId_15    string
	ScalingGroupId_16    string
	ScalingGroupId_17    string
	ScalingGroupId_18    string
	ScalingGroupId_19    string
	ScalingGroupId_20    string
	ScalingGroupName_1   string
	ScalingGroupName_2   string
	ScalingGroupName_3   string
	ScalingGroupName_4   string
	ScalingGroupName_5   string
	ScalingGroupName_6   string
	ScalingGroupName_7   string
	ScalingGroupName_8   string
	ScalingGroupName_9   string
	ScalingGroupName_10  string
	ScalingGroupName_11  string
	ScalingGroupName_12  string
	ScalingGroupName_13  string
	ScalingGroupName_14  string
	ScalingGroupName_15  string
	ScalingGroupName_16  string
	ScalingGroupName_17  string
	ScalingGroupName_18  string
	ScalingGroupName_19  string
	ScalingGroupName_20  string
	OwnerAccount         string
}

func (r *DescribeScalingGroupsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingGroupsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeScalingGroupsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeScalingGroupsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeScalingGroupsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingGroupsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeScalingGroupsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeScalingGroupsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeScalingGroupsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeScalingGroupsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_1(value string) {
	r.ScalingGroupId_1 = value
	r.QueryParams.Set("ScalingGroupId.1", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_1() string {
	return r.ScalingGroupId_1
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_2(value string) {
	r.ScalingGroupId_2 = value
	r.QueryParams.Set("ScalingGroupId.2", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_2() string {
	return r.ScalingGroupId_2
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_3(value string) {
	r.ScalingGroupId_3 = value
	r.QueryParams.Set("ScalingGroupId.3", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_3() string {
	return r.ScalingGroupId_3
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_4(value string) {
	r.ScalingGroupId_4 = value
	r.QueryParams.Set("ScalingGroupId.4", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_4() string {
	return r.ScalingGroupId_4
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_5(value string) {
	r.ScalingGroupId_5 = value
	r.QueryParams.Set("ScalingGroupId.5", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_5() string {
	return r.ScalingGroupId_5
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_6(value string) {
	r.ScalingGroupId_6 = value
	r.QueryParams.Set("ScalingGroupId.6", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_6() string {
	return r.ScalingGroupId_6
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_7(value string) {
	r.ScalingGroupId_7 = value
	r.QueryParams.Set("ScalingGroupId.7", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_7() string {
	return r.ScalingGroupId_7
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_8(value string) {
	r.ScalingGroupId_8 = value
	r.QueryParams.Set("ScalingGroupId.8", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_8() string {
	return r.ScalingGroupId_8
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_9(value string) {
	r.ScalingGroupId_9 = value
	r.QueryParams.Set("ScalingGroupId.9", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_9() string {
	return r.ScalingGroupId_9
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_10(value string) {
	r.ScalingGroupId_10 = value
	r.QueryParams.Set("ScalingGroupId.10", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_10() string {
	return r.ScalingGroupId_10
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_12(value string) {
	r.ScalingGroupId_12 = value
	r.QueryParams.Set("ScalingGroupId.12", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_12() string {
	return r.ScalingGroupId_12
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_13(value string) {
	r.ScalingGroupId_13 = value
	r.QueryParams.Set("ScalingGroupId.13", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_13() string {
	return r.ScalingGroupId_13
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_14(value string) {
	r.ScalingGroupId_14 = value
	r.QueryParams.Set("ScalingGroupId.14", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_14() string {
	return r.ScalingGroupId_14
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_15(value string) {
	r.ScalingGroupId_15 = value
	r.QueryParams.Set("ScalingGroupId.15", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_15() string {
	return r.ScalingGroupId_15
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_16(value string) {
	r.ScalingGroupId_16 = value
	r.QueryParams.Set("ScalingGroupId.16", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_16() string {
	return r.ScalingGroupId_16
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_17(value string) {
	r.ScalingGroupId_17 = value
	r.QueryParams.Set("ScalingGroupId.17", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_17() string {
	return r.ScalingGroupId_17
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_18(value string) {
	r.ScalingGroupId_18 = value
	r.QueryParams.Set("ScalingGroupId.18", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_18() string {
	return r.ScalingGroupId_18
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_19(value string) {
	r.ScalingGroupId_19 = value
	r.QueryParams.Set("ScalingGroupId.19", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_19() string {
	return r.ScalingGroupId_19
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupId_20(value string) {
	r.ScalingGroupId_20 = value
	r.QueryParams.Set("ScalingGroupId.20", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupId_20() string {
	return r.ScalingGroupId_20
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_1(value string) {
	r.ScalingGroupName_1 = value
	r.QueryParams.Set("ScalingGroupName.1", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_1() string {
	return r.ScalingGroupName_1
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_2(value string) {
	r.ScalingGroupName_2 = value
	r.QueryParams.Set("ScalingGroupName.2", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_2() string {
	return r.ScalingGroupName_2
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_3(value string) {
	r.ScalingGroupName_3 = value
	r.QueryParams.Set("ScalingGroupName.3", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_3() string {
	return r.ScalingGroupName_3
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_4(value string) {
	r.ScalingGroupName_4 = value
	r.QueryParams.Set("ScalingGroupName.4", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_4() string {
	return r.ScalingGroupName_4
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_5(value string) {
	r.ScalingGroupName_5 = value
	r.QueryParams.Set("ScalingGroupName.5", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_5() string {
	return r.ScalingGroupName_5
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_6(value string) {
	r.ScalingGroupName_6 = value
	r.QueryParams.Set("ScalingGroupName.6", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_6() string {
	return r.ScalingGroupName_6
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_7(value string) {
	r.ScalingGroupName_7 = value
	r.QueryParams.Set("ScalingGroupName.7", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_7() string {
	return r.ScalingGroupName_7
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_8(value string) {
	r.ScalingGroupName_8 = value
	r.QueryParams.Set("ScalingGroupName.8", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_8() string {
	return r.ScalingGroupName_8
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_9(value string) {
	r.ScalingGroupName_9 = value
	r.QueryParams.Set("ScalingGroupName.9", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_9() string {
	return r.ScalingGroupName_9
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_10(value string) {
	r.ScalingGroupName_10 = value
	r.QueryParams.Set("ScalingGroupName.10", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_10() string {
	return r.ScalingGroupName_10
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_11(value string) {
	r.ScalingGroupName_11 = value
	r.QueryParams.Set("ScalingGroupName.11", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_11() string {
	return r.ScalingGroupName_11
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_12(value string) {
	r.ScalingGroupName_12 = value
	r.QueryParams.Set("ScalingGroupName.12", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_12() string {
	return r.ScalingGroupName_12
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_13(value string) {
	r.ScalingGroupName_13 = value
	r.QueryParams.Set("ScalingGroupName.13", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_13() string {
	return r.ScalingGroupName_13
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_14(value string) {
	r.ScalingGroupName_14 = value
	r.QueryParams.Set("ScalingGroupName.14", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_14() string {
	return r.ScalingGroupName_14
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_15(value string) {
	r.ScalingGroupName_15 = value
	r.QueryParams.Set("ScalingGroupName.15", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_15() string {
	return r.ScalingGroupName_15
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_16(value string) {
	r.ScalingGroupName_16 = value
	r.QueryParams.Set("ScalingGroupName.16", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_16() string {
	return r.ScalingGroupName_16
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_17(value string) {
	r.ScalingGroupName_17 = value
	r.QueryParams.Set("ScalingGroupName.17", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_17() string {
	return r.ScalingGroupName_17
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_18(value string) {
	r.ScalingGroupName_18 = value
	r.QueryParams.Set("ScalingGroupName.18", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_18() string {
	return r.ScalingGroupName_18
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_19(value string) {
	r.ScalingGroupName_19 = value
	r.QueryParams.Set("ScalingGroupName.19", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_19() string {
	return r.ScalingGroupName_19
}
func (r *DescribeScalingGroupsRequest) SetScalingGroupName_20(value string) {
	r.ScalingGroupName_20 = value
	r.QueryParams.Set("ScalingGroupName.20", value)
}
func (r *DescribeScalingGroupsRequest) GetScalingGroupName_20() string {
	return r.ScalingGroupName_20
}
func (r *DescribeScalingGroupsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeScalingGroupsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeScalingGroupsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeScalingGroups")
	r.SetProduct(Product)
}

type DescribeScalingGroupsResponse struct {
	TotalCount    int `xml:"TotalCount" json:"TotalCount"`
	PageNumber    int `xml:"PageNumber" json:"PageNumber"`
	PageSize      int `xml:"PageSize" json:"PageSize"`
	ScalingGroups struct {
		ScalingGroup []struct {
			DefaultCooldown              int    `xml:"DefaultCooldown" json:"DefaultCooldown"`
			MaxSize                      int    `xml:"MaxSize" json:"MaxSize"`
			PendingCapacity              int    `xml:"PendingCapacity" json:"PendingCapacity"`
			LoadBalancerId               string `xml:"LoadBalancerId" json:"LoadBalancerId"`
			RemovingCapacity             int    `xml:"RemovingCapacity" json:"RemovingCapacity"`
			ScalingGroupName             string `xml:"ScalingGroupName" json:"ScalingGroupName"`
			ActiveCapacity               int    `xml:"ActiveCapacity" json:"ActiveCapacity"`
			ActiveScalingConfigurationId string `xml:"ActiveScalingConfigurationId" json:"ActiveScalingConfigurationId"`
			ScalingGroupId               string `xml:"ScalingGroupId" json:"ScalingGroupId"`
			RegionId                     string `xml:"RegionId" json:"RegionId"`
			TotalCapacity                int    `xml:"TotalCapacity" json:"TotalCapacity"`
			MinSize                      int    `xml:"MinSize" json:"MinSize"`
			LifecycleState               string `xml:"LifecycleState" json:"LifecycleState"`
			CreationTime                 string `xml:"CreationTime" json:"CreationTime"`
			RemovalPolicies              struct {
				RemovalPolicy []string `xml:"RemovalPolicy" json:"RemovalPolicy"`
			} `xml:"RemovalPolicies" json:"RemovalPolicies"`
			DBInstanceIds struct {
				DBInstanceId []string `xml:"DBInstanceId" json:"DBInstanceId"`
			} `xml:"DBInstanceIds" json:"DBInstanceIds"`
		} `xml:"ScalingGroup" json:"ScalingGroup"`
	} `xml:"ScalingGroups" json:"ScalingGroups"`
}

func DescribeScalingGroups(req *DescribeScalingGroupsRequest, accessId, accessSecret string) (*DescribeScalingGroupsResponse, error) {
	var pResponse DescribeScalingGroupsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
