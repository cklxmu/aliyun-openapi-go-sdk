package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateReadOnlyDBInstanceRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ClientToken           string
	ZoneId                string
	DBInstanceId          string
	DBInstanceClass       string
	DBInstanceStorage     int
	EngineVersion         string
	PayType               string
	DBInstanceDescription string
	InstanceNetworkType   string
	VPCId                 string
	VSwitchId             string
	PrivateIpAddress      string
	OwnerAccount          string
}

func (r *CreateReadOnlyDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateReadOnlyDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateReadOnlyDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateReadOnlyDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateReadOnlyDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateReadOnlyDBInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreateReadOnlyDBInstanceRequest) SetZoneId(value string) {
	r.ZoneId = value
	r.QueryParams.Set("ZoneId", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetZoneId() string {
	return r.ZoneId
}
func (r *CreateReadOnlyDBInstanceRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateReadOnlyDBInstanceRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *CreateReadOnlyDBInstanceRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *CreateReadOnlyDBInstanceRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *CreateReadOnlyDBInstanceRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *CreateReadOnlyDBInstanceRequest) SetPayType(value string) {
	r.PayType = value
	r.QueryParams.Set("PayType", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetPayType() string {
	return r.PayType
}
func (r *CreateReadOnlyDBInstanceRequest) SetDBInstanceDescription(value string) {
	r.DBInstanceDescription = value
	r.QueryParams.Set("DBInstanceDescription", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetDBInstanceDescription() string {
	return r.DBInstanceDescription
}
func (r *CreateReadOnlyDBInstanceRequest) SetInstanceNetworkType(value string) {
	r.InstanceNetworkType = value
	r.QueryParams.Set("InstanceNetworkType", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetInstanceNetworkType() string {
	return r.InstanceNetworkType
}
func (r *CreateReadOnlyDBInstanceRequest) SetVPCId(value string) {
	r.VPCId = value
	r.QueryParams.Set("VPCId", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetVPCId() string {
	return r.VPCId
}
func (r *CreateReadOnlyDBInstanceRequest) SetVSwitchId(value string) {
	r.VSwitchId = value
	r.QueryParams.Set("VSwitchId", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetVSwitchId() string {
	return r.VSwitchId
}
func (r *CreateReadOnlyDBInstanceRequest) SetPrivateIpAddress(value string) {
	r.PrivateIpAddress = value
	r.QueryParams.Set("PrivateIpAddress", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetPrivateIpAddress() string {
	return r.PrivateIpAddress
}
func (r *CreateReadOnlyDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateReadOnlyDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateReadOnlyDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateReadOnlyDBInstance")
	r.SetProduct(Product)
}

type CreateReadOnlyDBInstanceResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	OrderId          string `xml:"OrderId" json:"OrderId"`
	ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
	Port             string `xml:"Port" json:"Port"`
}

func CreateReadOnlyDBInstance(req *CreateReadOnlyDBInstanceRequest, accessId, accessSecret string) (*CreateReadOnlyDBInstanceResponse, error) {
	var pResponse CreateReadOnlyDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
