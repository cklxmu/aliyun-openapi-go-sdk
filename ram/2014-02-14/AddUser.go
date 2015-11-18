package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type AddUserRequest struct {
	RpcRequest
	UserName string
	Comments string
}

func (r *AddUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *AddUserRequest) GetUserName() string {
	return r.UserName
}
func (r *AddUserRequest) SetComments(value string) {
	r.Comments = value
	r.QueryParams.Set("Comments", value)
}
func (r *AddUserRequest) GetComments() string {
	return r.Comments
}

func (r *AddUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AddUserResponse struct {
}

func AddUser(req *AddUserRequest, accessId, accessSecret string) (*AddUserResponse, error) {
	var pResponse AddUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
