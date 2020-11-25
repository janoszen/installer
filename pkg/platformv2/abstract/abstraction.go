package abstract

import (
	"errors"

	machineapi "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/machines"
	"github.com/openshift/installer/pkg/asset/rhcos"
	"github.com/openshift/installer/pkg/types"
)

// NotAnIPIPlatform is returned when the platform does not support IPI.
var NotAnIPIPlatform = errors.New("this is not an IPI platform")

// PlatformV2 is an abstraction for supported platforms.
type PlatformV2 interface {
	// Name returns the name of the platform.
	Name() string

	// Metadata translates the installConfig to the appropriate entry in clusterMetadata.
	Metadata(clusterMetadata *types.ClusterMetadata, installConfig *installconfig.InstallConfig) error

	// SupportsIPI returns true if the platform supports IPI
	SupportsIPI() bool

	// AddToInstallConfigPlatform adds the current platform to the installconfig.
	AddToInstallConfigPlatform(p *types.Platform) error

	// Validate validates the install config
	Validate(ic *installconfig.InstallConfig) error

	// GetNoProxyIPs returns a list of IPs that should not be routed through the proxy.
	GetNoProxyIPs() []string

	// GetNoProxyHostnames returns a list of host names that should not be routed through the proxy.
	GetNoProxyHostnames() []string

	// GetIPI returns the handler for IPI-specific calls. Returns NotAnIPIPlatform if the platform does not support IPI.
	GetIPI() (IPIPlatform, error)
}

type IPIPlatform interface {
	// CreateTFVars is responsible for creating a .tfvars
	// file for the installation
	CreateTFVars(
		installConfig *installconfig.InstallConfig,
		mastersAsset *machines.Master,
		rhcosImage *rhcos.Image,
		clusterID *installconfig.ClusterID,
	) ([]byte, error)

	// GetMasterMachines returns a pool of machine configurations to create for the master nodes.
	GetMasterMachines(
		ic *types.InstallConfig,
		pool *types.MachinePool,
		rhcosImage *rhcos.Image,
		clusterID *installconfig.ClusterID,
	) ([]machineapi.Machine, error)

	// AddWorkerMachines adds a set of machines to the worker pool.
	AddWorkerMachines(
		ic *types.InstallConfig,
		pool *types.MachinePool,
		rhcosImage *rhcos.Image,
		clusterID *installconfig.ClusterID,
		machineSets []runtime.Object,
	) error
}
