package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetPolicyVersionRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
	VersionId  string
}

func (r *GetPolicyVersionRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *GetPolicyVersionRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *GetPolicyVersionRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *GetPolicyVersionRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *GetPolicyVersionRequest) SetVersionId(value string) {
	r.VersionId = value
	r.QueryParams.Set("VersionId", value)
}
func (r *GetPolicyVersionRequest) GetVersionId() string {
	return r.VersionId
}

func (r *GetPolicyVersionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetPolicyVersion")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetPolicyVersionResponse struct {
	PolicyVersion struct {
		VersionId        string `xml:"VersionId" json:"VersionId"`
		IsDefaultVersion bool   `xml:"IsDefaultVersion" json:"IsDefaultVersion"`
		PolicyDocument   string `xml:"PolicyDocument" json:"PolicyDocument"`
		CreateDate       string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"PolicyVersion" json:"PolicyVersion"`
}

func GetPolicyVersion(req *GetPolicyVersionRequest, accessId, accessSecret string) (*GetPolicyVersionResponse, error) {
	var pResponse GetPolicyVersionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
