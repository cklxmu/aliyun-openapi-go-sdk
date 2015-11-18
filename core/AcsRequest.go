package core

import (
	"net/http"
	"net/url"
)

type ApiRequest interface {
	GetQueryParams() url.Values
	GetParameter(string) string
	SetParameter(string, string)
	GetMethod() string
	SetMethod(string)
	GetProtocol() string
	SetProtocol(string)
	GetDomain() string
	SetDomain(string)
	GetProduct() string
	SetProduct(string)
	GetAction() string
	SetAction(string)
	GenSignature(accessId, accessSecret string) string
	DoAction(accessId, accessSecret string) (*http.Response, error)
}

type AcsRequest struct {
	QueryParams url.Values
	Method      string
	Protocol    string
	Domain      string
	Product     string
}

func (r *AcsRequest) GetQueryParams() url.Values {
	return r.QueryParams
}
func (r *AcsRequest) SetParameter(key, value string) {
	r.QueryParams.Set(key, value)
}
func (r *AcsRequest) GetParameter(key string) string {
	return r.QueryParams.Get(key)
}
func (r *AcsRequest) SetMethod(value string) {
	r.Method = value
}
func (r *AcsRequest) GetMethod() string {
	return r.Method
}
func (r *AcsRequest) SetProtocol(value string) {
	r.Protocol = value
}
func (r *AcsRequest) GetProtocol() string {
	return r.Protocol
}
func (r *AcsRequest) SetDomain(value string) {
	r.Domain = value
}
func (r *AcsRequest) GetDomain() string {
	return r.Domain
}
func (r *AcsRequest) SetProduct(value string) {
	r.Product = value
}
func (r *AcsRequest) GetProduct() string {
	return r.Product
}

func (r *AcsRequest) Init() {
	r.QueryParams = make(url.Values)
	r.SetFormat(FORMAT_XML)
	r.SetMethod(METHOD_GET)
	r.SetProtocol(PROTOCOL_HTTPS)
}
func (r *AcsRequest) SetRegionId(value string) {
	r.QueryParams.Set("RegionId", value)
}
func (r *AcsRequest) GetRegionId() string {
	return r.QueryParams.Get("RegionId")
}
func (r *AcsRequest) SetAction(value string) {
	r.QueryParams.Set("Action", value)
}
func (r *AcsRequest) GetAction() string {
	return r.QueryParams.Get("Action")
}
func (r *AcsRequest) SetFormat(value string) {
	r.QueryParams.Set("Format", value)
}
func (r *AcsRequest) GetFormat() string {
	return r.QueryParams.Get("Format")
}
func (r *AcsRequest) SetVersion(value string) {
	r.QueryParams.Set("Version", value)
}
func (r *AcsRequest) GetVersion() string {
	return r.QueryParams.Get("Version")
}
func (r *AcsRequest) SetSignatureVersion(value string) {
	r.QueryParams.Set("SignatureVersion", value)
}
func (r *AcsRequest) GetSignatureVersion() string {
	return r.QueryParams.Get("SignatureVersion")
}
func (r *AcsRequest) SetSignatureMethod(value string) {
	r.QueryParams.Set("SignatureMethod", value)
}
func (r *AcsRequest) GetSignatureMethod() string {
	return r.QueryParams.Get("SignatureMethod")
}
