package risk

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SendVerifyCodeRequest struct {
	RpcRequest
	RequestId         string
	MteeCode          string
	CodeType          string
	IdType            string
	UserId            string
	ChannelType       string
	BizId             string
	EventId           string
	MessageReiver     string
	TimeInterval      int
	MessageParameters string
	Extend            string
}

func (r *SendVerifyCodeRequest) SetRequestId(value string) {
	r.RequestId = value
	r.QueryParams.Set("RequestId", value)
}
func (r *SendVerifyCodeRequest) GetRequestId() string {
	return r.RequestId
}
func (r *SendVerifyCodeRequest) SetMteeCode(value string) {
	r.MteeCode = value
	r.QueryParams.Set("MteeCode", value)
}
func (r *SendVerifyCodeRequest) GetMteeCode() string {
	return r.MteeCode
}
func (r *SendVerifyCodeRequest) SetCodeType(value string) {
	r.CodeType = value
	r.QueryParams.Set("CodeType", value)
}
func (r *SendVerifyCodeRequest) GetCodeType() string {
	return r.CodeType
}
func (r *SendVerifyCodeRequest) SetIdType(value string) {
	r.IdType = value
	r.QueryParams.Set("IdType", value)
}
func (r *SendVerifyCodeRequest) GetIdType() string {
	return r.IdType
}
func (r *SendVerifyCodeRequest) SetUserId(value string) {
	r.UserId = value
	r.QueryParams.Set("UserId", value)
}
func (r *SendVerifyCodeRequest) GetUserId() string {
	return r.UserId
}
func (r *SendVerifyCodeRequest) SetChannelType(value string) {
	r.ChannelType = value
	r.QueryParams.Set("ChannelType", value)
}
func (r *SendVerifyCodeRequest) GetChannelType() string {
	return r.ChannelType
}
func (r *SendVerifyCodeRequest) SetBizId(value string) {
	r.BizId = value
	r.QueryParams.Set("BizId", value)
}
func (r *SendVerifyCodeRequest) GetBizId() string {
	return r.BizId
}
func (r *SendVerifyCodeRequest) SetEventId(value string) {
	r.EventId = value
	r.QueryParams.Set("EventId", value)
}
func (r *SendVerifyCodeRequest) GetEventId() string {
	return r.EventId
}
func (r *SendVerifyCodeRequest) SetMessageReiver(value string) {
	r.MessageReiver = value
	r.QueryParams.Set("MessageReiver", value)
}
func (r *SendVerifyCodeRequest) GetMessageReiver() string {
	return r.MessageReiver
}
func (r *SendVerifyCodeRequest) SetTimeInterval(value int) {
	r.TimeInterval = value
	r.QueryParams.Set("TimeInterval", strconv.Itoa(value))
}
func (r *SendVerifyCodeRequest) GetTimeInterval() int {
	return r.TimeInterval
}
func (r *SendVerifyCodeRequest) SetMessageParameters(value string) {
	r.MessageParameters = value
	r.QueryParams.Set("MessageParameters", value)
}
func (r *SendVerifyCodeRequest) GetMessageParameters() string {
	return r.MessageParameters
}
func (r *SendVerifyCodeRequest) SetExtend(value string) {
	r.Extend = value
	r.QueryParams.Set("Extend", value)
}
func (r *SendVerifyCodeRequest) GetExtend() string {
	return r.Extend
}

func (r *SendVerifyCodeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SendVerifyCode")
	r.SetProduct(Product)
}

type SendVerifyCodeResponse struct {
	Code string `xml:"Code" json:"Code"`
}

func SendVerifyCode(req *SendVerifyCodeRequest, accessId, accessSecret string) (*SendVerifyCodeResponse, error) {
	var pResponse SendVerifyCodeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
