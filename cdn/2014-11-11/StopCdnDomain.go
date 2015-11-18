package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type StopCdnDomainRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
}

func (r *StopCdnDomainRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *StopCdnDomainRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *StopCdnDomainRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *StopCdnDomainRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *StopCdnDomainRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *StopCdnDomainRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *StopCdnDomainRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *StopCdnDomainRequest) GetDomainName() string {
	return r.DomainName
}

func (r *StopCdnDomainRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopCdnDomain")
	r.SetProduct(Product)
}

type StopCdnDomainResponse struct {
}

func StopCdnDomain(req *StopCdnDomainRequest, accessId, accessSecret string) (*StopCdnDomainResponse, error) {
	var pResponse StopCdnDomainResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
