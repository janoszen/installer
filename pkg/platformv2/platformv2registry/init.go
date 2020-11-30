package platformv2registry

import (
	nonePlatform "github.com/openshift/installer/pkg/platformv2/platforms/none"
	ovirtPlatform "github.com/openshift/installer/pkg/platformv2/platforms/ovirt"
	"github.com/openshift/installer/pkg/types/none"
	"github.com/openshift/installer/pkg/types/ovirt"
)

func init()  {
	Register(none.Name, &nonePlatform.Factory{})
	Register(ovirt.Name, &ovirtPlatform.Factory{})
}
