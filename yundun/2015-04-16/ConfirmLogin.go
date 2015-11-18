package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type ConfirmLoginRequest struct {
	RpcRequest
	InstanceId string
	SourceIp   string
	Time       string
}

func (r *ConfirmLoginRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ConfirmLoginRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ConfirmLoginRequest) SetSourceIp(value string) {
	r.SourceIp = value
	r.QueryParams.Set("SourceIp", value)
}
func (r *ConfirmLoginRequest) GetSourceIp() string {
	return r.SourceIp
}
func (r *ConfirmLoginRequest) SetTime(value string) {
	r.Time = value
	r.QueryParams.Set("Time", value)
}
func (r *ConfirmLoginRequest) GetTime() string {
	return r.Time
}

func (r *ConfirmLoginRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ConfirmLogin")
	r.SetProduct(Product)
}

type ConfirmLoginResponse struct {
}

func ConfirmLogin(req *ConfirmLoginRequest, accessId, accessSecret string) (*ConfirmLoginResponse, error) {
	var pResponse ConfirmLoginResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
