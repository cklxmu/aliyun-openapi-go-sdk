package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DetachPolicyFromUserRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	UserName   string
}

func (r *DetachPolicyFromUserRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *DetachPolicyFromUserRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *DetachPolicyFromUserRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *DetachPolicyFromUserRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *DetachPolicyFromUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *DetachPolicyFromUserRequest) GetUserName() string {
	return r.UserName
}

func (r *DetachPolicyFromUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DetachPolicyFromUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DetachPolicyFromUserResponse struct {
}

func DetachPolicyFromUser(req *DetachPolicyFromUserRequest, accessId, accessSecret string) (*DetachPolicyFromUserResponse, error) {
	var pResponse DetachPolicyFromUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
