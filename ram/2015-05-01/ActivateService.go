package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ActivateServiceRequest struct {
	RpcRequest
	AccountId string
}

func (r *ActivateServiceRequest) SetAccountId(value string) {
	r.AccountId = value
	r.QueryParams.Set("AccountId", value)
}
func (r *ActivateServiceRequest) GetAccountId() string {
	return r.AccountId
}

func (r *ActivateServiceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ActivateService")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ActivateServiceResponse struct {
}

func ActivateService(req *ActivateServiceRequest, accessId, accessSecret string) (*ActivateServiceResponse, error) {
	var pResponse ActivateServiceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
