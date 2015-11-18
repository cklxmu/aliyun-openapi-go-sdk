package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StartCdnDomainRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
}

func (r *StartCdnDomainRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StartCdnDomainRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StartCdnDomainRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StartCdnDomainRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StartCdnDomainRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StartCdnDomainRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StartCdnDomainRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *StartCdnDomainRequest) GetDomainName() string {
	return r.DomainName
}

func (r *StartCdnDomainRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StartCdnDomain")
	r.SetProduct(Product)
}

type StartCdnDomainResponse struct {
}

func StartCdnDomain(req *StartCdnDomainRequest, accessId, accessSecret string) (*StartCdnDomainResponse, error) {
	var pResponse StartCdnDomainResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
