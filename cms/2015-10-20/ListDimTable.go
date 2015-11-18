package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ListDimTableRequest struct {
	RpcRequest
	DimTableName string
	Page         int
	PageSize     int
}

func (r *ListDimTableRequest) SetDimTableName(value string) {
	r.DimTableName = value
	r.QueryParams.Set("DimTableName", value)
}
func (r *ListDimTableRequest) GetDimTableName() string {
	return r.DimTableName
}
func (r *ListDimTableRequest) SetPage(value int) {
	r.Page = value
	r.QueryParams.Set("Page", strconv.Itoa(value))
}
func (r *ListDimTableRequest) GetPage() int {
	return r.Page
}
func (r *ListDimTableRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *ListDimTableRequest) GetPageSize() int {
	return r.PageSize
}

func (r *ListDimTableRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListDimTable")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type ListDimTableResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	NextToken  int    `xml:"NextToken" json:"NextToken"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func ListDimTable(req *ListDimTableRequest, accessId, accessSecret string) (*ListDimTableResponse, error) {
	var pResponse ListDimTableResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
