package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type UpdateAlertRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
	Alert       string
}

func (r *UpdateAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *UpdateAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateAlertRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *UpdateAlertRequest) GetAlertName() string {
	return r.AlertName
}
func (r *UpdateAlertRequest) SetAlert(value string) {
	r.Alert = value
	r.QueryParams.Set("Alert", value)
}
func (r *UpdateAlertRequest) GetAlert() string {
	return r.Alert
}

func (r *UpdateAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateAlertResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func UpdateAlert(req *UpdateAlertRequest, accessId, accessSecret string) (*UpdateAlertResponse, error) {
	var pResponse UpdateAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
