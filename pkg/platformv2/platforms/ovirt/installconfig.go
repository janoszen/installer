package ovirt

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig/aws"
	icazure "github.com/openshift/installer/pkg/asset/installconfig/azure"
	ovirtconfig "github.com/openshift/installer/pkg/platformv2/platforms/ovirt/installconfig"
	"github.com/openshift/installer/pkg/types"
)

func (p *ovirtPlatform) AddToInstallConfigPlatform(a *types.Platform) (err error) {
	a.Ovirt, err = ovirtconfig.Platform()
	return err
}

func (p *ovirtPlatform) Validate(
	Config *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *icazure.Metadata,
) error {
	return ovirtconfig.Validate(Config)
}
