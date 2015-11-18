package ace

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetMonitorDataRequest struct {
	RpcRequest
	AppId     int
	Item      string
	StartTime string
	EndTime   string
	CurPage   int
	PageSize  int
}

func (r *GetMonitorDataRequest) SetAppId(value int) {
	r.AppId = value
	r.QueryParams.Set("AppId", strconv.Itoa(value))
}
func (r *GetMonitorDataRequest) GetAppId() int {
	return r.AppId
}
func (r *GetMonitorDataRequest) SetItem(value string) {
	r.Item = value
	r.QueryParams.Set("Item", value)
}
func (r *GetMonitorDataRequest) GetItem() string {
	return r.Item
}
func (r *GetMonitorDataRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *GetMonitorDataRequest) GetStartTime() string {
	return r.StartTime
}
func (r *GetMonitorDataRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *GetMonitorDataRequest) GetEndTime() string {
	return r.EndTime
}
func (r *GetMonitorDataRequest) SetCurPage(value int) {
	r.CurPage = value
	r.QueryParams.Set("CurPage", strconv.Itoa(value))
}
func (r *GetMonitorDataRequest) GetCurPage() int {
	return r.CurPage
}
func (r *GetMonitorDataRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *GetMonitorDataRequest) GetPageSize() int {
	return r.PageSize
}

func (r *GetMonitorDataRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetMonitorData")
	r.SetProduct(Product)
}

type GetMonitorDataResponse struct {
	NextPageNum string `xml:"NextPageNum" json:"NextPageNum"`
	Data        struct {
		Item []struct {
			Timestamp int     `xml:"Timestamp" json:"Timestamp"`
			Value     float32 `xml:"Value" json:"Value"`
		} `xml:"Item" json:"Item"`
	} `xml:"Data" json:"Data"`
}

func GetMonitorData(req *GetMonitorDataRequest, accessId, accessSecret string) (*GetMonitorDataResponse, error) {
	var pResponse GetMonitorDataResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
