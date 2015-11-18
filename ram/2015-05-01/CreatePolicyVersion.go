package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreatePolicyVersionRequest struct {
	RpcRequest
	PolicyName     string
	PolicyDocument string
	SetAsDefault   bool
}

func (r *CreatePolicyVersionRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *CreatePolicyVersionRequest) GetPolicyName() string {
	return r.PolicyName
}
func (r *CreatePolicyVersionRequest) SetPolicyDocument(value string) {
	r.PolicyDocument = value
	r.QueryParams.Set("PolicyDocument", value)
}
func (r *CreatePolicyVersionRequest) GetPolicyDocument() string {
	return r.PolicyDocument
}
func (r *CreatePolicyVersionRequest) SetSetAsDefault(value bool) {
	r.SetAsDefault = value
	r.QueryParams.Set("SetAsDefault", strconv.FormatBool(value))
}
func (r *CreatePolicyVersionRequest) GetSetAsDefault() bool {
	return r.SetAsDefault
}

func (r *CreatePolicyVersionRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreatePolicyVersion")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreatePolicyVersionResponse struct {
	PolicyVersion struct {
		VersionId        string `xml:"VersionId" json:"VersionId"`
		IsDefaultVersion bool   `xml:"IsDefaultVersion" json:"IsDefaultVersion"`
		PolicyDocument   string `xml:"PolicyDocument" json:"PolicyDocument"`
		CreateDate       string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"PolicyVersion" json:"PolicyVersion"`
}

func CreatePolicyVersion(req *CreatePolicyVersionRequest, accessId, accessSecret string) (*CreatePolicyVersionResponse, error) {
	var pResponse CreatePolicyVersionResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
