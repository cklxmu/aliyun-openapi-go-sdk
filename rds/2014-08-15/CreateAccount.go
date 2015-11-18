package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateAccountRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	AccountName          string
	AccountPassword      string
	AccountDescription   string
	OwnerAccount         string
}

func (r *CreateAccountRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateAccountRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateAccountRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateAccountRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateAccountRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateAccountRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateAccountRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateAccountRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateAccountRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *CreateAccountRequest) GetAccountName() string {
	return r.AccountName
}
func (r *CreateAccountRequest) SetAccountPassword(value string) {
	r.AccountPassword = value
	r.QueryParams.Set("AccountPassword", value)
}
func (r *CreateAccountRequest) GetAccountPassword() string {
	return r.AccountPassword
}
func (r *CreateAccountRequest) SetAccountDescription(value string) {
	r.AccountDescription = value
	r.QueryParams.Set("AccountDescription", value)
}
func (r *CreateAccountRequest) GetAccountDescription() string {
	return r.AccountDescription
}
func (r *CreateAccountRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateAccountRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateAccount")
	r.SetProduct(Product)
}

type CreateAccountResponse struct {
}

func CreateAccount(req *CreateAccountRequest, accessId, accessSecret string) (*CreateAccountResponse, error) {
	var pResponse CreateAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
