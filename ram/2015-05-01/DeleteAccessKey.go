package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteAccessKeyRequest struct {
	RpcRequest
	UserName        string
	UserAccessKeyId string
}

func (r *DeleteAccessKeyRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *DeleteAccessKeyRequest) GetUserName() string {
	return r.UserName
}
func (r *DeleteAccessKeyRequest) SetUserAccessKeyId(value string) {
	r.UserAccessKeyId = value
	r.QueryParams.Set("UserAccessKeyId", value)
}
func (r *DeleteAccessKeyRequest) GetUserAccessKeyId() string {
	return r.UserAccessKeyId
}

func (r *DeleteAccessKeyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteAccessKey")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteAccessKeyResponse struct {
}

func DeleteAccessKey(req *DeleteAccessKeyRequest, accessId, accessSecret string) (*DeleteAccessKeyResponse, error) {
	var pResponse DeleteAccessKeyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
