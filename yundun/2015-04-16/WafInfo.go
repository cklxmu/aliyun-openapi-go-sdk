package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type WafInfoRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
}

func (r *WafInfoRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *WafInfoRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *WafInfoRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *WafInfoRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *WafInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("WafInfo")
	r.SetProduct(Product)
}

type WafInfoResponse struct {
	WafDomainNum int `xml:"WafDomainNum" json:"WafDomainNum"`
	WafInfos     struct {
		WafInfo []struct {
			Id     int    `xml:"Id" json:"Id"`
			Domain string `xml:"Domain" json:"Domain"`
			Cname  string `xml:"Cname" json:"Cname"`
			Status int    `xml:"Status" json:"Status"`
		} `xml:"WafInfo" json:"WafInfo"`
	} `xml:"WafInfos" json:"WafInfos"`
}

func WafInfo(req *WafInfoRequest, accessId, accessSecret string) (*WafInfoResponse, error) {
	var pResponse WafInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
