package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeSQLLogReportsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	StartTime            string
	EndTime              string
	ReportType           string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeSQLLogReportsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeSQLLogReportsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeSQLLogReportsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeSQLLogReportsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeSQLLogReportsRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeSQLLogReportsRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeSQLLogReportsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeSQLLogReportsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeSQLLogReportsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeSQLLogReportsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeSQLLogReportsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeSQLLogReportsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeSQLLogReportsRequest) SetReportType(value string) {
	r.ReportType = value
	r.QueryParams.Set("ReportType", value)
}
func (r *DescribeSQLLogReportsRequest) GetReportType() string {
	return r.ReportType
}
func (r *DescribeSQLLogReportsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeSQLLogReportsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeSQLLogReportsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeSQLLogReportsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeSQLLogReportsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeSQLLogReportsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeSQLLogReports")
	r.SetProduct(Product)
}

type DescribeSQLLogReportsResponse struct {
	TotalRecordCount string `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       string `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  string `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		SQLReport []struct {
			SQLText              string `xml:"SQLText" json:"SQLText"`
			TotalExecutionCounts string `xml:"TotalExecutionCounts" json:"TotalExecutionCounts"`
			ReturnTotalRowCounts string `xml:"ReturnTotalRowCounts" json:"ReturnTotalRowCounts"`
			TotalExecutionTimes  string `xml:"TotalExecutionTimes" json:"TotalExecutionTimes"`
		} `xml:"SQLReport" json:"SQLReport"`
	} `xml:"Items" json:"Items"`
}

func DescribeSQLLogReports(req *DescribeSQLLogReportsRequest, accessId, accessSecret string) (*DescribeSQLLogReportsResponse, error) {
	var pResponse DescribeSQLLogReportsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
