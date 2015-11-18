package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteDBSourceRequest struct {
	RoaRequest
	ProjectName string
	SourceName  string
}

func (r *DeleteDBSourceRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteDBSourceRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteDBSourceRequest) SetSourceName(value string) {
	r.SourceName = value
	r.PathParams.Set("SourceName", value)
}
func (r *DeleteDBSourceRequest) GetSourceName() string {
	return r.SourceName
}

func (r *DeleteDBSourceRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/sources/SourceName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteDBSourceResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteDBSource(req *DeleteDBSourceRequest, accessId, accessSecret string) (*DeleteDBSourceResponse, error) {
	var pResponse DeleteDBSourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
