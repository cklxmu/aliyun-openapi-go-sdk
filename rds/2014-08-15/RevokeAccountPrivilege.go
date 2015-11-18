package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RevokeAccountPrivilegeRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	DBName               string
	OwnerAccount         string
}

func (r *RevokeAccountPrivilegeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *RevokeAccountPrivilegeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *RevokeAccountPrivilegeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *RevokeAccountPrivilegeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *RevokeAccountPrivilegeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *RevokeAccountPrivilegeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *RevokeAccountPrivilegeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *RevokeAccountPrivilegeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *RevokeAccountPrivilegeRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *RevokeAccountPrivilegeRequest) GetAccountName() string {
	return r.AccountName
}
func (r *RevokeAccountPrivilegeRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *RevokeAccountPrivilegeRequest) GetDBName() string {
	return r.DBName
}
func (r *RevokeAccountPrivilegeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *RevokeAccountPrivilegeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *RevokeAccountPrivilegeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RevokeAccountPrivilege")
	r.SetProduct(Product)
}

type RevokeAccountPrivilegeResponse struct {
}

func RevokeAccountPrivilege(req *RevokeAccountPrivilegeRequest, accessId, accessSecret string) (*RevokeAccountPrivilegeResponse, error) {
	var pResponse RevokeAccountPrivilegeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
