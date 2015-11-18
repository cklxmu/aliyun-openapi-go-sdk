package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetServerCertificateNameRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ServerCertificateId   string
	ServerCertificateName string
	OwnerAccount          string
}

func (r *SetServerCertificateNameRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *SetServerCertificateNameRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *SetServerCertificateNameRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *SetServerCertificateNameRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *SetServerCertificateNameRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *SetServerCertificateNameRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *SetServerCertificateNameRequest) SetServerCertificateId(value string) {
	r.ServerCertificateId = value
	r.QueryParams.Set("ServerCertificateId", value)
}
func (r *SetServerCertificateNameRequest) GetServerCertificateId() string {
	return r.ServerCertificateId
}
func (r *SetServerCertificateNameRequest) SetServerCertificateName(value string) {
	r.ServerCertificateName = value
	r.QueryParams.Set("ServerCertificateName", value)
}
func (r *SetServerCertificateNameRequest) GetServerCertificateName() string {
	return r.ServerCertificateName
}
func (r *SetServerCertificateNameRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *SetServerCertificateNameRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *SetServerCertificateNameRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetServerCertificateName")
	r.SetProduct(Product)
}

type SetServerCertificateNameResponse struct {
}

func SetServerCertificateName(req *SetServerCertificateNameRequest, accessId, accessSecret string) (*SetServerCertificateNameResponse, error) {
	var pResponse SetServerCertificateNameResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
