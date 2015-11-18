package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetMetricStreamStatusRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
}

func (r *GetMetricStreamStatusRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetMetricStreamStatusRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetMetricStreamStatusRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *GetMetricStreamStatusRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}

func (r *GetMetricStreamStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMetricStreamStatus")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetMetricStreamStatusResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
	Success string `xml:"Success" json:"Success"`
	TraceId string `xml:"TraceId" json:"TraceId"`
}

func GetMetricStreamStatus(req *GetMetricStreamStatusRequest, accessId, accessSecret string) (*GetMetricStreamStatusResponse, error) {
	var pResponse GetMetricStreamStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
