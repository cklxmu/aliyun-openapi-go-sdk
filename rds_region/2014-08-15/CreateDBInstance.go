package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDBInstanceRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ClientToken           string
	Engine                string
	EngineVersion         string
	DBInstanceClass       string
	DBInstanceStorage     int
	SystemDBCharset       string
	DBInstanceNetType     string
	DBInstanceDescription string
	SecurityIPList        string
	PayType               string
	ZoneId                string
	InstanceNetworkType   string
	ConnectionMode        string
	VPCId                 string
	VSwitchId             string
	PrivateIpAddress      string
	OwnerAccount          string
}

func (r *CreateDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateDBInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateDBInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateDBInstanceRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *CreateDBInstanceRequest) GetEngine() string {
	return r.Engine
}
func (r *CreateDBInstanceRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *CreateDBInstanceRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *CreateDBInstanceRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *CreateDBInstanceRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *CreateDBInstanceRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *CreateDBInstanceRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *CreateDBInstanceRequest) SetSystemDBCharset(value string) {
	r.SystemDBCharset = value
	r.QueryParams.Set("SystemDBCharset", value)
}
func (r *CreateDBInstanceRequest) GetSystemDBCharset() string {
	return r.SystemDBCharset
}
func (r *CreateDBInstanceRequest) SetDBInstanceNetType(value string) {
	r.DBInstanceNetType = value
	r.QueryParams.Set("DBInstanceNetType", value)
}
func (r *CreateDBInstanceRequest) GetDBInstanceNetType() string {
	return r.DBInstanceNetType
}
func (r *CreateDBInstanceRequest) SetDBInstanceDescription(value string) {
	r.DBInstanceDescription = value
	r.QueryParams.Set("DBInstanceDescription", value)
}
func (r *CreateDBInstanceRequest) GetDBInstanceDescription() string {
	return r.DBInstanceDescription
}
func (r *CreateDBInstanceRequest) SetSecurityIPList(value string) {
	r.SecurityIPList = value
	r.QueryParams.Set("SecurityIPList", value)
}
func (r *CreateDBInstanceRequest) GetSecurityIPList() string {
	return r.SecurityIPList
}
func (r *CreateDBInstanceRequest) SetPayType(value string) {
	r.PayType = value
	r.QueryParams.Set("PayType", value)
}
func (r *CreateDBInstanceRequest) GetPayType() string {
	return r.PayType
}
func (r *CreateDBInstanceRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateDBInstanceRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateDBInstanceRequest) SetInstanceNetworkType(value string) {
	r.InstanceNetworkType = value
	r.QueryParams.Set("InstanceNetworkType", value)
}
func (r *CreateDBInstanceRequest) GetInstanceNetworkType() string {
	return r.InstanceNetworkType
}
func (r *CreateDBInstanceRequest) SetConnectionMode(value string) {
	r.ConnectionMode = value
	r.QueryParams.Set("ConnectionMode", value)
}
func (r *CreateDBInstanceRequest) GetConnectionMode() string {
	return r.ConnectionMode
}
func (r *CreateDBInstanceRequest) SetVPCId(value string) {
	r.VPCId = value
	r.QueryParams.Set("VPCId", value)
}
func (r *CreateDBInstanceRequest) GetVPCId() string {
	return r.VPCId
}
func (r *CreateDBInstanceRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *CreateDBInstanceRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *CreateDBInstanceRequest) SetPrivateIpAddress(value string) {
	r.PrivateIpAddress = value
	r.QueryParams.Set("PrivateIpAddress", value)
}
func (r *CreateDBInstanceRequest) GetPrivateIpAddress() string {
	return r.PrivateIpAddress
}
func (r *CreateDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDBInstance")
	r.SetProduct(Product)
}

type CreateDBInstanceResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	OrderId          string `xml:"OrderId" json:"OrderId"`
	ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
	Port             string `xml:"Port" json:"Port"`
}

func CreateDBInstance(req *CreateDBInstanceRequest, accessId, accessSecret string) (*CreateDBInstanceResponse, error) {
	var pResponse CreateDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
