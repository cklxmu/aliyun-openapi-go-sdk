package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListAccessKeysRequest struct {
	RpcRequest
	UserName string
}

func (r *ListAccessKeysRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *ListAccessKeysRequest) GetUserName() string {
	return r.UserName
}

func (r *ListAccessKeysRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListAccessKeys")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListAccessKeysResponse struct {
	AccessKeys struct {
		AccessKey []struct {
			AccessKeyId string `xml:"AccessKeyId" json:"AccessKeyId"`
			Status      string `xml:"Status" json:"Status"`
			CreateDate  string `xml:"CreateDate" json:"CreateDate"`
		} `xml:"AccessKey" json:"AccessKey"`
	} `xml:"AccessKeys" json:"AccessKeys"`
}

func ListAccessKeys(req *ListAccessKeysRequest, accessId, accessSecret string) (*ListAccessKeysResponse, error) {
	var pResponse ListAccessKeysResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
