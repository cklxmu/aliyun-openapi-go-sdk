package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BatchGrantAccountPrivilegeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	DBName               string
	AccountPrivilege     string
	OwnerAccount         string
}

func (r *BatchGrantAccountPrivilegeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *BatchGrantAccountPrivilegeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *BatchGrantAccountPrivilegeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *BatchGrantAccountPrivilegeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *BatchGrantAccountPrivilegeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *BatchGrantAccountPrivilegeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *BatchGrantAccountPrivilegeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *BatchGrantAccountPrivilegeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *BatchGrantAccountPrivilegeRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *BatchGrantAccountPrivilegeRequest) GetAccountName() string {
	return r.AccountName
}
func (r *BatchGrantAccountPrivilegeRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *BatchGrantAccountPrivilegeRequest) GetDBName() string {
	return r.DBName
}
func (r *BatchGrantAccountPrivilegeRequest) SetAccountPrivilege(value string) {
	r.AccountPrivilege = value
	r.QueryParams.Set("AccountPrivilege", value)
}
func (r *BatchGrantAccountPrivilegeRequest) GetAccountPrivilege() string {
	return r.AccountPrivilege
}
func (r *BatchGrantAccountPrivilegeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *BatchGrantAccountPrivilegeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *BatchGrantAccountPrivilegeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchGrantAccountPrivilege")
	r.SetProduct(Product)
}

type BatchGrantAccountPrivilegeResponse struct {
}

func BatchGrantAccountPrivilege(req *BatchGrantAccountPrivilegeRequest, accessId, accessSecret string) (*BatchGrantAccountPrivilegeResponse, error) {
	var pResponse BatchGrantAccountPrivilegeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
