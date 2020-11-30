package platformv2registry

import (
	"github.com/openshift/installer/pkg/platformv2"
)

// Register registers a platform factory with the given name.
func Register(name string, factory platformv2.PlatformFactory) {
	platforms[name] = factory
}
