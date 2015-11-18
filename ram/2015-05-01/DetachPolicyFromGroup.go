package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DetachPolicyFromGroupRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	GroupName  string
}

func (r *DetachPolicyFromGroupRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *DetachPolicyFromGroupRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *DetachPolicyFromGroupRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *DetachPolicyFromGroupRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *DetachPolicyFromGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *DetachPolicyFromGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *DetachPolicyFromGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DetachPolicyFromGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DetachPolicyFromGroupResponse struct {
}

func DetachPolicyFromGroup(req *DetachPolicyFromGroupRequest, accessId, accessSecret string) (*DetachPolicyFromGroupResponse, error) {
	var pResponse DetachPolicyFromGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
