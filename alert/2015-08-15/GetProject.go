package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetProjectRequest struct {
	RoaRequest
	ProjectName string
}

func (r *GetProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetProjectRequest) GetProjectName() string {
	return r.ProjectName
}

func (r *GetProjectRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetProjectResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetProject(req *GetProjectRequest, accessId, accessSecret string) (*GetProjectResponse, error) {
	var pResponse GetProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
