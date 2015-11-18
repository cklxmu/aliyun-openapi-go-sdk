package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateAccountForInnerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	AccountPassword      string
	DBName               string
	AccountPrivilege     string
	AccountDescription   string
	OwnerAccount         string
}

func (r *CreateAccountForInnerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateAccountForInnerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateAccountForInnerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateAccountForInnerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateAccountForInnerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateAccountForInnerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateAccountForInnerRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateAccountForInnerRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateAccountForInnerRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *CreateAccountForInnerRequest) GetAccountName() string {
	return r.AccountName
}
func (r *CreateAccountForInnerRequest) SetAccountPassword(value string) {
	r.AccountPassword = value
	r.QueryParams.Set("AccountPassword", value)
}
func (r *CreateAccountForInnerRequest) GetAccountPassword() string {
	return r.AccountPassword
}
func (r *CreateAccountForInnerRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *CreateAccountForInnerRequest) GetDBName() string {
	return r.DBName
}
func (r *CreateAccountForInnerRequest) SetAccountPrivilege(value string) {
	r.AccountPrivilege = value
	r.QueryParams.Set("AccountPrivilege", value)
}
func (r *CreateAccountForInnerRequest) GetAccountPrivilege() string {
	return r.AccountPrivilege
}
func (r *CreateAccountForInnerRequest) SetAccountDescription(value string) {
	r.AccountDescription = value
	r.QueryParams.Set("AccountDescription", value)
}
func (r *CreateAccountForInnerRequest) GetAccountDescription() string {
	return r.AccountDescription
}
func (r *CreateAccountForInnerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateAccountForInnerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateAccountForInnerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateAccountForInner")
	r.SetProduct(Product)
}

type CreateAccountForInnerResponse struct {
}

func CreateAccountForInner(req *CreateAccountForInnerRequest, accessId, accessSecret string) (*CreateAccountForInnerResponse, error) {
	var pResponse CreateAccountForInnerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
