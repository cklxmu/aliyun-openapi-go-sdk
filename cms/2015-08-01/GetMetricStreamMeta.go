package cms

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetMetricStreamMetaRequest struct {
	RpcRequest
	ProjectName      string
	MetricStreamName string
}

func (r *GetMetricStreamMetaRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.QueryParams.Set("ProjectName", value)
}
func (r *GetMetricStreamMetaRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetMetricStreamMetaRequest) SetMetricStreamName(value string) {
	r.MetricStreamName = value
	r.QueryParams.Set("MetricStreamName", value)
}
func (r *GetMetricStreamMetaRequest) GetMetricStreamName() string {
	return r.MetricStreamName
}

func (r *GetMetricStreamMetaRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMetricStreamMeta")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetMetricStreamMetaResponse struct {
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Success    string `xml:"Success" json:"Success"`
	TraceId    string `xml:"TraceId" json:"TraceId"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func GetMetricStreamMeta(req *GetMetricStreamMetaRequest, accessId, accessSecret string) (*GetMetricStreamMetaResponse, error) {
	var pResponse GetMetricStreamMetaResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
