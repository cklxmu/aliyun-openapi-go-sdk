package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListGroupsForUserRequest struct {
	RpcRequest
	UserName string
}

func (r *ListGroupsForUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *ListGroupsForUserRequest) GetUserName() string {
	return r.UserName
}

func (r *ListGroupsForUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListGroupsForUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListGroupsForUserResponse struct {
	Groups struct {
		Group []struct {
			GroupName string `xml:"GroupName" json:"GroupName"`
			Comments  string `xml:"Comments" json:"Comments"`
			JoinDate  string `xml:"JoinDate" json:"JoinDate"`
		} `xml:"Group" json:"Group"`
	} `xml:"Groups" json:"Groups"`
}

func ListGroupsForUser(req *ListGroupsForUserRequest, accessId, accessSecret string) (*ListGroupsForUserResponse, error) {
	var pResponse ListGroupsForUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
