package ubsms

import (
	. "aliyun-openapi-go-sdk/core"
)

type SetUserBusinessStatusRequest struct {
	RpcRequest
	Uid         string
	Service     string
	StatusKey   string
	StatusValue string
}

func (r *SetUserBusinessStatusRequest) SetUid(value string) {
	r.Uid = value
	r.QueryParams.Set("Uid", value)
}
func (r *SetUserBusinessStatusRequest) GetUid() string {
	return r.Uid
}
func (r *SetUserBusinessStatusRequest) SetService(value string) {
	r.Service = value
	r.QueryParams.Set("Service", value)
}
func (r *SetUserBusinessStatusRequest) GetService() string {
	return r.Service
}
func (r *SetUserBusinessStatusRequest) SetStatusKey(value string) {
	r.StatusKey = value
	r.QueryParams.Set("StatusKey", value)
}
func (r *SetUserBusinessStatusRequest) GetStatusKey() string {
	return r.StatusKey
}
func (r *SetUserBusinessStatusRequest) SetStatusValue(value string) {
	r.StatusValue = value
	r.QueryParams.Set("StatusValue", value)
}
func (r *SetUserBusinessStatusRequest) GetStatusValue() string {
	return r.StatusValue
}

func (r *SetUserBusinessStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetUserBusinessStatus")
	r.SetProduct(Product)
}

type SetUserBusinessStatusResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func SetUserBusinessStatus(req *SetUserBusinessStatusRequest, accessId, accessSecret string) (*SetUserBusinessStatusResponse, error) {
	var pResponse SetUserBusinessStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
