package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListProjectRequest struct {
	RoaRequest
	ProjectOwner string
	Page         int
	PageSize     int
}

func (r *ListProjectRequest) SetProjectOwner(value string) {
	r.ProjectOwner = value
	r.QueryParams.Set("ProjectOwner", value)
}
func (r *ListProjectRequest) GetProjectOwner() string {
	return r.ProjectOwner
}
func (r *ListProjectRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListProjectRequest) GetPage() int {
	return r.Page
}
func (r *ListProjectRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListProjectRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListProjectRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListProjectResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListProject(req *ListProjectRequest, accessId, accessSecret string) (*ListProjectResponse, error) {
	var pResponse ListProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
