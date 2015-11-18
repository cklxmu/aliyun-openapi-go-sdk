package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type CurrentDdosAttackNumRequest struct {
	RpcRequest
}

func (r *CurrentDdosAttackNumRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("CurrentDdosAttackNum")
	r.SetProduct(Product)
}

type CurrentDdosAttackNumResponse struct {
	CurrentDdosAttackNum int `xml:"CurrentDdosAttackNum" json:"CurrentDdosAttackNum"`
}

func CurrentDdosAttackNum(req *CurrentDdosAttackNumRequest, accessId, accessSecret string) (*CurrentDdosAttackNumResponse, error) {
	var pResponse CurrentDdosAttackNumResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
