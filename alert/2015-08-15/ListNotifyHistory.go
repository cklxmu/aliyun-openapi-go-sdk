package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListNotifyHistoryRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
	Dimensions  string
	StartTime   string
	EndTime     string
	Page        int
	PageSize    int
}

func (r *ListNotifyHistoryRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListNotifyHistoryRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListNotifyHistoryRequest) SetAlertName(value string) {
	r.AlertName = value
	r.QueryParams.Set("AlertName", value)
}
func (r *ListNotifyHistoryRequest) GetAlertName() string {
	return r.AlertName
}
func (r *ListNotifyHistoryRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *ListNotifyHistoryRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *ListNotifyHistoryRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *ListNotifyHistoryRequest) GetStartTime() string {
	return r.StartTime
}
func (r *ListNotifyHistoryRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *ListNotifyHistoryRequest) GetEndTime() string {
	return r.EndTime
}
func (r *ListNotifyHistoryRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListNotifyHistoryRequest) GetPage() int {
	return r.Page
}
func (r *ListNotifyHistoryRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListNotifyHistoryRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListNotifyHistoryRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/notify_history"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListNotifyHistoryResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListNotifyHistory(req *ListNotifyHistoryRequest, accessId, accessSecret string) (*ListNotifyHistoryResponse, error) {
	var pResponse ListNotifyHistoryResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
