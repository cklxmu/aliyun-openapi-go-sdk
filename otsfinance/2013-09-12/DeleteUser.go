package otsfinance

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteUserRequest struct {
	RpcRequest
}

func (r *DeleteUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteUser")
	r.SetMethod("POST")
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
