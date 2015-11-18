package ubsms

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeBusinessStatusRequest struct {
	RpcRequest
	callerBid string
	Password  string
}

func (r *DescribeBusinessStatusRequest) SetcallerBid(value string) {
	r.callerBid = value
	r.QueryParams.Set("callerBid", value)
}
func (r *DescribeBusinessStatusRequest) GetcallerBid() string {
	return r.callerBid
}
func (r *DescribeBusinessStatusRequest) SetPassword(value string) {
	r.Password = value
	r.QueryParams.Set("Password", value)
}
func (r *DescribeBusinessStatusRequest) GetPassword() string {
	return r.Password
}

func (r *DescribeBusinessStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeBusinessStatus")
	r.SetProduct(Product)
}

type DescribeBusinessStatusResponse struct {
	Success                bool `xml:"Success" json:"Success"`
	UserBusinessStatusList struct {
		UserBusinessStatus []struct {
			Uid         string `xml:"Uid" json:"Uid"`
			ServiceCode string `xml:"ServiceCode" json:"ServiceCode"`
			Statuses    struct {
				Status []struct {
					StatusKey   string `xml:"StatusKey" json:"StatusKey"`
					StatusValue string `xml:"StatusValue" json:"StatusValue"`
				} `xml:"Status" json:"Status"`
			} `xml:"Statuses" json:"Statuses"`
		} `xml:"UserBusinessStatus" json:"UserBusinessStatus"`
	} `xml:"UserBusinessStatusList" json:"UserBusinessStatusList"`
}

func DescribeBusinessStatus(req *DescribeBusinessStatusRequest, accessId, accessSecret string) (*DescribeBusinessStatusResponse, error) {
	var pResponse DescribeBusinessStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
