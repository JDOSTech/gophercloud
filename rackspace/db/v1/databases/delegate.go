package databases

import (
	"github.com/JDOSTech/gophercloud"
	os "github.com/JDOSTech/gophercloud/openstack/db/v1/databases"
	"github.com/JDOSTech/gophercloud/pagination"
)

func Create(client *gophercloud.ServiceClient, instanceID string, opts os.CreateOptsBuilder) os.CreateResult {
	return os.Create(client, instanceID, opts)
}

func List(client *gophercloud.ServiceClient, instanceID string) pagination.Pager {
	return os.List(client, instanceID)
}

func Delete(client *gophercloud.ServiceClient, instanceID, dbName string) os.DeleteResult {
	return os.Delete(client, instanceID, dbName)
}
