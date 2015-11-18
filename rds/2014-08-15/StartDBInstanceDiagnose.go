package rds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StartDBInstanceDiagnoseRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	proxyId              string
	DBInstanceId         string
	OwnerAccount         string
}

func (r *StartDBInstanceDiagnoseRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StartDBInstanceDiagnoseRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StartDBInstanceDiagnoseRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StartDBInstanceDiagnoseRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StartDBInstanceDiagnoseRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StartDBInstanceDiagnoseRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StartDBInstanceDiagnoseRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *StartDBInstanceDiagnoseRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *StartDBInstanceDiagnoseRequest) SetproxyId(value string) {
	r.proxyId = value
	r.QueryParams.Set("proxyId", value)
}
func (r *StartDBInstanceDiagnoseRequest) GetproxyId() string {
	return r.proxyId
}
func (r *StartDBInstanceDiagnoseRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *StartDBInstanceDiagnoseRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *StartDBInstanceDiagnoseRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *StartDBInstanceDiagnoseRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *StartDBInstanceDiagnoseRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StartDBInstanceDiagnose")
	r.SetProduct(Product)
}

type StartDBInstanceDiagnoseResponse struct {
	DBInstanceName string `xml:"DBInstanceName" json:"DBInstanceName"`
	DBInstanceId   string `xml:"DBInstanceId" json:"DBInstanceId"`
}

func StartDBInstanceDiagnose(req *StartDBInstanceDiagnoseRequest, accessId, accessSecret string) (*StartDBInstanceDiagnoseResponse, error) {
	var pResponse StartDBInstanceDiagnoseResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
