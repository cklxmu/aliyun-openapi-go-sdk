package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSlowLogsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	StartTime            string
	EndTime              string
	DBName               string
	SortKey              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeSlowLogsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSlowLogsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSlowLogsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSlowLogsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSlowLogsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSlowLogsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSlowLogsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeSlowLogsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeSlowLogsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeSlowLogsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeSlowLogsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeSlowLogsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeSlowLogsRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *DescribeSlowLogsRequest) GetDBName() string {
	return r.DBName
}
func (r *DescribeSlowLogsRequest) SetSortKey(value string) {
	r.SortKey = value
	r.QueryParams.Set("SortKey", value)
}
func (r *DescribeSlowLogsRequest) GetSortKey() string {
	return r.SortKey
}
func (r *DescribeSlowLogsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSlowLogsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSlowLogsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSlowLogsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSlowLogsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSlowLogsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeSlowLogsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSlowLogs")
	r.SetProduct(Product)
}

type DescribeSlowLogsResponse struct {
	Engine           string `xml:"Engine" json:"Engine"`
	StartTime        string `xml:"StartTime" json:"StartTime"`
	EndTime          string `xml:"EndTime" json:"EndTime"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		SQLSlowLog []struct {
			SlowLogId                     int    `xml:"SlowLogId" json:"SlowLogId"`
			SQLId                         int    `xml:"SQLId" json:"SQLId"`
			DBName                        string `xml:"DBName" json:"DBName"`
			SQLText                       string `xml:"SQLText" json:"SQLText"`
			MySQLTotalExecutionCounts     int    `xml:"MySQLTotalExecutionCounts" json:"MySQLTotalExecutionCounts"`
			MySQLTotalExecutionTimes      int    `xml:"MySQLTotalExecutionTimes" json:"MySQLTotalExecutionTimes"`
			TotalLockTimes                int    `xml:"TotalLockTimes" json:"TotalLockTimes"`
			MaxLockTime                   int    `xml:"MaxLockTime" json:"MaxLockTime"`
			ParseTotalRowCounts           int    `xml:"ParseTotalRowCounts" json:"ParseTotalRowCounts"`
			ParseMaxRowCount              int    `xml:"ParseMaxRowCount" json:"ParseMaxRowCount"`
			ReturnTotalRowCounts          int    `xml:"ReturnTotalRowCounts" json:"ReturnTotalRowCounts"`
			ReturnMaxRowCount             int    `xml:"ReturnMaxRowCount" json:"ReturnMaxRowCount"`
			CreateTime                    string `xml:"CreateTime" json:"CreateTime"`
			SQLServerTotalExecutionCounts int    `xml:"SQLServerTotalExecutionCounts" json:"SQLServerTotalExecutionCounts"`
			SQLServerTotalExecutionTimes  int    `xml:"SQLServerTotalExecutionTimes" json:"SQLServerTotalExecutionTimes"`
			TotalLogicalReadCounts        int    `xml:"TotalLogicalReadCounts" json:"TotalLogicalReadCounts"`
			TotalPhysicalReadCounts       int    `xml:"TotalPhysicalReadCounts" json:"TotalPhysicalReadCounts"`
			ReportTime                    string `xml:"ReportTime" json:"ReportTime"`
			MaxExecutionTime              int    `xml:"MaxExecutionTime" json:"MaxExecutionTime"`
			AvgExecutionTime              int    `xml:"AvgExecutionTime" json:"AvgExecutionTime"`
		} `xml:"SQLSlowLog" json:"SQLSlowLog"`
	} `xml:"Items" json:"Items"`
}

func DescribeSlowLogs(req *DescribeSlowLogsRequest, accessId, accessSecret string) (*DescribeSlowLogsResponse, error) {
	var pResponse DescribeSlowLogsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
