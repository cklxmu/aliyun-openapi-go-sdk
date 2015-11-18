package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type AttachPolicyToUserRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	UserName   string
}

func (r *AttachPolicyToUserRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *AttachPolicyToUserRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *AttachPolicyToUserRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *AttachPolicyToUserRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *AttachPolicyToUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *AttachPolicyToUserRequest) GetUserName() string {
	return r.UserName
}

func (r *AttachPolicyToUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AttachPolicyToUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AttachPolicyToUserResponse struct {
}

func AttachPolicyToUser(req *AttachPolicyToUserRequest, accessId, accessSecret string) (*AttachPolicyToUserResponse, error) {
	var pResponse AttachPolicyToUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
