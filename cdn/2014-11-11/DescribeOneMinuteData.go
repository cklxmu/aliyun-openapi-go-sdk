package cdn

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeOneMinuteDataRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DomainName           string
	DataTime             string
}

func (r *DescribeOneMinuteDataRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeOneMinuteDataRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeOneMinuteDataRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeOneMinuteDataRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeOneMinuteDataRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeOneMinuteDataRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeOneMinuteDataRequest) SetDomainName(value string) {
	r.DomainName = value
	r.QueryParams.Set("DomainName", value)
}
func (r *DescribeOneMinuteDataRequest) GetDomainName() string {
	return r.DomainName
}
func (r *DescribeOneMinuteDataRequest) SetDataTime(value string) {
	r.DataTime = value
	r.QueryParams.Set("DataTime", value)
}
func (r *DescribeOneMinuteDataRequest) GetDataTime() string {
	return r.DataTime
}

func (r *DescribeOneMinuteDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeOneMinuteData")
	r.SetProduct(Product)
}

type DescribeOneMinuteDataResponse struct {
	Bps string `xml:"Bps" json:"Bps"`
	Qps string `xml:"Qps" json:"Qps"`
}

func DescribeOneMinuteData(req *DescribeOneMinuteDataRequest, accessId, accessSecret string) (*DescribeOneMinuteDataResponse, error) {
	var pResponse DescribeOneMinuteDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
