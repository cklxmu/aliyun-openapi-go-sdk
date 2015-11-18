package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListVirtualMFADevicesRequest struct {
	RpcRequest
}

func (r *ListVirtualMFADevicesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListVirtualMFADevices")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type ListVirtualMFADevicesResponse struct {
	VirtualMFADevices struct {
		VirtualMFADevice []struct {
			SerialNumber string `xml:"SerialNumber" json:"SerialNumber"`
			ActivateDate string `xml:"ActivateDate" json:"ActivateDate"`
			User         struct {
				UserId      string `xml:"UserId" json:"UserId"`
				UserName    string `xml:"UserName" json:"UserName"`
				DisplayName string `xml:"DisplayName" json:"DisplayName"`
			} `xml:"User" json:"User"`
		} `xml:"VirtualMFADevice" json:"VirtualMFADevice"`
	} `xml:"VirtualMFADevices" json:"VirtualMFADevices"`
}

func ListVirtualMFADevices(req *ListVirtualMFADevicesRequest, accessId, accessSecret string) (*ListVirtualMFADevicesResponse, error) {
	var pResponse ListVirtualMFADevicesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
