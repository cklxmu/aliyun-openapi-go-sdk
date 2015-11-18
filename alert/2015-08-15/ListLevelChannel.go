package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListLevelChannelRequest struct {
	RoaRequest
	ProjectName string
	Level       string
	Page        int
	PageSize    int
}

func (r *ListLevelChannelRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *ListLevelChannelRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *ListLevelChannelRequest) SetLevel(value string) {
	r.Level = value
	r.QueryParams.Set("Level", value)
}
func (r *ListLevelChannelRequest) GetLevel() string {
	return r.Level
}
func (r *ListLevelChannelRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListLevelChannelRequest) GetPage() int {
	return r.Page
}
func (r *ListLevelChannelRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListLevelChannelRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListLevelChannelRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/level_channels"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListLevelChannelResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func ListLevelChannel(req *ListLevelChannelRequest, accessId, accessSecret string) (*ListLevelChannelResponse, error) {
	var pResponse ListLevelChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
