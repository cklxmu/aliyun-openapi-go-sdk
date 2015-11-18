package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetMetricStreamRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
}

func (r *GetMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetMetricStreamRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *GetMetricStreamRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}

func (r *GetMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMetricStream")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetMetricStreamResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
	Result  string `xml:"Result" json:"Result"`
}

func GetMetricStream(req *GetMetricStreamRequest, accessId, accessSecret string) (*GetMetricStreamResponse, error) {
	var pResponse GetMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
