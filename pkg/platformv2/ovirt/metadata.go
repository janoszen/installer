package ovirt

import (
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ovirt"
)

func (p *ovirtPlatform) Metadata(clusterMetadata *types.ClusterMetadata, installConfig *installconfig.InstallConfig) error {
	clusterMetadata.ClusterPlatformMetadata.Ovirt = &ovirt.Metadata{
		ClusterID:      installConfig.Config.Ovirt.ClusterID,
		RemoveTemplate: p.removeTemplate,
	}
	return nil
}
