package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetDimensionsRequest struct {
	RoaRequest
	ProjectName  string
	AlertName    string
	DimensionsId string
}

func (r *GetDimensionsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetDimensionsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetDimensionsRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *GetDimensionsRequest) GetAlertName() string {
	return r.AlertName
}
func (r *GetDimensionsRequest) SetDimensionsId(value string) {
	r.DimensionsId = value
	r.PathParams.Set("DimensionsId", value)
}
func (r *GetDimensionsRequest) GetDimensionsId() string {
	return r.DimensionsId
}

func (r *GetDimensionsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/dimensions/DimensionsId"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetDimensionsResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetDimensions(req *GetDimensionsRequest, accessId, accessSecret string) (*GetDimensionsResponse, error) {
	var pResponse GetDimensionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
