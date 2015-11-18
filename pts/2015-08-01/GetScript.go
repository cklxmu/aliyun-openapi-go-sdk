package pts

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type GetScriptRequest struct {
	RpcRequest
	ScriptId int
	Tfsname  string
}

func (r *GetScriptRequest) SetScriptId(value int) {
	r.ScriptId = value
	r.QueryParams.Set("ScriptId", strconv.Itoa(value))
}
func (r *GetScriptRequest) GetScriptId() int {
	return r.ScriptId
}
func (r *GetScriptRequest) SetTfsname(value string) {
	r.Tfsname = value
	r.QueryParams.Set("Tfsname", value)
}
func (r *GetScriptRequest) GetTfsname() string {
	return r.Tfsname
}

func (r *GetScriptRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("GetScript")
	r.SetProduct(Product)
}

type GetScriptResponse struct {
	Script string `xml:"Script" json:"Script"`
}

func GetScript(req *GetScriptRequest, accessId, accessSecret string) (*GetScriptResponse, error) {
	var pResponse GetScriptResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
