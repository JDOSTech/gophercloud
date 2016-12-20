package images

import (
	"testing"

	"github.com/JDOSTech/gophercloud"
	th "github.com/JDOSTech/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/v2/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestListURL(t *testing.T) {
	th.AssertEquals(t, endpoint+"images", listURL(endpointClient()))
}
