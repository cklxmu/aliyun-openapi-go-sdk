package ram

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetSecurityPreferenceRequest struct {
	RpcRequest
	EnableSaveMFATicket bool
}

func (r *SetSecurityPreferenceRequest) SetEnableSaveMFATicket(value bool) {
	r.EnableSaveMFATicket = value
	r.QueryParams.Set("EnableSaveMFATicket", strconv.FormatBool(value))
}
func (r *SetSecurityPreferenceRequest) GetEnableSaveMFATicket() bool {
	return r.EnableSaveMFATicket
}

func (r *SetSecurityPreferenceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetSecurityPreference")
	r.SetProtocol("HTTPS")
	r.SetProduct(Product)
}

type SetSecurityPreferenceResponse struct {
	SecurityPreference struct {
		LoginProfilePreference struct {
			EnableSaveMFATicket bool `xml:"EnableSaveMFATicket" json:"EnableSaveMFATicket"`
		} `xml:"LoginProfilePreference" json:"LoginProfilePreference"`
	} `xml:"SecurityPreference" json:"SecurityPreference"`
}

func SetSecurityPreference(req *SetSecurityPreferenceRequest, accessId, accessSecret string) (*SetSecurityPreferenceResponse, error) {
	var pResponse SetSecurityPreferenceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
