package ram

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetSecurityPreferenceRequest struct {
	RpcRequest
}

func (r *GetSecurityPreferenceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetSecurityPreference")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type GetSecurityPreferenceResponse struct {
	SecurityPreference struct {
		LoginProfilePreference struct {
			EnableSaveMFATicket bool `xml:"EnableSaveMFATicket" json:"EnableSaveMFATicket"`
		} `xml:"LoginProfilePreference" json:"LoginProfilePreference"`
	} `xml:"SecurityPreference" json:"SecurityPreference"`
}

func GetSecurityPreference(req *GetSecurityPreferenceRequest, accessId, accessSecret string) (*GetSecurityPreferenceResponse, error) {
	var pResponse GetSecurityPreferenceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
