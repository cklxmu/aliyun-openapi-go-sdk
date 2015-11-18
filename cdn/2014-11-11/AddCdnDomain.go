package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AddCdnDomainRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
	SslFlag              string
	SourceType           string
	CdnType              string
	Sources              string
}

func (r *AddCdnDomainRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *AddCdnDomainRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *AddCdnDomainRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *AddCdnDomainRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *AddCdnDomainRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *AddCdnDomainRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *AddCdnDomainRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *AddCdnDomainRequest) GetDomainName() string {
	return r.DomainName
}
func (r *AddCdnDomainRequest) SetSslFlag(value string) {
	r.SslFlag = value
	r.QueryParams.Set("SslFlag", value)
}
func (r *AddCdnDomainRequest) GetSslFlag() string {
	return r.SslFlag
}
func (r *AddCdnDomainRequest) SetSourceType(value string) {
	r.SourceType = value
	r.QueryParams.Set("SourceType", value)
}
func (r *AddCdnDomainRequest) GetSourceType() string {
	return r.SourceType
}
func (r *AddCdnDomainRequest) SetCdnType(value string) {
	r.CdnType = value
	r.QueryParams.Set("CdnType", value)
}
func (r *AddCdnDomainRequest) GetCdnType() string {
	return r.CdnType
}
func (r *AddCdnDomainRequest) SetSources(value string) {
	r.Sources = value
	r.QueryParams.Set("Sources", value)
}
func (r *AddCdnDomainRequest) GetSources() string {
	return r.Sources
}

func (r *AddCdnDomainRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddCdnDomain")
	r.SetProduct(Product)
}

type AddCdnDomainResponse struct {
}

func AddCdnDomain(req *AddCdnDomainRequest, accessId, accessSecret string) (*AddCdnDomainResponse, error) {
	var pResponse AddCdnDomainResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
