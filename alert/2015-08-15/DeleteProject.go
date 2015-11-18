package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteProjectRequest struct {
	RoaRequest
	ProjectName string
}

func (r *DeleteProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteProjectRequest) GetProjectName() string {
	return r.ProjectName
}

func (r *DeleteProjectRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteProjectResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteProject(req *DeleteProjectRequest, accessId, accessSecret string) (*DeleteProjectResponse, error) {
	var pResponse DeleteProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
