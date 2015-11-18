package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListUserPoliciesRequest struct {
	RpcRequest
	UserName string
}

func (r *ListUserPoliciesRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *ListUserPoliciesRequest) GetUserName() string {
	return r.UserName
}

func (r *ListUserPoliciesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListUserPolicies")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListUserPoliciesResponse struct {
	Policies struct {
		Policy []string `xml:"Policy" json:"Policy"`
	} `xml:"Policies" json:"Policies"`
}

func ListUserPolicies(req *ListUserPoliciesRequest, accessId, accessSecret string) (*ListUserPoliciesResponse, error) {
	var pResponse ListUserPoliciesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
