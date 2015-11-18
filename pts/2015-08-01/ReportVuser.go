package pts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ReportVuserRequest struct {
	RpcRequest
	Wskey      string
	ScenarioId int
	Vuser      int
	GmtCreated int
}

func (r *ReportVuserRequest) SetWskey(value string) {
	r.Wskey = value
	r.QueryParams.Set("Wskey", value)
}
func (r *ReportVuserRequest) GetWskey() string {
	return r.Wskey
}
func (r *ReportVuserRequest) SetScenarioId(value int) {
	r.ScenarioId = value
	r.QueryParams.Set("ScenarioId", strconv.Itoa(value))
}
func (r *ReportVuserRequest) GetScenarioId() int {
	return r.ScenarioId
}
func (r *ReportVuserRequest) SetVuser(value int) {
	r.Vuser = value
	r.QueryParams.Set("Vuser", strconv.Itoa(value))
}
func (r *ReportVuserRequest) GetVuser() int {
	return r.Vuser
}
func (r *ReportVuserRequest) SetGmtCreated(value int) {
	r.GmtCreated = value
	r.QueryParams.Set("GmtCreated", strconv.Itoa(value))
}
func (r *ReportVuserRequest) GetGmtCreated() int {
	return r.GmtCreated
}

func (r *ReportVuserRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ReportVuser")
	r.SetProduct(Product)
}

type ReportVuserResponse struct {
}

func ReportVuser(req *ReportVuserRequest, accessId, accessSecret string) (*ReportVuserResponse, error) {
	var pResponse ReportVuserResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
