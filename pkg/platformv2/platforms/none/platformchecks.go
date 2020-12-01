package none

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig/aws"
	"github.com/openshift/installer/pkg/asset/installconfig/azure"
	"github.com/openshift/installer/pkg/types"
)

func (p *nonePlatform) PlatformCredsCheck(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *azure.Metadata,
) error {
	return nil
}

func (p *nonePlatform) PlatformPermsCheck(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *azure.Metadata,
) error {
	return nil
}

func (p *nonePlatform) PlatformProvisionCheck(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *azure.Metadata,
) error {
	return nil
}
