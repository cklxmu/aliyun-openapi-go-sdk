package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteServerCertificateRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ServerCertificateId  string
	OwnerAccount         string
}

func (r *DeleteServerCertificateRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteServerCertificateRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteServerCertificateRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteServerCertificateRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteServerCertificateRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteServerCertificateRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteServerCertificateRequest) SetServerCertificateId(value string) {
	r.ServerCertificateId = value
	r.QueryParams.Set("ServerCertificateId", value)
}
func (r *DeleteServerCertificateRequest) GetServerCertificateId() string {
	return r.ServerCertificateId
}
func (r *DeleteServerCertificateRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DeleteServerCertificateRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DeleteServerCertificateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteServerCertificate")
	r.SetProduct(Product)
}

type DeleteServerCertificateResponse struct {
}

func DeleteServerCertificate(req *DeleteServerCertificateRequest, accessId, accessSecret string) (*DeleteServerCertificateResponse, error) {
	var pResponse DeleteServerCertificateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
