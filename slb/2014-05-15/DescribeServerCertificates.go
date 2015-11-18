package slb

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeServerCertificatesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ServerCertificateId  string
	OwnerAccount         string
}

func (r *DescribeServerCertificatesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeServerCertificatesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeServerCertificatesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeServerCertificatesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeServerCertificatesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeServerCertificatesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeServerCertificatesRequest) SetServerCertificateId(value string) {
	r.ServerCertificateId = value
	r.QueryParams.Set("ServerCertificateId", value)
}
func (r *DescribeServerCertificatesRequest) GetServerCertificateId() string {
	return r.ServerCertificateId
}
func (r *DescribeServerCertificatesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeServerCertificatesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeServerCertificatesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeServerCertificates")
	r.SetProduct(Product)
}

type DescribeServerCertificatesResponse struct {
	ServerCertificates struct {
		ServerCertificate []struct {
			ServerCertificateId   string `xml:"ServerCertificateId" json:"ServerCertificateId"`
			Fingerprint           string `xml:"Fingerprint" json:"Fingerprint"`
			ServerCertificateName string `xml:"ServerCertificateName" json:"ServerCertificateName"`
			RegionId              string `xml:"RegionId" json:"RegionId"`
			RegionIdAlias         string `xml:"RegionIdAlias" json:"RegionIdAlias"`
		} `xml:"ServerCertificate" json:"ServerCertificate"`
	} `xml:"ServerCertificates" json:"ServerCertificates"`
}

func DescribeServerCertificates(req *DescribeServerCertificatesRequest, accessId, accessSecret string) (*DescribeServerCertificatesResponse, error) {
	var pResponse DescribeServerCertificatesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
