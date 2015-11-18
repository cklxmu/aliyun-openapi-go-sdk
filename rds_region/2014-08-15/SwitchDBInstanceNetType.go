package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SwitchDBInstanceNetTypeRequest struct {
	RpcRequest
	OwnerId                int
	ResourceOwnerAccount   string
	ResourceOwnerId        int
	ClientToken            string
	DBInstanceId           string
	ConnectionStringPrefix string
	Port                   int
	OwnerAccount           string
}

func (r *SwitchDBInstanceNetTypeRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SwitchDBInstanceNetTypeRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SwitchDBInstanceNetTypeRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SwitchDBInstanceNetTypeRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SwitchDBInstanceNetTypeRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SwitchDBInstanceNetTypeRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SwitchDBInstanceNetTypeRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *SwitchDBInstanceNetTypeRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *SwitchDBInstanceNetTypeRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(value string) {
	r.ConnectionStringPrefix = value
	r.QueryParams.Set("ConnectionStringPrefix", value)
}
func (r *SwitchDBInstanceNetTypeRequest) GetConnectionStringPrefix() string {
	return r.ConnectionStringPrefix
}
func (r *SwitchDBInstanceNetTypeRequest) SetPort(value int) {
	r.Port = value
	r.QueryParams.Set("Port", strconv.Itoa(value))
}
func (r *SwitchDBInstanceNetTypeRequest) GetPort() int {
	return r.Port
}
func (r *SwitchDBInstanceNetTypeRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SwitchDBInstanceNetTypeRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SwitchDBInstanceNetTypeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SwitchDBInstanceNetType")
	r.SetProduct(Product)
}

type SwitchDBInstanceNetTypeResponse struct {
}

func SwitchDBInstanceNetType(req *SwitchDBInstanceNetTypeRequest, accessId, accessSecret string) (*SwitchDBInstanceNetTypeResponse, error) {
	var pResponse SwitchDBInstanceNetTypeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
