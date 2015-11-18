package ubsms_inner

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetUserSecurityStatusRequest struct {
	RpcRequest
	Uid           string
	StatusKey1    string
	StatusValue1  string
	StatusKey2    string
	StatusValue2  string
	StatusKey3    string
	StatusValue3  string
	StatusKey4    string
	StatusValue4  string
	StatusKey5    string
	StatusValue5  string
	StatusKey6    string
	StatusValue6  string
	StatusKey7    string
	StatusValue7  string
	StatusKey8    string
	StatusValue8  string
	StatusKey9    string
	StatusValue9  string
	StatusKey10   string
	StatusValue10 string
	Password      string
}

func (r *SetUserSecurityStatusRequest) SetUid(value string) {
	r.Uid = value
	r.QueryParams.Set("Uid", value)
}
func (r *SetUserSecurityStatusRequest) GetUid() string {
	return r.Uid
}
func (r *SetUserSecurityStatusRequest) SetStatusKey1(value string) {
	r.StatusKey1 = value
	r.QueryParams.Set("StatusKey1", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey1() string {
	return r.StatusKey1
}
func (r *SetUserSecurityStatusRequest) SetStatusValue1(value string) {
	r.StatusValue1 = value
	r.QueryParams.Set("StatusValue1", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue1() string {
	return r.StatusValue1
}
func (r *SetUserSecurityStatusRequest) SetStatusKey2(value string) {
	r.StatusKey2 = value
	r.QueryParams.Set("StatusKey2", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey2() string {
	return r.StatusKey2
}
func (r *SetUserSecurityStatusRequest) SetStatusValue2(value string) {
	r.StatusValue2 = value
	r.QueryParams.Set("StatusValue2", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue2() string {
	return r.StatusValue2
}
func (r *SetUserSecurityStatusRequest) SetStatusKey3(value string) {
	r.StatusKey3 = value
	r.QueryParams.Set("StatusKey3", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey3() string {
	return r.StatusKey3
}
func (r *SetUserSecurityStatusRequest) SetStatusValue3(value string) {
	r.StatusValue3 = value
	r.QueryParams.Set("StatusValue3", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue3() string {
	return r.StatusValue3
}
func (r *SetUserSecurityStatusRequest) SetStatusKey4(value string) {
	r.StatusKey4 = value
	r.QueryParams.Set("StatusKey4", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey4() string {
	return r.StatusKey4
}
func (r *SetUserSecurityStatusRequest) SetStatusValue4(value string) {
	r.StatusValue4 = value
	r.QueryParams.Set("StatusValue4", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue4() string {
	return r.StatusValue4
}
func (r *SetUserSecurityStatusRequest) SetStatusKey5(value string) {
	r.StatusKey5 = value
	r.QueryParams.Set("StatusKey5", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey5() string {
	return r.StatusKey5
}
func (r *SetUserSecurityStatusRequest) SetStatusValue5(value string) {
	r.StatusValue5 = value
	r.QueryParams.Set("StatusValue5", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue5() string {
	return r.StatusValue5
}
func (r *SetUserSecurityStatusRequest) SetStatusKey6(value string) {
	r.StatusKey6 = value
	r.QueryParams.Set("StatusKey6", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey6() string {
	return r.StatusKey6
}
func (r *SetUserSecurityStatusRequest) SetStatusValue6(value string) {
	r.StatusValue6 = value
	r.QueryParams.Set("StatusValue6", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue6() string {
	return r.StatusValue6
}
func (r *SetUserSecurityStatusRequest) SetStatusKey7(value string) {
	r.StatusKey7 = value
	r.QueryParams.Set("StatusKey7", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey7() string {
	return r.StatusKey7
}
func (r *SetUserSecurityStatusRequest) SetStatusValue7(value string) {
	r.StatusValue7 = value
	r.QueryParams.Set("StatusValue7", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue7() string {
	return r.StatusValue7
}
func (r *SetUserSecurityStatusRequest) SetStatusKey8(value string) {
	r.StatusKey8 = value
	r.QueryParams.Set("StatusKey8", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey8() string {
	return r.StatusKey8
}
func (r *SetUserSecurityStatusRequest) SetStatusValue8(value string) {
	r.StatusValue8 = value
	r.QueryParams.Set("StatusValue8", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue8() string {
	return r.StatusValue8
}
func (r *SetUserSecurityStatusRequest) SetStatusKey9(value string) {
	r.StatusKey9 = value
	r.QueryParams.Set("StatusKey9", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey9() string {
	return r.StatusKey9
}
func (r *SetUserSecurityStatusRequest) SetStatusValue9(value string) {
	r.StatusValue9 = value
	r.QueryParams.Set("StatusValue9", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue9() string {
	return r.StatusValue9
}
func (r *SetUserSecurityStatusRequest) SetStatusKey10(value string) {
	r.StatusKey10 = value
	r.QueryParams.Set("StatusKey10", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusKey10() string {
	return r.StatusKey10
}
func (r *SetUserSecurityStatusRequest) SetStatusValue10(value string) {
	r.StatusValue10 = value
	r.QueryParams.Set("StatusValue10", value)
}
func (r *SetUserSecurityStatusRequest) GetStatusValue10() string {
	return r.StatusValue10
}
func (r *SetUserSecurityStatusRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *SetUserSecurityStatusRequest) GetPassword() string {
	return r.Password
}

func (r *SetUserSecurityStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetUserSecurityStatus")
	r.SetProduct(Product)
}

type SetUserSecurityStatusResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func SetUserSecurityStatus(req *SetUserSecurityStatusRequest, accessId, accessSecret string) (*SetUserSecurityStatusResponse, error) {
	var pResponse SetUserSecurityStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
