package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeDrdsDBRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
}

func (r *DescribeDrdsDBRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DescribeDrdsDBRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *DescribeDrdsDBRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *DescribeDrdsDBRequest) GetDbName() string {
	return r.DbName
}

func (r *DescribeDrdsDBRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDrdsDB")
	r.SetProduct(Product)
}

type DescribeDrdsDBResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		DbName     string `xml:"DbName" json:"DbName"`
		Status     int    `xml:"Status" json:"Status"`
		CreateTime string `xml:"CreateTime" json:"CreateTime"`
		Msg        string `xml:"Msg" json:"Msg"`
		Mode       string `xml:"Mode" json:"Mode"`
	} `xml:"Data" json:"Data"`
}

func DescribeDrdsDB(req *DescribeDrdsDBRequest, accessId, accessSecret string) (*DescribeDrdsDBResponse, error) {
	var pResponse DescribeDrdsDBResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
