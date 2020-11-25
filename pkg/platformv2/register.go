package platformv2

import (
	"github.com/openshift/installer/pkg/platformv2/abstract"

	// Importing platforms for registration
	_ "github.com/openshift/installer/pkg/platformv2/none"
	_ "github.com/openshift/installer/pkg/platformv2/ovirt"
)

// Register registers a platform factory with the given name.
func Register(name string, factory abstract.PlatformFactory) {
	platforms[name] = factory
}
