package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type UnbindMFADeviceRequest struct {
	RpcRequest
	UserName string
}

func (r *UnbindMFADeviceRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *UnbindMFADeviceRequest) GetUserName() string {
	return r.UserName
}

func (r *UnbindMFADeviceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UnbindMFADevice")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type UnbindMFADeviceResponse struct {
}

func UnbindMFADevice(req *UnbindMFADeviceRequest, accessId, accessSecret string) (*UnbindMFADeviceResponse, error) {
	var pResponse UnbindMFADeviceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
