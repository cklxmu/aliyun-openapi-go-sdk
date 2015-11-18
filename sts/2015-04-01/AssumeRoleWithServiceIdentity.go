package sts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AssumeRoleWithServiceIdentityRequest struct {
	RpcRequest
	DurationSeconds int
	Policy          string
	RoleArn         string
	RoleSessionName string
	AssumeRoleFor   string
}

func (r *AssumeRoleWithServiceIdentityRequest) SetDurationSeconds(value int) {
	r.DurationSeconds = value
	r.QueryParams.Set("DurationSeconds", strconv.Itoa(value))
}
func (r *AssumeRoleWithServiceIdentityRequest) GetDurationSeconds() int {
	return r.DurationSeconds
}
func (r *AssumeRoleWithServiceIdentityRequest) SetPolicy(value string) {
	r.Policy = value
	r.QueryParams.Set("Policy", value)
}
func (r *AssumeRoleWithServiceIdentityRequest) GetPolicy() string {
	return r.Policy
}
func (r *AssumeRoleWithServiceIdentityRequest) SetRoleArn(value string) {
	r.RoleArn = value
	r.QueryParams.Set("RoleArn", value)
}
func (r *AssumeRoleWithServiceIdentityRequest) GetRoleArn() string {
	return r.RoleArn
}
func (r *AssumeRoleWithServiceIdentityRequest) SetRoleSessionName(value string) {
	r.RoleSessionName = value
	r.QueryParams.Set("RoleSessionName", value)
}
func (r *AssumeRoleWithServiceIdentityRequest) GetRoleSessionName() string {
	return r.RoleSessionName
}
func (r *AssumeRoleWithServiceIdentityRequest) SetAssumeRoleFor(value string) {
	r.AssumeRoleFor = value
	r.QueryParams.Set("AssumeRoleFor", value)
}
func (r *AssumeRoleWithServiceIdentityRequest) GetAssumeRoleFor() string {
	return r.AssumeRoleFor
}

func (r *AssumeRoleWithServiceIdentityRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AssumeRoleWithServiceIdentity")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AssumeRoleWithServiceIdentityResponse struct {
	Credentials struct {
		SecurityToken   string `xml:"SecurityToken" json:"SecurityToken"`
		AccessKeySecret string `xml:"AccessKeySecret" json:"AccessKeySecret"`
		AccessKeyId     string `xml:"AccessKeyId" json:"AccessKeyId"`
		Expiration      string `xml:"Expiration" json:"Expiration"`
	} `xml:"Credentials" json:"Credentials"`
	AssumedRoleUser struct {
		Arn           string `xml:"Arn" json:"Arn"`
		AssumedRoleId string `xml:"AssumedRoleId" json:"AssumedRoleId"`
	} `xml:"AssumedRoleUser" json:"AssumedRoleUser"`
}

func AssumeRoleWithServiceIdentity(req *AssumeRoleWithServiceIdentityRequest, accessId, accessSecret string) (*AssumeRoleWithServiceIdentityResponse, error) {
	var pResponse AssumeRoleWithServiceIdentityResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
