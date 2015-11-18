package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetLoginProfileRequest struct {
	RpcRequest
	UserName string
}

func (r *GetLoginProfileRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *GetLoginProfileRequest) GetUserName() string {
	return r.UserName
}

func (r *GetLoginProfileRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetLoginProfile")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetLoginProfileResponse struct {
	LoginProfile struct {
		UserName              string `xml:"UserName" json:"UserName"`
		PasswordResetRequired bool   `xml:"PasswordResetRequired" json:"PasswordResetRequired"`
		MFABindRequired       bool   `xml:"MFABindRequired" json:"MFABindRequired"`
		CreateDate            string `xml:"CreateDate" json:"CreateDate"`
	} `xml:"LoginProfile" json:"LoginProfile"`
}

func GetLoginProfile(req *GetLoginProfileRequest, accessId, accessSecret string) (*GetLoginProfileResponse, error) {
	var pResponse GetLoginProfileResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
