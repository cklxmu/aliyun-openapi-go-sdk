package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DeleteCNameWafRequest struct {
	RpcRequest
	InstanceId   string
	Domain       string
	CnameId      int
	InstanceType string
}

func (r *DeleteCNameWafRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DeleteCNameWafRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DeleteCNameWafRequest) SetDomain(value string) {
	r.Domain = value
	r.QueryParams.Set("Domain", value)
}
func (r *DeleteCNameWafRequest) GetDomain() string {
	return r.Domain
}
func (r *DeleteCNameWafRequest) SetCnameId(value int) {
	r.CnameId = value
	r.QueryParams.Set("CnameId", strconv.Itoa(value))
}
func (r *DeleteCNameWafRequest) GetCnameId() int {
	return r.CnameId
}
func (r *DeleteCNameWafRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *DeleteCNameWafRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *DeleteCNameWafRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteCNameWaf")
	r.SetProduct(Product)
}

type DeleteCNameWafResponse struct {
	WafInfoList struct {
		WafInfo []struct {
			Id     int    `xml:"Id" json:"Id"`
			Domain string `xml:"Domain" json:"Domain"`
			Cname  string `xml:"Cname" json:"Cname"`
			Status int    `xml:"Status" json:"Status"`
		} `xml:"WafInfo" json:"WafInfo"`
	} `xml:"WafInfoList" json:"WafInfoList"`
}

func DeleteCNameWaf(req *DeleteCNameWafRequest, accessId, accessSecret string) (*DeleteCNameWafResponse, error) {
	var pResponse DeleteCNameWafResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
