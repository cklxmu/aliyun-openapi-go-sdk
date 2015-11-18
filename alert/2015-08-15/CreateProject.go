package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateProjectRequest struct {
	RoaRequest
	Project string
}

func (r *CreateProjectRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *CreateProjectRequest) GetProject() string {
	return r.Project
}

func (r *CreateProjectRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateProjectResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateProject(req *CreateProjectRequest, accessId, accessSecret string) (*CreateProjectResponse, error) {
	var pResponse CreateProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
