package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type updateSQLMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
	Sql         string
	IsPublic    int
}

func (r *updateSQLMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *updateSQLMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *updateSQLMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *updateSQLMetricsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *updateSQLMetricsRequest) SetSql(value string) {
	r.Sql = value
	r.QueryParams.Set("Sql", value)
}
func (r *updateSQLMetricsRequest) GetSql() string {
	return r.Sql
}
func (r *updateSQLMetricsRequest) SetIsPublic(value int) {
	r.IsPublic = value
	r.QueryParams.Set("IsPublic", strconv.Itoa(value))
}
func (r *updateSQLMetricsRequest) GetIsPublic() int {
	return r.IsPublic
}

func (r *updateSQLMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("updateSQLMetrics")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type updateSQLMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func updateSQLMetrics(req *updateSQLMetricsRequest, accessId, accessSecret string) (*updateSQLMetricsResponse, error) {
	var pResponse updateSQLMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
