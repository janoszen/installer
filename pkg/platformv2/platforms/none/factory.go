package none

import (
	"github.com/openshift/installer/pkg/platformv2"
)

// NoneFactory creates a "none" platformv2 representation.
type Factory struct {
}

func (n *Factory) Create() (platformv2.PlatformV2, error) {
	return &nonePlatform{}, nil
}
