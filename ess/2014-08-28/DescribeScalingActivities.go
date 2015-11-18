package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeScalingActivitiesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageNumber           int
	PageSize             int
	ScalingGroupId       string
	StatusCode           string
	ScalingActivityId_1  string
	ScalingActivityId_2  string
	ScalingActivityId_3  string
	ScalingActivityId_4  string
	ScalingActivityId_5  string
	ScalingActivityId_6  string
	ScalingActivityId_7  string
	ScalingActivityId_8  string
	ScalingActivityId_9  string
	ScalingActivityId_10 string
	ScalingActivityId_11 string
	ScalingActivityId_12 string
	ScalingActivityId_13 string
	ScalingActivityId_14 string
	ScalingActivityId_15 string
	ScalingActivityId_16 string
	ScalingActivityId_17 string
	ScalingActivityId_18 string
	ScalingActivityId_19 string
	ScalingActivityId_20 string
	OwnerAccount         string
}

func (r *DescribeScalingActivitiesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingActivitiesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeScalingActivitiesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeScalingActivitiesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeScalingActivitiesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeScalingActivitiesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeScalingActivitiesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeScalingActivitiesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeScalingActivitiesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeScalingActivitiesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeScalingActivitiesRequest) SetScalingGroupId(value string) {
	r.ScalingGroupId = value
	r.QueryParams.Set("ScalingGroupId", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingGroupId() string {
	return r.ScalingGroupId
}
func (r *DescribeScalingActivitiesRequest) SetStatusCode(value string) {
	r.StatusCode = value
	r.QueryParams.Set("StatusCode", value)
}
func (r *DescribeScalingActivitiesRequest) GetStatusCode() string {
	return r.StatusCode
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_1(value string) {
	r.ScalingActivityId_1 = value
	r.QueryParams.Set("ScalingActivityId.1", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_1() string {
	return r.ScalingActivityId_1
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_2(value string) {
	r.ScalingActivityId_2 = value
	r.QueryParams.Set("ScalingActivityId.2", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_2() string {
	return r.ScalingActivityId_2
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_3(value string) {
	r.ScalingActivityId_3 = value
	r.QueryParams.Set("ScalingActivityId.3", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_3() string {
	return r.ScalingActivityId_3
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_4(value string) {
	r.ScalingActivityId_4 = value
	r.QueryParams.Set("ScalingActivityId.4", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_4() string {
	return r.ScalingActivityId_4
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_5(value string) {
	r.ScalingActivityId_5 = value
	r.QueryParams.Set("ScalingActivityId.5", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_5() string {
	return r.ScalingActivityId_5
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_6(value string) {
	r.ScalingActivityId_6 = value
	r.QueryParams.Set("ScalingActivityId.6", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_6() string {
	return r.ScalingActivityId_6
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_7(value string) {
	r.ScalingActivityId_7 = value
	r.QueryParams.Set("ScalingActivityId.7", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_7() string {
	return r.ScalingActivityId_7
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_8(value string) {
	r.ScalingActivityId_8 = value
	r.QueryParams.Set("ScalingActivityId.8", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_8() string {
	return r.ScalingActivityId_8
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_9(value string) {
	r.ScalingActivityId_9 = value
	r.QueryParams.Set("ScalingActivityId.9", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_9() string {
	return r.ScalingActivityId_9
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_10(value string) {
	r.ScalingActivityId_10 = value
	r.QueryParams.Set("ScalingActivityId.10", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_10() string {
	return r.ScalingActivityId_10
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_11(value string) {
	r.ScalingActivityId_11 = value
	r.QueryParams.Set("ScalingActivityId.11", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_11() string {
	return r.ScalingActivityId_11
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_12(value string) {
	r.ScalingActivityId_12 = value
	r.QueryParams.Set("ScalingActivityId.12", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_12() string {
	return r.ScalingActivityId_12
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_13(value string) {
	r.ScalingActivityId_13 = value
	r.QueryParams.Set("ScalingActivityId.13", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_13() string {
	return r.ScalingActivityId_13
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_14(value string) {
	r.ScalingActivityId_14 = value
	r.QueryParams.Set("ScalingActivityId.14", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_14() string {
	return r.ScalingActivityId_14
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_15(value string) {
	r.ScalingActivityId_15 = value
	r.QueryParams.Set("ScalingActivityId.15", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_15() string {
	return r.ScalingActivityId_15
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_16(value string) {
	r.ScalingActivityId_16 = value
	r.QueryParams.Set("ScalingActivityId.16", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_16() string {
	return r.ScalingActivityId_16
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_17(value string) {
	r.ScalingActivityId_17 = value
	r.QueryParams.Set("ScalingActivityId.17", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_17() string {
	return r.ScalingActivityId_17
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_18(value string) {
	r.ScalingActivityId_18 = value
	r.QueryParams.Set("ScalingActivityId.18", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_18() string {
	return r.ScalingActivityId_18
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_19(value string) {
	r.ScalingActivityId_19 = value
	r.QueryParams.Set("ScalingActivityId.19", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_19() string {
	return r.ScalingActivityId_19
}
func (r *DescribeScalingActivitiesRequest) SetScalingActivityId_20(value string) {
	r.ScalingActivityId_20 = value
	r.QueryParams.Set("ScalingActivityId.20", value)
}
func (r *DescribeScalingActivitiesRequest) GetScalingActivityId_20() string {
	return r.ScalingActivityId_20
}
func (r *DescribeScalingActivitiesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeScalingActivitiesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeScalingActivitiesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeScalingActivities")
	r.SetProduct(Product)
}

type DescribeScalingActivitiesResponse struct {
	TotalCount        int `xml:"TotalCount" json:"TotalCount"`
	PageNumber        int `xml:"PageNumber" json:"PageNumber"`
	PageSize          int `xml:"PageSize" json:"PageSize"`
	ScalingActivities struct {
		ScalingActivity []struct {
			ScalingActivityId string `xml:"ScalingActivityId" json:"ScalingActivityId"`
			ScalingGroupId    string `xml:"ScalingGroupId" json:"ScalingGroupId"`
			Description       string `xml:"Description" json:"Description"`
			Cause             string `xml:"Cause" json:"Cause"`
			StartTime         string `xml:"StartTime" json:"StartTime"`
			EndTime           string `xml:"EndTime" json:"EndTime"`
			Progress          int    `xml:"Progress" json:"Progress"`
			StatusCode        string `xml:"StatusCode" json:"StatusCode"`
			StatusMessage     string `xml:"StatusMessage" json:"StatusMessage"`
		} `xml:"ScalingActivity" json:"ScalingActivity"`
	} `xml:"ScalingActivities" json:"ScalingActivities"`
}

func DescribeScalingActivities(req *DescribeScalingActivitiesRequest, accessId, accessSecret string) (*DescribeScalingActivitiesResponse, error) {
	var pResponse DescribeScalingActivitiesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
