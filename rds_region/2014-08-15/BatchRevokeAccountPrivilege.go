package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BatchRevokeAccountPrivilegeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	DBName               string
	OwnerAccount         string
}

func (r *BatchRevokeAccountPrivilegeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *BatchRevokeAccountPrivilegeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *BatchRevokeAccountPrivilegeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *BatchRevokeAccountPrivilegeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *BatchRevokeAccountPrivilegeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *BatchRevokeAccountPrivilegeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *BatchRevokeAccountPrivilegeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *BatchRevokeAccountPrivilegeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *BatchRevokeAccountPrivilegeRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *BatchRevokeAccountPrivilegeRequest) GetAccountName() string {
	return r.AccountName
}
func (r *BatchRevokeAccountPrivilegeRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *BatchRevokeAccountPrivilegeRequest) GetDBName() string {
	return r.DBName
}
func (r *BatchRevokeAccountPrivilegeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *BatchRevokeAccountPrivilegeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *BatchRevokeAccountPrivilegeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchRevokeAccountPrivilege")
	r.SetProduct(Product)
}

type BatchRevokeAccountPrivilegeResponse struct {
}

func BatchRevokeAccountPrivilege(req *BatchRevokeAccountPrivilegeRequest, accessId, accessSecret string) (*BatchRevokeAccountPrivilegeResponse, error) {
	var pResponse BatchRevokeAccountPrivilegeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
