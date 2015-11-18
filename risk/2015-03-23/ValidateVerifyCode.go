package risk

import (
	. "aliyun-openapi-go-sdk/core"
)

type ValidateVerifyCodeRequest struct {
	RpcRequest
	RequestId   string
	MteeCode    string
	CodeType    string
	IdType      string
	UserId      string
	ChannelType string
	VerifyCode  string
	UmidToken   string
	Collina     string
	Ip          string
}

func (r *ValidateVerifyCodeRequest) SetRequestId(value string) {
	r.RequestId = value
	r.QueryParams.Set("RequestId", value)
}
func (r *ValidateVerifyCodeRequest) GetRequestId() string {
	return r.RequestId
}
func (r *ValidateVerifyCodeRequest) SetMteeCode(value string) {
	r.MteeCode = value
	r.QueryParams.Set("MteeCode", value)
}
func (r *ValidateVerifyCodeRequest) GetMteeCode() string {
	return r.MteeCode
}
func (r *ValidateVerifyCodeRequest) SetCodeType(value string) {
	r.CodeType = value
	r.QueryParams.Set("CodeType", value)
}
func (r *ValidateVerifyCodeRequest) GetCodeType() string {
	return r.CodeType
}
func (r *ValidateVerifyCodeRequest) SetIdType(value string) {
	r.IdType = value
	r.QueryParams.Set("IdType", value)
}
func (r *ValidateVerifyCodeRequest) GetIdType() string {
	return r.IdType
}
func (r *ValidateVerifyCodeRequest) SetUserId(value string) {
	r.UserId = value
	r.QueryParams.Set("UserId", value)
}
func (r *ValidateVerifyCodeRequest) GetUserId() string {
	return r.UserId
}
func (r *ValidateVerifyCodeRequest) SetChannelType(value string) {
	r.ChannelType = value
	r.QueryParams.Set("ChannelType", value)
}
func (r *ValidateVerifyCodeRequest) GetChannelType() string {
	return r.ChannelType
}
func (r *ValidateVerifyCodeRequest) SetVerifyCode(value string) {
	r.VerifyCode = value
	r.QueryParams.Set("VerifyCode", value)
}
func (r *ValidateVerifyCodeRequest) GetVerifyCode() string {
	return r.VerifyCode
}
func (r *ValidateVerifyCodeRequest) SetUmidToken(value string) {
	r.UmidToken = value
	r.QueryParams.Set("UmidToken", value)
}
func (r *ValidateVerifyCodeRequest) GetUmidToken() string {
	return r.UmidToken
}
func (r *ValidateVerifyCodeRequest) SetCollina(value string) {
	r.Collina = value
	r.QueryParams.Set("Collina", value)
}
func (r *ValidateVerifyCodeRequest) GetCollina() string {
	return r.Collina
}
func (r *ValidateVerifyCodeRequest) SetIp(value string) {
	r.Ip = value
	r.QueryParams.Set("Ip", value)
}
func (r *ValidateVerifyCodeRequest) GetIp() string {
	return r.Ip
}

func (r *ValidateVerifyCodeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ValidateVerifyCode")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ValidateVerifyCodeResponse struct {
	Code string `xml:"Code" json:"Code"`
}

func ValidateVerifyCode(req *ValidateVerifyCodeRequest, accessId, accessSecret string) (*ValidateVerifyCodeResponse, error) {
	var pResponse ValidateVerifyCodeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
