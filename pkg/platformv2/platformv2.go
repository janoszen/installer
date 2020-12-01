package platformv2

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig/aws"
	icazure "github.com/openshift/installer/pkg/asset/installconfig/azure"
	"github.com/openshift/installer/pkg/types"
)

// PlatformV2 is an abstraction for supported platforms.
type PlatformV2 interface {
	// Name returns the name of the platform.
	Name() string

	// SupportsIPI returns true if the platform supports IPI.
	SupportsIPI() bool

	// GetIPI returns the handler for IPI-specific calls.
	//
	// Returns platformv2errors.ErrPlatformDoesNotSupportIPI if the platform does not support IPI.
	GetIPI() (IPIPlatform, error)

	// AddToInstallConfigPlatform adds the current platform to the installconfig.
	AddToInstallConfigPlatform(p *types.Platform) error

	// Validate validates the install config.
	//
	// Note: this func accepts the components of the install config directly because the install config
	// interface and implementation was not cleanly separated and the implementation needs to depend on the
	// platformv2 package.
	Validate(
		Config *types.InstallConfig,
		File *asset.File,
		AWS *aws.Metadata,
		Azure *icazure.Metadata,
	) error

	// PlatformCredsCheck validates the platform credentials.
	//
	// Note: this func accepts the components of the install config directly because the install config
	// interface and implementation was not cleanly separated and the implementation needs to depend on the
	// platformv2 package.
	PlatformCredsCheck(
		Config *types.InstallConfig,
		File *asset.File,
		AWS *aws.Metadata,
		Azure *icazure.Metadata,
	) error

	// PlatformPermsCheck validates the platform permissions.
	//
	// Note: this func accepts the components of the install config directly because the install config
	// interface and implementation was not cleanly separated and the implementation needs to depend on the
	// platformv2 package.
	PlatformPermsCheck(
		Config *types.InstallConfig,
		File *asset.File,
		AWS *aws.Metadata,
		Azure *icazure.Metadata,
	) error

	// PlatformProvisionCheck validates the if provisioning can commence on the platform.
	//
	// Note: this func accepts the components of the install config directly because the install config
	// interface and implementation was not cleanly separated and the implementation needs to depend on the
	// platformv2 package.
	PlatformProvisionCheck(
		Config *types.InstallConfig,
		File *asset.File,
		AWS *aws.Metadata,
		Azure *icazure.Metadata,
	) error
}
