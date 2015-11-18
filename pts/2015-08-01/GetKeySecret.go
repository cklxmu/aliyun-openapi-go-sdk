package pts

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetKeySecretRequest struct {
	RpcRequest
}

func (r *GetKeySecretRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetKeySecret")
	r.SetProduct(Product)
}

type GetKeySecretResponse struct {
	Key    string `xml:"Key" json:"Key"`
	Secret string `xml:"Secret" json:"Secret"`
}

func GetKeySecret(req *GetKeySecretRequest, accessId, accessSecret string) (*GetKeySecretResponse, error) {
	var pResponse GetKeySecretResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
