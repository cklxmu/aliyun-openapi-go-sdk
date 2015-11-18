package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetUserMFAInfoRequest struct {
	RpcRequest
	UserName string
}

func (r *GetUserMFAInfoRequest) SetUserName(value string) {
	r.UserName = value
	r.QueryParams.Set("UserName", value)
}
func (r *GetUserMFAInfoRequest) GetUserName() string {
	return r.UserName
}

func (r *GetUserMFAInfoRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetUserMFAInfo")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetUserMFAInfoResponse struct {
	MFADevice struct {
		SerialNumber string `xml:"SerialNumber" json:"SerialNumber"`
	} `xml:"MFADevice" json:"MFADevice"`
}

func GetUserMFAInfo(req *GetUserMFAInfoRequest, accessId, accessSecret string) (*GetUserMFAInfoResponse, error) {
	var pResponse GetUserMFAInfoResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
