package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetDBSourceRequest struct {
	RoaRequest
	ProjectName string
	SourceName  string
}

func (r *GetDBSourceRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetDBSourceRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetDBSourceRequest) SetSourceName(value string) {
	r.SourceName = value
	r.PathParams.Set("SourceName", value)
}
func (r *GetDBSourceRequest) GetSourceName() string {
	return r.SourceName
}

func (r *GetDBSourceRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/sources/SourceName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetDBSourceResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetDBSource(req *GetDBSourceRequest, accessId, accessSecret string) (*GetDBSourceResponse, error) {
	var pResponse GetDBSourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
