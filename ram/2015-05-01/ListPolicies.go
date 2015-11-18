package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListPoliciesRequest struct {
	RpcRequest
	PolicyType string
	Marker     string
	MaxItems   int
}

func (r *ListPoliciesRequest) SetPolicyType(value string) {
	r.PolicyType = value
	r.QueryParams.Set("PolicyType", value)
}
func (r *ListPoliciesRequest) GetPolicyType() string {
	return r.PolicyType
}
func (r *ListPoliciesRequest) SetMarker(value string) {
	r.Marker = value
	r.QueryParams.Set("Marker", value)
}
func (r *ListPoliciesRequest) GetMarker() string {
	return r.Marker
}
func (r *ListPoliciesRequest) SetMaxItems(value int) {
	r.MaxItems = value
	r.QueryParams.Set("MaxItems", strconv.Itoa(value))
}
func (r *ListPoliciesRequest) GetMaxItems() int {
	return r.MaxItems
}

func (r *ListPoliciesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListPolicies")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListPoliciesResponse struct {
	IsTruncated bool   `xml:"IsTruncated" json:"IsTruncated"`
	Marker      string `xml:"Marker" json:"Marker"`
	Policies    struct {
		Policy []struct {
			PolicyName      string `xml:"PolicyName" json:"PolicyName"`
			PolicyType      string `xml:"PolicyType" json:"PolicyType"`
			Description     string `xml:"Description" json:"Description"`
			DefaultVersion  string `xml:"DefaultVersion" json:"DefaultVersion"`
			CreateDate      string `xml:"CreateDate" json:"CreateDate"`
			UpdateDate      string `xml:"UpdateDate" json:"UpdateDate"`
			AttachmentCount int    `xml:"AttachmentCount" json:"AttachmentCount"`
		} `xml:"Policy" json:"Policy"`
	} `xml:"Policies" json:"Policies"`
}

func ListPolicies(req *ListPoliciesRequest, accessId, accessSecret string) (*ListPoliciesResponse, error) {
	var pResponse ListPoliciesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
