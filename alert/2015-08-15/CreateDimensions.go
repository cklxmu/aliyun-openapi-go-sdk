package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateDimensionsRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
	Dimensions  string
}

func (r *CreateDimensionsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateDimensionsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateDimensionsRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *CreateDimensionsRequest) GetAlertName() string {
	return r.AlertName
}
func (r *CreateDimensionsRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *CreateDimensionsRequest) GetDimensions() string {
	return r.Dimensions
}

func (r *CreateDimensionsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/dimensions"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateDimensionsResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateDimensions(req *CreateDimensionsRequest, accessId, accessSecret string) (*CreateDimensionsResponse, error) {
	var pResponse CreateDimensionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
