package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ExtractBackupFromOASRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	BackupId             string
	OwnerAccount         string
}

func (r *ExtractBackupFromOASRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ExtractBackupFromOASRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ExtractBackupFromOASRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ExtractBackupFromOASRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ExtractBackupFromOASRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ExtractBackupFromOASRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ExtractBackupFromOASRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ExtractBackupFromOASRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ExtractBackupFromOASRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ExtractBackupFromOASRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ExtractBackupFromOASRequest) SetBackupId(value string) {
	r.BackupId = value
	r.QueryParams.Set("BackupId", value)
}
func (r *ExtractBackupFromOASRequest) GetBackupId() string {
	return r.BackupId
}
func (r *ExtractBackupFromOASRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ExtractBackupFromOASRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ExtractBackupFromOASRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ExtractBackupFromOAS")
	r.SetProduct(Product)
}

type ExtractBackupFromOASResponse struct {
	DataExtractionLastTime string `xml:"DataExtractionLastTime" json:"DataExtractionLastTime"`
}

func ExtractBackupFromOAS(req *ExtractBackupFromOASRequest, accessId, accessSecret string) (*ExtractBackupFromOASResponse, error) {
	var pResponse ExtractBackupFromOASResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
