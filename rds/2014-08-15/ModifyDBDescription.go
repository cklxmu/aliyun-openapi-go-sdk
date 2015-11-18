package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDBDescriptionRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	DBName               string
	DBDescription        string
	OwnerAccount         string
}

func (r *ModifyDBDescriptionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyDBDescriptionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyDBDescriptionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyDBDescriptionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyDBDescriptionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyDBDescriptionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyDBDescriptionRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyDBDescriptionRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyDBDescriptionRequest) SetDBName(value string) {
	r.DBName = value
	r.QueryParams.Set("DBName", value)
}
func (r *ModifyDBDescriptionRequest) GetDBName() string {
	return r.DBName
}
func (r *ModifyDBDescriptionRequest) SetDBDescription(value string) {
	r.DBDescription = value
	r.QueryParams.Set("DBDescription", value)
}
func (r *ModifyDBDescriptionRequest) GetDBDescription() string {
	return r.DBDescription
}
func (r *ModifyDBDescriptionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyDBDescriptionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyDBDescriptionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDBDescription")
	r.SetProduct(Product)
}

type ModifyDBDescriptionResponse struct {
}

func ModifyDBDescription(req *ModifyDBDescriptionRequest, accessId, accessSecret string) (*ModifyDBDescriptionResponse, error) {
	var pResponse ModifyDBDescriptionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
