package ace

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeAppLogsRequest struct {
	RpcRequest
	AppId      string
	StartTime  string
	EndTime    string
	PageSize   int
	PageNumber int
}

func (r *DescribeAppLogsRequest) SetAppId(value string) {
	r.AppId = value
	r.QueryParams.Set("AppId", value)
}
func (r *DescribeAppLogsRequest) GetAppId() string {
	return r.AppId
}
func (r *DescribeAppLogsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeAppLogsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeAppLogsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeAppLogsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeAppLogsRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeAppLogsRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeAppLogsRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeAppLogsRequest) GetPageNumber() int {
	return r.PageNumber
}

func (r *DescribeAppLogsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeAppLogs")
	r.SetProduct(Product)
}

type DescribeAppLogsResponse struct {
	PageRecordCount int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items           struct {
		AppLog []struct {
			LogTime    string `xml:"LogTime" json:"LogTime"`
			LogContent string `xml:"LogContent" json:"LogContent"`
		} `xml:"AppLog" json:"AppLog"`
	} `xml:"Items" json:"Items"`
}

func DescribeAppLogs(req *DescribeAppLogsRequest, accessId, accessSecret string) (*DescribeAppLogsResponse, error) {
	var pResponse DescribeAppLogsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
