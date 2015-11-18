package alert

import (
	. "aliyun-openapi-go-sdk/core"
)

type CreateLevelChannelRequest struct {
	RoaRequest
	ProjectName         string
	LevelChannelSetting string
}

func (r *CreateLevelChannelRequest) SetProjectName(value string) {
	r.ProjectName = value
	r.PathParams.Set("ProjectName", value)
}
func (r *CreateLevelChannelRequest) GetProjectName() string {
	return r.ProjectName
}
func (r *CreateLevelChannelRequest) SetLevelChannelSetting(value string) {
	r.LevelChannelSetting = value
	r.QueryParams.Set("LevelChannelSetting", value)
}
func (r *CreateLevelChannelRequest) GetLevelChannelSetting() string {
	return r.LevelChannelSetting
}

func (r *CreateLevelChannelRequest) Init() {
	r.RoaRequest.Init()
	r.PathPattern = "/projects/ProjectName/level_channels"
	r.SetMethod("POST")
	r.SetProtocol("HTTP")
	r.SetProduct(Product)
}

type CreateLevelChannelResponse struct {
	code    string `xml:"code" json:"code"`
	message string `xml:"message" json:"message"`
	success string `xml:"success" json:"success"`
	traceId string `xml:"traceId" json:"traceId"`
	result  string `xml:"result" json:"result"`
}

func CreateLevelChannel(req *CreateLevelChannelRequest, accessId, accessSecret string) (*CreateLevelChannelResponse, error) {
	var pResponse CreateLevelChannelResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
