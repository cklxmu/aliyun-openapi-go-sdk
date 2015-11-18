package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListContactGroupRequest struct {
	RoaRequest
	ProjectName string
	GroupName   string
	Page        int
	PageSize    int
}

func (r *ListContactGroupRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListContactGroupRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListContactGroupRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *ListContactGroupRequest) GetGroupName() string {
	return r.GroupName
}
func (r *ListContactGroupRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListContactGroupRequest) GetPage() int {
	return r.Page
}
func (r *ListContactGroupRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListContactGroupRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListContactGroupRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/groups"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListContactGroupResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListContactGroup(req *ListContactGroupRequest, accessId, accessSecret string) (*ListContactGroupResponse, error) {
	var pResponse ListContactGroupResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
