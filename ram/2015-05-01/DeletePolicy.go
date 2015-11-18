package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeletePolicyRequest struct {
	RpcRequest
	PolicyName string
}

func (r *DeletePolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *DeletePolicyRequest) GetPolicyName() string {
	return r.PolicyName
}

func (r *DeletePolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeletePolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeletePolicyResponse struct {
}

func DeletePolicy(req *DeletePolicyRequest, accessId, accessSecret string) (*DeletePolicyResponse, error) {
	var pResponse DeletePolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
