package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetUserPolicyRequest struct {
	RpcRequest
	UserName   string
	PolicyName string
}

func (r *GetUserPolicyRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *GetUserPolicyRequest) GetUserName() string {
	return r.UserName
}
func (r *GetUserPolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *GetUserPolicyRequest) GetPolicyName() string {
	return r.PolicyName
}

func (r *GetUserPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetUserPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetUserPolicyResponse struct {
	UserName       string `xml:"UserName" json:"UserName"`
	PolicyName     string `xml:"PolicyName" json:"PolicyName"`
	PolicyDocument string `xml:"PolicyDocument" json:"PolicyDocument"`
}

func GetUserPolicy(req *GetUserPolicyRequest, accessId, accessSecret string) (*GetUserPolicyResponse, error) {
	var pResponse GetUserPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
