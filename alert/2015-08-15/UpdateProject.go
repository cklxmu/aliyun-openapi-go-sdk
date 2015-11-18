package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateProjectRequest struct {
	RoaRequest
	ProjectName string
	Project     string
}

func (r *UpdateProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateProjectRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateProjectRequest) SetProject(value string) {
	r.Project = value
	r.QueryParams.Set("Project", value)
}
func (r *UpdateProjectRequest) GetProject() string {
	return r.Project
}

func (r *UpdateProjectRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateProjectResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateProject(req *UpdateProjectRequest, accessId, accessSecret string) (*UpdateProjectResponse, error) {
	var pResponse UpdateProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
