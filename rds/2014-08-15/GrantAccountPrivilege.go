package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GrantAccountPrivilegeRequest struct {
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

func (r *GrantAccountPrivilegeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *GrantAccountPrivilegeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *GrantAccountPrivilegeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *GrantAccountPrivilegeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *GrantAccountPrivilegeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *GrantAccountPrivilegeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *GrantAccountPrivilegeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *GrantAccountPrivilegeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *GrantAccountPrivilegeRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *GrantAccountPrivilegeRequest) GetAccountName() string {
	return r.AccountName
}
func (r *GrantAccountPrivilegeRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *GrantAccountPrivilegeRequest) GetDBName() string {
	return r.DBName
}
func (r *GrantAccountPrivilegeRequest) SetAccountPrivilege(value string) {
	r.AccountPrivilege = value
	r.QueryParams.Set("AccountPrivilege", value)
}
func (r *GrantAccountPrivilegeRequest) GetAccountPrivilege() string {
	return r.AccountPrivilege
}
func (r *GrantAccountPrivilegeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *GrantAccountPrivilegeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *GrantAccountPrivilegeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GrantAccountPrivilege")
	r.SetProduct(Product)
}

type GrantAccountPrivilegeResponse struct {
}

func GrantAccountPrivilege(req *GrantAccountPrivilegeRequest, accessId, accessSecret string) (*GrantAccountPrivilegeResponse, error) {
	var pResponse GrantAccountPrivilegeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
