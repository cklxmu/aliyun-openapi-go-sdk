package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetPolicyRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
}

func (r *GetPolicyRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *GetPolicyRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *GetPolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *GetPolicyRequest) GetPolicyName() string {
	return r.PolicyName
}

func (r *GetPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetPolicyResponse struct {
	Policy struct {
		PolicyName      string `xml:"PolicyName" json:"PolicyName"`
		PolicyType      string `xml:"PolicyType" json:"PolicyType"`
		Description     string `xml:"Description" json:"Description"`
		DefaultVersion  string `xml:"DefaultVersion" json:"DefaultVersion"`
		PolicyDocument  string `xml:"PolicyDocument" json:"PolicyDocument"`
		CreateDate      string `xml:"CreateDate" json:"CreateDate"`
		UpdateDate      string `xml:"UpdateDate" json:"UpdateDate"`
		AttachmentCount int    `xml:"AttachmentCount" json:"AttachmentCount"`
	} `xml:"Policy" json:"Policy"`
}

func GetPolicy(req *GetPolicyRequest, accessId, accessSecret string) (*GetPolicyResponse, error) {
	var pResponse GetPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
