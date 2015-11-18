package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateAccessKeyRequest struct {
	RpcRequest
	UserName string
}

func (r *CreateAccessKeyRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *CreateAccessKeyRequest) GetUserName() string {
	return r.UserName
}

func (r *CreateAccessKeyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateAccessKey")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreateAccessKeyResponse struct {
	AccessKey struct {
		AccessKeyId     string `xml:"AccessKeyId" json:"AccessKeyId"`
		AccessKeySecret string `xml:"AccessKeySecret" json:"AccessKeySecret"`
		Status          string `xml:"Status" json:"Status"`
		CreateDate      string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"AccessKey" json:"AccessKey"`
}

func CreateAccessKey(req *CreateAccessKeyRequest, accessId, accessSecret string) (*CreateAccessKeyResponse, error) {
	var pResponse CreateAccessKeyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
