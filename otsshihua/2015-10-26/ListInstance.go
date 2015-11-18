package otsshihua

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListInstanceRequest struct {
	RpcRequest
}

func (r *ListInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListInstance")
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type ListInstanceResponse struct {
	InstanceInfos struct {
		InstanceInfo []struct {
			InstanceName string `xml:"InstanceName" json:"InstanceName"`
			Timestamp    string `xml:"Timestamp" json:"Timestamp"`
		} `xml:"InstanceInfo" json:"InstanceInfo"`
	} `xml:"InstanceInfos" json:"InstanceInfos"`
}

func ListInstance(req *ListInstanceRequest, accessId, accessSecret string) (*ListInstanceResponse, error) {
	var pResponse ListInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
