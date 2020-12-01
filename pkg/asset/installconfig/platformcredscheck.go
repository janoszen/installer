package installconfig

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/openshift/installer/pkg/asset"
	gcpconfig "github.com/openshift/installer/pkg/asset/installconfig/gcp"
	openstackconfig "github.com/openshift/installer/pkg/asset/installconfig/openstack"
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

// PlatformCredsCheck is an asset that checks the platform credentials, asks for them or errors out if invalid
// the cluster.
type PlatformCredsCheck struct {
}

var _ asset.Asset = (*PlatformCredsCheck)(nil)

// Dependencies returns the dependencies for PlatformCredsCheck
func (a *PlatformCredsCheck) Dependencies() []asset.Asset {
	return []asset.Asset{
		&InstallConfig{},
	}
}

// Generate queries for input from the user.
func (a *PlatformCredsCheck) Generate(dependencies asset.Parents) error {
	ctx := context.TODO()
	ic := &InstallConfig{}
	dependencies.Get(ic)

	platform := ic.Config.Platform.Name()

	p, err := platformv2registry.GetByName(platform)
	if err == nil {
		return p.PlatformCredsCheck(ic.Config, ic.File, ic.AWS, ic.Azure)
	} else if !errors.Is(err, platformv2errors.ErrPlatformNotRegistered) {
		return err
	}

	switch platform {
	case aws.Name:
		_, err := ic.AWS.Session(ctx)
		if err != nil {
			return err
		}
	case gcp.Name:
		_, err = gcpconfig.GetSession(ctx)
		if err != nil {
			return errors.Wrap(err, "creating GCP session")
		}
	case openstack.Name:
		_, err = openstackconfig.GetSession(ic.Config.Platform.OpenStack.Cloud)
		if err != nil {
			return errors.Wrap(err, "creating OpenStack session")
		}
	case baremetal.Name, libvirt.Name, vsphere.Name:
		// no creds to check
	case azure.Name:
		_, err = ic.Azure.Session()
		if err != nil {
			return errors.Wrap(err, "creating Azure session")
		}
	default:
		err = fmt.Errorf("unknown platform type %q", platform)
	}

	return err
}

// Name returns the human-friendly name of the asset.
func (a *PlatformCredsCheck) Name() string {
	return "Platform Credentials Check"
}
