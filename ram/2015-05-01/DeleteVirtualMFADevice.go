package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type DeleteVirtualMFADeviceRequest struct {
	RpcRequest
	SerialNumber string
}

func (r *DeleteVirtualMFADeviceRequest) SetSerialNumber(value string) {
	r.SerialNumber = value
	r.QueryParams.Set("SerialNumber", value)
}
func (r *DeleteVirtualMFADeviceRequest) GetSerialNumber() string {
	return r.SerialNumber
}

func (r *DeleteVirtualMFADeviceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DeleteVirtualMFADevice")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type DeleteVirtualMFADeviceResponse struct {
}

func DeleteVirtualMFADevice(req *DeleteVirtualMFADeviceRequest, accessId, accessSecret string) (*DeleteVirtualMFADeviceResponse, error) {
	var pResponse DeleteVirtualMFADeviceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
