package risk

import (
	. "aliyun-openapi-go-sdk/core"
)

type FindRiskRequest struct {
	RpcRequest
	MteeCode  string
	CodeType  string
	IdType    string
	UserId    string
	Collina   string
	UmidToken string
	Ip        string
	Email     string
	Phone     string
}

func (r *FindRiskRequest) SetMteeCode(value string) {
	r.MteeCode = value
	r.QueryParams.Set("MteeCode", value)
}
func (r *FindRiskRequest) GetMteeCode() string {
	return r.MteeCode
}
func (r *FindRiskRequest) SetCodeType(value string) {
	r.CodeType = value
	r.QueryParams.Set("CodeType", value)
}
func (r *FindRiskRequest) GetCodeType() string {
	return r.CodeType
}
func (r *FindRiskRequest) SetIdType(value string) {
	r.IdType = value
	r.QueryParams.Set("IdType", value)
}
func (r *FindRiskRequest) GetIdType() string {
	return r.IdType
}
func (r *FindRiskRequest) SetUserId(value string) {
	r.UserId = value
	r.QueryParams.Set("UserId", value)
}
func (r *FindRiskRequest) GetUserId() string {
	return r.UserId
}
func (r *FindRiskRequest) SetCollina(value string) {
	r.Collina = value
	r.QueryParams.Set("Collina", value)
}
func (r *FindRiskRequest) GetCollina() string {
	return r.Collina
}
func (r *FindRiskRequest) SetUmidToken(value string) {
	r.UmidToken = value
	r.QueryParams.Set("UmidToken", value)
}
func (r *FindRiskRequest) GetUmidToken() string {
	return r.UmidToken
}
func (r *FindRiskRequest) SetIp(value string) {
	r.Ip = value
	r.QueryParams.Set("Ip", value)
}
func (r *FindRiskRequest) GetIp() string {
	return r.Ip
}
func (r *FindRiskRequest) SetEmail(value string) {
	r.Email = value
	r.QueryParams.Set("Email", value)
}
func (r *FindRiskRequest) GetEmail() string {
	return r.Email
}
func (r *FindRiskRequest) SetPhone(value string) {
	r.Phone = value
	r.QueryParams.Set("Phone", value)
}
func (r *FindRiskRequest) GetPhone() string {
	return r.Phone
}

func (r *FindRiskRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("FindRisk")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type FindRiskResponse struct {
	Code string `xml:"Code" json:"Code"`
	Data struct {
		NoRisk  bool   `xml:"NoRisk" json:"NoRisk"`
		Action  string `xml:"Action" json:"Action"`
		Tag     string `xml:"Tag" json:"Tag"`
		Message string `xml:"Message" json:"Message"`
	} `xml:"Data" json:"Data"`
}

func FindRisk(req *FindRiskRequest, accessId, accessSecret string) (*FindRiskResponse, error) {
	var pResponse FindRiskResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
