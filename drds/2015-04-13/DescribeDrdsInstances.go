package drds

import (
	. "aliyun-openapi-go-sdk/core"
)

type DescribeDrdsInstancesRequest struct {
	RpcRequest
	Type string
}

func (r *DescribeDrdsInstancesRequest) SetType(value string) {
	r.Type = value
	r.QueryParams.Set("Type", value)
}
func (r *DescribeDrdsInstancesRequest) GetType() string {
	return r.Type
}

func (r *DescribeDrdsInstancesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeDrdsInstances")
	r.SetProduct(Product)
}

type DescribeDrdsInstancesResponse struct {
	Success bool `xml:"Success" json:"Success"`
	Data    struct {
		Instance []struct {
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
		} `xml:"Instance" json:"Instance"`
	} `xml:"Data" json:"Data"`
}

func DescribeDrdsInstances(req *DescribeDrdsInstancesRequest, accessId, accessSecret string) (*DescribeDrdsInstancesResponse, error) {
	var pResponse DescribeDrdsInstancesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
