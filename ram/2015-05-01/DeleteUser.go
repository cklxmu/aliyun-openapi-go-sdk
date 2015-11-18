package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteUserRequest struct {
	RpcRequest
	UserName string
}

func (r *DeleteUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *DeleteUserRequest) GetUserName() string {
	return r.UserName
}

func (r *DeleteUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteUserResponse struct {
}

func DeleteUser(req *DeleteUserRequest, accessId, accessSecret string) (*DeleteUserResponse, error) {
	var pResponse DeleteUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
