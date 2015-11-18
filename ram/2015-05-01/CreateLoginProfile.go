package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateLoginProfileRequest struct {
	RpcRequest
	UserName              string
	Password              string
	PasswordResetRequired bool
	MFABindRequired       bool
}

func (r *CreateLoginProfileRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *CreateLoginProfileRequest) GetUserName() string {
	return r.UserName
}
func (r *CreateLoginProfileRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *CreateLoginProfileRequest) GetPassword() string {
	return r.Password
}
func (r *CreateLoginProfileRequest) SetPasswordResetRequired(value bool) {
	r.PasswordResetRequired = value
	r.QueryParams.Set("PasswordResetRequired", strconv.FormatBool(value))
}
func (r *CreateLoginProfileRequest) GetPasswordResetRequired() bool {
	return r.PasswordResetRequired
}
func (r *CreateLoginProfileRequest) SetMFABindRequired(value bool) {
	r.MFABindRequired = value
	r.QueryParams.Set("MFABindRequired", strconv.FormatBool(value))
}
func (r *CreateLoginProfileRequest) GetMFABindRequired() bool {
	return r.MFABindRequired
}

func (r *CreateLoginProfileRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateLoginProfile")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreateLoginProfileResponse struct {
	LoginProfile struct {
		UserName              string `xml:"UserName" json:"UserName"`
		PasswordResetRequired bool   `xml:"PasswordResetRequired" json:"PasswordResetRequired"`
		MFABindRequired       bool   `xml:"MFABindRequired" json:"MFABindRequired"`
		CreateDate            string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"LoginProfile" json:"LoginProfile"`
}

func CreateLoginProfile(req *CreateLoginProfileRequest, accessId, accessSecret string) (*CreateLoginProfileResponse, error) {
	var pResponse CreateLoginProfileResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
