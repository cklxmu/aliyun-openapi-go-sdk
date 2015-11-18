package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UpdateLoginProfileRequest struct {
	RpcRequest
	UserName              string
	Password              string
	PasswordResetRequired bool
	MFABindRequired       bool
}

func (r *UpdateLoginProfileRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *UpdateLoginProfileRequest) GetUserName() string {
	return r.UserName
}
func (r *UpdateLoginProfileRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *UpdateLoginProfileRequest) GetPassword() string {
	return r.Password
}
func (r *UpdateLoginProfileRequest) SetPasswordResetRequired(value bool) {
	r.PasswordResetRequired = value
	r.QueryParams.Set("PasswordResetRequired", strconv.FormatBool(value))
}
func (r *UpdateLoginProfileRequest) GetPasswordResetRequired() bool {
	return r.PasswordResetRequired
}
func (r *UpdateLoginProfileRequest) SetMFABindRequired(value bool) {
	r.MFABindRequired = value
	r.QueryParams.Set("MFABindRequired", strconv.FormatBool(value))
}
func (r *UpdateLoginProfileRequest) GetMFABindRequired() bool {
	return r.MFABindRequired
}

func (r *UpdateLoginProfileRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateLoginProfile")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type UpdateLoginProfileResponse struct {
}

func UpdateLoginProfile(req *UpdateLoginProfileRequest, accessId, accessSecret string) (*UpdateLoginProfileResponse, error) {
	var pResponse UpdateLoginProfileResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
