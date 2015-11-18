package aas

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetBasicInfoForAccountRequest struct {
	RpcRequest
	AliyunId string
}

func (r *GetBasicInfoForAccountRequest) SetAliyunId(value string) {
	r.AliyunId = value
	r.QueryParams.Set("AliyunId", value)
}
func (r *GetBasicInfoForAccountRequest) GetAliyunId() string {
	return r.AliyunId
}

func (r *GetBasicInfoForAccountRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetBasicInfoForAccount")
	r.SetProduct(Product)
}

type GetBasicInfoForAccountResponse struct {
	PK            string `xml:"PK" json:"PK"`
	LastLoginDate string `xml:"LastLoginDate" json:"LastLoginDate"`
}

func GetBasicInfoForAccount(req *GetBasicInfoForAccountRequest, accessId, accessSecret string) (*GetBasicInfoForAccountResponse, error) {
	var pResponse GetBasicInfoForAccountResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
