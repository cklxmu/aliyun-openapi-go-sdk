package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListDBSourceRequest struct {
	RoaRequest
	ProjectName string
	SourceName  string
	Page        int
	PageSize    int
}

func (r *ListDBSourceRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListDBSourceRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListDBSourceRequest) SetSourceName(value string) {
	r.SourceName = value
	r.QueryParams.Set("SourceName", value)
}
func (r *ListDBSourceRequest) GetSourceName() string {
	return r.SourceName
}
func (r *ListDBSourceRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListDBSourceRequest) GetPage() int {
	return r.Page
}
func (r *ListDBSourceRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListDBSourceRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListDBSourceRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/sources"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListDBSourceResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListDBSource(req *ListDBSourceRequest, accessId, accessSecret string) (*ListDBSourceResponse, error) {
	var pResponse ListDBSourceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
