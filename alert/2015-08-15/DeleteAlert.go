package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteAlertRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
}

func (r *DeleteAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *DeleteAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteAlertRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *DeleteAlertRequest) GetAlertName() string {
	return r.AlertName
}

func (r *DeleteAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName"
	r.SetMethod("DELETE")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteAlertResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
}

func DeleteAlert(req *DeleteAlertRequest, accessId, accessSecret string) (*DeleteAlertResponse, error) {
	var pResponse DeleteAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
