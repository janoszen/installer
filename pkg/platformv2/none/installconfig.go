package ovirt

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/none"
)

func (p *nonePlatform) AddToInstallConfigPlatform(a *types.Platform) error {
	a.None = &none.Platform{}
	return nil
}

func (p *nonePlatform) Validate(ic *installconfig.InstallConfig) error {
	return field.ErrorList{}.ToAggregate()
}
