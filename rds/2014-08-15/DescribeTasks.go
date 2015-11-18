package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeTasksRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	Status               string
	TaskAction           string
	OwnerAccount         string
}

func (r *DescribeTasksRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeTasksRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeTasksRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeTasksRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeTasksRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeTasksRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeTasksRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeTasksRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeTasksRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeTasksRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeTasksRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeTasksRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeTasksRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeTasksRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeTasksRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeTasksRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeTasksRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *DescribeTasksRequest) GetStatus() string {
	return r.Status
}
func (r *DescribeTasksRequest) SetTaskAction(value string) {
	r.TaskAction = value
	r.QueryParams.Set("TaskAction", value)
}
func (r *DescribeTasksRequest) GetTaskAction() string {
	return r.TaskAction
}
func (r *DescribeTasksRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeTasksRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeTasksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeTasks")
	r.SetProduct(Product)
}

type DescribeTasksResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		TaskProgressInfo []struct {
			DBName             string `xml:"DBName" json:"DBName"`
			BeginTime          string `xml:"BeginTime" json:"BeginTime"`
			ProgressInfo       string `xml:"ProgressInfo" json:"ProgressInfo"`
			FinishTime         string `xml:"FinishTime" json:"FinishTime"`
			TaskAction         string `xml:"TaskAction" json:"TaskAction"`
			TaskId             string `xml:"TaskId" json:"TaskId"`
			Progress           string `xml:"Progress" json:"Progress"`
			ExpectedFinishTime string `xml:"ExpectedFinishTime" json:"ExpectedFinishTime"`
			Status             string `xml:"Status" json:"Status"`
			TaskErrorCode      string `xml:"TaskErrorCode" json:"TaskErrorCode"`
			TaskErrorMessage   string `xml:"TaskErrorMessage" json:"TaskErrorMessage"`
		} `xml:"TaskProgressInfo" json:"TaskProgressInfo"`
	} `xml:"Items" json:"Items"`
}

func DescribeTasks(req *DescribeTasksRequest, accessId, accessSecret string) (*DescribeTasksResponse, error) {
	var pResponse DescribeTasksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
