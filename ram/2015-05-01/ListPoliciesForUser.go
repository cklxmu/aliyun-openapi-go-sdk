package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListPoliciesForUserRequest struct {
	RpcRequest
	UserName string
}

func (r *ListPoliciesForUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *ListPoliciesForUserRequest) GetUserName() string {
	return r.UserName
}

func (r *ListPoliciesForUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListPoliciesForUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListPoliciesForUserResponse struct {
	Policies struct {
		Policy []struct {
			PolicyName     string `xml:"PolicyName" json:"PolicyName"`
			PolicyType     string `xml:"PolicyType" json:"PolicyType"`
			Description    string `xml:"Description" json:"Description"`
			DefaultVersion string `xml:"DefaultVersion" json:"DefaultVersion"`
			AttachDate     string `xml:"AttachDate" json:"AttachDate"`
		} `xml:"Policy" json:"Policy"`
	} `xml:"Policies" json:"Policies"`
}

func ListPoliciesForUser(req *ListPoliciesForUserRequest, accessId, accessSecret string) (*ListPoliciesForUserResponse, error) {
	var pResponse ListPoliciesForUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
