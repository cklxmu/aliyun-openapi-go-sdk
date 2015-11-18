package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type LogineventLogRequest struct {
	RpcRequest
	JstOwnerId int
	PageNumber int
	PageSize   int
	InstanceId string
	RecordType int
}

func (r *LogineventLogRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *LogineventLogRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *LogineventLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *LogineventLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *LogineventLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *LogineventLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *LogineventLogRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *LogineventLogRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *LogineventLogRequest) SetRecordType(value int) {
	r.RecordType = value
	r.QueryParams.Set("RecordType", strconv.Itoa(value))
}
func (r *LogineventLogRequest) GetRecordType() int {
	return r.RecordType
}

func (r *LogineventLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("LogineventLog")
	r.SetProduct(Product)
}

type LogineventLogResponse struct {
	StartTime  string `xml:"StartTime" json:"StartTime"`
	EndTime    string `xml:"EndTime" json:"EndTime"`
	PageNumber int    `xml:"PageNumber" json:"PageNumber"`
	PageSize   int    `xml:"PageSize" json:"PageSize"`
	TotalCount int    `xml:"TotalCount" json:"TotalCount"`
	LogList    struct {
		LoginEventLog []struct {
			BlockTimes int    `xml:"BlockTimes" json:"BlockTimes"`
			SourceIp   string `xml:"SourceIp" json:"SourceIp"`
			Status     int    `xml:"Status" json:"Status"`
			Time       string `xml:"Time" json:"Time"`
			Location   string `xml:"Location" json:"Location"`
		} `xml:"LoginEventLog" json:"LoginEventLog"`
	} `xml:"LogList" json:"LogList"`
}

func LogineventLog(req *LogineventLogRequest, accessId, accessSecret string) (*LogineventLogResponse, error) {
	var pResponse LogineventLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
