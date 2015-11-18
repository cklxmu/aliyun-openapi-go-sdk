package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PreCheckBeforeImportDataRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	ClientToken            string
	DBInstanceId           string
	SourceDatabaseIp       string
	SourceDatabasePort     string
	SourceDatabaseUserName string
	SourceDatabasePassword string
	ImportDataType         string
	SourceDatabaseDBNames  string
	OwnerAccount           string
}

func (r *PreCheckBeforeImportDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *PreCheckBeforeImportDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *PreCheckBeforeImportDataRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *PreCheckBeforeImportDataRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *PreCheckBeforeImportDataRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *PreCheckBeforeImportDataRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *PreCheckBeforeImportDataRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *PreCheckBeforeImportDataRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *PreCheckBeforeImportDataRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *PreCheckBeforeImportDataRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *PreCheckBeforeImportDataRequest) SetSourceDatabaseIp(value string) {
	r.SourceDatabaseIp = value
	r.QueryParams.Set("SourceDatabaseIp", value)
}
func (r *PreCheckBeforeImportDataRequest) GetSourceDatabaseIp() string {
	return r.SourceDatabaseIp
}
func (r *PreCheckBeforeImportDataRequest) SetSourceDatabasePort(value string) {
	r.SourceDatabasePort = value
	r.QueryParams.Set("SourceDatabasePort", value)
}
func (r *PreCheckBeforeImportDataRequest) GetSourceDatabasePort() string {
	return r.SourceDatabasePort
}
func (r *PreCheckBeforeImportDataRequest) SetSourceDatabaseUserName(value string) {
	r.SourceDatabaseUserName = value
	r.QueryParams.Set("SourceDatabaseUserName", value)
}
func (r *PreCheckBeforeImportDataRequest) GetSourceDatabaseUserName() string {
	return r.SourceDatabaseUserName
}
func (r *PreCheckBeforeImportDataRequest) SetSourceDatabasePassword(value string) {
	r.SourceDatabasePassword = value
	r.QueryParams.Set("SourceDatabasePassword", value)
}
func (r *PreCheckBeforeImportDataRequest) GetSourceDatabasePassword() string {
	return r.SourceDatabasePassword
}
func (r *PreCheckBeforeImportDataRequest) SetImportDataType(value string) {
	r.ImportDataType = value
	r.QueryParams.Set("ImportDataType", value)
}
func (r *PreCheckBeforeImportDataRequest) GetImportDataType() string {
	return r.ImportDataType
}
func (r *PreCheckBeforeImportDataRequest) SetSourceDatabaseDBNames(value string) {
	r.SourceDatabaseDBNames = value
	r.QueryParams.Set("SourceDatabaseDBNames", value)
}
func (r *PreCheckBeforeImportDataRequest) GetSourceDatabaseDBNames() string {
	return r.SourceDatabaseDBNames
}
func (r *PreCheckBeforeImportDataRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *PreCheckBeforeImportDataRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *PreCheckBeforeImportDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PreCheckBeforeImportData")
	r.SetProduct(Product)
}

type PreCheckBeforeImportDataResponse struct {
	PreCheckId string `xml:"PreCheckId" json:"PreCheckId"`
}

func PreCheckBeforeImportData(req *PreCheckBeforeImportDataRequest, accessId, accessSecret string) (*PreCheckBeforeImportDataResponse, error) {
	var pResponse PreCheckBeforeImportDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
