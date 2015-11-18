package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateProjectRequest struct {
	RpcRequest
	ProjectName string
	Project     string
}

func (r *UpdateProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
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
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateProject")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateProjectResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
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
