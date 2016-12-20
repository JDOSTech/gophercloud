package backups

import "github.com/JDOSTech/gophercloud"

func baseURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("backups")
}

func resourceURL(c *gophercloud.ServiceClient, backupID string) string {
	return c.ServiceURL("backups", backupID)
}
