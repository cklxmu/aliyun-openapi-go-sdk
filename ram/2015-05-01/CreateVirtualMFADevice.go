package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateVirtualMFADeviceRequest struct {
	RpcRequest
	VirtualMFADeviceName string
}

func (r *CreateVirtualMFADeviceRequest) SetVirtualMFADeviceName(value string) {
	r.VirtualMFADeviceName = value
	r.QueryParams.Set("VirtualMFADeviceName", value)
}
func (r *CreateVirtualMFADeviceRequest) GetVirtualMFADeviceName() string {
	return r.VirtualMFADeviceName
}

func (r *CreateVirtualMFADeviceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CreateVirtualMFADevice")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type CreateVirtualMFADeviceResponse struct {
	VirtualMFADevice struct {
		SerialNumber     string `xml:"SerialNumber" json:"SerialNumber"`
		Base32StringSeed string `xml:"Base32StringSeed" json:"Base32StringSeed"`
		QRCodePNG        string `xml:"QRCodePNG" json:"QRCodePNG"`
	} `xml:"VirtualMFADevice" json:"VirtualMFADevice"`
}

func CreateVirtualMFADevice(req *CreateVirtualMFADeviceRequest, accessId, accessSecret string) (*CreateVirtualMFADeviceResponse, error) {
	var pResponse CreateVirtualMFADeviceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
