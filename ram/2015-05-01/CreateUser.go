package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateUserRequest struct {
	RpcRequest
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
}

func (r *CreateUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *CreateUserRequest) GetUserName() string {
	return r.UserName
}
func (r *CreateUserRequest) SetDisplayName(value string) {
	r.DisplayName = value
	r.QueryParams.Set("DisplayName", value)
}
func (r *CreateUserRequest) GetDisplayName() string {
	return r.DisplayName
}
func (r *CreateUserRequest) SetMobilePhone(value string) {
	r.MobilePhone = value
	r.QueryParams.Set("MobilePhone", value)
}
func (r *CreateUserRequest) GetMobilePhone() string {
	return r.MobilePhone
}
func (r *CreateUserRequest) SetEmail(value string) {
	r.Email = value
	r.QueryParams.Set("Email", value)
}
func (r *CreateUserRequest) GetEmail() string {
	return r.Email
}
func (r *CreateUserRequest) SetComments(value string) {
	r.Comments = value
	r.QueryParams.Set("Comments", value)
}
func (r *CreateUserRequest) GetComments() string {
	return r.Comments
}

func (r *CreateUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreateUserResponse struct {
	User struct {
		UserId      string `xml:"UserId" json:"UserId"`
		UserName    string `xml:"UserName" json:"UserName"`
		DisplayName string `xml:"DisplayName" json:"DisplayName"`
		MobilePhone string `xml:"MobilePhone" json:"MobilePhone"`
		Email       string `xml:"Email" json:"Email"`
		Comments    string `xml:"Comments" json:"Comments"`
		CreateDate  string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"User" json:"User"`
}

func CreateUser(req *CreateUserRequest, accessId, accessSecret string) (*CreateUserResponse, error) {
	var pResponse CreateUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
