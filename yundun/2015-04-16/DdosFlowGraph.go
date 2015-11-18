package yundun

import (
	. "aliyun-openapi-go-sdk/core"
)

type DdosFlowGraphRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
}

func (r *DdosFlowGraphRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DdosFlowGraphRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DdosFlowGraphRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *DdosFlowGraphRequest) GetInstanceType() string {
	return r.InstanceType
}

func (r *DdosFlowGraphRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DdosFlowGraph")
	r.SetProduct(Product)
}

type DdosFlowGraphResponse struct {
	NormalFlows struct {
		NormalFlow []struct {
			time    int `xml:"time" json:"time"`
			BitRecv int `xml:"BitRecv" json:"BitRecv"`
			BitSend int `xml:"BitSend" json:"BitSend"`
			PktRecv int `xml:"PktRecv" json:"PktRecv"`
			PktSend int `xml:"PktSend" json:"PktSend"`
		} `xml:"NormalFlow" json:"NormalFlow"`
	} `xml:"NormalFlows" json:"NormalFlows"`
	TotalFlows struct {
		TotalFlow []struct {
			time    int `xml:"time" json:"time"`
			BitRecv int `xml:"BitRecv" json:"BitRecv"`
			PktRecv int `xml:"PktRecv" json:"PktRecv"`
		} `xml:"TotalFlow" json:"TotalFlow"`
	} `xml:"TotalFlows" json:"TotalFlows"`
}

func DdosFlowGraph(req *DdosFlowGraphRequest, accessId, accessSecret string) (*DdosFlowGraphResponse, error) {
	var pResponse DdosFlowGraphResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
