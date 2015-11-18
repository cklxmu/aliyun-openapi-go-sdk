package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetMetricsMetaRequest struct {
	RpcRequest
	ProjectName string
	MetricName  string
}

func (r *GetMetricsMetaRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetMetricsMetaRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetMetricsMetaRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *GetMetricsMetaRequest) GetMetricName() string {
	return r.MetricName
}

func (r *GetMetricsMetaRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMetricsMeta")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetMetricsMetaResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func GetMetricsMeta(req *GetMetricsMetaRequest, accessId, accessSecret string) (*GetMetricsMetaResponse, error) {
	var pResponse GetMetricsMetaResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
