package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ConfigDdosRequest struct {
	RpcRequest
	InstanceId       string
	InstanceType     string
	FlowPosition     int
	StrategyPosition int
	Level            int
}

func (r *ConfigDdosRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *ConfigDdosRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *ConfigDdosRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *ConfigDdosRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *ConfigDdosRequest) SetFlowPosition(value int) {
	r.FlowPosition = value
	r.QueryParams.Set("FlowPosition", strconv.Itoa(value))
}
func (r *ConfigDdosRequest) GetFlowPosition() int {
	return r.FlowPosition
}
func (r *ConfigDdosRequest) SetStrategyPosition(value int) {
	r.StrategyPosition = value
	r.QueryParams.Set("StrategyPosition", strconv.Itoa(value))
}
func (r *ConfigDdosRequest) GetStrategyPosition() int {
	return r.StrategyPosition
}
func (r *ConfigDdosRequest) SetLevel(value int) {
	r.Level = value
	r.QueryParams.Set("Level", strconv.Itoa(value))
}
func (r *ConfigDdosRequest) GetLevel() int {
	return r.Level
}

func (r *ConfigDdosRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ConfigDdos")
	r.SetProduct(Product)
}

type ConfigDdosResponse struct {
}

func ConfigDdos(req *ConfigDdosRequest, accessId, accessSecret string) (*ConfigDdosResponse, error) {
	var pResponse ConfigDdosResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
