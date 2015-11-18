package drds

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type ModifyDrdsIpWhiteListRequest struct {
	RpcRequest
	DrdsInstanceId string
	DbName         string
	IpWhiteList    string
	Mode           bool
}

func (r *ModifyDrdsIpWhiteListRequest) SetDrdsInstanceId(value string) {
	r.DrdsInstanceId = value
	r.QueryParams.Set("DrdsInstanceId", value)
}
func (r *ModifyDrdsIpWhiteListRequest) GetDrdsInstanceId() string {
	return r.DrdsInstanceId
}
func (r *ModifyDrdsIpWhiteListRequest) SetDbName(value string) {
	r.DbName = value
	r.QueryParams.Set("DbName", value)
}
func (r *ModifyDrdsIpWhiteListRequest) GetDbName() string {
	return r.DbName
}
func (r *ModifyDrdsIpWhiteListRequest) SetIpWhiteList(value string) {
	r.IpWhiteList = value
	r.QueryParams.Set("IpWhiteList", value)
}
func (r *ModifyDrdsIpWhiteListRequest) GetIpWhiteList() string {
	return r.IpWhiteList
}
func (r *ModifyDrdsIpWhiteListRequest) SetMode(value bool) {
	r.Mode = value
	r.QueryParams.Set("Mode", strconv.FormatBool(value))
}
func (r *ModifyDrdsIpWhiteListRequest) GetMode() bool {
	return r.Mode
}

func (r *ModifyDrdsIpWhiteListRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ModifyDrdsIpWhiteList")
	r.SetProduct(Product)
}

type ModifyDrdsIpWhiteListResponse struct {
	Success bool `xml:"Success" json:"Success"`
}

func ModifyDrdsIpWhiteList(req *ModifyDrdsIpWhiteListRequest, accessId, accessSecret string) (*ModifyDrdsIpWhiteListResponse, error) {
	var pResponse ModifyDrdsIpWhiteListResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
