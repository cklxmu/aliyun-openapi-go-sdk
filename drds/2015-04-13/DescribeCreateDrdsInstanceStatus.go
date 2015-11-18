package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeCreateDrdsInstanceStatusRequest struct {
	RpcRequest
	DrdsInstanceId string
}

func (r *DescribeCreateDrdsInstanceStatusRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DescribeCreateDrdsInstanceStatusRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}

func (r *DescribeCreateDrdsInstanceStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeCreateDrdsInstanceStatus")
	r.SetProduct(Product)
}

type DescribeCreateDrdsInstanceStatusResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		Status string `xml:"Status" json:"Status"`
	} `xml:"Data" json:"Data"`
}

func DescribeCreateDrdsInstanceStatus(req *DescribeCreateDrdsInstanceStatusRequest, accessId, accessSecret string) (*DescribeCreateDrdsInstanceStatusResponse, error) {
	var pResponse DescribeCreateDrdsInstanceStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
