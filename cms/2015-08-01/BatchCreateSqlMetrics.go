package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BatchCreateSqlMetricsRequest struct {
	RpcRequest
	ProjectName string
	Sqls        string
	IsPublic    int
}

func (r *BatchCreateSqlMetricsRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *BatchCreateSqlMetricsRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *BatchCreateSqlMetricsRequest) SetSqls(value string) {
	r.Sqls = value
	r.QueryParams.Set("Sqls", value)
}
func (r *BatchCreateSqlMetricsRequest) GetSqls() string {
	return r.Sqls
}
func (r *BatchCreateSqlMetricsRequest) SetIsPublic(value int) {
	r.IsPublic = value
	r.QueryParams.Set("IsPublic", strconv.Itoa(value))
}
func (r *BatchCreateSqlMetricsRequest) GetIsPublic() int {
	return r.IsPublic
}

func (r *BatchCreateSqlMetricsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchCreateSqlMetrics")
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type BatchCreateSqlMetricsResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func BatchCreateSqlMetrics(req *BatchCreateSqlMetricsRequest, accessId, accessSecret string) (*BatchCreateSqlMetricsResponse, error) {
	var pResponse BatchCreateSqlMetricsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
