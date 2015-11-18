package otsshihua

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetInstanceRequest struct {
	RpcRequest
	InstanceName string
}

func (r *GetInstanceRequest) SetInstanceName(value string) {
	r.InstanceName = value
	r.QueryParams.Set("InstanceName", value)
}
func (r *GetInstanceRequest) GetInstanceName() string {
	return r.InstanceName
}

func (r *GetInstanceRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetInstance")
	r.SetMethod("GET")
	r.SetProduct(Product)
}

type GetInstanceResponse struct {
	InstanceInfo struct {
		InstanceName string `xml:"InstanceName" json:"InstanceName"`
		Status       int    `xml:"Status" json:"Status"`
		Description  string `xml:"Description" json:"Description"`
		Timestamp    string `xml:"Timestamp" json:"Timestamp"`
	} `xml:"InstanceInfo" json:"InstanceInfo"`
}

func GetInstance(req *GetInstanceRequest, accessId, accessSecret string) (*GetInstanceResponse, error) {
	var pResponse GetInstanceResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
