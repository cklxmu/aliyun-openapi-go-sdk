package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeParameterTemplatesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	ClientToken          string
	Engine               string
	EngineVersion        string
	OwnerAccount         string
}

func (r *DescribeParameterTemplatesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeParameterTemplatesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeParameterTemplatesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeParameterTemplatesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeParameterTemplatesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeParameterTemplatesRequest) SetClientToken(value string) {
	r.ClientToken = value
	r.QueryParams.Set("ClientToken", value)
}
func (r *DescribeParameterTemplatesRequest) GetClientToken() string {
	return r.ClientToken
}
func (r *DescribeParameterTemplatesRequest) SetEngine(value string) {
	r.Engine = value
	r.QueryParams.Set("Engine", value)
}
func (r *DescribeParameterTemplatesRequest) GetEngine() string {
	return r.Engine
}
func (r *DescribeParameterTemplatesRequest) SetEngineVersion(value string) {
	r.EngineVersion = value
	r.QueryParams.Set("EngineVersion", value)
}
func (r *DescribeParameterTemplatesRequest) GetEngineVersion() string {
	return r.EngineVersion
}
func (r *DescribeParameterTemplatesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeParameterTemplatesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeParameterTemplatesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeParameterTemplates")
	r.SetProduct(Product)
}

type DescribeParameterTemplatesResponse struct {
	Engine         string `xml:"Engine" json:"Engine"`
	EngineVersion  string `xml:"EngineVersion" json:"EngineVersion"`
	ParameterCount string `xml:"ParameterCount" json:"ParameterCount"`
	Parameters     struct {
		TemplateRecord []struct {
			ParameterName        string `xml:"ParameterName" json:"ParameterName"`
			ParameterValue       string `xml:"ParameterValue" json:"ParameterValue"`
			ForceModify          string `xml:"ForceModify" json:"ForceModify"`
			ForceRestart         string `xml:"ForceRestart" json:"ForceRestart"`
			ForceModify          bool   `xml:"ForceModify" json:"ForceModify"`
			ForceRestart         bool   `xml:"ForceRestart" json:"ForceRestart"`
			CheckingCode         string `xml:"CheckingCode" json:"CheckingCode"`
			ParameterDescription string `xml:"ParameterDescription" json:"ParameterDescription"`
		} `xml:"TemplateRecord" json:"TemplateRecord"`
	} `xml:"Parameters" json:"Parameters"`
}

func DescribeParameterTemplates(req *DescribeParameterTemplatesRequest, accessId, accessSecret string) (*DescribeParameterTemplatesResponse, error) {
	var pResponse DescribeParameterTemplatesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
