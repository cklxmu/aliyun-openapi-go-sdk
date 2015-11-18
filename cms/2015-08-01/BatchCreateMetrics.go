package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type BatchCreateMetricsRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
	Sqls             string
}

func (r *BatchCreateMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *BatchCreateMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *BatchCreateMetricsRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *BatchCreateMetricsRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}
func (r *BatchCreateMetricsRequest) SetSqls(value string) {
	r.Sqls = value
	r.QueryParams.Set("Sqls", value)
}
func (r *BatchCreateMetricsRequest) GetSqls() string {
	return r.Sqls
}

func (r *BatchCreateMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchCreateMetrics")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type BatchCreateMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func BatchCreateMetrics(req *BatchCreateMetricsRequest, accessId, accessSecret string) (*BatchCreateMetricsResponse, error) {
	var pResponse BatchCreateMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
