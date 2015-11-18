package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescibeImportsFromDatabaseRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	Engine               string
	ImportId             int
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescibeImportsFromDatabaseRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescibeImportsFromDatabaseRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescibeImportsFromDatabaseRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescibeImportsFromDatabaseRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescibeImportsFromDatabaseRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescibeImportsFromDatabaseRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescibeImportsFromDatabaseRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescibeImportsFromDatabaseRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetEngine() string {
	return r.Engine
}
func (r *DescibeImportsFromDatabaseRequest) SetImportId(value int) {
	r.ImportId = value
	r.QueryParams.Set("ImportId", strconv.Itoa(value))
}
func (r *DescibeImportsFromDatabaseRequest) GetImportId() int {
	return r.ImportId
}
func (r *DescibeImportsFromDatabaseRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescibeImportsFromDatabaseRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescibeImportsFromDatabaseRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescibeImportsFromDatabaseRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescibeImportsFromDatabaseRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescibeImportsFromDatabaseRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescibeImportsFromDatabaseRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescibeImportsFromDatabaseRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescibeImportsFromDatabaseRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescibeImportsFromDatabase")
	r.SetProduct(Product)
}

type DescibeImportsFromDatabaseResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		ImportResultFromDB []struct {
			ImportId                    int    `xml:"ImportId" json:"ImportId"`
			ImportDataType              string `xml:"ImportDataType" json:"ImportDataType"`
			ImportDataStatus            string `xml:"ImportDataStatus" json:"ImportDataStatus"`
			ImportDataStatusDescription string `xml:"ImportDataStatusDescription" json:"ImportDataStatusDescription"`
			IncrementalImportingTime    string `xml:"IncrementalImportingTime" json:"IncrementalImportingTime"`
		} `xml:"ImportResultFromDB" json:"ImportResultFromDB"`
	} `xml:"Items" json:"Items"`
}

func DescibeImportsFromDatabase(req *DescibeImportsFromDatabaseRequest, accessId, accessSecret string) (*DescibeImportsFromDatabaseResponse, error) {
	var pResponse DescibeImportsFromDatabaseResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
