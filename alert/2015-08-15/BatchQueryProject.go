package alert

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BatchQueryProjectRequest struct {
	RoaRequest
	Names    string
	Page     int
	PageSize int
}

func (r *BatchQueryProjectRequest) SetNames(value string) {
	r.Names = value
	r.QueryParams.Set("Names", value)
}
func (r *BatchQueryProjectRequest) GetNames() string {
	return r.Names
}
func (r *BatchQueryProjectRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *BatchQueryProjectRequest) GetPage() int {
	return r.Page
}
func (r *BatchQueryProjectRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *BatchQueryProjectRequest) GetPageSize() int {
	return r.PageSize
}

func (r *BatchQueryProjectRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/batchQuery"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type BatchQueryProjectResponse struct {
	code       string `xml:"code" json:"code"`
	message    string `xml:"message" json:"message"`
	success    string `xml:"success" json:"success"`
	traceId    string `xml:"traceId" json:"traceId"`
	datapoints string `xml:"datapoints" json:"datapoints"`
	total      string `xml:"total" json:"total"`
}

func BatchQueryProject(req *BatchQueryProjectRequest, accessId, accessSecret string) (*BatchQueryProjectResponse, error) {
	var pResponse BatchQueryProjectResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
