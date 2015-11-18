package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetDefaultPolicyVersionRequest struct {
	RpcRequest
	PolicyName string
	VersionId  string
}

func (r *SetDefaultPolicyVersionRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *SetDefaultPolicyVersionRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *SetDefaultPolicyVersionRequest) SetVersionId(value string) {
	r.VersionId = value
	r.QueryParams.Set("VersionId", value)
}
func (r *SetDefaultPolicyVersionRequest) GetVersionId() string {
	return r.VersionId
}

func (r *SetDefaultPolicyVersionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetDefaultPolicyVersion")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type SetDefaultPolicyVersionResponse struct {
}

func SetDefaultPolicyVersion(req *SetDefaultPolicyVersionRequest, accessId, accessSecret string) (*SetDefaultPolicyVersionResponse, error) {
	var pResponse SetDefaultPolicyVersionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
