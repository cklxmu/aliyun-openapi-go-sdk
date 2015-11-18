package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteProjectRequest struct {
	RpcRequest
	ProjectName string
}

func (r *DeleteProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *DeleteProjectRequest) GetProjectName() string {
	return r.ProjectName
}

func (r *DeleteProjectRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteProject")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteProjectResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
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
