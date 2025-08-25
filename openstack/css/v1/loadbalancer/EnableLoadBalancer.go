package loadbalancer

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableLoadBalancerOpts struct {
	// These parameters are passed to the loadbalancer.EnableLoadBalancer function.
	// Enable or Disable the loadbalancer for the css cluster.
	Enable bool `json:"enable" required:"false`
	// These parameters are passed to the loadbalancer.EnableLoadBalancer function.
	// ID of the loadbalancer of the css cluster.
	ElbId string `json:"elb_id" required:"true"`
	// These parameters are passed to the loadbalancer.EnableLoadBalancer function.
	// Agency is the agency name used for the css cluster.
	Agency string `json:"agency" required:"true"`
}

// EnableLoadBalancer function is used to enable the loadbalancer switch of a CSS cluster base on EnableLoadBalancerOpts.
func EnableLoadBalancer(client *golangsdk.ServiceClient, clusterID string, opts EnableLoadBalancerOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "loadbalancers", "es-switch")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
