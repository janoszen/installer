package ovirt

import (
	"github.com/openshift/installer/pkg/platformv2"
	"github.com/openshift/installer/pkg/types/ovirt"
)

func init() {
	platformv2.Register(ovirt.Name, &ovirtFactory{})
}
