package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type EnableAlertRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
}

func (r *EnableAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *EnableAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *EnableAlertRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *EnableAlertRequest) GetAlertName() string {
	return r.AlertName
}

func (r *EnableAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/enable"
	r.SetMethod(PUT)
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type EnableAlertResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func EnableAlert(req *EnableAlertRequest, accessId, accessSecret string) (*EnableAlertResponse, error) {
	var pResponse EnableAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
