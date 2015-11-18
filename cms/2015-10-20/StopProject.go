package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type StopProjectRequest struct {
	RpcRequest
	ProjectName string
}

func (r *StopProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *StopProjectRequest) GetProjectName() string {
	return r.ProjectName
}

func (r *StopProjectRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StopProject")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type StopProjectResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func StopProject(req *StopProjectRequest, accessId, accessSecret string) (*StopProjectResponse, error) {
	var pResponse StopProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
