package ovirt

import (
	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig/aws"
	"github.com/openshift/installer/pkg/asset/installconfig/azure"
	ovirtconfig "github.com/openshift/installer/pkg/platformv2/platforms/ovirt/installconfig"
	"github.com/openshift/installer/pkg/types"
)

func (p *ovirtPlatform) PlatformCredsCheck(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *azure.Metadata,
) error {
	con, err := ovirtconfig.NewConnection()
	if err != nil {
		return errors.Wrap(err, "creating Engine connection")
	}
	err = con.Test()
	if err != nil {
		return errors.Wrap(err, "testing Engine connection")
	}
	return nil
}

func (p *ovirtPlatform) PlatformPermsCheck(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *azure.Metadata,
) error {
	return nil
}

func (p *ovirtPlatform) PlatformProvisionCheck(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *azure.Metadata,
) error {
	return nil
}
