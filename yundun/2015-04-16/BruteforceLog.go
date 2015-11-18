package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BruteforceLogRequest struct {
	RpcRequest
	JstOwnerId int
	PageNumber int
	PageSize   int
	InstanceId string
	RecordType int
}

func (r *BruteforceLogRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *BruteforceLogRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *BruteforceLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *BruteforceLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *BruteforceLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *BruteforceLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *BruteforceLogRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *BruteforceLogRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *BruteforceLogRequest) SetRecordType(value int) {
	r.RecordType = value
	r.QueryParams.Set("RecordType", strconv.Itoa(value))
}
func (r *BruteforceLogRequest) GetRecordType() int {
	return r.RecordType
}

func (r *BruteforceLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BruteforceLog")
	r.SetProduct(Product)
}

type BruteforceLogResponse struct {
	StartTime  string `xml:"StartTime" json:"StartTime"`
	EndTime    string `xml:"EndTime" json:"EndTime"`
	PageNumber int    `xml:"PageNumber" json:"PageNumber"`
	PageSize   int    `xml:"PageSize" json:"PageSize"`
	TotalCount int    `xml:"TotalCount" json:"TotalCount"`
	LogList    struct {
		BruteforceLog []struct {
			BlockTimes int    `xml:"BlockTimes" json:"BlockTimes"`
			SourceIp   string `xml:"SourceIp" json:"SourceIp"`
			Status     int    `xml:"Status" json:"Status"`
			Time       string `xml:"Time" json:"Time"`
			Location   string `xml:"Location" json:"Location"`
		} `xml:"BruteforceLog" json:"BruteforceLog"`
	} `xml:"LogList" json:"LogList"`
}

func BruteforceLog(req *BruteforceLogRequest, accessId, accessSecret string) (*BruteforceLogResponse, error) {
	var pResponse BruteforceLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
