package aas

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateAccessKeyStatusForAccountRequest struct {
	RpcRequest
	PK       string
	AKId     string
	AKStatus string
}

func (r *UpdateAccessKeyStatusForAccountRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *UpdateAccessKeyStatusForAccountRequest) GetPK() string {
	return r.PK
}
func (r *UpdateAccessKeyStatusForAccountRequest) SetAKId(value string) {
	r.AKId = value
	r.QueryParams.Set("AKId", value)
}
func (r *UpdateAccessKeyStatusForAccountRequest) GetAKId() string {
	return r.AKId
}
func (r *UpdateAccessKeyStatusForAccountRequest) SetAKStatus(value string) {
	r.AKStatus = value
	r.QueryParams.Set("AKStatus", value)
}
func (r *UpdateAccessKeyStatusForAccountRequest) GetAKStatus() string {
	return r.AKStatus
}

func (r *UpdateAccessKeyStatusForAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateAccessKeyStatusForAccount")
	r.SetProduct(Product)
}

type UpdateAccessKeyStatusForAccountResponse struct {
	PK     string `xml:"PK" json:"PK"`
	Result string `xml:"Result" json:"Result"`
}

func UpdateAccessKeyStatusForAccount(req *UpdateAccessKeyStatusForAccountRequest, accessId, accessSecret string) (*UpdateAccessKeyStatusForAccountResponse, error) {
	var pResponse UpdateAccessKeyStatusForAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
