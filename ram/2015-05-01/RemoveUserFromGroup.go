package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type RemoveUserFromGroupRequest struct {
	RpcRequest
	UserName  string
	GroupName string
}

func (r *RemoveUserFromGroupRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *RemoveUserFromGroupRequest) GetUserName() string {
	return r.UserName
}
func (r *RemoveUserFromGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *RemoveUserFromGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *RemoveUserFromGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RemoveUserFromGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type RemoveUserFromGroupResponse struct {
}

func RemoveUserFromGroup(req *RemoveUserFromGroupRequest, accessId, accessSecret string) (*RemoveUserFromGroupResponse, error) {
	var pResponse RemoveUserFromGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
