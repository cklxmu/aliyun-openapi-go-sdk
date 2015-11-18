package ubsms_inner

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeUserBusinessStatusRequest struct {
	RpcRequest
	callerBid   string
	Uid         string
	ServiceCode string
	Password    string
}

func (r *DescribeUserBusinessStatusRequest) SetcallerBid(value string) {
	r.callerBid = value
	r.QueryParams.Set("callerBid", value)
}
func (r *DescribeUserBusinessStatusRequest) GetcallerBid() string {
	return r.callerBid
}
func (r *DescribeUserBusinessStatusRequest) SetUid(value string) {
	r.Uid = value
	r.QueryParams.Set("Uid", value)
}
func (r *DescribeUserBusinessStatusRequest) GetUid() string {
	return r.Uid
}
func (r *DescribeUserBusinessStatusRequest) SetServiceCode(value string) {
	r.ServiceCode = value
	r.QueryParams.Set("ServiceCode", value)
}
func (r *DescribeUserBusinessStatusRequest) GetServiceCode() string {
	return r.ServiceCode
}
func (r *DescribeUserBusinessStatusRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *DescribeUserBusinessStatusRequest) GetPassword() string {
	return r.Password
}

func (r *DescribeUserBusinessStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeUserBusinessStatus")
	r.SetProduct(Product)
}

type DescribeUserBusinessStatusResponse struct {
	Success     bool   `xml:"Success" json:"Success"`
	Uid         string `xml:"Uid" json:"Uid"`
	ServiceCode string `xml:"ServiceCode" json:"ServiceCode"`
	Statuses    struct {
		Status []struct {
			StatusKey   string `xml:"StatusKey" json:"StatusKey"`
			StatusValue string `xml:"StatusValue" json:"StatusValue"`
		} `xml:"Status" json:"Status"`
	} `xml:"Statuses" json:"Statuses"`
}

func DescribeUserBusinessStatus(req *DescribeUserBusinessStatusRequest, accessId, accessSecret string) (*DescribeUserBusinessStatusResponse, error) {
	var pResponse DescribeUserBusinessStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
