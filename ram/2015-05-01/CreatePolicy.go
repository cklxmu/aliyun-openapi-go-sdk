package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreatePolicyRequest struct {
	RpcRequest
	PolicyName     string
	Description    string
	PolicyDocument string
}

func (r *CreatePolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *CreatePolicyRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *CreatePolicyRequest) SetDescription(value string) {
	r.Description = value
	r.QueryParams.Set("Description", value)
}
func (r *CreatePolicyRequest) GetDescription() string {
	return r.Description
}
func (r *CreatePolicyRequest) SetPolicyDocument(value string) {
	r.PolicyDocument = value
	r.QueryParams.Set("PolicyDocument", value)
}
func (r *CreatePolicyRequest) GetPolicyDocument() string {
	return r.PolicyDocument
}

func (r *CreatePolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreatePolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreatePolicyResponse struct {
	Policy struct {
		PolicyName     string `xml:"PolicyName" json:"PolicyName"`
		PolicyType     string `xml:"PolicyType" json:"PolicyType"`
		Description    string `xml:"Description" json:"Description"`
		DefaultVersion string `xml:"DefaultVersion" json:"DefaultVersion"`
		CreateDate     string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"Policy" json:"Policy"`
}

func CreatePolicy(req *CreatePolicyRequest, accessId, accessSecret string) (*CreatePolicyResponse, error) {
	var pResponse CreatePolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
