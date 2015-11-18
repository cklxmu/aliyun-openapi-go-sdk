package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteMetricStreamRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
}

func (r *DeleteMetricStreamRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *DeleteMetricStreamRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *DeleteMetricStreamRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *DeleteMetricStreamRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}

func (r *DeleteMetricStreamRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteMetricStream")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DeleteMetricStreamResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func DeleteMetricStream(req *DeleteMetricStreamRequest, accessId, accessSecret string) (*DeleteMetricStreamResponse, error) {
	var pResponse DeleteMetricStreamResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
