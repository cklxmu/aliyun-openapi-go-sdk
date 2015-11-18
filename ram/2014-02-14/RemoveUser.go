package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type RemoveUserRequest struct {
	RpcRequest
	UserName string
}

func (r *RemoveUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *RemoveUserRequest) GetUserName() string {
	return r.UserName
}

func (r *RemoveUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type RemoveUserResponse struct {
}

func RemoveUser(req *RemoveUserRequest, accessId, accessSecret string) (*RemoveUserResponse, error) {
	var pResponse RemoveUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
