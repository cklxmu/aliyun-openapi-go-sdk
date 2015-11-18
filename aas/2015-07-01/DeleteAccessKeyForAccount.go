package aas

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteAccessKeyForAccountRequest struct {
	RpcRequest
	PK   string
	AKId string
}

func (r *DeleteAccessKeyForAccountRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *DeleteAccessKeyForAccountRequest) GetPK() string {
	return r.PK
}
func (r *DeleteAccessKeyForAccountRequest) SetAKId(value string) {
	r.AKId = value
	r.QueryParams.Set("AKId", value)
}
func (r *DeleteAccessKeyForAccountRequest) GetAKId() string {
	return r.AKId
}

func (r *DeleteAccessKeyForAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteAccessKeyForAccount")
	r.SetProduct(Product)
}

type DeleteAccessKeyForAccountResponse struct {
	PK     string `xml:"PK" json:"PK"`
	Result string `xml:"Result" json:"Result"`
}

func DeleteAccessKeyForAccount(req *DeleteAccessKeyForAccountRequest, accessId, accessSecret string) (*DeleteAccessKeyForAccountResponse, error) {
	var pResponse DeleteAccessKeyForAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
