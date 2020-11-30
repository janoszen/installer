package platformv2registry

import (
	"github.com/openshift/installer/pkg/platformv2"

	// Importing platforms for registration
	_ "github.com/openshift/installer/pkg/platformv2/platforms/none"
	_ "github.com/openshift/installer/pkg/platformv2/platforms/ovirt"
)

var platforms = map[string]platformv2.PlatformFactory{}
