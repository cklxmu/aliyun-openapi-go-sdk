package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type AttachPolicyToGroupRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	GroupName  string
}

func (r *AttachPolicyToGroupRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *AttachPolicyToGroupRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *AttachPolicyToGroupRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *AttachPolicyToGroupRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *AttachPolicyToGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *AttachPolicyToGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *AttachPolicyToGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AttachPolicyToGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AttachPolicyToGroupResponse struct {
}

func AttachPolicyToGroup(req *AttachPolicyToGroupRequest, accessId, accessSecret string) (*AttachPolicyToGroupResponse, error) {
	var pResponse AttachPolicyToGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
