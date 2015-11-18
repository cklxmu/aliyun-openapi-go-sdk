package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateDimensionsRequest struct {
	RoaRequest
	ProjectName  string
	AlertName    string
	DimensionsId string
	Dimensions   string
}

func (r *UpdateDimensionsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateDimensionsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateDimensionsRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *UpdateDimensionsRequest) GetAlertName() string {
	return r.AlertName
}
func (r *UpdateDimensionsRequest) SetDimensionsId(value string) {
	r.DimensionsId = value
	r.PathParams.Set("DimensionsId", value)
}
func (r *UpdateDimensionsRequest) GetDimensionsId() string {
	return r.DimensionsId
}
func (r *UpdateDimensionsRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *UpdateDimensionsRequest) GetDimensions() string {
	return r.Dimensions
}

func (r *UpdateDimensionsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/dimensions/DimensionsId"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateDimensionsResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateDimensions(req *UpdateDimensionsRequest, accessId, accessSecret string) (*UpdateDimensionsResponse, error) {
	var pResponse UpdateDimensionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
