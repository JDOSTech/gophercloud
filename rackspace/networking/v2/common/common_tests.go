package common

import (
	"github.com/JDOSTech/gophercloud"
	"github.com/JDOSTech/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}
