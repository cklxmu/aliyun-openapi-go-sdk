package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeRealtimeDiagnosesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeRealtimeDiagnosesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeRealtimeDiagnosesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeRealtimeDiagnosesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeRealtimeDiagnosesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeRealtimeDiagnosesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeRealtimeDiagnosesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeRealtimeDiagnosesRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeRealtimeDiagnosesRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeRealtimeDiagnosesRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeRealtimeDiagnosesRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeRealtimeDiagnosesRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeRealtimeDiagnosesRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeRealtimeDiagnosesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeRealtimeDiagnosesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeRealtimeDiagnosesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeRealtimeDiagnosesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeRealtimeDiagnosesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeRealtimeDiagnosesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeRealtimeDiagnosesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeRealtimeDiagnoses")
	r.SetProduct(Product)
}

type DescribeRealtimeDiagnosesResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Tasks            struct {
		RealtimeDiagnoseTasks []struct {
			CreateTime  string `xml:"CreateTime" json:"CreateTime"`
			TaskId      string `xml:"TaskId" json:"TaskId"`
			HealthScore string `xml:"HealthScore" json:"HealthScore"`
			Status      string `xml:"Status" json:"Status"`
		} `xml:"RealtimeDiagnoseTasks" json:"RealtimeDiagnoseTasks"`
	} `xml:"Tasks" json:"Tasks"`
}

func DescribeRealtimeDiagnoses(req *DescribeRealtimeDiagnosesRequest, accessId, accessSecret string) (*DescribeRealtimeDiagnosesResponse, error) {
	var pResponse DescribeRealtimeDiagnosesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
