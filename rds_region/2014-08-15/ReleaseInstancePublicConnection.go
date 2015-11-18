package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReleaseInstancePublicConnectionRequest struct {
	RpcRequest
	OwnerId                 int
	ResourceOwnerAccount    string
	ResourceOwnerId         int
	DBInstanceId            string
	CurrentConnectionString string
	OwnerAccount            string
}

func (r *ReleaseInstancePublicConnectionRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ReleaseInstancePublicConnectionRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ReleaseInstancePublicConnectionRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ReleaseInstancePublicConnectionRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ReleaseInstancePublicConnectionRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ReleaseInstancePublicConnectionRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ReleaseInstancePublicConnectionRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(value string) {
	r.CurrentConnectionString = value
	r.QueryParams.Set("CurrentConnectionString", value)
}
func (r *ReleaseInstancePublicConnectionRequest) GetCurrentConnectionString() string {
	return r.CurrentConnectionString
}
func (r *ReleaseInstancePublicConnectionRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ReleaseInstancePublicConnectionRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ReleaseInstancePublicConnectionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReleaseInstancePublicConnection")
	r.SetProduct(Product)
}

type ReleaseInstancePublicConnectionResponse struct {
}

func ReleaseInstancePublicConnection(req *ReleaseInstancePublicConnectionRequest, accessId, accessSecret string) (*ReleaseInstancePublicConnectionResponse, error) {
	var pResponse ReleaseInstancePublicConnectionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
