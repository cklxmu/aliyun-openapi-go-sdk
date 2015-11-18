package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetAccountAliasRequest struct {
	RpcRequest
	AccountAlias string
}

func (r *SetAccountAliasRequest) SetAccountAlias(value string) {
	r.AccountAlias = value
	r.QueryParams.Set("AccountAlias", value)
}
func (r *SetAccountAliasRequest) GetAccountAlias() string {
	return r.AccountAlias
}

func (r *SetAccountAliasRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetAccountAlias")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type SetAccountAliasResponse struct {
}

func SetAccountAlias(req *SetAccountAliasRequest, accessId, accessSecret string) (*SetAccountAliasResponse, error) {
	var pResponse SetAccountAliasResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
