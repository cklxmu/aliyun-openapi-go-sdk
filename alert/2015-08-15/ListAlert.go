package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListAlertRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
	Page        int
	PageSize    int
}

func (r *ListAlertRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListAlertRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListAlertRequest) SetAlertName(value string) {
	r.AlertName = value
	r.QueryParams.Set("AlertName", value)
}
func (r *ListAlertRequest) GetAlertName() string {
	return r.AlertName
}
func (r *ListAlertRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListAlertRequest) GetPage() int {
	return r.Page
}
func (r *ListAlertRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListAlertRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListAlertRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListAlertResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListAlert(req *ListAlertRequest, accessId, accessSecret string) (*ListAlertResponse, error) {
	var pResponse ListAlertResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
