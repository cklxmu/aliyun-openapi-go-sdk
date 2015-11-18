package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDBInstanceforFirstPayRequest struct {
	RpcRequest
	uid               int
	bid               string
	uidLoginEmail     string
	bidLoginEmail     string
	Engine            string
	EngineVersion     string
	DBInstanceClass   string
	DBInstanceStorage int
	DBInstanceNetType string
	CharacterSetName  string
	DBInstanceRemarks string
	ClientToken       string
	OwnerAccount      string
}

func (r *CreateDBInstanceforFirstPayRequest) Setuid(value int) {
	r.uid = value
	r.QueryParams.Set("uid", strconv.Itoa(value))
}
func (r *CreateDBInstanceforFirstPayRequest) Getuid() int {
	return r.uid
}
func (r *CreateDBInstanceforFirstPayRequest) Setbid(value string) {
	r.bid = value
	r.QueryParams.Set("bid", value)
}
func (r *CreateDBInstanceforFirstPayRequest) Getbid() string {
	return r.bid
}
func (r *CreateDBInstanceforFirstPayRequest) SetuidLoginEmail(value string) {
	r.uidLoginEmail = value
	r.QueryParams.Set("uidLoginEmail", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetuidLoginEmail() string {
	return r.uidLoginEmail
}
func (r *CreateDBInstanceforFirstPayRequest) SetbidLoginEmail(value string) {
	r.bidLoginEmail = value
	r.QueryParams.Set("bidLoginEmail", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetbidLoginEmail() string {
	return r.bidLoginEmail
}
func (r *CreateDBInstanceforFirstPayRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetEngine() string {
	return r.Engine
}
func (r *CreateDBInstanceforFirstPayRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *CreateDBInstanceforFirstPayRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *CreateDBInstanceforFirstPayRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *CreateDBInstanceforFirstPayRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *CreateDBInstanceforFirstPayRequest) SetDBInstanceNetType(value string) {
	r.DBInstanceNetType = value
	r.QueryParams.Set("DBInstanceNetType", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetDBInstanceNetType() string {
	return r.DBInstanceNetType
}
func (r *CreateDBInstanceforFirstPayRequest) SetCharacterSetName(value string) {
	r.CharacterSetName = value
	r.QueryParams.Set("CharacterSetName", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetCharacterSetName() string {
	return r.CharacterSetName
}
func (r *CreateDBInstanceforFirstPayRequest) SetDBInstanceRemarks(value string) {
	r.DBInstanceRemarks = value
	r.QueryParams.Set("DBInstanceRemarks", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetDBInstanceRemarks() string {
	return r.DBInstanceRemarks
}
func (r *CreateDBInstanceforFirstPayRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateDBInstanceforFirstPayRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateDBInstanceforFirstPayRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateDBInstanceforFirstPayRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDBInstanceforFirstPay")
	r.SetProduct(Product)
}

type CreateDBInstanceforFirstPayResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
	Port             string `xml:"Port" json:"Port"`
}

func CreateDBInstanceforFirstPay(req *CreateDBInstanceforFirstPayRequest, accessId, accessSecret string) (*CreateDBInstanceforFirstPayResponse, error) {
	var pResponse CreateDBInstanceforFirstPayResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
