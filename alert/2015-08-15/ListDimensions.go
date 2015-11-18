package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListDimensionsRequest struct {
	RoaRequest
	ProjectName string
	AlertName   string
	Page        int
	PageSize    int
}

func (r *ListDimensionsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListDimensionsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListDimensionsRequest) SetAlertName(value string) {
	r.AlertName = value
	r.PathParams.Set("AlertName", value)
}
func (r *ListDimensionsRequest) GetAlertName() string {
	return r.AlertName
}
func (r *ListDimensionsRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListDimensionsRequest) GetPage() int {
	return r.Page
}
func (r *ListDimensionsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListDimensionsRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListDimensionsRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/alerts/AlertName/dimensions"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListDimensionsResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListDimensions(req *ListDimensionsRequest, accessId, accessSecret string) (*ListDimensionsResponse, error) {
	var pResponse ListDimensionsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
