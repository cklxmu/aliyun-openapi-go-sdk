package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeServerCertificateRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ServerCertificateId  string
	OwnerAccount         string
}

func (r *DescribeServerCertificateRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeServerCertificateRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeServerCertificateRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeServerCertificateRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeServerCertificateRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeServerCertificateRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeServerCertificateRequest) SetServerCertificateId(value string) {
	r.ServerCertificateId = value
	r.QueryParams.Set("ServerCertificateId", value)
}
func (r *DescribeServerCertificateRequest) GetServerCertificateId() string {
	return r.ServerCertificateId
}
func (r *DescribeServerCertificateRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeServerCertificateRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeServerCertificateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeServerCertificate")
	r.SetProduct(Product)
}

type DescribeServerCertificateResponse struct {
	ServerCertificateId   string `xml:"ServerCertificateId" json:"ServerCertificateId"`
	Fingerprint           string `xml:"Fingerprint" json:"Fingerprint"`
	RegionId              string `xml:"RegionId" json:"RegionId"`
	RegionIdAlias         string `xml:"RegionIdAlias" json:"RegionIdAlias"`
	ServerCertificateName string `xml:"ServerCertificateName" json:"ServerCertificateName"`
}

func DescribeServerCertificate(req *DescribeServerCertificateRequest, accessId, accessSecret string) (*DescribeServerCertificateResponse, error) {
	var pResponse DescribeServerCertificateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
