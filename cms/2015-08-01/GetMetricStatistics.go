package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetMetricStatisticsRequest struct {
	RpcRequest
	Namespace  string
	MetricName string
	StartTime  string
	EndTime    string
	Interval   string
	Dimensions string
	NextToken  int
	Length     int
}

func (r *GetMetricStatisticsRequest) SetNamespace(value string) {
	r.Namespace = value
	r.QueryParams.Set("Namespace", value)
}
func (r *GetMetricStatisticsRequest) GetNamespace() string {
	return r.Namespace
}
func (r *GetMetricStatisticsRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *GetMetricStatisticsRequest) GetMetricName() string {
	return r.MetricName
}
func (r *GetMetricStatisticsRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *GetMetricStatisticsRequest) GetStartTime() string {
	return r.StartTime
}
func (r *GetMetricStatisticsRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *GetMetricStatisticsRequest) GetEndTime() string {
	return r.EndTime
}
func (r *GetMetricStatisticsRequest) SetInterval(value string) {
	r.Interval = value
	r.QueryParams.Set("Interval", value)
}
func (r *GetMetricStatisticsRequest) GetInterval() string {
	return r.Interval
}
func (r *GetMetricStatisticsRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *GetMetricStatisticsRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *GetMetricStatisticsRequest) SetNextToken(value int) {
	r.NextToken = value
	r.QueryParams.Set("NextToken", strconv.Itoa(value))
}
func (r *GetMetricStatisticsRequest) GetNextToken() int {
	return r.NextToken
}
func (r *GetMetricStatisticsRequest) SetLength(value int) {
	r.Length = value
	r.QueryParams.Set("Length", strconv.Itoa(value))
}
func (r *GetMetricStatisticsRequest) GetLength() int {
	return r.Length
}

func (r *GetMetricStatisticsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMetricStatistics")
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetMetricStatisticsResponse struct {
	NextToken  string `xml:"NextToken" json:"NextToken"`
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func GetMetricStatistics(req *GetMetricStatisticsRequest, accessId, accessSecret string) (*GetMetricStatisticsResponse, error) {
	var pResponse GetMetricStatisticsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
