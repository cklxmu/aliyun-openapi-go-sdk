package emr

import (
	. "aliyun-openapi-go-sdk/core"
)

type ListConfigTypeRequest struct {
	RpcRequest
}

func (r *ListConfigTypeRequest) Init() {
	r.RpcRequest.Init()
	r.SetVersion(Version)
	r.SetAction("ListConfigType")
	r.SetProduct(Product)
}

type ListConfigTypeResponse struct {
	securityGroupTypes struct {
		SecurityGroupType []struct {
			Name  string `xml:"Name" json:"Name"`
			State string `xml:"State" json:"State"`
			Id    string `xml:"Id" json:"Id"`
		} `xml:"SecurityGroupType" json:"SecurityGroupType"`
	} `xml:"securityGroupTypes" json:"securityGroupTypes"`
	instanceTypes struct {
		InstanceType []struct {
			Classify     string `xml:"Classify" json:"Classify"`
			State        string `xml:"State" json:"State"`
			Type         string `xml:"Type" json:"Type"`
			CpuCore      string `xml:"CpuCore" json:"CpuCore"`
			MemSize      string `xml:"MemSize" json:"MemSize"`
			HasCloudDisk bool   `xml:"HasCloudDisk" json:"HasCloudDisk"`
			HasLocalDisk bool   `xml:"HasLocalDisk" json:"HasLocalDisk"`
			HasLocalSSD  bool   `xml:"HasLocalSSD" json:"HasLocalSSD"`
			HasCloudSSD  bool   `xml:"HasCloudSSD" json:"HasCloudSSD"`
		} `xml:"InstanceType" json:"InstanceType"`
	} `xml:"instanceTypes" json:"instanceTypes"`
	EmrVerTypes struct {
		EmrVerType []struct {
			Name       string `xml:"Name" json:"Name"`
			SubModules struct {
				SubModule []struct {
					Name     string `xml:"Name" json:"Name"`
					Required string `xml:"Required" json:"Required"`
					Optional string `xml:"Optional" json:"Optional"`
				} `xml:"SubModule" json:"SubModule"`
			} `xml:"SubModules" json:"SubModules"`
		} `xml:"EmrVerType" json:"EmrVerType"`
	} `xml:"EmrVerTypes" json:"EmrVerTypes"`
	ZoneTypes struct {
		ZoneType []struct {
			Name                           string `xml:"Name" json:"Name"`
			Id                             string `xml:"Id" json:"Id"`
			AvailableResourceCreateionList struct {
				AvailableResourceCreateion []string `xml:"AvailableResourceCreateion" json:"AvailableResourceCreateion"`
			} `xml:"AvailableResourceCreateionList" json:"AvailableResourceCreateionList"`
			AvailableDiskCategoriesList struct {
				AvailableDiskCategories []string `xml:"AvailableDiskCategories" json:"AvailableDiskCategories"`
			} `xml:"AvailableDiskCategoriesList" json:"AvailableDiskCategoriesList"`
		} `xml:"ZoneType" json:"ZoneType"`
	} `xml:"ZoneTypes" json:"ZoneTypes"`
}

func ListConfigType(req *ListConfigTypeRequest, accessId, accessSecret string) (*ListConfigTypeResponse, error) {
	var pResponse ListConfigTypeResponse
	body, err := ApiHttpRequest(accessId, accessSecret, req)
	if err != nil {
		return nil, err
	}
	ApiUnmarshalResponse(req.GetFormat(), body, &pResponse)
	return &pResponse, err
}
