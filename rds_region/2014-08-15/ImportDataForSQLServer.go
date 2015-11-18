package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ImportDataForSQLServerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	FileName             string
	OwnerAccount         string
}

func (r *ImportDataForSQLServerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ImportDataForSQLServerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ImportDataForSQLServerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ImportDataForSQLServerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ImportDataForSQLServerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ImportDataForSQLServerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ImportDataForSQLServerRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ImportDataForSQLServerRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ImportDataForSQLServerRequest) SetFileName(value string) {
	r.FileName = value
	r.QueryParams.Set("FileName", value)
}
func (r *ImportDataForSQLServerRequest) GetFileName() string {
	return r.FileName
}
func (r *ImportDataForSQLServerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ImportDataForSQLServerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ImportDataForSQLServerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ImportDataForSQLServer")
	r.SetProduct(Product)
}

type ImportDataForSQLServerResponse struct {
	ImportID int `xml:"ImportID" json:"ImportID"`
}

func ImportDataForSQLServer(req *ImportDataForSQLServerRequest, accessId, accessSecret string) (*ImportDataForSQLServerResponse, error) {
	var pResponse ImportDataForSQLServerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
