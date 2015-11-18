package risk

import (
	. "aliyun-openapi-go-sdk/core"
)

type WriteUsscRequest struct {
	RpcRequest
	AppKey       string
	SignTime     string
	Sign         string
	MteeCode     string
	CodeType     string
	IdType       string
	UserId       string
	ChannelType  string
	VerifyResult string
	UmidToken    string
	Collina      string
	Ip           string
	Extend       string
}

func (r *WriteUsscRequest) SetAppKey(value string) {
	r.AppKey = value
	r.QueryParams.Set("AppKey", value)
}
func (r *WriteUsscRequest) GetAppKey() string {
	return r.AppKey
}
func (r *WriteUsscRequest) SetSignTime(value string) {
	r.SignTime = value
	r.QueryParams.Set("SignTime", value)
}
func (r *WriteUsscRequest) GetSignTime() string {
	return r.SignTime
}
func (r *WriteUsscRequest) SetSign(value string) {
	r.Sign = value
	r.QueryParams.Set("Sign", value)
}
func (r *WriteUsscRequest) GetSign() string {
	return r.Sign
}
func (r *WriteUsscRequest) SetMteeCode(value string) {
	r.MteeCode = value
	r.QueryParams.Set("MteeCode", value)
}
func (r *WriteUsscRequest) GetMteeCode() string {
	return r.MteeCode
}
func (r *WriteUsscRequest) SetCodeType(value string) {
	r.CodeType = value
	r.QueryParams.Set("CodeType", value)
}
func (r *WriteUsscRequest) GetCodeType() string {
	return r.CodeType
}
func (r *WriteUsscRequest) SetIdType(value string) {
	r.IdType = value
	r.QueryParams.Set("IdType", value)
}
func (r *WriteUsscRequest) GetIdType() string {
	return r.IdType
}
func (r *WriteUsscRequest) SetUserId(value string) {
	r.UserId = value
	r.QueryParams.Set("UserId", value)
}
func (r *WriteUsscRequest) GetUserId() string {
	return r.UserId
}
func (r *WriteUsscRequest) SetChannelType(value string) {
	r.ChannelType = value
	r.QueryParams.Set("ChannelType", value)
}
func (r *WriteUsscRequest) GetChannelType() string {
	return r.ChannelType
}
func (r *WriteUsscRequest) SetVerifyResult(value string) {
	r.VerifyResult = value
	r.QueryParams.Set("VerifyResult", value)
}
func (r *WriteUsscRequest) GetVerifyResult() string {
	return r.VerifyResult
}
func (r *WriteUsscRequest) SetUmidToken(value string) {
	r.UmidToken = value
	r.QueryParams.Set("UmidToken", value)
}
func (r *WriteUsscRequest) GetUmidToken() string {
	return r.UmidToken
}
func (r *WriteUsscRequest) SetCollina(value string) {
	r.Collina = value
	r.QueryParams.Set("Collina", value)
}
func (r *WriteUsscRequest) GetCollina() string {
	return r.Collina
}
func (r *WriteUsscRequest) SetIp(value string) {
	r.Ip = value
	r.QueryParams.Set("Ip", value)
}
func (r *WriteUsscRequest) GetIp() string {
	return r.Ip
}
func (r *WriteUsscRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *WriteUsscRequest) GetExtend() string {
	return r.Extend
}

func (r *WriteUsscRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("WriteUssc")
	r.SetProduct(Product)
}

type WriteUsscResponse struct {
	Code string `xml:"Code" json:"Code"`
}

func WriteUssc(req *WriteUsscRequest, accessId, accessSecret string) (*WriteUsscResponse, error) {
	var pResponse WriteUsscResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
