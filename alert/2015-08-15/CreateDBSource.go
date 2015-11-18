package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateDBSourceRequest struct {
	RoaRequest
	ProjectName string
	Source      string
}

func (r *CreateDBSourceRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateDBSourceRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateDBSourceRequest) SetSource(value string) {
	r.Source = value
	r.QueryParams.Set("Source", value)
}
func (r *CreateDBSourceRequest) GetSource() string {
	return r.Source
}

func (r *CreateDBSourceRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/sources"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateDBSourceResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateDBSource(req *CreateDBSourceRequest, accessId, accessSecret string) (*CreateDBSourceResponse, error) {
	var pResponse CreateDBSourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
