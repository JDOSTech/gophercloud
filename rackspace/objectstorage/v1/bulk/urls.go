package bulk

import "github.com/JDOSTech/gophercloud"

func deleteURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint + "?bulk-delete"
}

func extractURL(c *gophercloud.ServiceClient, ext string) string {
	return c.Endpoint + "?extract-archive=" + ext
}
