package cms

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type PutMetricDatumRequest struct {
	RpcRequest
	ResourceOwnerId int
	Namespace       string
	Metrics         string
}

func (r *PutMetricDatumRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *PutMetricDatumRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *PutMetricDatumRequest) SetNamespace(value string) {
	r.Namespace = value
	r.QueryParams.Set("Namespace", value)
}
func (r *PutMetricDatumRequest) GetNamespace() string {
	return r.Namespace
}
func (r *PutMetricDatumRequest) SetMetrics(value string) {
	r.Metrics = value
	r.QueryParams.Set("Metrics", value)
}
func (r *PutMetricDatumRequest) GetMetrics() string {
	return r.Metrics
}

func (r *PutMetricDatumRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("PutMetricDatum")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type PutMetricDatumResponse struct {
	Code    string `xml:"Code" json:"Code"`
	Message string `xml:"Message" json:"Message"`
}

func PutMetricDatum(req *PutMetricDatumRequest, accessId, accessSecret string) (*PutMetricDatumResponse, error) {
	var pResponse PutMetricDatumResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
