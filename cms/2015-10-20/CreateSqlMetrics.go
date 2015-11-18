package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateSqlMetricsRequest struct {
	RpcRequest
	ProjectName string
	Sql         string
	IsPublic    int
}

func (r *CreateSqlMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *CreateSqlMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateSqlMetricsRequest) SetSql(value string) {
	r.Sql = value
	r.QueryParams.Set("Sql", value)
}
func (r *CreateSqlMetricsRequest) GetSql() string {
	return r.Sql
}
func (r *CreateSqlMetricsRequest) SetIsPublic(value int) {
	r.IsPublic = value
	r.QueryParams.Set("IsPublic", strconv.Itoa(value))
}
func (r *CreateSqlMetricsRequest) GetIsPublic() int {
	return r.IsPublic
}

func (r *CreateSqlMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateSqlMetrics")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateSqlMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func CreateSqlMetrics(req *CreateSqlMetricsRequest, accessId, accessSecret string) (*CreateSqlMetricsResponse, error) {
	var pResponse CreateSqlMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
