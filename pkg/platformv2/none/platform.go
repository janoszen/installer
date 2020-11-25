package ovirt

import (
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/platformv2/abstract"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/none"
)

type nonePlatform struct {
}

func (p *nonePlatform) Metadata(_ *types.ClusterMetadata, _ *installconfig.InstallConfig) error {
	return nil
}

func (p *nonePlatform) Name() string {
	return none.Name
}

func (p *nonePlatform) SupportsIPI() bool {
	return false
}

func (p *nonePlatform) GetIPI() (abstract.IPIPlatform, error) {
	return nil, abstract.NotAnIPIPlatform
}
