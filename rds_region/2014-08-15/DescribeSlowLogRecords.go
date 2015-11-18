package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSlowLogRecordsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	SQLId                int
	StartTime            string
	EndTime              string
	DBName               string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeSlowLogRecordsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSlowLogRecordsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSlowLogRecordsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSlowLogRecordsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSlowLogRecordsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSlowLogRecordsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeSlowLogRecordsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeSlowLogRecordsRequest) SetSQLId(value int) {
	r.SQLId = value
	r.QueryParams.Set("SQLId", strconv.Itoa(value))
}
func (r *DescribeSlowLogRecordsRequest) GetSQLId() int {
	return r.SQLId
}
func (r *DescribeSlowLogRecordsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeSlowLogRecordsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeSlowLogRecordsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeSlowLogRecordsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeSlowLogRecordsRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *DescribeSlowLogRecordsRequest) GetDBName() string {
	return r.DBName
}
func (r *DescribeSlowLogRecordsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSlowLogRecordsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSlowLogRecordsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSlowLogRecordsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSlowLogRecordsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSlowLogRecordsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeSlowLogRecordsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSlowLogRecords")
	r.SetProduct(Product)
}

type DescribeSlowLogRecordsResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		SQLSlowRecord []struct {
			HostAddress        string `xml:"HostAddress" json:"HostAddress"`
			DBName             string `xml:"DBName" json:"DBName"`
			SQLText            string `xml:"SQLText" json:"SQLText"`
			QueryTimes         int    `xml:"QueryTimes" json:"QueryTimes"`
			LockTimes          int    `xml:"LockTimes" json:"LockTimes"`
			ParseRowCounts     int    `xml:"ParseRowCounts" json:"ParseRowCounts"`
			ReturnRowCounts    int    `xml:"ReturnRowCounts" json:"ReturnRowCounts"`
			ExecutionStartTime string `xml:"ExecutionStartTime" json:"ExecutionStartTime"`
		} `xml:"SQLSlowRecord" json:"SQLSlowRecord"`
	} `xml:"Items" json:"Items"`
}

func DescribeSlowLogRecords(req *DescribeSlowLogRecordsRequest, accessId, accessSecret string) (*DescribeSlowLogRecordsResponse, error) {
	var pResponse DescribeSlowLogRecordsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
