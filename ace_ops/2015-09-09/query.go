package ace_ops

import (
	. "aliyun-openapi-go-sdk/core"
)

type queryRequest struct {
	RpcRequest
	group       string
	name        string
	ip          string
	softversion string
}

func (r *queryRequest) Setgroup(value string) {
	r.group = value
	r.QueryParams.Set("group", value)
}
func (r *queryRequest) Getgroup() string {
	return r.group
}
func (r *queryRequest) Setname(value string) {
	r.name = value
	r.QueryParams.Set("name", value)
}
func (r *queryRequest) Getname() string {
	return r.name
}
func (r *queryRequest) Setip(value string) {
	r.ip = value
	r.QueryParams.Set("ip", value)
}
func (r *queryRequest) Getip() string {
	return r.ip
}
func (r *queryRequest) Setsoftversion(value string) {
	r.softversion = value
	r.QueryParams.Set("softversion", value)
}
func (r *queryRequest) Getsoftversion() string {
	return r.softversion
}

func (r *queryRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("query")
	r.SetProduct(Product)
}

type queryResponse struct {
	url         string `xml:"url" json:"url"`
	softversion string `xml:"softversion" json:"softversion"`
}

func query(req *queryRequest, accessId, accessSecret string) (*queryResponse, error) {
	var pResponse queryResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
