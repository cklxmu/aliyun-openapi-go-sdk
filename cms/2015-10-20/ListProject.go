package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListProjectRequest struct {
	RpcRequest
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
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListProject")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListProjectResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
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
