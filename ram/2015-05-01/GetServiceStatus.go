package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetServiceStatusRequest struct {
	RpcRequest
	AccountId string
}

func (r *GetServiceStatusRequest) SetAccountId(value string) {
	r.AccountId = value
	r.QueryParams.Set("AccountId", value)
}
func (r *GetServiceStatusRequest) GetAccountId() string {
	return r.AccountId
}

func (r *GetServiceStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetServiceStatus")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetServiceStatusResponse struct {
	Status string `xml:"Status" json:"Status"`
}

func GetServiceStatus(req *GetServiceStatusRequest, accessId, accessSecret string) (*GetServiceStatusResponse, error) {
	var pResponse GetServiceStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
