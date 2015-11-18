package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteUserPolicyRequest struct {
	RpcRequest
	UserName   string
	PolicyName string
}

func (r *DeleteUserPolicyRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *DeleteUserPolicyRequest) GetUserName() string {
	return r.UserName
}
func (r *DeleteUserPolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *DeleteUserPolicyRequest) GetPolicyName() string {
	return r.PolicyName
}

func (r *DeleteUserPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteUserPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteUserPolicyResponse struct {
}

func DeleteUserPolicy(req *DeleteUserPolicyRequest, accessId, accessSecret string) (*DeleteUserPolicyResponse, error) {
	var pResponse DeleteUserPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
