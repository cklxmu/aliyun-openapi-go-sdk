package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsTopicSearchRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	Search       string
}

func (r *OnsTopicSearchRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsTopicSearchRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsTopicSearchRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsTopicSearchRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsTopicSearchRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsTopicSearchRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsTopicSearchRequest) SetSearch(value string) {
	r.Search = value
	r.QueryParams.Set("Search", value)
}
func (r *OnsTopicSearchRequest) GetSearch() string {
	return r.Search
}

func (r *OnsTopicSearchRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsTopicSearch")
	r.SetProduct(Product)
}

type OnsTopicSearchResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		PublishInfoDo []struct {
			Id           int    `xml:"Id" json:"Id"`
			ChannelId    int    `xml:"ChannelId" json:"ChannelId"`
			ChannelName  string `xml:"ChannelName" json:"ChannelName"`
			RegionId     string `xml:"RegionId" json:"RegionId"`
			RegionName   string `xml:"RegionName" json:"RegionName"`
			Topic        string `xml:"Topic" json:"Topic"`
			Owner        string `xml:"Owner" json:"Owner"`
			Relation     int    `xml:"Relation" json:"Relation"`
			RelationName string `xml:"RelationName" json:"RelationName"`
			Status       int    `xml:"Status" json:"Status"`
			StatusName   string `xml:"StatusName" json:"StatusName"`
			Appkey       int    `xml:"Appkey" json:"Appkey"`
			CreateTime   int    `xml:"CreateTime" json:"CreateTime"`
			UpdateTime   int    `xml:"UpdateTime" json:"UpdateTime"`
			Remark       string `xml:"Remark" json:"Remark"`
		} `xml:"PublishInfoDo" json:"PublishInfoDo"`
	} `xml:"Data" json:"Data"`
}

func OnsTopicSearch(req *OnsTopicSearchRequest, accessId, accessSecret string) (*OnsTopicSearchResponse, error) {
	var pResponse OnsTopicSearchResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
