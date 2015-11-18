package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetUserRequest struct {
	RpcRequest
	UserName string
}

func (r *GetUserRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *GetUserRequest) GetUserName() string {
	return r.UserName
}

func (r *GetUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetUser")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetUserResponse struct {
	User struct {
		UserId        string `xml:"UserId" json:"UserId"`
		UserName      string `xml:"UserName" json:"UserName"`
		DisplayName   string `xml:"DisplayName" json:"DisplayName"`
		MobilePhone   string `xml:"MobilePhone" json:"MobilePhone"`
		Email         string `xml:"Email" json:"Email"`
		Comments      string `xml:"Comments" json:"Comments"`
		CreateDate    string `xml:"CreateDate" json:"CreateDate"`
		UpdateDate    string `xml:"UpdateDate" json:"UpdateDate"`
		LastLoginDate string `xml:"LastLoginDate" json:"LastLoginDate"`
	} `xml:"User" json:"User"`
}

func GetUser(req *GetUserRequest, accessId, accessSecret string) (*GetUserResponse, error) {
	var pResponse GetUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
