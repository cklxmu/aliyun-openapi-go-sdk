package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateAccessKeyRequest struct {
	RpcRequest
	UserName        string
	UserAccessKeyId string
	Status          string
}

func (r *UpdateAccessKeyRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *UpdateAccessKeyRequest) GetUserName() string {
	return r.UserName
}
func (r *UpdateAccessKeyRequest) SetUserAccessKeyId(value string) {
	r.UserAccessKeyId = value
	r.QueryParams.Set("UserAccessKeyId", value)
}
func (r *UpdateAccessKeyRequest) GetUserAccessKeyId() string {
	return r.UserAccessKeyId
}
func (r *UpdateAccessKeyRequest) SetStatus(value string) {
	r.Status = value
	r.QueryParams.Set("Status", value)
}
func (r *UpdateAccessKeyRequest) GetStatus() string {
	return r.Status
}

func (r *UpdateAccessKeyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateAccessKey")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type UpdateAccessKeyResponse struct {
}

func UpdateAccessKey(req *UpdateAccessKeyRequest, accessId, accessSecret string) (*UpdateAccessKeyResponse, error) {
	var pResponse UpdateAccessKeyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
