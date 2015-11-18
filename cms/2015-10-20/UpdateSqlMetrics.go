package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type UpdateSqlMetricsRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
	Sql         string
	IsPublic    int
}

func (r *UpdateSqlMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *UpdateSqlMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *UpdateSqlMetricsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *UpdateSqlMetricsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *UpdateSqlMetricsRequest) SetSql(value string) {
	r.Sql = value
	r.QueryParams.Set("Sql", value)
}
func (r *UpdateSqlMetricsRequest) GetSql() string {
	return r.Sql
}
func (r *UpdateSqlMetricsRequest) SetIsPublic(value int) {
	r.IsPublic = value
	r.QueryParams.Set("IsPublic", strconv.Itoa(value))
}
func (r *UpdateSqlMetricsRequest) GetIsPublic() int {
	return r.IsPublic
}

func (r *UpdateSqlMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("UpdateSqlMetrics")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type UpdateSqlMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func UpdateSqlMetrics(req *UpdateSqlMetricsRequest, accessId, accessSecret string) (*UpdateSqlMetricsResponse, error) {
	var pResponse UpdateSqlMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
