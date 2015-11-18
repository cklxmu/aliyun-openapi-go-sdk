package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateUploadPathForSQLServerRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	DBName               string
	OwnerAccount         string
}

func (r *CreateUploadPathForSQLServerRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateUploadPathForSQLServerRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateUploadPathForSQLServerRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateUploadPathForSQLServerRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateUploadPathForSQLServerRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateUploadPathForSQLServerRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateUploadPathForSQLServerRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateUploadPathForSQLServerRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateUploadPathForSQLServerRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *CreateUploadPathForSQLServerRequest) GetDBName() string {
	return r.DBName
}
func (r *CreateUploadPathForSQLServerRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateUploadPathForSQLServerRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateUploadPathForSQLServerRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateUploadPathForSQLServer")
	r.SetProduct(Product)
}

type CreateUploadPathForSQLServerResponse struct {
	InternetFtpServer string `xml:"InternetFtpServer" json:"InternetFtpServer"`
	InternetPort      int    `xml:"InternetPort" json:"InternetPort"`
	IntranetFtpserver string `xml:"IntranetFtpserver" json:"IntranetFtpserver"`
	Intranetport      int    `xml:"Intranetport" json:"Intranetport"`
	UserName          string `xml:"UserName" json:"UserName"`
	Password          string `xml:"Password" json:"Password"`
	FileName          string `xml:"FileName" json:"FileName"`
}

func CreateUploadPathForSQLServer(req *CreateUploadPathForSQLServerRequest, accessId, accessSecret string) (*CreateUploadPathForSQLServerResponse, error) {
	var pResponse CreateUploadPathForSQLServerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
