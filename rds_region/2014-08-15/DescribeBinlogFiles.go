package rds_region

import (
	. "aliyun-openapi-go-sdk/core"
	"strconv"
)

type DescribeBinlogFilesRequest struct {
	RpcRequest
	OwnerId              int
	ResourceOwnerAccount string
	ResourceOwnerId      int
	DBInstanceId         string
	StartTime            string
	EndTime              string
	PageSize             int
	PageNumber           int
	OwnerAccount         string
}

func (r *DescribeBinlogFilesRequest) SetOwnerId(value int) {
	r.OwnerId = value
	r.QueryParams.Set("OwnerId", strconv.Itoa(value))
}
func (r *DescribeBinlogFilesRequest) GetOwnerId() int {
	return r.OwnerId
}
func (r *DescribeBinlogFilesRequest) SetResourceOwnerAccount(value string) {
	r.ResourceOwnerAccount = value
	r.QueryParams.Set("ResourceOwnerAccount", value)
}
func (r *DescribeBinlogFilesRequest) GetResourceOwnerAccount() string {
	return r.ResourceOwnerAccount
}
func (r *DescribeBinlogFilesRequest) SetResourceOwnerId(value int) {
	r.ResourceOwnerId = value
	r.QueryParams.Set("ResourceOwnerId", strconv.Itoa(value))
}
func (r *DescribeBinlogFilesRequest) GetResourceOwnerId() int {
	return r.ResourceOwnerId
}
func (r *DescribeBinlogFilesRequest) SetDBInstanceId(value string) {
	r.DBInstanceId = value
	r.QueryParams.Set("DBInstanceId", value)
}
func (r *DescribeBinlogFilesRequest) GetDBInstanceId() string {
	return r.DBInstanceId
}
func (r *DescribeBinlogFilesRequest) SetStartTime(value string) {
	r.StartTime = value
	r.QueryParams.Set("StartTime", value)
}
func (r *DescribeBinlogFilesRequest) GetStartTime() string {
	return r.StartTime
}
func (r *DescribeBinlogFilesRequest) SetEndTime(value string) {
	r.EndTime = value
	r.QueryParams.Set("EndTime", value)
}
func (r *DescribeBinlogFilesRequest) GetEndTime() string {
	return r.EndTime
}
func (r *DescribeBinlogFilesRequest) SetPageSize(value int) {
	r.PageSize = value
	r.QueryParams.Set("PageSize", strconv.Itoa(value))
}
func (r *DescribeBinlogFilesRequest) GetPageSize() int {
	return r.PageSize
}
func (r *DescribeBinlogFilesRequest) SetPageNumber(value int) {
	r.PageNumber = value
	r.QueryParams.Set("PageNumber", strconv.Itoa(value))
}
func (r *DescribeBinlogFilesRequest) GetPageNumber() int {
	return r.PageNumber
}
func (r *DescribeBinlogFilesRequest) SetOwnerAccount(value string) {
	r.OwnerAccount = value
	r.QueryParams.Set("OwnerAccount", value)
}
func (r *DescribeBinlogFilesRequest) GetOwnerAccount() string {
	return r.OwnerAccount
}

func (r *DescribeBinlogFilesRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("DescribeBinlogFiles")
	r.SetProduct(Product)
}

type DescribeBinlogFilesResponse struct {
	TotalRecordCount int `xml:"TotalRecordCount" json:"TotalRecordCount"`
	PageNumber       int `xml:"PageNumber" json:"PageNumber"`
	PageRecordCount  int `xml:"PageRecordCount" json:"PageRecordCount"`
	Items            struct {
		BinLogFile []struct {
			FileSize        int    `xml:"FileSize" json:"FileSize"`
			LogBeginTime    string `xml:"LogBeginTime" json:"LogBeginTime"`
			LogEndTime      string `xml:"LogEndTime" json:"LogEndTime"`
			DownloadLink    string `xml:"DownloadLink" json:"DownloadLink"`
			LinkExpiredTime string `xml:"LinkExpiredTime" json:"LinkExpiredTime"`
		} `xml:"BinLogFile" json:"BinLogFile"`
	} `xml:"Items" json:"Items"`
}

func DescribeBinlogFiles(req *DescribeBinlogFilesRequest, accessId, accessSecret string) (*DescribeBinlogFilesResponse, error) {
	var pResponse DescribeBinlogFilesResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
