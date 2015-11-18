package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListUsersRequest struct {
	RpcRequest
}

func (r *ListUsersRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListUsers")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListUsersResponse struct {
	Users struct {
		User []struct {
			UserName string `xml:"UserName" json:"UserName"`
			Comments string `xml:"Comments" json:"Comments"`
		} `xml:"User" json:"User"`
	} `xml:"Users" json:"Users"`
}

func ListUsers(req *ListUsersRequest, accessId, accessSecret string) (*ListUsersResponse, error) {
	var pResponse ListUsersResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
