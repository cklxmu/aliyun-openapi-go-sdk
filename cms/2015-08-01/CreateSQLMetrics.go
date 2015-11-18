package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type CreateSQLMetricsRequest struct {
	RpcRequest
	ProjectName string
	Sql         string
	IsPublic    int
}

func (r *CreateSQLMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *CreateSQLMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateSQLMetricsRequest) SetSql(value string) {
	r.Sql = value
	r.QueryParams.Set("Sql", value)
}
func (r *CreateSQLMetricsRequest) GetSql() string {
	return r.Sql
}
func (r *CreateSQLMetricsRequest) SetIsPublic(value int) {
	r.IsPublic = value
	r.QueryParams.Set("IsPublic", strconv.Itoa(value))
}
func (r *CreateSQLMetricsRequest) GetIsPublic() int {
	return r.IsPublic
}

func (r *CreateSQLMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateSQLMetrics")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateSQLMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func CreateSQLMetrics(req *CreateSQLMetricsRequest, accessId, accessSecret string) (*CreateSQLMetricsResponse, error) {
	var pResponse CreateSQLMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
