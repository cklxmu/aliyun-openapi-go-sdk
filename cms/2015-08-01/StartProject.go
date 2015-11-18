package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type StartProjectRequest struct {
	RpcRequest
	ProjectName string
}

func (r *StartProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *StartProjectRequest) GetProjectName() string {
	return r.ProjectName
}

func (r *StartProjectRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StartProject")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type StartProjectResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func StartProject(req *StartProjectRequest, accessId, accessSecret string) (*StartProjectResponse, error) {
	var pResponse StartProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
