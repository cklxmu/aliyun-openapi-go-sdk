package ons

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type OnsConsumerStatusRequest struct {
	RpcRequest
	OnsRegionId  string
	OnsPlatform  string
	PreventCache int
	ConsumerId   string
	Detail       bool
	NeedJstack   bool
}

func (r *OnsConsumerStatusRequest) SetOnsRegionId(value string) {
	r.OnsRegionId = value
	r.QueryParams.Set("OnsRegionId", value)
}
func (r *OnsConsumerStatusRequest) GetOnsRegionId() string {
	return r.OnsRegionId
}
func (r *OnsConsumerStatusRequest) SetOnsPlatform(value string) {
	r.OnsPlatform = value
	r.QueryParams.Set("OnsPlatform", value)
}
func (r *OnsConsumerStatusRequest) GetOnsPlatform() string {
	return r.OnsPlatform
}
func (r *OnsConsumerStatusRequest) SetPreventCache(value int) {
	r.PreventCache = value
	r.QueryParams.Set("PreventCache", strconv.Itoa(value))
}
func (r *OnsConsumerStatusRequest) GetPreventCache() int {
	return r.PreventCache
}
func (r *OnsConsumerStatusRequest) SetConsumerId(value string) {
	r.ConsumerId = value
	r.QueryParams.Set("ConsumerId", value)
}
func (r *OnsConsumerStatusRequest) GetConsumerId() string {
	return r.ConsumerId
}
func (r *OnsConsumerStatusRequest) SetDetail(value bool) {
	r.Detail = value
	r.QueryParams.Set("Detail", strconv.FormatBool(value))
}
func (r *OnsConsumerStatusRequest) GetDetail() bool {
	return r.Detail
}
func (r *OnsConsumerStatusRequest) SetNeedJstack(value bool) {
	r.NeedJstack = value
	r.QueryParams.Set("NeedJstack", strconv.FormatBool(value))
}
func (r *OnsConsumerStatusRequest) GetNeedJstack() bool {
	return r.NeedJstack
}

func (r *OnsConsumerStatusRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("OnsConsumerStatus")
	r.SetProduct(Product)
}

type OnsConsumerStatusResponse struct {
	HelpUrl string `xml:"HelpUrl" json:"HelpUrl"`
	Data    struct {
		Online           bool    `xml:"Online" json:"Online"`
		TotalDiff        int     `xml:"TotalDiff" json:"TotalDiff"`
		ConsumeTps       float32 `xml:"ConsumeTps" json:"ConsumeTps"`
		LastTimestamp    int     `xml:"LastTimestamp" json:"LastTimestamp"`
		DelayTime        int     `xml:"DelayTime" json:"DelayTime"`
		ConsumeModel     string  `xml:"ConsumeModel" json:"ConsumeModel"`
		SubscriptionSame bool    `xml:"SubscriptionSame" json:"SubscriptionSame"`
		RebalanceOK      bool    `xml:"RebalanceOK" json:"RebalanceOK"`
		ConnectionSet    struct {
			ConnectionDo []struct {
				ClientId   string `xml:"ClientId" json:"ClientId"`
				ClientAddr string `xml:"ClientAddr" json:"ClientAddr"`
				Language   string `xml:"Language" json:"Language"`
				Version    string `xml:"Version" json:"Version"`
			} `xml:"ConnectionDo" json:"ConnectionDo"`
		} `xml:"ConnectionSet" json:"ConnectionSet"`
		DetailInTopicList struct {
			DetailInTopicDo []struct {
				Topic         string  `xml:"Topic" json:"Topic"`
				ConsumeTps    float32 `xml:"ConsumeTps" json:"ConsumeTps"`
				LastTimestamp int     `xml:"LastTimestamp" json:"LastTimestamp"`
				DelayTime     int     `xml:"DelayTime" json:"DelayTime"`
				OffsetList    struct {
					ConsumeQueueOffset []struct {
						Topic          string `xml:"Topic" json:"Topic"`
						BrokerName     string `xml:"BrokerName" json:"BrokerName"`
						QueueId        int    `xml:"QueueId" json:"QueueId"`
						BrokerOffset   int    `xml:"BrokerOffset" json:"BrokerOffset"`
						ConsumerOffset int    `xml:"ConsumerOffset" json:"ConsumerOffset"`
						LastTimestamp  int    `xml:"LastTimestamp" json:"LastTimestamp"`
					} `xml:"ConsumeQueueOffset" json:"ConsumeQueueOffset"`
				} `xml:"OffsetList" json:"OffsetList"`
			} `xml:"DetailInTopicDo" json:"DetailInTopicDo"`
		} `xml:"DetailInTopicList" json:"DetailInTopicList"`
		ConsumerConnectionInfoList struct {
			ConsumerConnectionInfoDo []struct {
				InstanceId      string `xml:"InstanceId" json:"InstanceId"`
				Connection      string `xml:"Connection" json:"Connection"`
				Language        string `xml:"Language" json:"Language"`
				Version         string `xml:"Version" json:"Version"`
				ConsumeModel    string `xml:"ConsumeModel" json:"ConsumeModel"`
				ConsumeType     string `xml:"ConsumeType" json:"ConsumeType"`
				ThreadCount     int    `xml:"ThreadCount" json:"ThreadCount"`
				StartTimeStamp  int    `xml:"StartTimeStamp" json:"StartTimeStamp"`
				LastTimeStamp   int    `xml:"LastTimeStamp" json:"LastTimeStamp"`
				SubscriptionSet struct {
					SubscriptionData []struct {
						ClassFilterMode bool   `xml:"ClassFilterMode" json:"ClassFilterMode"`
						Topic           string `xml:"Topic" json:"Topic"`
						SubString       string `xml:"SubString" json:"SubString"`
						SubVersion      int    `xml:"SubVersion" json:"SubVersion"`
						TagsSet         struct {
							Tag []string `xml:"Tag" json:"Tag"`
						} `xml:"TagsSet" json:"TagsSet"`
						CodeSet struct {
							Code []string `xml:"Code" json:"Code"`
						} `xml:"CodeSet" json:"CodeSet"`
					} `xml:"SubscriptionData" json:"SubscriptionData"`
				} `xml:"SubscriptionSet" json:"SubscriptionSet"`
				RunningDataList struct {
					ConsumerRunningDataDo []struct {
						ConsumerId         string  `xml:"ConsumerId" json:"ConsumerId"`
						Topic              string  `xml:"Topic" json:"Topic"`
						Rt                 float32 `xml:"Rt" json:"Rt"`
						OkTps              float32 `xml:"OkTps" json:"OkTps"`
						FailedTps          int     `xml:"FailedTps" json:"FailedTps"`
						FailedCountPerHour int     `xml:"FailedCountPerHour" json:"FailedCountPerHour"`
					} `xml:"ConsumerRunningDataDo" json:"ConsumerRunningDataDo"`
				} `xml:"RunningDataList" json:"RunningDataList"`
				Jstack struct {
					ThreadTrackDo []struct {
						Thread    string `xml:"Thread" json:"Thread"`
						TrackList struct {
							Track []string `xml:"Track" json:"Track"`
						} `xml:"TrackList" json:"TrackList"`
					} `xml:"ThreadTrackDo" json:"ThreadTrackDo"`
				} `xml:"Jstack" json:"Jstack"`
				ProcessQueueInfoDoList struct {
					ProcessQueueInfoDo []struct {
						topic      string `xml:"topic" json:"topic"`
						BrokerName string `xml:"BrokerName" json:"BrokerName"`
						QueueId    int    `xml:"QueueId" json:"QueueId"`
						WarnMap    struct {
							LOCK   string `xml:"LOCK" json:"LOCK"`
							UNLOCK string `xml:"UNLOCK" json:"UNLOCK"`
							BLOCK  string `xml:"BLOCK" json:"BLOCK"`
						} `xml:"WarnMap" json:"WarnMap"`
					} `xml:"ProcessQueueInfoDo" json:"ProcessQueueInfoDo"`
				} `xml:"ProcessQueueInfoDoList" json:"ProcessQueueInfoDoList"`
			} `xml:"ConsumerConnectionInfoDo" json:"ConsumerConnectionInfoDo"`
		} `xml:"ConsumerConnectionInfoList" json:"ConsumerConnectionInfoList"`
	} `xml:"Data" json:"Data"`
}

func OnsConsumerStatus(req *OnsConsumerStatusRequest, accessId, accessSecret string) (*OnsConsumerStatusResponse, error) {
	var pResponse OnsConsumerStatusResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
