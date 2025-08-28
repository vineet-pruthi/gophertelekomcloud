package v1

import (
	"log"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/clusters"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/css/v1/loadbalancer"

	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestCSSLoadBalancerFullLifecycle(t *testing.T) {
	clusterID := clients.EnvOS.GetEnv("CSS_CLUSTER_ID")
	if clusterID == "" {
		t.Skip("`OS_CSS_CLUSTER_ID` must be defined")
	}
	agency := clients.EnvOS.GetEnv("AGENCY_NAME")
	if agency == "" {
		t.Skipf("OS_AGENCY_NAME is required for this test")
	}
	elbid := clients.EnvOS.GetEnv("ELB_ID")
	if elbid == "" {
		t.Skipf("OS_ELB_ID is required for this test")
	}

	client, err := clients.NewCssV1Client()
	th.AssertNoErr(t, err)

	basicOptsEnable := loadbalancer.EnableLoadBalancerOpts{
		ElbId:  elbid,
		Agency: agency,
	}
	elbID, err := loadbalancer.EnableLoadBalancer(client, clusterID, basicOptsEnable)
	th.AssertNoErr(t, err)
	log.Println("CSS loadbalancer id")
	tools.PrintResource(t, elbID)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

	err = loadbalancer.DisableLoadBalancer(client, clusterID)

	th.AssertNoErr(t, err)

	th.AssertNoErr(t, clusters.WaitForCluster(client, clusterID, timeout))

}
