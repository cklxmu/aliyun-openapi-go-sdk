package slb

import (
	"fmt"
	"testing"
)

func TestSetLoadBalancerTCPListenerAttribute(t *testing.T) {
	var req SetLoadBalancerTCPListenerAttributeRequest
	req.Init()
	req.SetFormat("JSON")
	req.SetRegionId("cn-shenzhen")
	var accessId = "Ie65kUInu5GeAsma"
	var accessSecret = "8cCqoxdYU9zKUihwXFXiN1HEACBDwB"
	resp, err := SetLoadBalancerTCPListenerAttribute(&req, accessId, accessSecret)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	fmt.Printf("Success: %v\n", resp)
}
