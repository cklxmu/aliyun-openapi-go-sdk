package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeletePolicyVersionRequest struct {
	RpcRequest
	PolicyName string
	VersionId  string
}

func (r *DeletePolicyVersionRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *DeletePolicyVersionRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *DeletePolicyVersionRequest) SetVersionId(value string) {
	r.VersionId = value
	r.QueryParams.Set("VersionId", value)
}
func (r *DeletePolicyVersionRequest) GetVersionId() string {
	return r.VersionId
}

func (r *DeletePolicyVersionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeletePolicyVersion")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeletePolicyVersionResponse struct {
}

func DeletePolicyVersion(req *DeletePolicyVersionRequest, accessId, accessSecret string) (*DeletePolicyVersionResponse, error) {
	var pResponse DeletePolicyVersionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
