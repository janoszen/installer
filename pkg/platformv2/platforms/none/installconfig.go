package none

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig/aws"
	icazure "github.com/openshift/installer/pkg/asset/installconfig/azure"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/none"
)

func (p *nonePlatform) AddToInstallConfigPlatform(a *types.Platform) error {
	a.None = &none.Platform{}
	return nil
}

func (p *nonePlatform) Validate(
	_ *types.InstallConfig,
	_ *asset.File,
	_ *aws.Metadata,
	_ *icazure.Metadata,
) error {
	return field.ErrorList{}.ToAggregate()
}
