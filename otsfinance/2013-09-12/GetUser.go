package otsfinance

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetUserRequest struct {
	RpcRequest
}

func (r *GetUserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetUser")
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetUserResponse struct {
	UserInfo struct {
		UserId      string `xml:"UserId" json:"UserId"`
		Description string `xml:"Description" json:"Description"`
		CreateTime  string `xml:"CreateTime" json:"CreateTime"`
		Quota       struct {
			InstanceQuota int `xml:"InstanceQuota" json:"InstanceQuota"`
		} `xml:"Quota" json:"Quota"`
	} `xml:"UserInfo" json:"UserInfo"`
}

func GetUser(req *GetUserRequest, accessId, accessSecret string) (*GetUserResponse, error) {
	var pResponse GetUserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
