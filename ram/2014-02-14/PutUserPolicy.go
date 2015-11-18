package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type PutUserPolicyRequest struct {
	RpcRequest
	UserName       string
	PolicyName     string
	PolicyDocument string
}

func (r *PutUserPolicyRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *PutUserPolicyRequest) GetUserName() string {
	return r.UserName
}
func (r *PutUserPolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *PutUserPolicyRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *PutUserPolicyRequest) SetPolicyDocument(value string) {
	r.PolicyDocument = value
	r.QueryParams.Set("PolicyDocument", value)
}
func (r *PutUserPolicyRequest) GetPolicyDocument() string {
	return r.PolicyDocument
}

func (r *PutUserPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PutUserPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type PutUserPolicyResponse struct {
}

func PutUserPolicy(req *PutUserPolicyRequest, accessId, accessSecret string) (*PutUserPolicyResponse, error) {
	var pResponse PutUserPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
