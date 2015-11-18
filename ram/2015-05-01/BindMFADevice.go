package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type BindMFADeviceRequest struct {
	RpcRequest
	SerialNumber        string
	UserName            string
	AuthenticationCode1 string
	AuthenticationCode2 string
}

func (r *BindMFADeviceRequest) SetSerialNumber(value string) {
	r.SerialNumber = value
	r.QueryParams.Set("SerialNumber", value)
}
func (r *BindMFADeviceRequest) GetSerialNumber() string {
	return r.SerialNumber
}
func (r *BindMFADeviceRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *BindMFADeviceRequest) GetUserName() string {
	return r.UserName
}
func (r *BindMFADeviceRequest) SetAuthenticationCode1(value string) {
	r.AuthenticationCode1 = value
	r.QueryParams.Set("AuthenticationCode1", value)
}
func (r *BindMFADeviceRequest) GetAuthenticationCode1() string {
	return r.AuthenticationCode1
}
func (r *BindMFADeviceRequest) SetAuthenticationCode2(value string) {
	r.AuthenticationCode2 = value
	r.QueryParams.Set("AuthenticationCode2", value)
}
func (r *BindMFADeviceRequest) GetAuthenticationCode2() string {
	return r.AuthenticationCode2
}

func (r *BindMFADeviceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BindMFADevice")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type BindMFADeviceResponse struct {
}

func BindMFADevice(req *BindMFADeviceRequest, accessId, accessSecret string) (*BindMFADeviceResponse, error) {
	var pResponse BindMFADeviceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
