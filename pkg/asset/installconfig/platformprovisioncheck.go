package installconfig

import (
	"context"
	"errors"
	"fmt"

	"github.com/openshift/installer/pkg/asset"
	azconfig "github.com/openshift/installer/pkg/asset/installconfig/azure"
	bmconfig "github.com/openshift/installer/pkg/asset/installconfig/baremetal"
	gcpconfig "github.com/openshift/installer/pkg/asset/installconfig/gcp"
	vsconfig "github.com/openshift/installer/pkg/asset/installconfig/vsphere"
	"github.com/openshift/installer/pkg/platformv2/platformv2errors"
	"github.com/openshift/installer/pkg/platformv2/platformv2registry"
	"github.com/openshift/installer/pkg/types/aws"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/baremetal"
	"github.com/openshift/installer/pkg/types/gcp"
	"github.com/openshift/installer/pkg/types/libvirt"
	"github.com/openshift/installer/pkg/types/openstack"
	"github.com/openshift/installer/pkg/types/vsphere"
)

// PlatformProvisionCheck is an asset that validates the install-config platform for
// any requirements specific for provisioning infrastructure.
type PlatformProvisionCheck struct {
}

var _ asset.Asset = (*PlatformProvisionCheck)(nil)

// Dependencies returns the dependencies for PlatformProvisionCheck
func (a *PlatformProvisionCheck) Dependencies() []asset.Asset {
	return []asset.Asset{
		&InstallConfig{},
	}
}

// Generate queries for input from the user.
func (a *PlatformProvisionCheck) Generate(dependencies asset.Parents) error {
	ic := &InstallConfig{}
	dependencies.Get(ic)

	var err error
	platform := ic.Config.Platform.Name()

	p, err := platformv2registry.GetByName(platform)
	if err == nil {
		return p.PlatformProvisionCheck(ic.Config, ic.File, ic.AWS, ic.Azure)
	} else if !errors.Is(err, platformv2errors.ErrPlatformNotRegistered) {
		return err
	}

	switch platform {
	case azure.Name:
		dnsConfig, err := ic.Azure.DNSConfig()
		if err != nil {
			return err
		}
		err = azconfig.ValidatePublicDNS(ic.Config, dnsConfig)
		if err != nil {
			return err
		}
		client, err := ic.Azure.Client()
		if err != nil {
			return err
		}
		return azconfig.ValidateForProvisioning(client, ic.Config)
	case baremetal.Name:
		err = bmconfig.ValidateProvisioning(ic.Config)
		if err != nil {
			return err
		}
	case gcp.Name:
		client, err := gcpconfig.NewClient(context.TODO())
		if err != nil {
			return err
		}
		err = gcpconfig.ValidatePreExitingPublicDNS(client, ic.Config)
		if err != nil {
			return err
		}
	case vsphere.Name:
		err = vsconfig.ValidateForProvisioning(ic.Config)
		if err != nil {
			return err
		}
	case aws.Name, libvirt.Name, openstack.Name:
		// no special provisioning requirements to check
	default:
		err = fmt.Errorf("unknown platform type %q", platform)
	}
	return err
}

// Name returns the human-friendly name of the asset.
func (a *PlatformProvisionCheck) Name() string {
	return "Platform Provisioning Check"
}
