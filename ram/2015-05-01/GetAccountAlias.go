package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetAccountAliasRequest struct {
	RpcRequest
}

func (r *GetAccountAliasRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetAccountAlias")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetAccountAliasResponse struct {
	AccountAlias string `xml:"AccountAlias" json:"AccountAlias"`
}

func GetAccountAlias(req *GetAccountAliasRequest, accessId, accessSecret string) (*GetAccountAliasResponse, error) {
	var pResponse GetAccountAliasResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
