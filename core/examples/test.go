package main

import (
	"aliyun-openapi-go-sdk/ecs/2014-05-26"
	"fmt"
)

func main() {
	var req ecs.DescribeRegionsRequest
	req.Init()
	req.SetFormat("JSON")
	req.SetRegionId("cn-shenzhen")
	var accessId = "Ie65kUInu5GeAsma"
	var accessSecret = "8cCqoxdYU9zKUihwXFXiN1HEACBDwB"
	resp, err := ecs.DescribeRegions(&req, accessId, accessSecret)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Success: %v\n", resp)
}
