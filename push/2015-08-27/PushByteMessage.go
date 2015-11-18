package push

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PushByteMessageRequest struct {
	RpcRequest
	AppId       int
	SendType    int
	Accounts    string
	DeviceIds   string
	PushContent string
}

func (r *PushByteMessageRequest) SetAppId(value int) {
	r.AppId = value
	r.QueryParams.Set("AppId", strconv.Itoa(value))
}
func (r *PushByteMessageRequest) GetAppId() int {
	return r.AppId
}
func (r *PushByteMessageRequest) SetSendType(value int) {
	r.SendType = value
	r.QueryParams.Set("SendType", strconv.Itoa(value))
}
func (r *PushByteMessageRequest) GetSendType() int {
	return r.SendType
}
func (r *PushByteMessageRequest) SetAccounts(value string) {
	r.Accounts = value
	r.QueryParams.Set("Accounts", value)
}
func (r *PushByteMessageRequest) GetAccounts() string {
	return r.Accounts
}
func (r *PushByteMessageRequest) SetDeviceIds(value string) {
	r.DeviceIds = value
	r.QueryParams.Set("DeviceIds", value)
}
func (r *PushByteMessageRequest) GetDeviceIds() string {
	return r.DeviceIds
}
func (r *PushByteMessageRequest) SetPushContent(value string) {
	r.PushContent = value
	r.QueryParams.Set("PushContent", value)
}
func (r *PushByteMessageRequest) GetPushContent() string {
	return r.PushContent
}

func (r *PushByteMessageRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PushByteMessage")
	r.SetProduct(Product)
}

type PushByteMessageResponse struct {
	ResponseId string `xml:"ResponseId" json:"ResponseId"`
}

func PushByteMessage(req *PushByteMessageRequest, accessId, accessSecret string) (*PushByteMessageResponse, error) {
	var pResponse PushByteMessageResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
