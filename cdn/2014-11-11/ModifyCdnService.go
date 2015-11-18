package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyCdnServiceRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	InternetChargeType   string
}

func (r *ModifyCdnServiceRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *ModifyCdnServiceRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *ModifyCdnServiceRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *ModifyCdnServiceRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *ModifyCdnServiceRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *ModifyCdnServiceRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *ModifyCdnServiceRequest) SetInternetChargeType(value string) {
	r.InternetChargeType = value
	r.QueryParams.Set("InternetChargeType", value)
}
func (r *ModifyCdnServiceRequest) GetInternetChargeType() string {
	return r.InternetChargeType
}

func (r *ModifyCdnServiceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyCdnService")
	r.SetProduct(Product)
}

type ModifyCdnServiceResponse struct {
}

func ModifyCdnService(req *ModifyCdnServiceRequest, accessId, accessSecret string) (*ModifyCdnServiceResponse, error) {
	var pResponse ModifyCdnServiceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
