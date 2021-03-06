package core

import (
	"fmt"
)

var endpointProductDomainMap = map[string]map[string]string{
	"cn-hangzhou": {
		"ROS":         "ros.aliyuncs.com",
		"Bss":         "bss.aliyuncs.com",
		"Ecs":         "ecs-cn-hangzhou.aliyuncs.com",
		"Oms":         "oms.aliyuncs.com",
		"Rds":         "rds.aliyuncs.com",
		"Slb":         "slb-pop.aliyuncs.com",
		"Oss":         "oss-cn-hangzhou.aliyuncs.com",
		"OssAdmin":    "oss-admin.aliyuncs.com",
		"Sts":         "sts.aliyuncs.com",
		"Yundun":      "yundun-cn-hangzhou.aliyuncs.com",
		"Risk":        "risk-cn-hangzhou.aliyuncs.com",
		"Drds":        "drds.aliyuncs.com",
		"M-kvstore":   "m-kvstore.aliyuncs.com",
		"Ram":         "ram.aliyuncs.com",
		"Cms":         "metrics.aliyuncs.com",
		"Crm":         "crm-cn-hangzhou.aliyuncs.com",
		"Ocs":         "pop-ocs.aliyuncs.com",
		"Ots":         "ots-pop.aliyuncs.com",
		"Dqs":         "dqs.aliyuncs.com",
		"Location":    "location.aliyuncs.com",
		"Ubsms":       "ubsms.aliyuncs.com",
		"Drc":         "drc.aliyuncs.com",
		"Ons":         "ons.aliyuncs.com",
		"Aas":         "aas.aliyuncs.com",
		"Ace":         "ace.cn-hangzhou.aliyuncs.com",
		"Ubsms-inner": "ubsms-inner.aliyuncs.com",
		"Dts":         "dts.aliyuncs.com",
		"R-kvstore":   "r-kvstore-cn-hangzhou.aliyuncs.com",
		"PTS":         "pts.aliyuncs.com",
		"Alert":       "alert.aliyuncs.com",
		"Push":        "cloudpush.aliyuncs.com",
		"Emr":         "emr.aliyuncs.com",
		"Cdn":         "cdn.aliyuncs.com",
	},
	"cn-shenzhen": {
		"ROS":          "ros.aliyuncs.com",
		"BatchCompute": "batchcompute.cn-shenzhen.aliyuncs.com",
		"Bss":          "bss.aliyuncs.com",
		"Alert":        "alert.aliyuncs.com",
		"Ecs":          "ecs-cn-hangzhou.aliyuncs.com",
		"Oms":          "oms.aliyuncs.com",
		"Rds":          "rds.aliyuncs.com",
		"Slb":          "slb-pop.aliyuncs.com",
		"Oss":          "oss-cn-hangzhou.aliyuncs.com",
		"OssAdmin":     "oss-admin.aliyuncs.com",
		"Sts":          "sts.aliyuncs.com",
		"Yundun":       "yundun-cn-hangzhou.aliyuncs.com",
		"Risk":         "risk-cn-hangzhou.aliyuncs.com",
		"Drds":         "drds.aliyuncs.com",
		"M-kvstore":    "m-kvstore.aliyuncs.com",
		"Ram":          "ram.aliyuncs.com",
		"Cms":          "metrics.aliyuncs.com",
		"Crm":          "crm-cn-hangzhou.aliyuncs.com",
		"Ocs":          "pop-ocs.aliyuncs.com",
		"Ots":          "ots-pop.aliyuncs.com",
		"Dqs":          "dqs.aliyuncs.com",
		"Location":     "location.aliyuncs.com",
		"Ubsms":        "ubsms.aliyuncs.com",
		"Drc":          "drc.aliyuncs.com",
		"Ons":          "ons.aliyuncs.com",
		"Aas":          "aas.aliyuncs.com",
		"Ace":          "ace.cn-hangzhou.aliyuncs.com",
		"Ubsms-inner":  "ubsms-inner.aliyuncs.com",
		"Dts":          "dts.aliyuncs.com",
		"R-kvstore":    "r-kvstore-cn-hangzhou.aliyuncs.com",
		"PTS":          "pts.aliyuncs.com",
		"Push":         "cloudpush.aliyuncs.com",
		"Emr":          "emr.aliyuncs.com",
		"Cdn":          "cdn.aliyuncs.com",
	},
	"cn-qingdao": {
		"ROS":          "ros.aliyuncs.com",
		"BatchCompute": "batchcompute.cn-qingdao.aliyuncs.com",
		"Bss":          "bss.aliyuncs.com",
		"Ecs":          "ecs-cn-hangzhou.aliyuncs.com",
		"Oms":          "oms.aliyuncs.com",
		"Rds":          "rds.aliyuncs.com",
		"Slb":          "slb-pop.aliyuncs.com",
		"Oss":          "oss-cn-hangzhou.aliyuncs.com",
		"OssAdmin":     "oss-admin.aliyuncs.com",
		"Sts":          "sts.aliyuncs.com",
		"Yundun":       "yundun-cn-hangzhou.aliyuncs.com",
		"Risk":         "risk-cn-hangzhou.aliyuncs.com",
		"Drds":         "drds.aliyuncs.com",
		"M-kvstore":    "m-kvstore.aliyuncs.com",
		"Ram":          "ram.aliyuncs.com",
		"Cms":          "metrics.aliyuncs.com",
		"Crm":          "crm-cn-hangzhou.aliyuncs.com",
		"Ocs":          "pop-ocs.aliyuncs.com",
		"Ots":          "ots-pop.aliyuncs.com",
		"Dqs":          "dqs.aliyuncs.com",
		"Location":     "location.aliyuncs.com",
		"Ubsms":        "ubsms.aliyuncs.com",
		"Drc":          "drc.aliyuncs.com",
		"Ons":          "ons.aliyuncs.com",
		"Aas":          "aas.aliyuncs.com",
		"Ace":          "ace.cn-hangzhou.aliyuncs.com",
		"Ubsms-inner":  "ubsms-inner.aliyuncs.com",
		"Dts":          "dts.aliyuncs.com",
		"R-kvstore":    "r-kvstore-cn-hangzhou.aliyuncs.com",
		"PTS":          "pts.aliyuncs.com",
		"Alert":        "alert.aliyuncs.com",
		"Push":         "cloudpush.aliyuncs.com",
		"Emr":          "emr.aliyuncs.com",
		"Cdn":          "cdn.aliyuncs.com",
	},
}

func FindProductDomain(product, regionId string) (string, error) {
	if endpoint, exist := endpointProductDomainMap[regionId]; !exist {
		return "", SdkError(fmt.Sprintf("Invalid RegionId %s", regionId))
	} else {
		if domain, exist := endpoint[product]; !exist {
			return "", SdkError(fmt.Sprintf("Invalid Product %s for RegionsId %s", product, endpoint))
		} else {
			return domain, nil
		}
	}
}
