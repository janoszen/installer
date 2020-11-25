package ovirt

import (
	"github.com/openshift/installer/pkg/platformv2/abstract"
	"github.com/openshift/installer/pkg/types/ovirt"
)

type ovirtPlatform struct {
	// removeTemplate is true if the template should be removed after installation.
	removeTemplate bool
}

func (p *ovirtPlatform) Name() string {
	return ovirt.Name
}

func (p *ovirtPlatform) SupportsIPI() bool {
	return true
}

func (p *ovirtPlatform) GetIPI() (abstract.IPIPlatform, error) {
	return p, nil
}
