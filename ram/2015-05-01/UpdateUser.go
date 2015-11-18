package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateUserRequest struct {
	RpcRequest
	UserName       string
	NewUserName    string
	NewDisplayName string
	NewMobilePhone string
	NewEmail       string
	NewComments    string
}

func (r *UpdateUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *UpdateUserRequest) GetUserName() string {
	return r.UserName
}
func (r *UpdateUserRequest) SetNewUserName(value string) {
	r.NewUserName = value
	r.QueryParams.Set("NewUserName", value)
}
func (r *UpdateUserRequest) GetNewUserName() string {
	return r.NewUserName
}
func (r *UpdateUserRequest) SetNewDisplayName(value string) {
	r.NewDisplayName = value
	r.QueryParams.Set("NewDisplayName", value)
}
func (r *UpdateUserRequest) GetNewDisplayName() string {
	return r.NewDisplayName
}
func (r *UpdateUserRequest) SetNewMobilePhone(value string) {
	r.NewMobilePhone = value
	r.QueryParams.Set("NewMobilePhone", value)
}
func (r *UpdateUserRequest) GetNewMobilePhone() string {
	return r.NewMobilePhone
}
func (r *UpdateUserRequest) SetNewEmail(value string) {
	r.NewEmail = value
	r.QueryParams.Set("NewEmail", value)
}
func (r *UpdateUserRequest) GetNewEmail() string {
	return r.NewEmail
}
func (r *UpdateUserRequest) SetNewComments(value string) {
	r.NewComments = value
	r.QueryParams.Set("NewComments", value)
}
func (r *UpdateUserRequest) GetNewComments() string {
	return r.NewComments
}

func (r *UpdateUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type UpdateUserResponse struct {
	User struct {
		UserId      string `xml:"UserId" json:"UserId"`
		UserName    string `xml:"UserName" json:"UserName"`
		DisplayName string `xml:"DisplayName" json:"DisplayName"`
		MobilePhone string `xml:"MobilePhone" json:"MobilePhone"`
		Email       string `xml:"Email" json:"Email"`
		Comments    string `xml:"Comments" json:"Comments"`
		CreateDate  string `xml:"CreateDate" json:"CreateDate"`
		UpdateDate  string `xml:"UpdateDate" json:"UpdateDate"`
	} `xml:"User" json:"User"`
}

func UpdateUser(req *UpdateUserRequest, accessId, accessSecret string) (*UpdateUserResponse, error) {
	var pResponse UpdateUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
