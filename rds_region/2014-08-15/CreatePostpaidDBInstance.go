package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreatePostpaidDBInstanceRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ClientToken           string
	Engine                string
	EngineVersion         string
	DBInstanceClass       string
	DBInstanceStorage     int
	DBInstanceNetType     string
	DBInstanceDescription string
	SecurityIPList        string
	OwnerAccount          string
}

func (r *CreatePostpaidDBInstanceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreatePostpaidDBInstanceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreatePostpaidDBInstanceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreatePostpaidDBInstanceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreatePostpaidDBInstanceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreatePostpaidDBInstanceRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreatePostpaidDBInstanceRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetEngine() string {
	return r.Engine
}
func (r *CreatePostpaidDBInstanceRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *CreatePostpaidDBInstanceRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *CreatePostpaidDBInstanceRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *CreatePostpaidDBInstanceRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *CreatePostpaidDBInstanceRequest) SetDBInstanceNetType(value string) {
	r.DBInstanceNetType = value
	r.QueryParams.Set("DBInstanceNetType", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetDBInstanceNetType() string {
	return r.DBInstanceNetType
}
func (r *CreatePostpaidDBInstanceRequest) SetDBInstanceDescription(value string) {
	r.DBInstanceDescription = value
	r.QueryParams.Set("DBInstanceDescription", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetDBInstanceDescription() string {
	return r.DBInstanceDescription
}
func (r *CreatePostpaidDBInstanceRequest) SetSecurityIPList(value string) {
	r.SecurityIPList = value
	r.QueryParams.Set("SecurityIPList", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetSecurityIPList() string {
	return r.SecurityIPList
}
func (r *CreatePostpaidDBInstanceRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreatePostpaidDBInstanceRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreatePostpaidDBInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreatePostpaidDBInstance")
	r.SetProduct(Product)
}

type CreatePostpaidDBInstanceResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	OrderId          string `xml:"OrderId" json:"OrderId"`
	ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
	Port             string `xml:"Port" json:"Port"`
}

func CreatePostpaidDBInstance(req *CreatePostpaidDBInstanceRequest, accessId, accessSecret string) (*CreatePostpaidDBInstanceResponse, error) {
	var pResponse CreatePostpaidDBInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
