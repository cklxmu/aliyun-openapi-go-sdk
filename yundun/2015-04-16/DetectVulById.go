package yundun

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DetectVulByIdRequest struct {
	RpcRequest
	InstanceId string
	VulId      int
}

func (r *DetectVulByIdRequest) SetInstanceId(value string) {
	r.InstanceId = value
	r.QueryParams.Set("InstanceId", value)
}
func (r *DetectVulByIdRequest) GetInstanceId() string {
	return r.InstanceId
}
func (r *DetectVulByIdRequest) SetVulId(value int) {
	r.VulId = value
	r.QueryParams.Set("VulId", strconv.Itoa(value))
}
func (r *DetectVulByIdRequest) GetVulId() int {
	return r.VulId
}

func (r *DetectVulByIdRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DetectVulById")
	r.SetProduct(Product)
}

type DetectVulByIdResponse struct {
}

func DetectVulById(req *DetectVulByIdRequest, accessId, accessSecret string) (*DetectVulByIdResponse, error) {
	var pResponse DetectVulByIdResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
