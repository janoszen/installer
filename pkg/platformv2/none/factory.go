package ovirt

import (
	"github.com/openshift/installer/pkg/platformv2/abstract"
)

type noneFactory struct {
}

func (n *noneFactory) Create() (abstract.PlatformV2, error) {
	return &nonePlatform{}, nil
}
