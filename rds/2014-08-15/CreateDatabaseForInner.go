package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDatabaseForInnerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	DBName               string
	CharacterSetName     string
	DBDescription        string
	AccountName          string
	AccountPrivilege     string
	AccountPassword      string
	OwnerAccount         string
}

func (r *CreateDatabaseForInnerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateDatabaseForInnerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateDatabaseForInnerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateDatabaseForInnerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateDatabaseForInnerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateDatabaseForInnerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateDatabaseForInnerRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateDatabaseForInnerRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateDatabaseForInnerRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *CreateDatabaseForInnerRequest) GetDBName() string {
	return r.DBName
}
func (r *CreateDatabaseForInnerRequest) SetCharacterSetName(value string) {
	r.CharacterSetName = value
	r.QueryParams.Set("CharacterSetName", value)
}
func (r *CreateDatabaseForInnerRequest) GetCharacterSetName() string {
	return r.CharacterSetName
}
func (r *CreateDatabaseForInnerRequest) SetDBDescription(value string) {
	r.DBDescription = value
	r.QueryParams.Set("DBDescription", value)
}
func (r *CreateDatabaseForInnerRequest) GetDBDescription() string {
	return r.DBDescription
}
func (r *CreateDatabaseForInnerRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *CreateDatabaseForInnerRequest) GetAccountName() string {
	return r.AccountName
}
func (r *CreateDatabaseForInnerRequest) SetAccountPrivilege(value string) {
	r.AccountPrivilege = value
	r.QueryParams.Set("AccountPrivilege", value)
}
func (r *CreateDatabaseForInnerRequest) GetAccountPrivilege() string {
	return r.AccountPrivilege
}
func (r *CreateDatabaseForInnerRequest) SetAccountPassword(value string) {
	r.AccountPassword = value
	r.QueryParams.Set("AccountPassword", value)
}
func (r *CreateDatabaseForInnerRequest) GetAccountPassword() string {
	return r.AccountPassword
}
func (r *CreateDatabaseForInnerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateDatabaseForInnerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateDatabaseForInnerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDatabaseForInner")
	r.SetProduct(Product)
}

type CreateDatabaseForInnerResponse struct {
}

func CreateDatabaseForInner(req *CreateDatabaseForInnerRequest, accessId, accessSecret string) (*CreateDatabaseForInnerResponse, error) {
	var pResponse CreateDatabaseForInnerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
