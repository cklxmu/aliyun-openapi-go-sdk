package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type AddUserToGroupRequest struct {
	RpcRequest
	UserName  string
	GroupName string
}

func (r *AddUserToGroupRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *AddUserToGroupRequest) GetUserName() string {
	return r.UserName
}
func (r *AddUserToGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *AddUserToGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *AddUserToGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddUserToGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AddUserToGroupResponse struct {
}

func AddUserToGroup(req *AddUserToGroupRequest, accessId, accessSecret string) (*AddUserToGroupResponse, error) {
	var pResponse AddUserToGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
