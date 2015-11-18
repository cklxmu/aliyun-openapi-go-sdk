package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeBackupPolicyRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	OwnerAccount         string
}

func (r *DescribeBackupPolicyRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeBackupPolicyRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeBackupPolicyRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeBackupPolicyRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeBackupPolicyRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeBackupPolicyRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeBackupPolicyRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeBackupPolicyRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeBackupPolicyRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeBackupPolicyRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeBackupPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeBackupPolicy")
	r.SetProduct(Product)
}

type DescribeBackupPolicyResponse struct {
	BackupRetentionPeriod   int    `xml:"BackupRetentionPeriod" json:"BackupRetentionPeriod"`
	PreferredNextBackupTime string `xml:"PreferredNextBackupTime" json:"PreferredNextBackupTime"`
	PreferredBackupTime     string `xml:"PreferredBackupTime" json:"PreferredBackupTime"`
	PreferredBackupPeriod   string `xml:"PreferredBackupPeriod" json:"PreferredBackupPeriod"`
}

func DescribeBackupPolicy(req *DescribeBackupPolicyRequest, accessId, accessSecret string) (*DescribeBackupPolicyResponse, error) {
	var pResponse DescribeBackupPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
