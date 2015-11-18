package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeDrdsInstanceRequest struct {
	RpcRequest
	DrdsInstanceId string
}

func (r *DescribeDrdsInstanceRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *DescribeDrdsInstanceRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}

func (r *DescribeDrdsInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDrdsInstance")
	r.SetProduct(Product)
}

type DescribeDrdsInstanceResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		DrdsInstanceId string `xml:"DrdsInstanceId" json:"DrdsInstanceId"`
		Type           string `xml:"Type" json:"Type"`
		RegionId       string `xml:"RegionId" json:"RegionId"`
		ZoneId         string `xml:"ZoneId" json:"ZoneId"`
		Description    string `xml:"Description" json:"Description"`
		NetworkType    string `xml:"NetworkType" json:"NetworkType"`
		Status         string `xml:"Status" json:"Status"`
		CreateTime     int    `xml:"CreateTime" json:"CreateTime"`
		Version        int    `xml:"Version" json:"Version"`
		Vips           struct {
			Vip []struct {
				IP   string `xml:"IP" json:"IP"`
				Port string `xml:"Port" json:"Port"`
				Type string `xml:"Type" json:"Type"`
			} `xml:"Vip" json:"Vip"`
		} `xml:"Vips" json:"Vips"`
	} `xml:"Data" json:"Data"`
}

func DescribeDrdsInstance(req *DescribeDrdsInstanceRequest, accessId, accessSecret string) (*DescribeDrdsInstanceResponse, error) {
	var pResponse DescribeDrdsInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
