package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeactivateServiceRequest struct {
	RpcRequest
	AccountId string
}

func (r *DeactivateServiceRequest) SetAccountId(value string) {
	r.AccountId = value
	r.QueryParams.Set("AccountId", value)
}
func (r *DeactivateServiceRequest) GetAccountId() string {
	return r.AccountId
}

func (r *DeactivateServiceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeactivateService")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeactivateServiceResponse struct {
}

func DeactivateService(req *DeactivateServiceRequest, accessId, accessSecret string) (*DeactivateServiceResponse, error) {
	var pResponse DeactivateServiceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
