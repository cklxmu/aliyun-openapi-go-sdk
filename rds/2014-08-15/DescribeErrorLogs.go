package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeErrorLogsRequest struct {
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

func (r *DescribeErrorLogsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeErrorLogsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeErrorLogsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeErrorLogsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeErrorLogsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeErrorLogsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeErrorLogsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeErrorLogsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeErrorLogsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeErrorLogsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeErrorLogsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeErrorLogsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeErrorLogsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeErrorLogsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeErrorLogsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeErrorLogsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeErrorLogsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeErrorLogsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeErrorLogsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeErrorLogs")
	r.SetProduct(Product)
}

type DescribeErrorLogsResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		ErrorLog []struct {
			ErrorInfo  string `xml:"ErrorInfo" json:"ErrorInfo"`
			CreateTime string `xml:"CreateTime" json:"CreateTime"`
		} `xml:"ErrorLog" json:"ErrorLog"`
	} `xml:"Items" json:"Items"`
}

func DescribeErrorLogs(req *DescribeErrorLogsRequest, accessId, accessSecret string) (*DescribeErrorLogsResponse, error) {
	var pResponse DescribeErrorLogsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
