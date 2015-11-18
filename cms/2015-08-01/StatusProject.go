package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type StatusProjectRequest struct {
	RpcRequest
	ProjectName string
}

func (r *StatusProjectRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *StatusProjectRequest) GetProjectName() string {
	return r.ProjectName
}

func (r *StatusProjectRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("StatusProject")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type StatusProjectResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func StatusProject(req *StatusProjectRequest, accessId, accessSecret string) (*StatusProjectResponse, error) {
	var pResponse StatusProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
