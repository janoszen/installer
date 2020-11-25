package ovirt

import (
	ovirtprovider "github.com/openshift/cluster-api-provider-ovirt/pkg/apis/ovirtprovider/v1beta1"
	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/machines"
	"github.com/openshift/installer/pkg/asset/rhcos"
	ovirtconfig "github.com/openshift/installer/pkg/platformv2/ovirt/installconfig"
	ovirttfvars "github.com/openshift/installer/pkg/tfvars/ovirt"
)

func (p *ovirtPlatform) CreateTFVars(
	installConfig *installconfig.InstallConfig,
	mastersAsset *machines.Master,
	rhcosImage *rhcos.Image,
	clusterID *installconfig.ClusterID,
) ([]byte, error) {
	config, err := ovirtconfig.NewConfig()
	if err != nil {
		return nil, err
	}
	con, err := ovirtconfig.NewConnection()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	if installConfig.Config.Platform.Ovirt.VNICProfileID == "" {
		profiles, err := ovirtconfig.FetchVNICProfileByClusterNetwork(
			con,
			installConfig.Config.Platform.Ovirt.ClusterID,
			installConfig.Config.Platform.Ovirt.NetworkName)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to compute values for Engine platform")
		}
		if len(profiles) != 1 {
			return nil, errors.Wrapf(err, "failed to compute values for Engine platform, there are multiple vNIC profiles.")
		}
		installConfig.Config.Platform.Ovirt.VNICProfileID = profiles[0].MustId()
	}

	masters, err := mastersAsset.Machines()
	if err != nil {
		return nil, err
	}

	data, err := ovirttfvars.TFVars(
		ovirttfvars.Auth{
			URL:      config.URL,
			Username: config.Username,
			Password: config.Password,
			Cafile:   config.CAFile,
		},
		installConfig.Config.Platform.Ovirt.ClusterID,
		installConfig.Config.Platform.Ovirt.StorageDomainID,
		installConfig.Config.Platform.Ovirt.NetworkName,
		installConfig.Config.Platform.Ovirt.VNICProfileID,
		string(*rhcosImage),
		clusterID.InfraID,
		masters[0].Spec.ProviderSpec.Value.Object.(*ovirtprovider.OvirtMachineProviderSpec),
	)
	if err != nil {
		return nil, err
	}
	return data, nil
}
