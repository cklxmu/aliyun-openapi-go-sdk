package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSQLLogRecordsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	SQLId                int
	QueryKeywords        string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeSQLLogRecordsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLLogRecordsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSQLLogRecordsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSQLLogRecordsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSQLLogRecordsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLLogRecordsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSQLLogRecordsRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeSQLLogRecordsRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeSQLLogRecordsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeSQLLogRecordsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeSQLLogRecordsRequest) SetSQLId(value int) {
	r.SQLId = value
	r.QueryParams.Set("SQLId", strconv.Itoa(value))
}
func (r *DescribeSQLLogRecordsRequest) GetSQLId() int {
	return r.SQLId
}
func (r *DescribeSQLLogRecordsRequest) SetQueryKeywords(value string) {
	r.QueryKeywords = value
	r.QueryParams.Set("QueryKeywords", value)
}
func (r *DescribeSQLLogRecordsRequest) GetQueryKeywords() string {
	return r.QueryKeywords
}
func (r *DescribeSQLLogRecordsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeSQLLogRecordsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeSQLLogRecordsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeSQLLogRecordsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeSQLLogRecordsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSQLLogRecordsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSQLLogRecordsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSQLLogRecordsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSQLLogRecordsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSQLLogRecordsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeSQLLogRecordsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSQLLogRecords")
	r.SetProduct(Product)
}

type DescribeSQLLogRecordsResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		SQLRecord []struct {
			DBName              string `xml:"DBName" json:"DBName"`
			AccountName         string `xml:"AccountName" json:"AccountName"`
			HostAddress         string `xml:"HostAddress" json:"HostAddress"`
			SQLText             string `xml:"SQLText" json:"SQLText"`
			TotalExecutionTimes int    `xml:"TotalExecutionTimes" json:"TotalExecutionTimes"`
			ReturnRowCounts     int    `xml:"ReturnRowCounts" json:"ReturnRowCounts"`
			ExecuteTime         string `xml:"ExecuteTime" json:"ExecuteTime"`
		} `xml:"SQLRecord" json:"SQLRecord"`
	} `xml:"Items" json:"Items"`
}

func DescribeSQLLogRecords(req *DescribeSQLLogRecordsRequest, accessId, accessSecret string) (*DescribeSQLLogRecordsResponse, error) {
	var pResponse DescribeSQLLogRecordsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
