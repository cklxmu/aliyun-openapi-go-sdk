package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ClearAccountAliasRequest struct {
	RpcRequest
}

func (r *ClearAccountAliasRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ClearAccountAlias")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ClearAccountAliasResponse struct {
}

func ClearAccountAlias(req *ClearAccountAliasRequest, accessId, accessSecret string) (*ClearAccountAliasResponse, error) {
	var pResponse ClearAccountAliasResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
