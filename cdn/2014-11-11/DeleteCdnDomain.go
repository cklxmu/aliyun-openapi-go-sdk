package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteCdnDomainRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
}

func (r *DeleteCdnDomainRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DeleteCdnDomainRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DeleteCdnDomainRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DeleteCdnDomainRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DeleteCdnDomainRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DeleteCdnDomainRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DeleteCdnDomainRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *DeleteCdnDomainRequest) GetDomainName() string {
	return r.DomainName
}

func (r *DeleteCdnDomainRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteCdnDomain")
	r.SetProduct(Product)
}

type DeleteCdnDomainResponse struct {
}

func DeleteCdnDomain(req *DeleteCdnDomainRequest, accessId, accessSecret string) (*DeleteCdnDomainResponse, error) {
	var pResponse DeleteCdnDomainResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
