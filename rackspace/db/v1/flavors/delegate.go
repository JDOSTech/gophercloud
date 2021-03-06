package flavors

import (
	"github.com/JDOSTech/gophercloud"
	os "github.com/JDOSTech/gophercloud/openstack/db/v1/flavors"
	"github.com/JDOSTech/gophercloud/pagination"
)

// List will list all available flavors.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return os.List(client)
}

// Get retrieves the details for a particular flavor.
func Get(client *gophercloud.ServiceClient, flavorID string) os.GetResult {
	return os.Get(client, flavorID)
}
