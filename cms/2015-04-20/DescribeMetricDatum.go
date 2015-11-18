package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeMetricDatumRequest struct {
	RpcRequest
	Namespace  string
	MetricName string
	StartTime  string
	EndTime    string
	Dimensions string
	Period     string
	Statistics string
	GroupName  string
	NextToken  string
	Length     int
}

func (r *DescribeMetricDatumRequest) SetNamespace(value string) {
	r.Namespace = value
	r.QueryParams.Set("Namespace", value)
}
func (r *DescribeMetricDatumRequest) GetNamespace() string {
	return r.Namespace
}
func (r *DescribeMetricDatumRequest) SetMetricName(value string) {
	r.MetricName = value
	r.QueryParams.Set("MetricName", value)
}
func (r *DescribeMetricDatumRequest) GetMetricName() string {
	return r.MetricName
}
func (r *DescribeMetricDatumRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeMetricDatumRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeMetricDatumRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeMetricDatumRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeMetricDatumRequest) SetDimensions(value string) {
	r.Dimensions = value
	r.QueryParams.Set("Dimensions", value)
}
func (r *DescribeMetricDatumRequest) GetDimensions() string {
	return r.Dimensions
}
func (r *DescribeMetricDatumRequest) SetPeriod(value string) {
	r.Period = value
	r.QueryParams.Set("Period", value)
}
func (r *DescribeMetricDatumRequest) GetPeriod() string {
	return r.Period
}
func (r *DescribeMetricDatumRequest) SetStatistics(value string) {
	r.Statistics = value
	r.QueryParams.Set("Statistics", value)
}
func (r *DescribeMetricDatumRequest) GetStatistics() string {
	return r.Statistics
}
func (r *DescribeMetricDatumRequest) SetGroupName(value string) {
	r.GroupName = value
	r.QueryParams.Set("GroupName", value)
}
func (r *DescribeMetricDatumRequest) GetGroupName() string {
	return r.GroupName
}
func (r *DescribeMetricDatumRequest) SetNextToken(value string) {
	r.NextToken = value
	r.QueryParams.Set("NextToken", value)
}
func (r *DescribeMetricDatumRequest) GetNextToken() string {
	return r.NextToken
}
func (r *DescribeMetricDatumRequest) SetLength(value int) {
	r.Length = value
	r.QueryParams.Set("Length", strconv.Itoa(value))
}
func (r *DescribeMetricDatumRequest) GetLength() int {
	return r.Length
}

func (r *DescribeMetricDatumRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeMetricDatum")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type DescribeMetricDatumResponse struct {
	NextToken  string `xml:"NextToken" json:"NextToken"`
	Code       string `xml:"Code" json:"Code"`
	Message    string `xml:"Message" json:"Message"`
	Datapoints struct {
		Datapoint []string `xml:"Datapoint" json:"Datapoint"`
	} `xml:"Datapoints" json:"Datapoints"`
}

func DescribeMetricDatum(req *DescribeMetricDatumRequest, accessId, accessSecret string) (*DescribeMetricDatumResponse, error) {
	var pResponse DescribeMetricDatumResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
