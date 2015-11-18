package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeResourceUsageRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *DescribeResourceUsageRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeResourceUsageRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeResourceUsageRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeResourceUsageRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeResourceUsageRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeResourceUsageRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeResourceUsageRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeResourceUsageRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeResourceUsageRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeResourceUsageRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeResourceUsageRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeResourceUsageRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeResourceUsageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeResourceUsage")
	r.SetProduct(Product)
}

type DescribeResourceUsageResponse struct {
	DBInstanceId   string `xml:"DBInstanceId" json:"DBInstanceId"`
	Engine         string `xml:"Engine" json:"Engine"`
	DiskUsed       int    `xml:"DiskUsed" json:"DiskUsed"`
	DataSize       int    `xml:"DataSize" json:"DataSize"`
	LogSize        int    `xml:"LogSize" json:"LogSize"`
	BackupSize     int    `xml:"BackupSize" json:"BackupSize"`
	SQLSize        int    `xml:"SQLSize" json:"SQLSize"`
	ColdBackupSize int    `xml:"ColdBackupSize" json:"ColdBackupSize"`
}

func DescribeResourceUsage(req *DescribeResourceUsageRequest, accessId, accessSecret string) (*DescribeResourceUsageResponse, error) {
	var pResponse DescribeResourceUsageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
