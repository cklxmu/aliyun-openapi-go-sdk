package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListAlertStateRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
	Dimensions  string
	StartTime   string
	EndTime     string
	Page        int
	PageSize    int
}

func (r *ListAlertStateRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListAlertStateRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListAlertStateRequest) SetAlertName(value string) {
	r.AlertName = value
	r.QueryParams.Set("AlertName", value)
}
func (r *ListAlertStateRequest) GetAlertName() string {
	return r.AlertName
}
func (r *ListAlertStateRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *ListAlertStateRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *ListAlertStateRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *ListAlertStateRequest) GetStartTime() string {
	return r.StartTime
}
func (r *ListAlertStateRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *ListAlertStateRequest) GetEndTime() string {
	return r.EndTime
}
func (r *ListAlertStateRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListAlertStateRequest) GetPage() int {
	return r.Page
}
func (r *ListAlertStateRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListAlertStateRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListAlertStateRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts_state"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListAlertStateResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListAlertState(req *ListAlertStateRequest, accessId, accessSecret string) (*ListAlertStateResponse, error) {
	var pResponse ListAlertStateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
