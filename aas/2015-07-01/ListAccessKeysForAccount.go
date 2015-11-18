package aas

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListAccessKeysForAccountRequest struct {
	RpcRequest
	PK       string
	AKStatus string
	AKType   string
}

func (r *ListAccessKeysForAccountRequest) SetPK(value string) {
	r.PK = value
	r.QueryParams.Set("PK", value)
}
func (r *ListAccessKeysForAccountRequest) GetPK() string {
	return r.PK
}
func (r *ListAccessKeysForAccountRequest) SetAKStatus(value string) {
	r.AKStatus = value
	r.QueryParams.Set("AKStatus", value)
}
func (r *ListAccessKeysForAccountRequest) GetAKStatus() string {
	return r.AKStatus
}
func (r *ListAccessKeysForAccountRequest) SetAKType(value string) {
	r.AKType = value
	r.QueryParams.Set("AKType", value)
}
func (r *ListAccessKeysForAccountRequest) GetAKType() string {
	return r.AKType
}

func (r *ListAccessKeysForAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListAccessKeysForAccount")
	r.SetProduct(Product)
}

type ListAccessKeysForAccountResponse struct {
	PK         string `xml:"PK" json:"PK"`
	AccessKeys struct {
		AccessKey []struct {
			CreateTime      string `xml:"CreateTime" json:"CreateTime"`
			AccessKeyId     string `xml:"AccessKeyId" json:"AccessKeyId"`
			AccessKeySecret string `xml:"AccessKeySecret" json:"AccessKeySecret"`
			AccessKeyStatus string `xml:"AccessKeyStatus" json:"AccessKeyStatus"`
			AccessKeyType   string `xml:"AccessKeyType" json:"AccessKeyType"`
		} `xml:"AccessKey" json:"AccessKey"`
	} `xml:"AccessKeys" json:"AccessKeys"`
}

func ListAccessKeysForAccount(req *ListAccessKeysForAccountRequest, accessId, accessSecret string) (*ListAccessKeysForAccountResponse, error) {
	var pResponse ListAccessKeysForAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
