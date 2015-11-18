package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type VulScanLogRequest struct {
	RpcRequest
	JstOwnerId int
	PageNumber int
	PageSize   int
	InstanceId string
	VulStatus  int
}

func (r *VulScanLogRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *VulScanLogRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *VulScanLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *VulScanLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *VulScanLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *VulScanLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *VulScanLogRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *VulScanLogRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *VulScanLogRequest) SetVulStatus(value int) {
	r.VulStatus = value
	r.QueryParams.Set("VulStatus", strconv.Itoa(value))
}
func (r *VulScanLogRequest) GetVulStatus() int {
	return r.VulStatus
}

func (r *VulScanLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("VulScanLog")
	r.SetProduct(Product)
}

type VulScanLogResponse struct {
	StartTime  string `xml:"StartTime" json:"StartTime"`
	EndTime    string `xml:"EndTime" json:"EndTime"`
	PageNumber int    `xml:"PageNumber" json:"PageNumber"`
	PageSize   int    `xml:"PageSize" json:"PageSize"`
	TotalCount int    `xml:"TotalCount" json:"TotalCount"`
	LogList    struct {
		VulScanLog []struct {
			Id           int    `xml:"Id" json:"Id"`
			Type         string `xml:"Type" json:"Type"`
			Url          string `xml:"Url" json:"Url"`
			HelpAddress  string `xml:"HelpAddress" json:"HelpAddress"`
			VulParameter string `xml:"VulParameter" json:"VulParameter"`
			Status       int    `xml:"Status" json:"Status"`
		} `xml:"VulScanLog" json:"VulScanLog"`
	} `xml:"LogList" json:"LogList"`
}

func VulScanLog(req *VulScanLogRequest, accessId, accessSecret string) (*VulScanLogResponse, error) {
	var pResponse VulScanLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
