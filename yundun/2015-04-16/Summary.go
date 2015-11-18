package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SummaryRequest struct {
	RpcRequest
	JstOwnerId  int
	InstanceIds string
}

func (r *SummaryRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *SummaryRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *SummaryRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *SummaryRequest) GetInstanceIds() string {
	return r.InstanceIds
}

func (r *SummaryRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("Summary")
	r.SetProduct(Product)
}

type SummaryResponse struct {
	Status            int `xml:"Status" json:"Status"`
	AbnormalHostCount int `xml:"AbnormalHostCount" json:"AbnormalHostCount"`
	Ddos              struct {
		Count     int `xml:"Count" json:"Count"`
		HostCount int `xml:"HostCount" json:"HostCount"`
	} `xml:"Ddos" json:"Ddos"`
	BruteForce struct {
		Count     int `xml:"Count" json:"Count"`
		HostCount int `xml:"HostCount" json:"HostCount"`
	} `xml:"BruteForce" json:"BruteForce"`
	Webshell struct {
		Count     int `xml:"Count" json:"Count"`
		HostCount int `xml:"HostCount" json:"HostCount"`
	} `xml:"Webshell" json:"Webshell"`
	RemoteLogin struct {
		Count     int `xml:"Count" json:"Count"`
		HostCount int `xml:"HostCount" json:"HostCount"`
	} `xml:"RemoteLogin" json:"RemoteLogin"`
	WebAttack struct {
		Count     int `xml:"Count" json:"Count"`
		HostCount int `xml:"HostCount" json:"HostCount"`
	} `xml:"WebAttack" json:"WebAttack"`
	WebLeak struct {
		Count     int `xml:"Count" json:"Count"`
		HostCount int `xml:"HostCount" json:"HostCount"`
	} `xml:"WebLeak" json:"WebLeak"`
}

func Summary(req *SummaryRequest, accessId, accessSecret string) (*SummaryResponse, error) {
	var pResponse SummaryResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
