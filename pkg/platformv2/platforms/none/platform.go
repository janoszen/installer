package none

import (
	"github.com/openshift/installer/pkg/platformv2"
	"github.com/openshift/installer/pkg/types/none"
)

type nonePlatform struct {
}

func (p *nonePlatform) Name() string {
	return none.Name
}

func (p *nonePlatform) SupportsIPI() bool {
	return false
}

func (p *nonePlatform) GetIPI() (platformv2.IPIPlatform, error) {
	return nil, platformv2.ErrPlatformDoesNotSupportIPI
}
