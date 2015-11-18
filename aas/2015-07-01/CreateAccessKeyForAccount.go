package aas

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateAccessKeyForAccountRequest struct {
	RpcRequest
	PK       string
	AKSecret string
}

func (r *CreateAccessKeyForAccountRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *CreateAccessKeyForAccountRequest) GetPK() string {
	return r.PK
}
func (r *CreateAccessKeyForAccountRequest) SetAKSecret(value string) {
	r.AKSecret = value
	r.QueryParams.Set("AKSecret", value)
}
func (r *CreateAccessKeyForAccountRequest) GetAKSecret() string {
	return r.AKSecret
}

func (r *CreateAccessKeyForAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateAccessKeyForAccount")
	r.SetProduct(Product)
}

type CreateAccessKeyForAccountResponse struct {
	PK        string `xml:"PK" json:"PK"`
	AccessKey struct {
		CreateTime      string `xml:"CreateTime" json:"CreateTime"`
		AccessKeyId     string `xml:"AccessKeyId" json:"AccessKeyId"`
		AccessKeySecret string `xml:"AccessKeySecret" json:"AccessKeySecret"`
		AccessKeyStatus string `xml:"AccessKeyStatus" json:"AccessKeyStatus"`
		AccessKeyType   string `xml:"AccessKeyType" json:"AccessKeyType"`
	} `xml:"AccessKey" json:"AccessKey"`
}

func CreateAccessKeyForAccount(req *CreateAccessKeyForAccountRequest, accessId, accessSecret string) (*CreateAccessKeyForAccountResponse, error) {
	var pResponse CreateAccessKeyForAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
