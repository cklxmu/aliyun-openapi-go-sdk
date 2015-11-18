package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GrantProjectOwnerRequest struct {
	RoaRequest
	ProjectName string
	Owners      string
}

func (r *GrantProjectOwnerRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GrantProjectOwnerRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GrantProjectOwnerRequest) SetOwners(value string) {
	r.Owners = value
	r.QueryParams.Set("Owners", value)
}
func (r *GrantProjectOwnerRequest) GetOwners() string {
	return r.Owners
}

func (r *GrantProjectOwnerRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/grantOwner"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GrantProjectOwnerResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func GrantProjectOwner(req *GrantProjectOwnerRequest, accessId, accessSecret string) (*GrantProjectOwnerResponse, error) {
	var pResponse GrantProjectOwnerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
