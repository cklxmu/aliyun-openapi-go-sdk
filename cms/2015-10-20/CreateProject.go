package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateProjectRequest struct {
	RpcRequest
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
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateProject")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateProjectResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
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
