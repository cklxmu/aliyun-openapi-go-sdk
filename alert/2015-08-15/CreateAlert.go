package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateAlertRequest struct {
	RoaRequest
	ProjectName string
	Alert       string
}

func (r *CreateAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateAlertRequest) SetAlert(value string) {
	r.Alert = value
	r.QueryParams.Set("Alert", value)
}
func (r *CreateAlertRequest) GetAlert() string {
	return r.Alert
}

func (r *CreateAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateAlertResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateAlert(req *CreateAlertRequest, accessId, accessSecret string) (*CreateAlertResponse, error) {
	var pResponse CreateAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
