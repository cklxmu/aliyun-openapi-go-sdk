package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetSqlMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
}

func (r *GetSqlMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetSqlMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetSqlMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *GetSqlMetricsRequest) GetMetricName() string {
	return r.MetricName
}

func (r *GetSqlMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetSqlMetrics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetSqlMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func GetSqlMetrics(req *GetSqlMetricsRequest, accessId, accessSecret string) (*GetSqlMetricsResponse, error) {
	var pResponse GetSqlMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
