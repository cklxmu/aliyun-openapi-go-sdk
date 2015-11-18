package pts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReportLogSampleRequest struct {
	RpcRequest
	Wskey      string
	ScenarioId int
	LogSample  string
}

func (r *ReportLogSampleRequest) SetWskey(value string) {
	r.Wskey = value
	r.QueryParams.Set("Wskey", value)
}
func (r *ReportLogSampleRequest) GetWskey() string {
	return r.Wskey
}
func (r *ReportLogSampleRequest) SetScenarioId(value int) {
	r.ScenarioId = value
	r.QueryParams.Set("ScenarioId", strconv.Itoa(value))
}
func (r *ReportLogSampleRequest) GetScenarioId() int {
	return r.ScenarioId
}
func (r *ReportLogSampleRequest) SetLogSample(value string) {
	r.LogSample = value
	r.QueryParams.Set("LogSample", value)
}
func (r *ReportLogSampleRequest) GetLogSample() string {
	return r.LogSample
}

func (r *ReportLogSampleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReportLogSample")
	r.SetProduct(Product)
}

type ReportLogSampleResponse struct {
}

func ReportLogSample(req *ReportLogSampleRequest, accessId, accessSecret string) (*ReportLogSampleResponse, error) {
	var pResponse ReportLogSampleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
