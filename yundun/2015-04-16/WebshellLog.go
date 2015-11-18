package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type WebshellLogRequest struct {
	RpcRequest
	JstOwnerId int
	PageNumber int
	PageSize   int
	InstanceId string
	RecordType int
}

func (r *WebshellLogRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *WebshellLogRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *WebshellLogRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *WebshellLogRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *WebshellLogRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *WebshellLogRequest) GetPageSize() int {
	return r.PageSize
}
func (r *WebshellLogRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *WebshellLogRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *WebshellLogRequest) SetRecordType(value int) {
	r.RecordType = value
	r.QueryParams.Set("RecordType", strconv.Itoa(value))
}
func (r *WebshellLogRequest) GetRecordType() int {
	return r.RecordType
}

func (r *WebshellLogRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("WebshellLog")
	r.SetProduct(Product)
}

type WebshellLogResponse struct {
	PageNumber int `xml:"PageNumber" json:"PageNumber"`
	PageSize   int `xml:"PageSize" json:"PageSize"`
	TotalCount int `xml:"TotalCount" json:"TotalCount"`
	LogList    struct {
		WebshellLog []struct {
			Id     string `xml:"Id" json:"Id"`
			Path   string `xml:"Path" json:"Path"`
			Status int    `xml:"Status" json:"Status"`
			Time   string `xml:"Time" json:"Time"`
		} `xml:"WebshellLog" json:"WebshellLog"`
	} `xml:"LogList" json:"LogList"`
}

func WebshellLog(req *WebshellLogRequest, accessId, accessSecret string) (*WebshellLogResponse, error) {
	var pResponse WebshellLogResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
