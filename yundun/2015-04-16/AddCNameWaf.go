package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type AddCNameWafRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
	Domain       string
}

func (r *AddCNameWafRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *AddCNameWafRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *AddCNameWafRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *AddCNameWafRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *AddCNameWafRequest) SetDomain(value string) {
	r.Domain = value
	r.QueryParams.Set("Domain", value)
}
func (r *AddCNameWafRequest) GetDomain() string {
	return r.Domain
}

func (r *AddCNameWafRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("AddCNameWaf")
	r.SetProduct(Product)
}

type AddCNameWafResponse struct {
	WafInfoList struct {
		WafInfo []struct {
			Id     int    `xml:"Id" json:"Id"`
			Domain string `xml:"Domain" json:"Domain"`
			Cname  string `xml:"Cname" json:"Cname"`
			Status int    `xml:"Status" json:"Status"`
		} `xml:"WafInfo" json:"WafInfo"`
	} `xml:"WafInfoList" json:"WafInfoList"`
}

func AddCNameWaf(req *AddCNameWafRequest, accessId, accessSecret string) (*AddCNameWafResponse, error) {
	var pResponse AddCNameWafResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
