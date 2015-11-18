package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyParameterRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	DBInstanceId         string
	Parameters           string
	Forcerestart         bool
	OwnerAccount         string
}

func (r *ModifyParameterRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyParameterRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyParameterRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyParameterRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyParameterRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyParameterRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyParameterRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *ModifyParameterRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *ModifyParameterRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *ModifyParameterRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *ModifyParameterRequest) SetParameters(value string) {
	r.Parameters = value
	r.QueryParams.Set("Parameters", value)
}
func (r *ModifyParameterRequest) GetParameters() string {
	return r.Parameters
}
func (r *ModifyParameterRequest) SetForcerestart(value bool) {
	r.Forcerestart = value
	r.QueryParams.Set("Forcerestart", strconv.FormatBool(value))
}
func (r *ModifyParameterRequest) GetForcerestart() bool {
	return r.Forcerestart
}
func (r *ModifyParameterRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *ModifyParameterRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *ModifyParameterRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyParameter")
	r.SetProduct(Product)
}

type ModifyParameterResponse struct {
}

func ModifyParameter(req *ModifyParameterRequest, accessId, accessSecret string) (*ModifyParameterResponse, error) {
	var pResponse ModifyParameterResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
