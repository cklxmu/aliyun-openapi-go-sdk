package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DdosLogRequest struct {
	RpcRequest
	JstOwnerId   int
	PageNumber   int
	PageSize     int
	InstanceId   string
	InstanceType string
}

func (r *DdosLogRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *DdosLogRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *DdosLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DdosLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DdosLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DdosLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DdosLogRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DdosLogRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DdosLogRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *DdosLogRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *DdosLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DdosLog")
	r.SetProduct(Product)
}

type DdosLogResponse struct {
	AttackStatus int    `xml:"AttackStatus" json:"AttackStatus"`
	StartTime    string `xml:"StartTime" json:"StartTime"`
	EndTime      string `xml:"EndTime" json:"EndTime"`
	PageNumber   int    `xml:"PageNumber" json:"PageNumber"`
	PageSize     int    `xml:"PageSize" json:"PageSize"`
	TotalCount   int    `xml:"TotalCount" json:"TotalCount"`
	LogList      struct {
		DdosLog []struct {
			StartTime    string `xml:"StartTime" json:"StartTime"`
			EndTime      string `xml:"EndTime" json:"EndTime"`
			Reason       string `xml:"Reason" json:"Reason"`
			Status       int    `xml:"Status" json:"Status"`
			Bps          int    `xml:"Bps" json:"Bps"`
			Pps          int    `xml:"Pps" json:"Pps"`
			Qps          int    `xml:"Qps" json:"Qps"`
			AttackType   string `xml:"AttackType" json:"AttackType"`
			AttackIpList string `xml:"AttackIpList" json:"AttackIpList"`
			Type         int    `xml:"Type" json:"Type"`
		} `xml:"DdosLog" json:"DdosLog"`
	} `xml:"LogList" json:"LogList"`
}

func DdosLog(req *DdosLogRequest, accessId, accessSecret string) (*DdosLogResponse, error) {
	var pResponse DdosLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
