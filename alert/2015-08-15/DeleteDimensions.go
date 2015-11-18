package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteDimensionsRequest struct {
	RoaRequest
	ProjectName  string
	AlertName    string
	DimensionsId string
}

func (r *DeleteDimensionsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteDimensionsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteDimensionsRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *DeleteDimensionsRequest) GetAlertName() string {
	return r.AlertName
}
func (r *DeleteDimensionsRequest) SetDimensionsId(value string) {
	r.DimensionsId = value
	r.PathParams.Set("DimensionsId", value)
}
func (r *DeleteDimensionsRequest) GetDimensionsId() string {
	return r.DimensionsId
}

func (r *DeleteDimensionsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/dimensions/DimensionsId"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteDimensionsResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteDimensions(req *DeleteDimensionsRequest, accessId, accessSecret string) (*DeleteDimensionsResponse, error) {
	var pResponse DeleteDimensionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
