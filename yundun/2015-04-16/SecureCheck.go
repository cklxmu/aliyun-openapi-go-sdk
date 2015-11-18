package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SecureCheckRequest struct {
	RpcRequest
	JstOwnerId  int
	InstanceIds string
}

func (r *SecureCheckRequest) SetJstOwnerId(value int) {
	r.JstOwnerId = value
	r.QueryParams.Set("JstOwnerId", strconv.Itoa(value))
}
func (r *SecureCheckRequest) GetJstOwnerId() int {
	return r.JstOwnerId
}
func (r *SecureCheckRequest) SetInstanceIds(value string) {
	r.InstanceIds = value
	r.QueryParams.Set("InstanceIds", value)
}
func (r *SecureCheckRequest) GetInstanceIds() string {
	return r.InstanceIds
}

func (r *SecureCheckRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SecureCheck")
	r.SetProduct(Product)
}

type SecureCheckResponse struct {
	RecentInstanceId string `xml:"RecentInstanceId" json:"RecentInstanceId"`
	ProblemList      struct {
		Info []struct {
			Ip         string `xml:"Ip" json:"Ip"`
			Status     string `xml:"Status" json:"Status"`
			VulNum     string `xml:"VulNum" json:"VulNum"`
			InstanceId string `xml:"InstanceId" json:"InstanceId"`
		} `xml:"Info" json:"Info"`
	} `xml:"ProblemList" json:"ProblemList"`
	NoProblemList struct {
		Info []struct {
			Ip         string `xml:"Ip" json:"Ip"`
			Status     string `xml:"Status" json:"Status"`
			VulNum     string `xml:"VulNum" json:"VulNum"`
			InstanceId string `xml:"InstanceId" json:"InstanceId"`
		} `xml:"Info" json:"Info"`
	} `xml:"NoProblemList" json:"NoProblemList"`
	NoScanList struct {
		Info []struct {
			Ip         string `xml:"Ip" json:"Ip"`
			Status     string `xml:"Status" json:"Status"`
			VulNum     string `xml:"VulNum" json:"VulNum"`
			InstanceId string `xml:"InstanceId" json:"InstanceId"`
		} `xml:"Info" json:"Info"`
	} `xml:"NoScanList" json:"NoScanList"`
	ScanningList struct {
		Info []struct {
			Ip         string `xml:"Ip" json:"Ip"`
			Status     string `xml:"Status" json:"Status"`
			VulNum     string `xml:"VulNum" json:"VulNum"`
			InstanceId string `xml:"InstanceId" json:"InstanceId"`
		} `xml:"Info" json:"Info"`
	} `xml:"ScanningList" json:"ScanningList"`
	InnerIpList struct {
		Info []struct {
			Ip         string `xml:"Ip" json:"Ip"`
			Status     string `xml:"Status" json:"Status"`
			VulNum     string `xml:"VulNum" json:"VulNum"`
			InstanceId string `xml:"InstanceId" json:"InstanceId"`
		} `xml:"Info" json:"Info"`
	} `xml:"InnerIpList" json:"InnerIpList"`
}

func SecureCheck(req *SecureCheckRequest, accessId, accessSecret string) (*SecureCheckResponse, error) {
	var pResponse SecureCheckResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
