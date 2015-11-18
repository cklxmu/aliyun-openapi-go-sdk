package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type AttachPolicyToRoleRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	RoleName   string
}

func (r *AttachPolicyToRoleRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *AttachPolicyToRoleRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *AttachPolicyToRoleRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *AttachPolicyToRoleRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *AttachPolicyToRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *AttachPolicyToRoleRequest) GetRoleName() string {
	return r.RoleName
}

func (r *AttachPolicyToRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AttachPolicyToRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AttachPolicyToRoleResponse struct {
}

func AttachPolicyToRole(req *AttachPolicyToRoleRequest, accessId, accessSecret string) (*AttachPolicyToRoleResponse, error) {
	var pResponse AttachPolicyToRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
