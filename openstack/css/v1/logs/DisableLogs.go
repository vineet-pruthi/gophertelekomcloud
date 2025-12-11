package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// DisableLogs function will disable the log option for a CSS cluster.
func DisableLogs(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Put(client.ServiceURL("clusters", clusterID, "logs", "close"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}

// DisableLogs function will disable the log option for a CSS cluster.
func DisableRealTimeLogs(client *golangsdk.ServiceClient, clusterID string) error {
	queryParam := getOpts{
		Action: "real_time_log_collect",
	}
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("clusters", clusterID, "logs", "close").
		WithQueryParams(&queryParam).Build()
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(url.String()), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
