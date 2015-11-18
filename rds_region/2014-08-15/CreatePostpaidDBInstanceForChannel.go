package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreatePostpaidDBInstanceForChannelRequest struct {
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

func (r *CreatePostpaidDBInstanceForChannelRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetEngine() string {
	return r.Engine
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetDBInstanceClass(value string) {
	r.DBInstanceClass = value
	r.QueryParams.Set("DBInstanceClass", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetDBInstanceClass() string {
	return r.DBInstanceClass
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetDBInstanceStorage(value int) {
	r.DBInstanceStorage = value
	r.QueryParams.Set("DBInstanceStorage", strconv.Itoa(value))
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetDBInstanceStorage() int {
	return r.DBInstanceStorage
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetDBInstanceNetType(value string) {
	r.DBInstanceNetType = value
	r.QueryParams.Set("DBInstanceNetType", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetDBInstanceNetType() string {
	return r.DBInstanceNetType
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetDBInstanceDescription(value string) {
	r.DBInstanceDescription = value
	r.QueryParams.Set("DBInstanceDescription", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetDBInstanceDescription() string {
	return r.DBInstanceDescription
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetSecurityIPList(value string) {
	r.SecurityIPList = value
	r.QueryParams.Set("SecurityIPList", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetSecurityIPList() string {
	return r.SecurityIPList
}
func (r *CreatePostpaidDBInstanceForChannelRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreatePostpaidDBInstanceForChannelRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreatePostpaidDBInstanceForChannelRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreatePostpaidDBInstanceForChannel")
	r.SetProduct(Product)
}

type CreatePostpaidDBInstanceForChannelResponse struct {
	DBInstanceId     string `xml:"DBInstanceId" json:"DBInstanceId"`
	OrderId          string `xml:"OrderId" json:"OrderId"`
	ConnectionString string `xml:"ConnectionString" json:"ConnectionString"`
	Port             string `xml:"Port" json:"Port"`
}

func CreatePostpaidDBInstanceForChannel(req *CreatePostpaidDBInstanceForChannelRequest, accessId, accessSecret string) (*CreatePostpaidDBInstanceForChannelResponse, error) {
	var pResponse CreatePostpaidDBInstanceForChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
