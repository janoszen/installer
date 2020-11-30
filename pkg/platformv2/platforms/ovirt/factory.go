package ovirt

import (
	"github.com/openshift/installer/pkg/platformv2"
)

type Factory struct {
}

func (o *Factory) Create() (platformv2.PlatformV2, error) {
	return &ovirtPlatform{}, nil
}
