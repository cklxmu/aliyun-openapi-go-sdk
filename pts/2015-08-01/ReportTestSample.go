package pts

import (
	. "aliyun-openapi-go-sdk/core"
)

type ReportTestSampleRequest struct {
	RpcRequest
	TestSample string
}

func (r *ReportTestSampleRequest) SetTestSample(value string) {
	r.TestSample = value
	r.QueryParams.Set("TestSample", value)
}
func (r *ReportTestSampleRequest) GetTestSample() string {
	return r.TestSample
}

func (r *ReportTestSampleRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReportTestSample")
	r.SetProduct(Product)
}

type ReportTestSampleResponse struct {
}

func ReportTestSample(req *ReportTestSampleRequest, accessId, accessSecret string) (*ReportTestSampleResponse, error) {
	var pResponse ReportTestSampleResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
