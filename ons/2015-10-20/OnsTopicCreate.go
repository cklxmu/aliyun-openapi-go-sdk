package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTopicCreateRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Topic        string
	Cluster      string
	QueueNum     int
	Order        bool
	Qps          int
	Status       int
	Remark       string
	Appkey       string
	AppName      string
}

func (r *OnsTopicCreateRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTopicCreateRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTopicCreateRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTopicCreateRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTopicCreateRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTopicCreateRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTopicCreateRequest) SetTopic(value string) {
	r.Topic = value
	r.QueryParams.Set("Topic", value)
}
func (r *OnsTopicCreateRequest) GetTopic() string {
	return r.Topic
}
func (r *OnsTopicCreateRequest) SetCluster(value string) {
	r.Cluster = value
	r.QueryParams.Set("Cluster", value)
}
func (r *OnsTopicCreateRequest) GetCluster() string {
	return r.Cluster
}
func (r *OnsTopicCreateRequest) SetQueueNum(value int) {
	r.QueueNum = value
	r.QueryParams.Set("QueueNum", strconv.Itoa(value))
}
func (r *OnsTopicCreateRequest) GetQueueNum() int {
	return r.QueueNum
}
func (r *OnsTopicCreateRequest) SetOrder(value bool) {
	r.Order = value
	r.QueryParams.Set("Order", strconv.FormatBool(value))
}
func (r *OnsTopicCreateRequest) GetOrder() bool {
	return r.Order
}
func (r *OnsTopicCreateRequest) SetQps(value int) {
	r.Qps = value
	r.QueryParams.Set("Qps", strconv.Itoa(value))
}
func (r *OnsTopicCreateRequest) GetQps() int {
	return r.Qps
}
func (r *OnsTopicCreateRequest) SetStatus(value int) {
	r.Status = value
	r.QueryParams.Set("Status", strconv.Itoa(value))
}
func (r *OnsTopicCreateRequest) GetStatus() int {
	return r.Status
}
func (r *OnsTopicCreateRequest) SetRemark(value string) {
	r.Remark = value
	r.QueryParams.Set("Remark", value)
}
func (r *OnsTopicCreateRequest) GetRemark() string {
	return r.Remark
}
func (r *OnsTopicCreateRequest) SetAppkey(value string) {
	r.Appkey = value
	r.QueryParams.Set("Appkey", value)
}
func (r *OnsTopicCreateRequest) GetAppkey() string {
	return r.Appkey
}
func (r *OnsTopicCreateRequest) SetAppName(value string) {
	r.AppName = value
	r.QueryParams.Set("AppName", value)
}
func (r *OnsTopicCreateRequest) GetAppName() string {
	return r.AppName
}

func (r *OnsTopicCreateRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTopicCreate")
	r.SetProduct(Product)
}

type OnsTopicCreateResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
}

func OnsTopicCreate(req *OnsTopicCreateRequest, accessId, accessSecret string) (*OnsTopicCreateResponse, error) {
	var pResponse OnsTopicCreateResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
