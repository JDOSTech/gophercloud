package tenants

import (
	"testing"

	"github.com/JDOSTech/gophercloud/pagination"
	th "github.com/JDOSTech/gophercloud/testhelper"
	"github.com/JDOSTech/gophercloud/testhelper/client"
)

func TestListTenants(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantsSuccessfully(t)

	count := 0
	err := List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := ExtractTenants(page)
		th.AssertNoErr(t, err)

		th.CheckDeepEquals(t, ExpectedTenantSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, count, 1)
}
