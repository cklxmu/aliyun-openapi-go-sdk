package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetAlertRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
}

func (r *GetAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetAlertRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *GetAlertRequest) GetAlertName() string {
	return r.AlertName
}

func (r *GetAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetAlertResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetAlert(req *GetAlertRequest, accessId, accessSecret string) (*GetAlertResponse, error) {
	var pResponse GetAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
