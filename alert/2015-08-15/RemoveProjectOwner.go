package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type RemoveProjectOwnerRequest struct {
	RoaRequest
	ProjectName string
	Owners      string
}

func (r *RemoveProjectOwnerRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *RemoveProjectOwnerRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *RemoveProjectOwnerRequest) SetOwners(value string) {
	r.Owners = value
	r.QueryParams.Set("Owners", value)
}
func (r *RemoveProjectOwnerRequest) GetOwners() string {
	return r.Owners
}

func (r *RemoveProjectOwnerRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/removeOwner"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type RemoveProjectOwnerResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func RemoveProjectOwner(req *RemoveProjectOwnerRequest, accessId, accessSecret string) (*RemoveProjectOwnerResponse, error) {
	var pResponse RemoveProjectOwnerResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
