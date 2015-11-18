package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type SetDdosQpsRequest struct {
	RpcRequest
	InstanceId   string
	InstanceType string
	QpsPosition  int
	Level        int
}

func (r *SetDdosQpsRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *SetDdosQpsRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *SetDdosQpsRequest) SetInstanceType(value string) {
	r.InstanceType = value
	r.QueryParams.Set("InstanceType", value)
}
func (r *SetDdosQpsRequest) GetInstanceType() string {
	return r.InstanceType
}
func (r *SetDdosQpsRequest) SetQpsPosition(value int) {
	r.QpsPosition = value
	r.QueryParams.Set("QpsPosition", strconv.Itoa(value))
}
func (r *SetDdosQpsRequest) GetQpsPosition() int {
	return r.QpsPosition
}
func (r *SetDdosQpsRequest) SetLevel(value int) {
	r.Level = value
	r.QueryParams.Set("Level", strconv.Itoa(value))
}
func (r *SetDdosQpsRequest) GetLevel() int {
	return r.Level
}

func (r *SetDdosQpsRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("SetDdosQps")
	r.SetProduct(Product)
}

type SetDdosQpsResponse struct {
}

func SetDdosQps(req *SetDdosQpsRequest, accessId, accessSecret string) (*SetDdosQpsResponse, error) {
	var pResponse SetDdosQpsResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
