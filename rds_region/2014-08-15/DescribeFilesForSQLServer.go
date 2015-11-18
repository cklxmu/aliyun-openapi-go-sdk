package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeFilesForSQLServerRequest struct {
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

func (r *DescribeFilesForSQLServerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeFilesForSQLServerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeFilesForSQLServerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeFilesForSQLServerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeFilesForSQLServerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeFilesForSQLServerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeFilesForSQLServerRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeFilesForSQLServerRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeFilesForSQLServerRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeFilesForSQLServerRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeFilesForSQLServerRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeFilesForSQLServerRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeFilesForSQLServerRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeFilesForSQLServerRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeFilesForSQLServerRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeFilesForSQLServerRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeFilesForSQLServerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeFilesForSQLServerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeFilesForSQLServerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeFilesForSQLServer")
	r.SetProduct(Product)
}

type DescribeFilesForSQLServerResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	TotalRecordCount int    `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int    `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int    `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		SQLServerUploadFile []struct {
			DBName            string `xml:"DBName" json:"DBName"`
			FileName          string `xml:"FileName" json:"FileName"`
			FileSize          int    `xml:"FileSize" json:"FileSize"`
			InternetFtpServer string `xml:"InternetFtpServer" json:"InternetFtpServer"`
			InternetPort      int    `xml:"InternetPort" json:"InternetPort"`
			IntranetFtpserver string `xml:"IntranetFtpserver" json:"IntranetFtpserver"`
			Intranetport      int    `xml:"Intranetport" json:"Intranetport"`
			UserName          string `xml:"UserName" json:"UserName"`
			Password          string `xml:"Password" json:"Password"`
			FileStatus        string `xml:"FileStatus" json:"FileStatus"`
			Description       string `xml:"Description" json:"Description"`
			CreationTime      string `xml:"CreationTime" json:"CreationTime"`
		} `xml:"SQLServerUploadFile" json:"SQLServerUploadFile"`
	} `xml:"Items" json:"Items"`
}

func DescribeFilesForSQLServer(req *DescribeFilesForSQLServerRequest, accessId, accessSecret string) (*DescribeFilesForSQLServerResponse, error) {
	var pResponse DescribeFilesForSQLServerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
