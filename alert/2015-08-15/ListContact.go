package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListContactRequest struct {
	RoaRequest
	ProjectName string
	ContactName string
	Page        int
	PageSize    int
}

func (r *ListContactRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListContactRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListContactRequest) SetContactName(value string) {
	r.ContactName = value
	r.QueryParams.Set("ContactName", value)
}
func (r *ListContactRequest) GetContactName() string {
	return r.ContactName
}
func (r *ListContactRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListContactRequest) GetPage() int {
	return r.Page
}
func (r *ListContactRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListContactRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListContactRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/contacts"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListContactResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListContact(req *ListContactRequest, accessId, accessSecret string) (*ListContactResponse, error) {
	var pResponse ListContactResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
