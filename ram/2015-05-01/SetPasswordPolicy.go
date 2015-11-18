package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetPasswordPolicyRequest struct {
	RpcRequest
	MinimumPasswordLength      int
	RequireLowercaseCharacters bool
	RequireUppercaseCharacters bool
	RequireNumbers             bool
	RequireSymbols             bool
}

func (r *SetPasswordPolicyRequest) SetMinimumPasswordLength(value int) {
	r.MinimumPasswordLength = value
	r.QueryParams.Set("MinimumPasswordLength", strconv.Itoa(value))
}
func (r *SetPasswordPolicyRequest) GetMinimumPasswordLength() int {
	return r.MinimumPasswordLength
}
func (r *SetPasswordPolicyRequest) SetRequireLowercaseCharacters(value bool) {
	r.RequireLowercaseCharacters = value
	r.QueryParams.Set("RequireLowercaseCharacters", strconv.FormatBool(value))
}
func (r *SetPasswordPolicyRequest) GetRequireLowercaseCharacters() bool {
	return r.RequireLowercaseCharacters
}
func (r *SetPasswordPolicyRequest) SetRequireUppercaseCharacters(value bool) {
	r.RequireUppercaseCharacters = value
	r.QueryParams.Set("RequireUppercaseCharacters", strconv.FormatBool(value))
}
func (r *SetPasswordPolicyRequest) GetRequireUppercaseCharacters() bool {
	return r.RequireUppercaseCharacters
}
func (r *SetPasswordPolicyRequest) SetRequireNumbers(value bool) {
	r.RequireNumbers = value
	r.QueryParams.Set("RequireNumbers", strconv.FormatBool(value))
}
func (r *SetPasswordPolicyRequest) GetRequireNumbers() bool {
	return r.RequireNumbers
}
func (r *SetPasswordPolicyRequest) SetRequireSymbols(value bool) {
	r.RequireSymbols = value
	r.QueryParams.Set("RequireSymbols", strconv.FormatBool(value))
}
func (r *SetPasswordPolicyRequest) GetRequireSymbols() bool {
	return r.RequireSymbols
}

func (r *SetPasswordPolicyRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetPasswordPolicy")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type SetPasswordPolicyResponse struct {
	PasswordPolicy struct {
		MinimumPasswordLength      int  `xml:"MinimumPasswordLength" json:"MinimumPasswordLength"`
		RequireLowercaseCharacters bool `xml:"RequireLowercaseCharacters" json:"RequireLowercaseCharacters"`
		RequireUppercaseCharacters bool `xml:"RequireUppercaseCharacters" json:"RequireUppercaseCharacters"`
		RequireNumbers             bool `xml:"RequireNumbers" json:"RequireNumbers"`
		RequireSymbols             bool `xml:"RequireSymbols" json:"RequireSymbols"`
	} `xml:"PasswordPolicy" json:"PasswordPolicy"`
}

func SetPasswordPolicy(req *SetPasswordPolicyRequest, accessId, accessSecret string) (*SetPasswordPolicyResponse, error) {
	var pResponse SetPasswordPolicyResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
