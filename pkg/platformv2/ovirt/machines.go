package ovirt

import (
	machineapi "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/machines/ovirt"
	"github.com/openshift/installer/pkg/asset/rhcos"
	rhcosutils "github.com/openshift/installer/pkg/rhcos"
	"github.com/openshift/installer/pkg/types"
	ovirttypes "github.com/openshift/installer/pkg/types/ovirt"
)

func (p *ovirtPlatform) defaultOvirtMachinePoolPlatform() ovirttypes.MachinePool {
	return ovirttypes.MachinePool{
		CPU: &ovirttypes.CPU{
			Cores:   4,
			Sockets: 1,
		},
		MemoryMB: 16348,
		OSDisk: &ovirttypes.Disk{
			SizeGB: 120,
		},
		VMType: ovirttypes.VMTypeServer,
	}
}

func (p *ovirtPlatform) GetMasterMachines(ic *types.InstallConfig, pool *types.MachinePool, rhcosImage *rhcos.Image, clusterID *installconfig.ClusterID) ([]machineapi.Machine, error) {
	mpool := p.defaultOvirtMachinePoolPlatform()
	mpool.VMType = ovirttypes.VMTypeHighPerformance
	mpool.Set(ic.Platform.Ovirt.DefaultMachinePlatform)
	mpool.Set(pool.Platform.Ovirt)
	pool.Platform.Ovirt = &mpool

	imageName, _ := rhcosutils.GenerateOpenStackImageName(string(*rhcosImage), clusterID.InfraID)

	var machines []machineapi.Machine
	var err error
	machines, err = ovirt.Machines(clusterID.InfraID, ic, pool, imageName, "master", "master-user-data")
	if err != nil {
		return nil, errors.Wrap(err, "failed to create master machine objects for ovirt provider")
	}
	return machines, err
}

func (p *ovirtPlatform) AddWorkerMachines(ic *types.InstallConfig, pool *types.MachinePool, rhcosImage *rhcos.Image, clusterID *installconfig.ClusterID, machineSets []runtime.Object) error {
	mpool := p.defaultOvirtMachinePoolPlatform()
	mpool.Set(ic.Platform.Ovirt.DefaultMachinePlatform)
	mpool.Set(pool.Platform.Ovirt)
	pool.Platform.Ovirt = &mpool

	imageName, _ := rhcosutils.GenerateOpenStackImageName(string(*rhcosImage), clusterID.InfraID)

	sets, err := ovirt.MachineSets(clusterID.InfraID, ic, pool, imageName, "worker", "worker-user-data")
	if err != nil {
		return errors.Wrap(err, "failed to create worker machine objects for ovirt provider")
	}
	for _, set := range sets {
		machineSets = append(machineSets, set)
	}
	return nil
}
