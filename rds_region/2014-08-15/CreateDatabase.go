package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateDatabaseRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	DBName               string
	CharacterSetName     string
	DBDescription        string
	OwnerAccount         string
}

func (r *CreateDatabaseRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *CreateDatabaseRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *CreateDatabaseRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *CreateDatabaseRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *CreateDatabaseRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *CreateDatabaseRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *CreateDatabaseRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *CreateDatabaseRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *CreateDatabaseRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *CreateDatabaseRequest) GetDBName() string {
	return r.DBName
}
func (r *CreateDatabaseRequest) SetCharacterSetName(value string) {
	r.CharacterSetName = value
	r.QueryParams.Set("CharacterSetName", value)
}
func (r *CreateDatabaseRequest) GetCharacterSetName() string {
	return r.CharacterSetName
}
func (r *CreateDatabaseRequest) SetDBDescription(value string) {
	r.DBDescription = value
	r.QueryParams.Set("DBDescription", value)
}
func (r *CreateDatabaseRequest) GetDBDescription() string {
	return r.DBDescription
}
func (r *CreateDatabaseRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *CreateDatabaseRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *CreateDatabaseRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateDatabase")
	r.SetProduct(Product)
}

type CreateDatabaseResponse struct {
}

func CreateDatabase(req *CreateDatabaseRequest, accessId, accessSecret string) (*CreateDatabaseResponse, error) {
	var pResponse CreateDatabaseResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
