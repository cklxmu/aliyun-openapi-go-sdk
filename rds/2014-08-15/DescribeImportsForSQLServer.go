package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeImportsForSQLServerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	ImportId             int
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeImportsForSQLServerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeImportsForSQLServerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeImportsForSQLServerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeImportsForSQLServerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeImportsForSQLServerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeImportsForSQLServerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeImportsForSQLServerRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeImportsForSQLServerRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeImportsForSQLServerRequest) SetImportId(value int) {
	r.ImportId = value
	r.QueryParams.Set("ImportId", strconv.Itoa(value))
}
func (r *DescribeImportsForSQLServerRequest) GetImportId() int {
	return r.ImportId
}
func (r *DescribeImportsForSQLServerRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeImportsForSQLServerRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeImportsForSQLServerRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeImportsForSQLServerRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeImportsForSQLServerRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeImportsForSQLServerRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeImportsForSQLServerRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeImportsForSQLServerRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeImportsForSQLServerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeImportsForSQLServerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeImportsForSQLServerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeImportsForSQLServer")
	r.SetProduct(Product)
}

type DescribeImportsForSQLServerResponse struct {
	TotalRecordCounts int `xml:"TotalRecordCounts" json:"TotalRecordCounts"`
	PageNumber        int `xml:"PageNumber" json:"PageNumber"`
	SQLItemsCounts    int `xml:"SQLItemsCounts" json:"SQLItemsCounts"`
	Items             struct {
		SQLServerImport []struct {
			ImportId     int    `xml:"ImportId" json:"ImportId"`
			FileName     string `xml:"FileName" json:"FileName"`
			DBName       string `xml:"DBName" json:"DBName"`
			ImportStatus string `xml:"ImportStatus" json:"ImportStatus"`
			StartTime    string `xml:"StartTime" json:"StartTime"`
		} `xml:"SQLServerImport" json:"SQLServerImport"`
	} `xml:"Items" json:"Items"`
}

func DescribeImportsForSQLServer(req *DescribeImportsForSQLServerRequest, accessId, accessSecret string) (*DescribeImportsForSQLServerResponse, error) {
	var pResponse DescribeImportsForSQLServerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
