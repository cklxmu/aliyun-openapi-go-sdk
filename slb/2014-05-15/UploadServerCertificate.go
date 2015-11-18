package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UploadServerCertificateRequest struct {
	RpcRequest
	OwnerId               int
	ResourceOwnerAccount  string
	ResourceOwnerId       int
	ServerCertificate     string
	PrivateKey            string
	ServerCertificateName string
	OwnerAccount          string
}

func (r *UploadServerCertificateRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *UploadServerCertificateRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *UploadServerCertificateRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *UploadServerCertificateRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *UploadServerCertificateRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *UploadServerCertificateRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *UploadServerCertificateRequest) SetServerCertificate(value string) {
	r.ServerCertificate = value
	r.QueryParams.Set("ServerCertificate", value)
}
func (r *UploadServerCertificateRequest) GetServerCertificate() string {
	return r.ServerCertificate
}
func (r *UploadServerCertificateRequest) SetPrivateKey(value string) {
	r.PrivateKey = value
	r.QueryParams.Set("PrivateKey", value)
}
func (r *UploadServerCertificateRequest) GetPrivateKey() string {
	return r.PrivateKey
}
func (r *UploadServerCertificateRequest) SetServerCertificateName(value string) {
	r.ServerCertificateName = value
	r.QueryParams.Set("ServerCertificateName", value)
}
func (r *UploadServerCertificateRequest) GetServerCertificateName() string {
	return r.ServerCertificateName
}
func (r *UploadServerCertificateRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *UploadServerCertificateRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *UploadServerCertificateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UploadServerCertificate")
	r.SetProduct(Product)
}

type UploadServerCertificateResponse struct {
	ServerCertificateId   string `xml:"ServerCertificateId" json:"ServerCertificateId"`
	Fingerprint           string `xml:"Fingerprint" json:"Fingerprint"`
	ServerCertificateName string `xml:"ServerCertificateName" json:"ServerCertificateName"`
	RegionId              string `xml:"RegionId" json:"RegionId"`
	RegionIdAlias         string `xml:"RegionIdAlias" json:"RegionIdAlias"`
}

func UploadServerCertificate(req *UploadServerCertificateRequest, accessId, accessSecret string) (*UploadServerCertificateResponse, error) {
	var pResponse UploadServerCertificateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
