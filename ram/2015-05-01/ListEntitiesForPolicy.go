package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListEntitiesForPolicyRequest struct {
	RpcRequest
	PolicyType string
	PolicyName string
}

func (r *ListEntitiesForPolicyRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *ListEntitiesForPolicyRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *ListEntitiesForPolicyRequest) SetPolicyName(value string) {
	r.PolicyName = value
	r.QueryParams.Set("PolicyName", value)
}
func (r *ListEntitiesForPolicyRequest) GetPolicyName() string {
	return r.PolicyName
}

func (r *ListEntitiesForPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListEntitiesForPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListEntitiesForPolicyResponse struct {
	Groups struct {
		Group []struct {
			GroupName  string `xml:"GroupName" json:"GroupName"`
			Comments   string `xml:"Comments" json:"Comments"`
			AttachDate string `xml:"AttachDate" json:"AttachDate"`
		} `xml:"Group" json:"Group"`
	} `xml:"Groups" json:"Groups"`
	Users struct {
		User []struct {
			UserId      string `xml:"UserId" json:"UserId"`
			UserName    string `xml:"UserName" json:"UserName"`
			DisplayName string `xml:"DisplayName" json:"DisplayName"`
			AttachDate  string `xml:"AttachDate" json:"AttachDate"`
		} `xml:"User" json:"User"`
	} `xml:"Users" json:"Users"`
	Roles struct {
		Role []struct {
			RoleId      string `xml:"RoleId" json:"RoleId"`
			RoleName    string `xml:"RoleName" json:"RoleName"`
			Arn         string `xml:"Arn" json:"Arn"`
			Description string `xml:"Description" json:"Description"`
			AttachDate  string `xml:"AttachDate" json:"AttachDate"`
		} `xml:"Role" json:"Role"`
	} `xml:"Roles" json:"Roles"`
}

func ListEntitiesForPolicy(req *ListEntitiesForPolicyRequest, accessId, accessSecret string) (*ListEntitiesForPolicyResponse, error) {
	var pResponse ListEntitiesForPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
