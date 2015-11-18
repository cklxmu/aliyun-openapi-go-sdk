package push

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type RevertRpcRequest struct {
	RpcRequest
	AppId      int
	DeviceId   string
	RpcContent string
	TimeOut    int
}

func (r *RevertRpcRequest) SetAppId(value int) {
	r.AppId = value
	r.QueryParams.Set("AppId", strconv.Itoa(value))
}
func (r *RevertRpcRequest) GetAppId() int {
	return r.AppId
}
func (r *RevertRpcRequest) SetDeviceId(value string) {
	r.DeviceId = value
	r.QueryParams.Set("DeviceId", value)
}
func (r *RevertRpcRequest) GetDeviceId() string {
	return r.DeviceId
}
func (r *RevertRpcRequest) SetRpcContent(value string) {
	r.RpcContent = value
	r.QueryParams.Set("RpcContent", value)
}
func (r *RevertRpcRequest) GetRpcContent() string {
	return r.RpcContent
}
func (r *RevertRpcRequest) SetTimeOut(value int) {
	r.TimeOut = value
	r.QueryParams.Set("TimeOut", strconv.Itoa(value))
}
func (r *RevertRpcRequest) GetTimeOut() int {
	return r.TimeOut
}

func (r *RevertRpcRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("RevertRpc")
	r.SetProduct(Product)
}

type RevertRpcResponse struct {
	RpcCode         string `xml:"RpcCode" json:"RpcCode"`
	ResponseContent string `xml:"ResponseContent" json:"ResponseContent"`
}

func RevertRpc(req *RevertRpcRequest, accessId, accessSecret string) (*RevertRpcResponse, error) {
	var pResponse RevertRpcResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
