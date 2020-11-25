package ovirt

import (
	"github.com/openshift/installer/pkg/asset/installconfig"
	ovirtconfig "github.com/openshift/installer/pkg/platformv2/ovirt/installconfig"
	"github.com/openshift/installer/pkg/types"
)

func (p *ovirtPlatform) AddToInstallConfigPlatform(a *types.Platform) (err error) {
	a.Ovirt, err = ovirtconfig.Platform()
	return err
}

func (p *ovirtPlatform) Validate(ic *installconfig.InstallConfig) error {
	return ovirtconfig.Validate(ic.Config)
}
