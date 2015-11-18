package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDBInstanceForChannelRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ClientToken           string
	SystemDBCharset       string
	Engine                string
	EngineVersion         string
	DBInstanceClass       string
	DBInstanceStorage     int
	DBInstanceNetType     string
	DBInstanceDescription string
	SecurityIPList        string
	AccountName           string
	AccountPassword       string
	PayType               string
	OwnerAccount          string
}

func (r *CreateDBInstanceForChannelRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateDBInstanceForChannelRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateDBInstanceForChannelRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateDBInstanceForChannelRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateDBInstanceForChannelRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateDBInstanceForChannelRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateDBInstanceForChannelRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateDBInstanceForChannelRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateDBInstanceForChannelRequest) SetSystemDBCharset(value string) {
	r.SystemDBCharset = value
	r.QueryParams.Set("SystemDBCharset", value)
}
func (r *CreateDBInstanceForChannelRequest) GetSystemDBCharset() string {
	return r.SystemDBCharset
}
func (r *CreateDBInstanceForChannelRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *CreateDBInstanceForChannelRequest) GetEngine() string {
	return r.Engine
}
func (r *CreateDBInstanceForChannelRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *CreateDBInstanceForChannelRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *CreateDBInstanceForChannelRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *CreateDBInstanceForChannelRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *CreateDBInstanceForChannelRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *CreateDBInstanceForChannelRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *CreateDBInstanceForChannelRequest) SetDBInstanceNetType(value string) {
	r.DBInstanceNetType = value
	r.QueryParams.Set("DBInstanceNetType", value)
}
func (r *CreateDBInstanceForChannelRequest) GetDBInstanceNetType() string {
	return r.DBInstanceNetType
}
func (r *CreateDBInstanceForChannelRequest) SetDBInstanceDescription(value string) {
	r.DBInstanceDescription = value
	r.QueryParams.Set("DBInstanceDescription", value)
}
func (r *CreateDBInstanceForChannelRequest) GetDBInstanceDescription() string {
	return r.DBInstanceDescription
}
func (r *CreateDBInstanceForChannelRequest) SetSecurityIPList(value string) {
	r.SecurityIPList = value
	r.QueryParams.Set("SecurityIPList", value)
}
func (r *CreateDBInstanceForChannelRequest) GetSecurityIPList() string {
	return r.SecurityIPList
}
func (r *CreateDBInstanceForChannelRequest) SetAccountName(value string) {
	r.AccountName = value
	r.QueryParams.Set("AccountName", value)
}
func (r *CreateDBInstanceForChannelRequest) GetAccountName() string {
	return r.AccountName
}
func (r *CreateDBInstanceForChannelRequest) SetAccountPassword(value string) {
	r.AccountPassword = value
	r.QueryParams.Set("AccountPassword", value)
}
func (r *CreateDBInstanceForChannelRequest) GetAccountPassword() string {
	return r.AccountPassword
}
func (r *CreateDBInstanceForChannelRequest) SetPayType(value string) {
	r.PayType = value
	r.QueryParams.Set("PayType", value)
}
func (r *CreateDBInstanceForChannelRequest) GetPayType() string {
	return r.PayType
}
func (r *CreateDBInstanceForChannelRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateDBInstanceForChannelRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateDBInstanceForChannelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDBInstanceForChannel")
	r.SetProduct(Product)
}

type CreateDBInstanceForChannelResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	OrderId          string `xml:"OrderId" json:"OrderId"`
	ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
	Port             string `xml:"Port" json:"Port"`
}

func CreateDBInstanceForChannel(req *CreateDBInstanceForChannelRequest, accessId, accessSecret string) (*CreateDBInstanceForChannelResponse, error) {
	var pResponse CreateDBInstanceForChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
