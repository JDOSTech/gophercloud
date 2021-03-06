package cdnobjects

import (
	"testing"

	os "github.com/JDOSTech/gophercloud/openstack/objectstorage/v1/objects"
	th "github.com/JDOSTech/gophercloud/testhelper"
	fake "github.com/JDOSTech/gophercloud/testhelper/client"
)

func TestDeleteCDNObject(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleDeleteObjectSuccessfully(t)

	res := Delete(fake.ServiceClient(), "testContainer", "testObject", nil)
	th.AssertNoErr(t, res.Err)

}
