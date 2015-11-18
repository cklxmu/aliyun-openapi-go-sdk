package pts

import (
	. "aliyun-openapi-go-sdk/core"
)

type SendWangWangRequest struct {
	RpcRequest
	To    string
	Title string
	Msg   string
}

func (r *SendWangWangRequest) SetTo(value string) {
	r.To = value
	r.QueryParams.Set("To", value)
}
func (r *SendWangWangRequest) GetTo() string {
	return r.To
}
func (r *SendWangWangRequest) SetTitle(value string) {
	r.Title = value
	r.QueryParams.Set("Title", value)
}
func (r *SendWangWangRequest) GetTitle() string {
	return r.Title
}
func (r *SendWangWangRequest) SetMsg(value string) {
	r.Msg = value
	r.QueryParams.Set("Msg", value)
}
func (r *SendWangWangRequest) GetMsg() string {
	return r.Msg
}

func (r *SendWangWangRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SendWangWang")
	r.SetProduct(Product)
}

type SendWangWangResponse struct {
}

func SendWangWang(req *SendWangWangRequest, accessId, accessSecret string) (*SendWangWangResponse, error) {
	var pResponse SendWangWangResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
