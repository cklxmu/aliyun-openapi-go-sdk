package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ImportDataFromDatabaseRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	DBInstanceId           string
	SourceDatabaseIp       string
	SourceDatabasePort     string
	SourceDatabaseUserName string
	SourceDatabasePassword string
	ImportDataType         string
	IsLockTable            string
	SourceDatabaseDBNames  string
	OwnerAccount           string
}

func (r *ImportDataFromDatabaseRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ImportDataFromDatabaseRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ImportDataFromDatabaseRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ImportDataFromDatabaseRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ImportDataFromDatabaseRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ImportDataFromDatabaseRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ImportDataFromDatabaseRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ImportDataFromDatabaseRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ImportDataFromDatabaseRequest) SetSourceDatabaseIp(value string) {
	r.SourceDatabaseIp = value
	r.QueryParams.Set("SourceDatabaseIp", value)
}
func (r *ImportDataFromDatabaseRequest) GetSourceDatabaseIp() string {
	return r.SourceDatabaseIp
}
func (r *ImportDataFromDatabaseRequest) SetSourceDatabasePort(value string) {
	r.SourceDatabasePort = value
	r.QueryParams.Set("SourceDatabasePort", value)
}
func (r *ImportDataFromDatabaseRequest) GetSourceDatabasePort() string {
	return r.SourceDatabasePort
}
func (r *ImportDataFromDatabaseRequest) SetSourceDatabaseUserName(value string) {
	r.SourceDatabaseUserName = value
	r.QueryParams.Set("SourceDatabaseUserName", value)
}
func (r *ImportDataFromDatabaseRequest) GetSourceDatabaseUserName() string {
	return r.SourceDatabaseUserName
}
func (r *ImportDataFromDatabaseRequest) SetSourceDatabasePassword(value string) {
	r.SourceDatabasePassword = value
	r.QueryParams.Set("SourceDatabasePassword", value)
}
func (r *ImportDataFromDatabaseRequest) GetSourceDatabasePassword() string {
	return r.SourceDatabasePassword
}
func (r *ImportDataFromDatabaseRequest) SetImportDataType(value string) {
	r.ImportDataType = value
	r.QueryParams.Set("ImportDataType", value)
}
func (r *ImportDataFromDatabaseRequest) GetImportDataType() string {
	return r.ImportDataType
}
func (r *ImportDataFromDatabaseRequest) SetIsLockTable(value string) {
	r.IsLockTable = value
	r.QueryParams.Set("IsLockTable", value)
}
func (r *ImportDataFromDatabaseRequest) GetIsLockTable() string {
	return r.IsLockTable
}
func (r *ImportDataFromDatabaseRequest) SetSourceDatabaseDBNames(value string) {
	r.SourceDatabaseDBNames = value
	r.QueryParams.Set("SourceDatabaseDBNames", value)
}
func (r *ImportDataFromDatabaseRequest) GetSourceDatabaseDBNames() string {
	return r.SourceDatabaseDBNames
}
func (r *ImportDataFromDatabaseRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ImportDataFromDatabaseRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ImportDataFromDatabaseRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ImportDataFromDatabase")
	r.SetProduct(Product)
}

type ImportDataFromDatabaseResponse struct {
	ImportId int `xml:"ImportId" json:"ImportId"`
}

func ImportDataFromDatabase(req *ImportDataFromDatabaseRequest, accessId, accessSecret string) (*ImportDataFromDatabaseResponse, error) {
	var pResponse ImportDataFromDatabaseResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
