package slb

import (
	"fmt"
	"testing"
)

func TestDescribeLoadBalancerUDPListenerAttribute(t *testing.T) {
	var req DescribeLoadBalancerUDPListenerAttributeRequest
	req.Init()
	req.SetFormat("JSON")
	req.SetRegionId("cn-shenzhen")
	var accessId = "Ie65kUInu5GeAsma"
	var accessSecret = "8cCqoxdYU9zKUihwXFXiN1HEACBDwB"
	resp, err := DescribeLoadBalancerUDPListenerAttribute(&req, accessId, accessSecret)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Success: %v\n", resp)
}
