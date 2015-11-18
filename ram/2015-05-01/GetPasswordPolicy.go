package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetPasswordPolicyRequest struct {
	RpcRequest
}

func (r *GetPasswordPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetPasswordPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetPasswordPolicyResponse struct {
	PasswordPolicy struct {
		MinimumPasswordLength      int  `xml:"MinimumPasswordLength" json:"MinimumPasswordLength"`
		RequireLowercaseCharacters bool `xml:"RequireLowercaseCharacters" json:"RequireLowercaseCharacters"`
		RequireUppercaseCharacters bool `xml:"RequireUppercaseCharacters" json:"RequireUppercaseCharacters"`
		RequireNumbers             bool `xml:"RequireNumbers" json:"RequireNumbers"`
		RequireSymbols             bool `xml:"RequireSymbols" json:"RequireSymbols"`
	} `xml:"PasswordPolicy" json:"PasswordPolicy"`
}

func GetPasswordPolicy(req *GetPasswordPolicyRequest, accessId, accessSecret string) (*GetPasswordPolicyResponse, error) {
	var pResponse GetPasswordPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
