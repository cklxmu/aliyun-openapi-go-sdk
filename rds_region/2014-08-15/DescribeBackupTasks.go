package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeBackupTasksRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	Flag                 string
	OwnerAccount         string
	DBInstanceId         string
	BackupJobId          string
	BackupMode           string
	BackupJobStatus      string
}

func (r *DescribeBackupTasksRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeBackupTasksRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeBackupTasksRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeBackupTasksRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeBackupTasksRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeBackupTasksRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeBackupTasksRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeBackupTasksRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeBackupTasksRequest) SetFlag(value string) {
	r.Flag = value
	r.QueryParams.Set("Flag", value)
}
func (r *DescribeBackupTasksRequest) GetFlag() string {
	return r.Flag
}
func (r *DescribeBackupTasksRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeBackupTasksRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}
func (r *DescribeBackupTasksRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeBackupTasksRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeBackupTasksRequest) SetBackupJobId(value string) {
	r.BackupJobId = value
	r.QueryParams.Set("BackupJobId", value)
}
func (r *DescribeBackupTasksRequest) GetBackupJobId() string {
	return r.BackupJobId
}
func (r *DescribeBackupTasksRequest) SetBackupMode(value string) {
	r.BackupMode = value
	r.QueryParams.Set("BackupMode", value)
}
func (r *DescribeBackupTasksRequest) GetBackupMode() string {
	return r.BackupMode
}
func (r *DescribeBackupTasksRequest) SetBackupJobStatus(value string) {
	r.BackupJobStatus = value
	r.QueryParams.Set("BackupJobStatus", value)
}
func (r *DescribeBackupTasksRequest) GetBackupJobStatus() string {
	return r.BackupJobStatus
}

func (r *DescribeBackupTasksRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeBackupTasks")
	r.SetProduct(Product)
}

type DescribeBackupTasksResponse struct {
	Items struct {
		BackupJob []struct {
			BackupProgressStatus string `xml:"BackupProgressStatus" json:"BackupProgressStatus"`
			JobMode              string `xml:"JobMode" json:"JobMode"`
			Process              string `xml:"Process" json:"Process"`
			TaskAction           string `xml:"TaskAction" json:"TaskAction"`
			BackupjobId          string `xml:"BackupjobId" json:"BackupjobId"`
		} `xml:"BackupJob" json:"BackupJob"`
	} `xml:"Items" json:"Items"`
}

func DescribeBackupTasks(req *DescribeBackupTasksRequest, accessId, accessSecret string) (*DescribeBackupTasksResponse, error) {
	var pResponse DescribeBackupTasksResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
