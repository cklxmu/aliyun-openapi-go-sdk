package push

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type BatchGetDevicesInfoRequest struct {
	RpcRequest
	Devices string
	AppId   int
}

func (r *BatchGetDevicesInfoRequest) SetDevices(value string) {
	r.Devices = value
	r.QueryParams.Set("Devices", value)
}
func (r *BatchGetDevicesInfoRequest) GetDevices() string {
	return r.Devices
}
func (r *BatchGetDevicesInfoRequest) SetAppId(value int) {
	r.AppId = value
	r.QueryParams.Set("AppId", strconv.Itoa(value))
}
func (r *BatchGetDevicesInfoRequest) GetAppId() int {
	return r.AppId
}

func (r *BatchGetDevicesInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("BatchGetDevicesInfo")
	r.SetProduct(Product)
}

type BatchGetDevicesInfoResponse struct {
	DeviceInfos struct {
		DeviceInfo []struct {
			DeviceId string `xml:"DeviceId" json:"DeviceId"`
			IsOnline bool   `xml:"IsOnline" json:"IsOnline"`
			Status   int    `xml:"Status" json:"Status"`
		} `xml:"DeviceInfo" json:"DeviceInfo"`
	} `xml:"DeviceInfos" json:"DeviceInfos"`
}

func BatchGetDevicesInfo(req *BatchGetDevicesInfoRequest, accessId, accessSecret string) (*BatchGetDevicesInfoResponse, error) {
	var pResponse BatchGetDevicesInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
