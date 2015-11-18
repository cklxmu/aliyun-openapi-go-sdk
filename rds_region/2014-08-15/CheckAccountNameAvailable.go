package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CheckAccountNameAvailableRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	AccountName          string
	OwnerAccount         string
}

func (r *CheckAccountNameAvailableRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CheckAccountNameAvailableRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CheckAccountNameAvailableRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CheckAccountNameAvailableRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CheckAccountNameAvailableRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CheckAccountNameAvailableRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CheckAccountNameAvailableRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CheckAccountNameAvailableRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CheckAccountNameAvailableRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CheckAccountNameAvailableRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CheckAccountNameAvailableRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *CheckAccountNameAvailableRequest) GetAccountName() string {
	return r.AccountName
}
func (r *CheckAccountNameAvailableRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CheckAccountNameAvailableRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CheckAccountNameAvailableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CheckAccountNameAvailable")
	r.SetProduct(Product)
}

type CheckAccountNameAvailableResponse struct {
}

func CheckAccountNameAvailable(req *CheckAccountNameAvailableRequest, accessId, accessSecret string) (*CheckAccountNameAvailableResponse, error) {
	var pResponse CheckAccountNameAvailableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
