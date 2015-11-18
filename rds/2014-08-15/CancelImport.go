package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CancelImportRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	ImportId             int
	OwnerAccount         string
}

func (r *CancelImportRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CancelImportRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CancelImportRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CancelImportRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CancelImportRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CancelImportRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CancelImportRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CancelImportRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CancelImportRequest) SetImportId(value int) {
	r.ImportId = value
	r.QueryParams.Set("ImportId", strconv.Itoa(value))
}
func (r *CancelImportRequest) GetImportId() int {
	return r.ImportId
}
func (r *CancelImportRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CancelImportRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CancelImportRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CancelImport")
	r.SetProduct(Product)
}

type CancelImportResponse struct {
}

func CancelImport(req *CancelImportRequest, accessId, accessSecret string) (*CancelImportResponse, error) {
	var pResponse CancelImportResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
