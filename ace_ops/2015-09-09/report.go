package ace_ops

import (
	. "aliyun-openapi-go-sdk/core"
)

type reportRequest struct {
	RpcRequest
	group       string
	name        string
	ip          string
	Type        string
	softversion string
	config      string
}

func (r *reportRequest) Setgroup(value string) {
	r.group = value
	r.QueryParams.Set("group", value)
}
func (r *reportRequest) Getgroup() string {
	return r.group
}
func (r *reportRequest) Setname(value string) {
	r.name = value
	r.QueryParams.Set("name", value)
}
func (r *reportRequest) Getname() string {
	return r.name
}
func (r *reportRequest) Setip(value string) {
	r.ip = value
	r.QueryParams.Set("ip", value)
}
func (r *reportRequest) Getip() string {
	return r.ip
}
func (r *reportRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("type", value)
}
func (r *reportRequest) GetType() string {
	return r.Type
}
func (r *reportRequest) Setsoftversion(value string) {
	r.softversion = value
	r.QueryParams.Set("softversion", value)
}
func (r *reportRequest) Getsoftversion() string {
	return r.softversion
}
func (r *reportRequest) Setconfig(value string) {
	r.config = value
	r.QueryParams.Set("config", value)
}
func (r *reportRequest) Getconfig() string {
	return r.config
}

func (r *reportRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("report")
	r.SetProduct(Product)
}

type reportResponse struct {
}

func report(req *reportRequest, accessId, accessSecret string) (*reportResponse, error) {
	var pResponse reportResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
