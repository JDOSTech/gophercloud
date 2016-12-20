package quotasets

import (
	"testing"

	th "github.com/JDOSTech/gophercloud/testhelper"
	"github.com/JDOSTech/gophercloud/testhelper/client"
)

func TestGetURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()

	th.CheckEquals(t, c.Endpoint+"os-quota-sets/wat", getURL(c, "wat"))
}
