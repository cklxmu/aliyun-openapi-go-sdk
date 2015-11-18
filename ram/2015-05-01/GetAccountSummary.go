package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetAccountSummaryRequest struct {
	RpcRequest
}

func (r *GetAccountSummaryRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetAccountSummary")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetAccountSummaryResponse struct {
	SummaryMap struct {
		UsersQuota                    int `xml:"UsersQuota" json:"UsersQuota"`
		Users                         int `xml:"Users" json:"Users"`
		AccessKeysPerUserQuota        int `xml:"AccessKeysPerUserQuota" json:"AccessKeysPerUserQuota"`
		VirtualMFADevicesQuota        int `xml:"VirtualMFADevicesQuota" json:"VirtualMFADevicesQuota"`
		MFADevices                    int `xml:"MFADevices" json:"MFADevices"`
		MFADevicesInUse               int `xml:"MFADevicesInUse" json:"MFADevicesInUse"`
		GroupsQuota                   int `xml:"GroupsQuota" json:"GroupsQuota"`
		Groups                        int `xml:"Groups" json:"Groups"`
		GroupsPerUserQuota            int `xml:"GroupsPerUserQuota" json:"GroupsPerUserQuota"`
		RolesQuota                    int `xml:"RolesQuota" json:"RolesQuota"`
		Roles                         int `xml:"Roles" json:"Roles"`
		PoliciesQuota                 int `xml:"PoliciesQuota" json:"PoliciesQuota"`
		Policies                      int `xml:"Policies" json:"Policies"`
		PolicySizeQuota               int `xml:"PolicySizeQuota" json:"PolicySizeQuota"`
		VersionsPerPolicyQuota        int `xml:"VersionsPerPolicyQuota" json:"VersionsPerPolicyQuota"`
		AttachedPoliciesPerUserQuota  int `xml:"AttachedPoliciesPerUserQuota" json:"AttachedPoliciesPerUserQuota"`
		AttachedPoliciesPerGroupQuota int `xml:"AttachedPoliciesPerGroupQuota" json:"AttachedPoliciesPerGroupQuota"`
		AttachedPoliciesPerRoleQuota  int `xml:"AttachedPoliciesPerRoleQuota" json:"AttachedPoliciesPerRoleQuota"`
	} `xml:"SummaryMap" json:"SummaryMap"`
}

func GetAccountSummary(req *GetAccountSummaryRequest, accessId, accessSecret string) (*GetAccountSummaryResponse, error) {
	var pResponse GetAccountSummaryResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
