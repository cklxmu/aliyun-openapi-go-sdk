package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateDBSourceRequest struct {
	RoaRequest
	ProjectName string
	SourceName  string
	Source      string
}

func (r *UpdateDBSourceRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateDBSourceRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateDBSourceRequest) SetSourceName(value string) {
	r.SourceName = value
	r.PathParams.Set("SourceName", value)
}
func (r *UpdateDBSourceRequest) GetSourceName() string {
	return r.SourceName
}
func (r *UpdateDBSourceRequest) SetSource(value string) {
	r.Source = value
	r.QueryParams.Set("Source", value)
}
func (r *UpdateDBSourceRequest) GetSource() string {
	return r.Source
}

func (r *UpdateDBSourceRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/sources/SourceName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateDBSourceResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateDBSource(req *UpdateDBSourceRequest, accessId, accessSecret string) (*UpdateDBSourceResponse, error) {
	var pResponse UpdateDBSourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
