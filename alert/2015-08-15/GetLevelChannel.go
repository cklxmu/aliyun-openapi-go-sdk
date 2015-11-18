package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type GetLevelChannelRequest struct {
	RoaRequest
	ProjectName string
	Level       string
}

func (r *GetLevelChannelRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *GetLevelChannelRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *GetLevelChannelRequest) SetLevel(value string) {
	r.Level = value
	r.PathParams.Set("Level", value)
}
func (r *GetLevelChannelRequest) GetLevel() string {
	return r.Level
}

func (r *GetLevelChannelRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/level_channels/Level"
	r.SetMethod("GET")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type GetLevelChannelResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func GetLevelChannel(req *GetLevelChannelRequest, accessId, accessSecret string) (*GetLevelChannelResponse, error) {
	var pResponse GetLevelChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
