package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListUsersForGroupRequest struct {
	RpcRequest
	GroupName string
}

func (r *ListUsersForGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *ListUsersForGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *ListUsersForGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListUsersForGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListUsersForGroupResponse struct {
	Users struct {
		User []struct {
			UserName    string `xml:"UserName" json:"UserName"`
			DisplayName string `xml:"DisplayName" json:"DisplayName"`
			JoinDate    string `xml:"JoinDate" json:"JoinDate"`
		} `xml:"User" json:"User"`
	} `xml:"Users" json:"Users"`
}

func ListUsersForGroup(req *ListUsersForGroupRequest, accessId, accessSecret string) (*ListUsersForGroupResponse, error) {
	var pResponse ListUsersForGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
