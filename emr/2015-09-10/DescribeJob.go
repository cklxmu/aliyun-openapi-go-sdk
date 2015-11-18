package emr

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeJobRequest struct {
	RpcRequest
	Id int
}

func (r *DescribeJobRequest) SetId(value int) {
	r.Id = value
	r.QueryParams.Set("Id", strconv.Itoa(value))
}
func (r *DescribeJobRequest) GetId() int {
	return r.Id
}

func (r *DescribeJobRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeJob")
	r.SetProduct(Product)
}

type DescribeJobResponse struct {
	Id       int    `xml:"Id" json:"Id"`
	Name     string `xml:"Name" json:"Name"`
	FailAct  int    `xml:"FailAct" json:"FailAct"`
	Type     int    `xml:"Type" json:"Type"`
	EnvParam string `xml:"EnvParam" json:"EnvParam"`
}

func DescribeJob(req *DescribeJobRequest, accessId, accessSecret string) (*DescribeJobResponse, error) {
	var pResponse DescribeJobResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
