package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListPoliciesForRoleRequest struct {
	RpcRequest
	RoleName string
}

func (r *ListPoliciesForRoleRequest) SetRoleName(value string) {
	r.RoleName = value
	r.QueryParams.Set("RoleName", value)
}
func (r *ListPoliciesForRoleRequest) GetRoleName() string {
	return r.RoleName
}

func (r *ListPoliciesForRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListPoliciesForRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListPoliciesForRoleResponse struct {
	Policies struct {
		Policy []struct {
			PolicyName     string `xml:"PolicyName" json:"PolicyName"`
			PolicyType     string `xml:"PolicyType" json:"PolicyType"`
			Description    string `xml:"Description" json:"Description"`
			DefaultVersion string `xml:"DefaultVersion" json:"DefaultVersion"`
			AttachDate     string `xml:"AttachDate" json:"AttachDate"`
		} `xml:"Policy" json:"Policy"`
	} `xml:"Policies" json:"Policies"`
}

func ListPoliciesForRole(req *ListPoliciesForRoleRequest, accessId, accessSecret string) (*ListPoliciesForRoleResponse, error) {
	var pResponse ListPoliciesForRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
