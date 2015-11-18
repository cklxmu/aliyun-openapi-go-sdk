package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DisableAlertRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
}

func (r *DisableAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DisableAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DisableAlertRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *DisableAlertRequest) GetAlertName() string {
	return r.AlertName
}

func (r *DisableAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/disable"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DisableAlertResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DisableAlert(req *DisableAlertRequest, accessId, accessSecret string) (*DisableAlertResponse, error) {
	var pResponse DisableAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
