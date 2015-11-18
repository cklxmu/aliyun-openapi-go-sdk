package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeDrdsDBsRequest struct {
	RpcRequest
	DrdsInstanceId string
}

func (r *DescribeDrdsDBsRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DescribeDrdsDBsRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}

func (r *DescribeDrdsDBsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDrdsDBs")
	r.SetProduct(Product)
}

type DescribeDrdsDBsResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		Db []struct {
			DbName     string `xml:"DbName" json:"DbName"`
			Status     int    `xml:"Status" json:"Status"`
			CreateTime string `xml:"CreateTime" json:"CreateTime"`
			Msg        string `xml:"Msg" json:"Msg"`
			Mode       string `xml:"Mode" json:"Mode"`
		} `xml:"Db" json:"Db"`
	} `xml:"Data" json:"Data"`
}

func DescribeDrdsDBs(req *DescribeDrdsDBsRequest, accessId, accessSecret string) (*DescribeDrdsDBsResponse, error) {
	var pResponse DescribeDrdsDBsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
