package ubsms_inner

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetUserBusinessStatusesRequest struct {
	RpcRequest
	Uid           string
	ServiceCode   string
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

func (r *SetUserBusinessStatusesRequest) SetUid(value string) {
	r.Uid = value
	r.QueryParams.Set("Uid", value)
}
func (r *SetUserBusinessStatusesRequest) GetUid() string {
	return r.Uid
}
func (r *SetUserBusinessStatusesRequest) SetServiceCode(value string) {
	r.ServiceCode = value
	r.QueryParams.Set("ServiceCode", value)
}
func (r *SetUserBusinessStatusesRequest) GetServiceCode() string {
	return r.ServiceCode
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey1(value string) {
	r.StatusKey1 = value
	r.QueryParams.Set("StatusKey1", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey1() string {
	return r.StatusKey1
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue1(value string) {
	r.StatusValue1 = value
	r.QueryParams.Set("StatusValue1", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue1() string {
	return r.StatusValue1
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey2(value string) {
	r.StatusKey2 = value
	r.QueryParams.Set("StatusKey2", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey2() string {
	return r.StatusKey2
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue2(value string) {
	r.StatusValue2 = value
	r.QueryParams.Set("StatusValue2", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue2() string {
	return r.StatusValue2
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey3(value string) {
	r.StatusKey3 = value
	r.QueryParams.Set("StatusKey3", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey3() string {
	return r.StatusKey3
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue3(value string) {
	r.StatusValue3 = value
	r.QueryParams.Set("StatusValue3", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue3() string {
	return r.StatusValue3
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey4(value string) {
	r.StatusKey4 = value
	r.QueryParams.Set("StatusKey4", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey4() string {
	return r.StatusKey4
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue4(value string) {
	r.StatusValue4 = value
	r.QueryParams.Set("StatusValue4", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue4() string {
	return r.StatusValue4
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey5(value string) {
	r.StatusKey5 = value
	r.QueryParams.Set("StatusKey5", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey5() string {
	return r.StatusKey5
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue5(value string) {
	r.StatusValue5 = value
	r.QueryParams.Set("StatusValue5", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue5() string {
	return r.StatusValue5
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey6(value string) {
	r.StatusKey6 = value
	r.QueryParams.Set("StatusKey6", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey6() string {
	return r.StatusKey6
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue6(value string) {
	r.StatusValue6 = value
	r.QueryParams.Set("StatusValue6", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue6() string {
	return r.StatusValue6
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey7(value string) {
	r.StatusKey7 = value
	r.QueryParams.Set("StatusKey7", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey7() string {
	return r.StatusKey7
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue7(value string) {
	r.StatusValue7 = value
	r.QueryParams.Set("StatusValue7", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue7() string {
	return r.StatusValue7
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey8(value string) {
	r.StatusKey8 = value
	r.QueryParams.Set("StatusKey8", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey8() string {
	return r.StatusKey8
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue8(value string) {
	r.StatusValue8 = value
	r.QueryParams.Set("StatusValue8", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue8() string {
	return r.StatusValue8
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey9(value string) {
	r.StatusKey9 = value
	r.QueryParams.Set("StatusKey9", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey9() string {
	return r.StatusKey9
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue9(value string) {
	r.StatusValue9 = value
	r.QueryParams.Set("StatusValue9", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue9() string {
	return r.StatusValue9
}
func (r *SetUserBusinessStatusesRequest) SetStatusKey10(value string) {
	r.StatusKey10 = value
	r.QueryParams.Set("StatusKey10", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusKey10() string {
	return r.StatusKey10
}
func (r *SetUserBusinessStatusesRequest) SetStatusValue10(value string) {
	r.StatusValue10 = value
	r.QueryParams.Set("StatusValue10", value)
}
func (r *SetUserBusinessStatusesRequest) GetStatusValue10() string {
	return r.StatusValue10
}
func (r *SetUserBusinessStatusesRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *SetUserBusinessStatusesRequest) GetPassword() string {
	return r.Password
}

func (r *SetUserBusinessStatusesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetUserBusinessStatuses")
	r.SetProduct(Product)
}

type SetUserBusinessStatusesResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func SetUserBusinessStatuses(req *SetUserBusinessStatusesRequest, accessId, accessSecret string) (*SetUserBusinessStatusesResponse, error) {
	var pResponse SetUserBusinessStatusesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
