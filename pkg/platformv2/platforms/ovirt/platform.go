package ovirt

import (
	"github.com/openshift/installer/pkg/platformv2"
	"github.com/openshift/installer/pkg/types/ovirt"
)

type ovirtPlatform struct {
}

func (p *ovirtPlatform) Name() string {
	return ovirt.Name
}

func (p *ovirtPlatform) SupportsIPI() bool {
	return true
}

func (p *ovirtPlatform) GetIPI() (platformv2.IPIPlatform, error) {
	return p, nil
}
