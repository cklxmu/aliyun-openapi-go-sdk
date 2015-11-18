package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DetachPolicyFromRoleRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	RoleName   string
}

func (r *DetachPolicyFromRoleRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *DetachPolicyFromRoleRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *DetachPolicyFromRoleRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *DetachPolicyFromRoleRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *DetachPolicyFromRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *DetachPolicyFromRoleRequest) GetRoleName() string {
	return r.RoleName
}

func (r *DetachPolicyFromRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DetachPolicyFromRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DetachPolicyFromRoleResponse struct {
}

func DetachPolicyFromRole(req *DetachPolicyFromRoleRequest, accessId, accessSecret string) (*DetachPolicyFromRoleResponse, error) {
	var pResponse DetachPolicyFromRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
