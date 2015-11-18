package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListPoliciesForGroupRequest struct {
	RpcRequest
	GroupName string
}

func (r *ListPoliciesForGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *ListPoliciesForGroupRequest) GetGroupName() string {
	return r.GroupName
}

func (r *ListPoliciesForGroupRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListPoliciesForGroup")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListPoliciesForGroupResponse struct {
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

func ListPoliciesForGroup(req *ListPoliciesForGroupRequest, accessId, accessSecret string) (*ListPoliciesForGroupResponse, error) {
	var pResponse ListPoliciesForGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
