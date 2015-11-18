package pts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetScenarioStatusRequest struct {
	RpcRequest
	Wskey      string
	ScenarioId int
	Status     int
	NodeIp     string
}

func (r *SetScenarioStatusRequest) SetWskey(value string) {
	r.Wskey = value
	r.QueryParams.Set("Wskey", value)
}
func (r *SetScenarioStatusRequest) GetWskey() string {
	return r.Wskey
}
func (r *SetScenarioStatusRequest) SetScenarioId(value int) {
	r.ScenarioId = value
	r.QueryParams.Set("ScenarioId", strconv.Itoa(value))
}
func (r *SetScenarioStatusRequest) GetScenarioId() int {
	return r.ScenarioId
}
func (r *SetScenarioStatusRequest) SetStatus(value int) {
	r.Status = value
	r.QueryParams.Set("Status", strconv.Itoa(value))
}
func (r *SetScenarioStatusRequest) GetStatus() int {
	return r.Status
}
func (r *SetScenarioStatusRequest) SetNodeIp(value string) {
	r.NodeIp = value
	r.QueryParams.Set("NodeIp", value)
}
func (r *SetScenarioStatusRequest) GetNodeIp() string {
	return r.NodeIp
}

func (r *SetScenarioStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetScenarioStatus")
	r.SetProduct(Product)
}

type SetScenarioStatusResponse struct {
}

func SetScenarioStatus(req *SetScenarioStatusRequest, accessId, accessSecret string) (*SetScenarioStatusResponse, error) {
	var pResponse SetScenarioStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
