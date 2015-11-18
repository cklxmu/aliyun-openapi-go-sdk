package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeDrdsDBIpWhiteListRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
}

func (r *DescribeDrdsDBIpWhiteListRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DescribeDrdsDBIpWhiteListRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *DescribeDrdsDBIpWhiteListRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *DescribeDrdsDBIpWhiteListRequest) GetDbName() string {
	return r.DbName
}

func (r *DescribeDrdsDBIpWhiteListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDrdsDBIpWhiteList")
	r.SetProduct(Product)
}

type DescribeDrdsDBIpWhiteListResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		IpWhiteList struct {
			Ip []string `xml:"Ip" json:"Ip"`
		} `xml:"IpWhiteList" json:"IpWhiteList"`
	} `xml:"Data" json:"Data"`
}

func DescribeDrdsDBIpWhiteList(req *DescribeDrdsDBIpWhiteListRequest, accessId, accessSecret string) (*DescribeDrdsDBIpWhiteListResponse, error) {
	var pResponse DescribeDrdsDBIpWhiteListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
