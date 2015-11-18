package aas

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateAliyunAccountRequest struct {
	RpcRequest
	AliyunId string
}

func (r *CreateAliyunAccountRequest) SetAliyunId(value string) {
	r.AliyunId = value
	r.QueryParams.Set("AliyunId", value)
}
func (r *CreateAliyunAccountRequest) GetAliyunId() string {
	return r.AliyunId
}

func (r *CreateAliyunAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateAliyunAccount")
	r.SetProduct(Product)
}

type CreateAliyunAccountResponse struct {
	PK       string `xml:"PK" json:"PK"`
	AliyunId string `xml:"AliyunId" json:"AliyunId"`
}

func CreateAliyunAccount(req *CreateAliyunAccountRequest, accessId, accessSecret string) (*CreateAliyunAccountResponse, error) {
	var pResponse CreateAliyunAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
