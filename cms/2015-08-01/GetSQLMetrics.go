package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetSQLMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
}

func (r *GetSQLMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetSQLMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetSQLMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *GetSQLMetricsRequest) GetMetricName() string {
	return r.MetricName
}

func (r *GetSQLMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetSQLMetrics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetSQLMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func GetSQLMetrics(req *GetSQLMetricsRequest, accessId, accessSecret string) (*GetSQLMetricsResponse, error) {
	var pResponse GetSQLMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
