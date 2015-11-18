package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeBackupsRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	BackupId             string
	BackupLocation       string
	BackupStatus         string
	BackupMode           string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeBackupsRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeBackupsRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeBackupsRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeBackupsRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeBackupsRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeBackupsRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeBackupsRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeBackupsRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeBackupsRequest) SetBackupId(value string) {
	r.BackupId = value
	r.QueryParams.Set("BackupId", value)
}
func (r *DescribeBackupsRequest) GetBackupId() string {
	return r.BackupId
}
func (r *DescribeBackupsRequest) SetBackupLocation(value string) {
	r.BackupLocation = value
	r.QueryParams.Set("BackupLocation", value)
}
func (r *DescribeBackupsRequest) GetBackupLocation() string {
	return r.BackupLocation
}
func (r *DescribeBackupsRequest) SetBackupStatus(value string) {
	r.BackupStatus = value
	r.QueryParams.Set("BackupStatus", value)
}
func (r *DescribeBackupsRequest) GetBackupStatus() string {
	return r.BackupStatus
}
func (r *DescribeBackupsRequest) SetBackupMode(value string) {
	r.BackupMode = value
	r.QueryParams.Set("BackupMode", value)
}
func (r *DescribeBackupsRequest) GetBackupMode() string {
	return r.BackupMode
}
func (r *DescribeBackupsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeBackupsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeBackupsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeBackupsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeBackupsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeBackupsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeBackupsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeBackupsRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeBackupsRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeBackupsRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeBackupsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeBackups")
	r.SetProduct(Product)
}

type DescribeBackupsResponse struct {
	TotalRecordCount string `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       string `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  string `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		Backup []struct {
			BackupId               string `xml:"BackupId" json:"BackupId"`
			DBInstanceId           string `xml:"DBInstanceId" json:"DBInstanceId"`
			BackupStatus           string `xml:"BackupStatus" json:"BackupStatus"`
			BackupStartTime        string `xml:"BackupStartTime" json:"BackupStartTime"`
			BackupEndTime          string `xml:"BackupEndTime" json:"BackupEndTime"`
			BackupType             string `xml:"BackupType" json:"BackupType"`
			BackupMode             string `xml:"BackupMode" json:"BackupMode"`
			BackupMethod           string `xml:"BackupMethod" json:"BackupMethod"`
			BackupDownloadURL      string `xml:"BackupDownloadURL" json:"BackupDownloadURL"`
			BackupLocation         string `xml:"BackupLocation" json:"BackupLocation"`
			BackupExtractionStatus string `xml:"BackupExtractionStatus" json:"BackupExtractionStatus"`
			BackupScale            string `xml:"BackupScale" json:"BackupScale"`
			BackupDBNames          string `xml:"BackupDBNames" json:"BackupDBNames"`
			BackupSize             int    `xml:"BackupSize" json:"BackupSize"`
		} `xml:"Backup" json:"Backup"`
	} `xml:"Items" json:"Items"`
}

func DescribeBackups(req *DescribeBackupsRequest, accessId, accessSecret string) (*DescribeBackupsResponse, error) {
	var pResponse DescribeBackupsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
