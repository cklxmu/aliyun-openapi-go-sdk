package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListPolicyVersionsRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
}

func (r *ListPolicyVersionsRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *ListPolicyVersionsRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *ListPolicyVersionsRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *ListPolicyVersionsRequest) GetPolicyName() string {
	return r.PolicyName
}

func (r *ListPolicyVersionsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListPolicyVersions")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListPolicyVersionsResponse struct {
	PolicyVersions struct {
		PolicyVersion []struct {
			VersionId        string `xml:"VersionId" json:"VersionId"`
			IsDefaultVersion bool   `xml:"IsDefaultVersion" json:"IsDefaultVersion"`
			PolicyDocument   string `xml:"PolicyDocument" json:"PolicyDocument"`
			CreateDate       string `xml:"CreateDate" json:"CreateDate"`
		} `xml:"PolicyVersion" json:"PolicyVersion"`
	} `xml:"PolicyVersions" json:"PolicyVersions"`
}

func ListPolicyVersions(req *ListPolicyVersionsRequest, accessId, accessSecret string) (*ListPolicyVersionsResponse, error) {
	var pResponse ListPolicyVersionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
