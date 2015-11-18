package ess

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeScheduledTasksRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	PageNumber           int
	PageSize             int
	ScheduledAction_1    string
	ScheduledAction_2    string
	ScheduledAction_3    string
	ScheduledAction_4    string
	ScheduledAction_5    string
	ScheduledAction_6    string
	ScheduledAction_7    string
	ScheduledAction_8    string
	ScheduledAction_9    string
	ScheduledAction_10   string
	ScheduledAction_11   string
	ScheduledAction_12   string
	ScheduledAction_13   string
	ScheduledAction_14   string
	ScheduledAction_15   string
	ScheduledAction_16   string
	ScheduledAction_17   string
	ScheduledAction_18   string
	ScheduledAction_19   string
	ScheduledAction_20   string
	ScheduledTaskId_1    string
	ScheduledTaskId_2    string
	ScheduledTaskId_3    string
	ScheduledTaskId_4    string
	ScheduledTaskId_5    string
	ScheduledTaskId_6    string
	ScheduledTaskId_7    string
	ScheduledTaskId_8    string
	ScheduledTaskId_9    string
	ScheduledTaskId_10   string
	ScheduledTaskId_11   string
	ScheduledTaskId_12   string
	ScheduledTaskId_13   string
	ScheduledTaskId_14   string
	ScheduledTaskId_15   string
	ScheduledTaskId_16   string
	ScheduledTaskId_17   string
	ScheduledTaskId_18   string
	ScheduledTaskId_19   string
	ScheduledTaskId_20   string
	ScheduledTaskName_1  string
	ScheduledTaskName_2  string
	ScheduledTaskName_3  string
	ScheduledTaskName_4  string
	ScheduledTaskName_5  string
	ScheduledTaskName_6  string
	ScheduledTaskName_7  string
	ScheduledTaskName_8  string
	ScheduledTaskName_9  string
	ScheduledTaskName_10 string
	ScheduledTaskName_11 string
	ScheduledTaskName_12 string
	ScheduledTaskName_13 string
	ScheduledTaskName_14 string
	ScheduledTaskName_15 string
	ScheduledTaskName_16 string
	ScheduledTaskName_17 string
	ScheduledTaskName_18 string
	ScheduledTaskName_19 string
	ScheduledTaskName_20 string
	OwnerAccount         string
}

func (r *DescribeScheduledTasksRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeScheduledTasksRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeScheduledTasksRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeScheduledTasksRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeScheduledTasksRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeScheduledTasksRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeScheduledTasksRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeScheduledTasksRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeScheduledTasksRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeScheduledTasksRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_1(value string) {
	r.ScheduledAction_1 = value
	r.QueryParams.Set("ScheduledAction.1", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_1() string {
	return r.ScheduledAction_1
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_2(value string) {
	r.ScheduledAction_2 = value
	r.QueryParams.Set("ScheduledAction.2", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_2() string {
	return r.ScheduledAction_2
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_3(value string) {
	r.ScheduledAction_3 = value
	r.QueryParams.Set("ScheduledAction.3", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_3() string {
	return r.ScheduledAction_3
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_4(value string) {
	r.ScheduledAction_4 = value
	r.QueryParams.Set("ScheduledAction.4", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_4() string {
	return r.ScheduledAction_4
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_5(value string) {
	r.ScheduledAction_5 = value
	r.QueryParams.Set("ScheduledAction.5", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_5() string {
	return r.ScheduledAction_5
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_6(value string) {
	r.ScheduledAction_6 = value
	r.QueryParams.Set("ScheduledAction.6", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_6() string {
	return r.ScheduledAction_6
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_7(value string) {
	r.ScheduledAction_7 = value
	r.QueryParams.Set("ScheduledAction.7", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_7() string {
	return r.ScheduledAction_7
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_8(value string) {
	r.ScheduledAction_8 = value
	r.QueryParams.Set("ScheduledAction.8", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_8() string {
	return r.ScheduledAction_8
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_9(value string) {
	r.ScheduledAction_9 = value
	r.QueryParams.Set("ScheduledAction.9", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_9() string {
	return r.ScheduledAction_9
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_10(value string) {
	r.ScheduledAction_10 = value
	r.QueryParams.Set("ScheduledAction.10", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_10() string {
	return r.ScheduledAction_10
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_11(value string) {
	r.ScheduledAction_11 = value
	r.QueryParams.Set("ScheduledAction.11", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_11() string {
	return r.ScheduledAction_11
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_12(value string) {
	r.ScheduledAction_12 = value
	r.QueryParams.Set("ScheduledAction.12", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_12() string {
	return r.ScheduledAction_12
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_13(value string) {
	r.ScheduledAction_13 = value
	r.QueryParams.Set("ScheduledAction.13", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_13() string {
	return r.ScheduledAction_13
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_14(value string) {
	r.ScheduledAction_14 = value
	r.QueryParams.Set("ScheduledAction.14", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_14() string {
	return r.ScheduledAction_14
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_15(value string) {
	r.ScheduledAction_15 = value
	r.QueryParams.Set("ScheduledAction.15", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_15() string {
	return r.ScheduledAction_15
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_16(value string) {
	r.ScheduledAction_16 = value
	r.QueryParams.Set("ScheduledAction.16", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_16() string {
	return r.ScheduledAction_16
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_17(value string) {
	r.ScheduledAction_17 = value
	r.QueryParams.Set("ScheduledAction.17", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_17() string {
	return r.ScheduledAction_17
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_18(value string) {
	r.ScheduledAction_18 = value
	r.QueryParams.Set("ScheduledAction.18", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_18() string {
	return r.ScheduledAction_18
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_19(value string) {
	r.ScheduledAction_19 = value
	r.QueryParams.Set("ScheduledAction.19", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_19() string {
	return r.ScheduledAction_19
}
func (r *DescribeScheduledTasksRequest) SetScheduledAction_20(value string) {
	r.ScheduledAction_20 = value
	r.QueryParams.Set("ScheduledAction.20", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledAction_20() string {
	return r.ScheduledAction_20
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_1(value string) {
	r.ScheduledTaskId_1 = value
	r.QueryParams.Set("ScheduledTaskId.1", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_1() string {
	return r.ScheduledTaskId_1
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_2(value string) {
	r.ScheduledTaskId_2 = value
	r.QueryParams.Set("ScheduledTaskId.2", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_2() string {
	return r.ScheduledTaskId_2
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_3(value string) {
	r.ScheduledTaskId_3 = value
	r.QueryParams.Set("ScheduledTaskId.3", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_3() string {
	return r.ScheduledTaskId_3
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_4(value string) {
	r.ScheduledTaskId_4 = value
	r.QueryParams.Set("ScheduledTaskId.4", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_4() string {
	return r.ScheduledTaskId_4
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_5(value string) {
	r.ScheduledTaskId_5 = value
	r.QueryParams.Set("ScheduledTaskId.5", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_5() string {
	return r.ScheduledTaskId_5
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_6(value string) {
	r.ScheduledTaskId_6 = value
	r.QueryParams.Set("ScheduledTaskId.6", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_6() string {
	return r.ScheduledTaskId_6
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_7(value string) {
	r.ScheduledTaskId_7 = value
	r.QueryParams.Set("ScheduledTaskId.7", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_7() string {
	return r.ScheduledTaskId_7
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_8(value string) {
	r.ScheduledTaskId_8 = value
	r.QueryParams.Set("ScheduledTaskId.8", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_8() string {
	return r.ScheduledTaskId_8
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_9(value string) {
	r.ScheduledTaskId_9 = value
	r.QueryParams.Set("ScheduledTaskId.9", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_9() string {
	return r.ScheduledTaskId_9
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_10(value string) {
	r.ScheduledTaskId_10 = value
	r.QueryParams.Set("ScheduledTaskId.10", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_10() string {
	return r.ScheduledTaskId_10
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_11(value string) {
	r.ScheduledTaskId_11 = value
	r.QueryParams.Set("ScheduledTaskId.11", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_11() string {
	return r.ScheduledTaskId_11
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_12(value string) {
	r.ScheduledTaskId_12 = value
	r.QueryParams.Set("ScheduledTaskId.12", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_12() string {
	return r.ScheduledTaskId_12
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_13(value string) {
	r.ScheduledTaskId_13 = value
	r.QueryParams.Set("ScheduledTaskId.13", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_13() string {
	return r.ScheduledTaskId_13
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_14(value string) {
	r.ScheduledTaskId_14 = value
	r.QueryParams.Set("ScheduledTaskId.14", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_14() string {
	return r.ScheduledTaskId_14
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_15(value string) {
	r.ScheduledTaskId_15 = value
	r.QueryParams.Set("ScheduledTaskId.15", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_15() string {
	return r.ScheduledTaskId_15
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_16(value string) {
	r.ScheduledTaskId_16 = value
	r.QueryParams.Set("ScheduledTaskId.16", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_16() string {
	return r.ScheduledTaskId_16
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_17(value string) {
	r.ScheduledTaskId_17 = value
	r.QueryParams.Set("ScheduledTaskId.17", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_17() string {
	return r.ScheduledTaskId_17
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_18(value string) {
	r.ScheduledTaskId_18 = value
	r.QueryParams.Set("ScheduledTaskId.18", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_18() string {
	return r.ScheduledTaskId_18
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_19(value string) {
	r.ScheduledTaskId_19 = value
	r.QueryParams.Set("ScheduledTaskId.19", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_19() string {
	return r.ScheduledTaskId_19
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskId_20(value string) {
	r.ScheduledTaskId_20 = value
	r.QueryParams.Set("ScheduledTaskId.20", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskId_20() string {
	return r.ScheduledTaskId_20
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_1(value string) {
	r.ScheduledTaskName_1 = value
	r.QueryParams.Set("ScheduledTaskName.1", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_1() string {
	return r.ScheduledTaskName_1
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_2(value string) {
	r.ScheduledTaskName_2 = value
	r.QueryParams.Set("ScheduledTaskName.2", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_2() string {
	return r.ScheduledTaskName_2
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_3(value string) {
	r.ScheduledTaskName_3 = value
	r.QueryParams.Set("ScheduledTaskName.3", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_3() string {
	return r.ScheduledTaskName_3
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_4(value string) {
	r.ScheduledTaskName_4 = value
	r.QueryParams.Set("ScheduledTaskName.4", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_4() string {
	return r.ScheduledTaskName_4
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_5(value string) {
	r.ScheduledTaskName_5 = value
	r.QueryParams.Set("ScheduledTaskName.5", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_5() string {
	return r.ScheduledTaskName_5
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_6(value string) {
	r.ScheduledTaskName_6 = value
	r.QueryParams.Set("ScheduledTaskName.6", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_6() string {
	return r.ScheduledTaskName_6
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_7(value string) {
	r.ScheduledTaskName_7 = value
	r.QueryParams.Set("ScheduledTaskName.7", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_7() string {
	return r.ScheduledTaskName_7
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_8(value string) {
	r.ScheduledTaskName_8 = value
	r.QueryParams.Set("ScheduledTaskName.8", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_8() string {
	return r.ScheduledTaskName_8
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_9(value string) {
	r.ScheduledTaskName_9 = value
	r.QueryParams.Set("ScheduledTaskName.9", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_9() string {
	return r.ScheduledTaskName_9
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_10(value string) {
	r.ScheduledTaskName_10 = value
	r.QueryParams.Set("ScheduledTaskName.10", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_10() string {
	return r.ScheduledTaskName_10
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_11(value string) {
	r.ScheduledTaskName_11 = value
	r.QueryParams.Set("ScheduledTaskName.11", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_11() string {
	return r.ScheduledTaskName_11
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_12(value string) {
	r.ScheduledTaskName_12 = value
	r.QueryParams.Set("ScheduledTaskName.12", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_12() string {
	return r.ScheduledTaskName_12
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_13(value string) {
	r.ScheduledTaskName_13 = value
	r.QueryParams.Set("ScheduledTaskName.13", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_13() string {
	return r.ScheduledTaskName_13
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_14(value string) {
	r.ScheduledTaskName_14 = value
	r.QueryParams.Set("ScheduledTaskName.14", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_14() string {
	return r.ScheduledTaskName_14
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_15(value string) {
	r.ScheduledTaskName_15 = value
	r.QueryParams.Set("ScheduledTaskName.15", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_15() string {
	return r.ScheduledTaskName_15
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_16(value string) {
	r.ScheduledTaskName_16 = value
	r.QueryParams.Set("ScheduledTaskName.16", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_16() string {
	return r.ScheduledTaskName_16
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_17(value string) {
	r.ScheduledTaskName_17 = value
	r.QueryParams.Set("ScheduledTaskName.17", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_17() string {
	return r.ScheduledTaskName_17
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_18(value string) {
	r.ScheduledTaskName_18 = value
	r.QueryParams.Set("ScheduledTaskName.18", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_18() string {
	return r.ScheduledTaskName_18
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_19(value string) {
	r.ScheduledTaskName_19 = value
	r.QueryParams.Set("ScheduledTaskName.19", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_19() string {
	return r.ScheduledTaskName_19
}
func (r *DescribeScheduledTasksRequest) SetScheduledTaskName_20(value string) {
	r.ScheduledTaskName_20 = value
	r.QueryParams.Set("ScheduledTaskName.20", value)
}
func (r *DescribeScheduledTasksRequest) GetScheduledTaskName_20() string {
	return r.ScheduledTaskName_20
}
func (r *DescribeScheduledTasksRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeScheduledTasksRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeScheduledTasksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeScheduledTasks")
	r.SetProduct(Product)
}

type DescribeScheduledTasksResponse struct {
	TotalCount     int `xml:"TotalCount" json:"TotalCount"`
	PageNumber     int `xml:"PageNumber" json:"PageNumber"`
	PageSize       int `xml:"PageSize" json:"PageSize"`
	ScheduledTasks struct {
		ScheduledTask []struct {
			ScheduledTaskId      string `xml:"ScheduledTaskId" json:"ScheduledTaskId"`
			ScheduledTaskName    string `xml:"ScheduledTaskName" json:"ScheduledTaskName"`
			Description          string `xml:"Description" json:"Description"`
			ScheduledAction      string `xml:"ScheduledAction" json:"ScheduledAction"`
			RecurrenceEndTime    string `xml:"RecurrenceEndTime" json:"RecurrenceEndTime"`
			LaunchTime           string `xml:"LaunchTime" json:"LaunchTime"`
			RecurrenceType       string `xml:"RecurrenceType" json:"RecurrenceType"`
			RecurrenceValue      string `xml:"RecurrenceValue" json:"RecurrenceValue"`
			LaunchExpirationTime int    `xml:"LaunchExpirationTime" json:"LaunchExpirationTime"`
			TaskEnabled          bool   `xml:"TaskEnabled" json:"TaskEnabled"`
		} `xml:"ScheduledTask" json:"ScheduledTask"`
	} `xml:"ScheduledTasks" json:"ScheduledTasks"`
}

func DescribeScheduledTasks(req *DescribeScheduledTasksRequest, accessId, accessSecret string) (*DescribeScheduledTasksResponse, error) {
	var pResponse DescribeScheduledTasksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
