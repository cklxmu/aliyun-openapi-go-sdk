package sts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type AssumeRoleRequest struct {
	RpcRequest
	DurationSeconds int
	Policy          string
	RoleArn         string
	RoleSessionName string
}

func (r *AssumeRoleRequest) SetDurationSeconds(value int) {
	r.DurationSeconds = value
	r.QueryParams.Set("DurationSeconds", strconv.Itoa(value))
}
func (r *AssumeRoleRequest) GetDurationSeconds() int {
	return r.DurationSeconds
}
func (r *AssumeRoleRequest) SetPolicy(value string) {
	r.Policy = value
	r.QueryParams.Set("Policy", value)
}
func (r *AssumeRoleRequest) GetPolicy() string {
	return r.Policy
}
func (r *AssumeRoleRequest) SetRoleArn(value string) {
	r.RoleArn = value
	r.QueryParams.Set("RoleArn", value)
}
func (r *AssumeRoleRequest) GetRoleArn() string {
	return r.RoleArn
}
func (r *AssumeRoleRequest) SetRoleSessionName(value string) {
	r.RoleSessionName = value
	r.QueryParams.Set("RoleSessionName", value)
}
func (r *AssumeRoleRequest) GetRoleSessionName() string {
	return r.RoleSessionName
}

func (r *AssumeRoleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AssumeRole")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type AssumeRoleResponse struct {
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

func AssumeRole(req *AssumeRoleRequest, accessId, accessSecret string) (*AssumeRoleResponse, error) {
	var pResponse AssumeRoleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
